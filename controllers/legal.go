package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"jjaa.me/util"
)

func LegalGdpr(c *gin.Context) {
	host := util.AllConfig.Http.Host
	c.SetCookie("gdpr_ok", "cookies, yes", 2147483647, "/", host, false, false)
	c.Redirect(http.StatusFound, "/")
	c.Abort()
}

func LegalPrivacy(c *gin.Context) {
	host := util.AllConfig.Http.Host
	c.HTML(http.StatusOK, "privacy.tmpl", gin.H{
		"flash": "",
		"name":  "jjaa.me", // hint: change me
		"host":  host,
	})

}
func LegalTerms(c *gin.Context) {
	host := util.AllConfig.Http.Host

	c.HTML(http.StatusOK, "terms.tmpl", gin.H{
		"flash": "",
		"name":  "jjaa.me", // hint: change me
		"host":  host,
	})

}
