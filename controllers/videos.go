
package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)
func VideosNew(c *gin.Context) {
	BeforeAll("", c)
	c.HTML(http.StatusOK, "videos__new.tmpl", gin.H{
		"flash": "",
	})

}
func VideosCreate(c *gin.Context) {
	BeforeAll("", c)
	c.Redirect(http.StatusFound, "/")
	c.Abort()
}
func VideosDestroy(c *gin.Context) {

	c.Redirect(http.StatusFound, "/")
	c.Abort()
}
