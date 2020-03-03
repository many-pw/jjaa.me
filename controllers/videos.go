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
		"flash": flash,
	})
}
func VideosIndex(c *gin.Context) {
	BeforeAll("", c)
	videos, _ := models.SelectVideos(Db, user.Id)
	c.HTML(http.StatusOK, "videos__index.tmpl", gin.H{
		"videos": videos,
		"user":   user,
		"flash":  flash,
	})
}
func VideosShow(c *gin.Context) {
	BeforeAll("", c)
	video, _ := models.SelectVideo(Db, c.Param("name"))
	c.HTML(http.StatusOK, "videos__show.tmpl", gin.H{
		"video": video,
		"user":  user,
		"flash": flash,
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
	title := strings.TrimSpace(c.PostForm("title"))
	if title == "" {
		SetFlash("title needed", c)
		c.Redirect(http.StatusFound, "/videos/new")
		c.Abort()
		return
	}

	babbler.Count = 4
	words := babbler.Babble()
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	safeName := reg.ReplaceAllString(strings.ToLower(words), "-")
	models.InsertVideo(Db, title, safeName, user.Id)
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
