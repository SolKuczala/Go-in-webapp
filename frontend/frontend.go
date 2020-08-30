package frontend

import (
	"fmt"
	"net/http"

	"golang.org/x/oauth2"

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

func RedirectHandler(c *gin.Context) {
	var token oauth2.Token
	if err := c.ShouldBindJSON(&token); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if token.Valid() {
		c.JSON(http.StatusOK, gin.H{"ok": "you are logged in"})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"error": "token is not valid"})
}

func SetCookie(c *gin.Context) {
	cookie, err := c.Cookie("so-ku_cookie")

	if err != nil {
		cookie = "NotSet"
		c.SetCookie("gin_cookie", "test", 3600, "/", HOST, false, true)
	}
	fmt.Printf("Cookie value: %s \n", cookie)
}
