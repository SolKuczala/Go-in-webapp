package frontend

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const HOST = "localhost"

//const HOST = "so-ku.herokuapp"

type WebUrls struct {
	Login            string
	SignUp           string
	EnterProfileInfo string
	MainProfile      string
	ForgotPassword   string
	ResetPassword    string
}

//ver como servir las paginas
func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.tmpl.html", gin.H{
		"title": "Home",
	})
}

func SignUp(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.tmpl.html", nil)
}

func EnterProfileInfo(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.tmpl.html", nil)
}

func MainProfile(c *gin.Context) {
	c.HTML(http.StatusOK, "main_profile.tmpl.html", nil)
}

func ForgotPassword(c *gin.Context) {
	c.HTML(http.StatusOK, "forgot-pass.tmpl.html", nil)
}

func ResetPassword(c *gin.Context) {
	c.HTML(http.StatusOK, "reset-pass.tmpl.html", nil)
}

func RedirectHandler(c *gin.Context) {
	state := c.Query("state")
	code := c.Query("code")

	fmt.Printf("%#v", state)
	fmt.Printf("%#v", code)

	// If no errors, show provider name
	//c.Writer.Write([]byte("Hi, " + user.FullName))
}

func SetCookie(c *gin.Context) {
	cookie, err := c.Cookie("so-ku_cookie")

	if err != nil {
		cookie = "NotSet"
		c.SetCookie("gin_cookie", "test", 3600, "/", HOST, false, true)
	}
	fmt.Printf("Cookie value: %s \n", cookie)
}
