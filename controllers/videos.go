package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func VideosNew(c *gin.Context) {
	BeforeAll("", c)
	c.HTML(http.StatusOK, "videos__new.tmpl", gin.H{
		"flash": "",
	})

}
func VideosUpload(c *gin.Context) {
	BeforeAll("", c)
	c.HTML(http.StatusOK, "videos__upload.tmpl", gin.H{
		"flash": "",
	})

}
func VideosCreate(c *gin.Context) {
	BeforeAll("", c)
	c.Redirect(http.StatusFound, "/videos/upload")
	c.Abort()
}
func VideosDestroy(c *gin.Context) {

	c.Redirect(http.StatusFound, "/")
	c.Abort()
}
