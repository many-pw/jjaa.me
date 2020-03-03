package controllers

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"jjaa.me/models"
	"jjaa.me/util"
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
	title := c.PostForm("title")
	models.InsertVideo(Db, title, user.Id)
	models.IncrementUserCount(Db, "videos", user.Id)
	c.Redirect(http.StatusFound, "/videos/upload")
	c.Abort()
}
func VideosDestroy(c *gin.Context) {

	c.Redirect(http.StatusFound, "/")
	c.Abort()
}
func VideosFile(c *gin.Context) {
	BeforeAll("", c)
	file, _ := c.FormFile("file")
	tokens := strings.Split(file.Filename, ".")
	fmt.Println("111111", tokens)
	ext := tokens[1]
	babbler.Count = 4
	filename := babbler.Babble()
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	filename = reg.ReplaceAllString(strings.ToLower(filename), "-")
	file_with_ext := filename + "." + ext
	c.SaveUploadedFile(file, util.AllConfig.Path.Videos+file_with_ext)
	c.Redirect(http.StatusFound, "/")
	c.Abort()
}
