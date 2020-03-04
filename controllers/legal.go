package controllers

import (
	"fmt"
	"net/http"

	"github.com/andrewarrow/feedback/util"
	"github.com/gin-gonic/gin"
)

func LegalGdpr(c *gin.Context) {
	host := util.AllConfig.Http.Host
	c.SetCookie("gdpr_ok", "cookies, yes", 0, "/", host, false, false)
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

	fmt.Println("1111", host)
	c.HTML(http.StatusOK, "terms.tmpl", gin.H{
		"flash": "",
		"name":  "jjaa.me", // hint: change me
		"host":  host,
	})

}
