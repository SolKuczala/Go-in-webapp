package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"so_ku/frontend"
	"so_ku/storage"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var DB *storage.Storage

func main() {

	DB = &storage.Storage{}

	DB.Init()
	defer DB.Close()

	port := os.Getenv("PORT") //5000
	if port == "" {
		port = "3000"
		//log.Fatal("$PORT must be set")
	}

	router := gin.Default()
	router.Use(gin.Logger())
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	apiPrefix := "/api/v1"
	router.GET(apiPrefix+"/user", userGetProfile)
	router.GET(apiPrefix+"/v1/auth/token", createToken)
	router.POST(apiPrefix+"/sing-up", signUp) // http -v --json POST 127.0.0.1:3000/api/v1/sing-up id:=0 email=pepe@email password=pass123
	router.POST(apiPrefix+"/login", login)
	router.POST(apiPrefix+"/user", userEditProfile)
	router.POST(apiPrefix+"/logout", logout)
	router.POST(apiPrefix+"/reset-pass", resetPassword)

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Running",
		})
	})

	webURL := frontend.WebUrls{
		Login:            "/login",
		SignUp:           "/sign-up",
		EnterProfileInfo: "/enter-profile-info",
		MainProfile:      "/main-profile",
		ForgotPassword:   "/forget-pass",
		ResetPassword:    "/reset-pass",
	}
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET(webURL.Login, frontend.Login)
	router.GET(webURL.SignUp, frontend.SignUp)
	router.GET(webURL.EnterProfileInfo, frontend.EnterProfileInfo)
	router.GET(webURL.MainProfile, frontend.MainProfile)
	router.GET(webURL.ForgotPassword, frontend.ForgotPassword)
	router.GET(webURL.ResetPassword, frontend.ResetPassword)
	router.POST("/auth/google", frontend.RedirectHandler)
	router.GET("/cookie", frontend.SetCookie)

	log.Fatal(router.Run(fmt.Sprintf(":%s", port)))
}

func logingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func createToken(c *gin.Context) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "auth-app",
		"sub": "medium",
		"aud": "any",
		"exp": time.Now().Add(time.Minute * 5).Unix(),
	})
	jwtToken, _ := token.SignedString([]byte("secret"))
	c.JSON(200, gin.H{
		"token": jwtToken,
	})
}

func signUp(c *gin.Context) {
	var user storage.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := DB.SaveNewUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

//func validateUser(ctx context.Context, r *http.Request, userName, password string) (auth.Info, error) {
//	if userName == "medium" && password == "medium" {
//		return auth.NewDefaultUser("medium", "1", nil, nil), nil
//	}
//	return nil, fmt.Errorf("Invalid credentials")
//}

//I can use gorilla/csrf toprotect tokens and submits requests
//func verifyToken(ctx context.Context, r *http.Request, tokenString string) (auth.Info, error) {
//	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
//		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
//		}
//		return []byte("secret"), nil
//	})
//	if err != nil {
//		return nil, err
//	}
//	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
//		user := auth.NewDefaultUser(claims["medium"].(string), "", nil, nil)
//		return user, nil
//	}
//	return nil, fmt.Errorf("Invaled token")
//}

func login(c *gin.Context) {
	var loginData storage.User
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := DB.GetUserAuth(&loginData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if loginData.ID == 0 {
		// user login data didn't match
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	// TODO: set session
	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
}

func userGetProfile(c *gin.Context) {

}

func userEditProfile(c *gin.Context) {
	//json: info: {full name, address, tel, email}
}

func logout(c *gin.Context) {
	//json: {email?}
}

func resetPassword(c *gin.Context) {
	//json: pass:{old, new}
}
