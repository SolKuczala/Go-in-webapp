package frontend

import (
	"fmt"
	"log"
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
	c.HTML(http.StatusOK, "enter_profile_info.tmpl.html", nil)
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

type GoogleLogin struct {
	IDToken string `json:"id_token,required"`
}

func RedirectHandler(c *gin.Context) {
	var googleLogin GoogleLogin
	if err := c.ShouldBindJSON(&googleLogin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("RECEIVED TOKEN: %s\n", googleLogin.IDToken[:10])
	c.Redirect(http.StatusTemporaryRedirect, "/enter-profile-info")
}

func SetCookie(c *gin.Context) {
	cookie, err := c.Cookie("so-ku_cookie")

	if err != nil {
		cookie = "NotSet"
		c.SetCookie("gin_cookie", "test", 3600, "/", HOST, false, true)
	}
	fmt.Printf("Cookie value: %s \n", cookie)
}
