package frontend

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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
