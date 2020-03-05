package controllers

import (
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"jjaa.me/models"
	"jjaa.me/util"
)

func VideosNew(c *gin.Context) {
	if !BeforeAll("user", c) {
		return
	}
	c.HTML(http.StatusOK, "videos__new.tmpl", gin.H{
		"flash": flash,
		"user":  user,
	})
}
func VideosAllIndex(c *gin.Context) {
	if !BeforeAll("", c) {
		return
	}
	videos, _ := models.SelectVideos(Db, 0)
	c.HTML(http.StatusOK, "videos__all_index.tmpl", gin.H{
		"videos": videos,
		"user":   user,
		"flash":  flash,
	})
}
func VideosIndex(c *gin.Context) {
	if !BeforeAll("", c) {
		return
	}
	videos, _ := models.SelectVideos(Db, user.Id)
	c.HTML(http.StatusOK, "videos__index.tmpl", gin.H{
		"videos": videos,
		"user":   user,
		"flash":  flash,
	})
}
func VideosShow(c *gin.Context) {
	if !BeforeAll("", c) {
		return
	}
	video, _ := models.SelectVideo(Db, c.Param("name"))
	c.HTML(http.StatusOK, "videos__show.tmpl", gin.H{
		"video": video,
		"user":  user,
		"flash": flash,
	})
}
func VideosUpload(c *gin.Context) {
	if !BeforeAll("", c) {
		return
	}
	video, _ := models.SelectVideo(Db, c.Param("name"))
	c.HTML(http.StatusOK, "videos__upload.tmpl", gin.H{
		"video": video,
		"flash": "",
		"user":  user,
	})

}
func VideosCreate(c *gin.Context) {
	if !BeforeAll("", c) {
		return
	}
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
	c.Redirect(http.StatusFound, "/videos/upload/"+safeName)
	c.Abort()
}
func VideosDestroy(c *gin.Context) {

	c.Redirect(http.StatusFound, "/")
	c.Abort()
}
func VideosFile(c *gin.Context) {
	BeforeAll("", c)
	video, _ := models.SelectVideo(Db, c.Param("name"))
	file, _ := c.FormFile("file")
	tokens := strings.Split(file.Filename, ".")
	ext := tokens[len(tokens)-1]
	fileWithExt := video.UrlSafeName + "." + ext
	c.SaveUploadedFile(file, util.AllConfig.Path.Videos+fileWithExt)
	models.UpdateVideo(Db, "uploaded", video.UrlSafeName)
	go convertVideoFile(fileWithExt, video.UrlSafeName)
	c.Redirect(http.StatusFound, "/")
	c.Abort()
}
func convertVideoFile(fileWithExt, filename string) {
	//ffmpeg -ss 00:00:03 -i input -vframes 1 -q:v 2 output.jpg
	exec.Command("ffmpeg", "-ss", "00:00:03", "-i",
		util.AllConfig.Path.Videos+fileWithExt,
		"-vframes", "1", "-q:v", "2",
		util.AllConfig.Path.Videos+filename+".jpg").Output()
	models.UpdateVideo(Db, "jpg_ready", filename)
	exec.Command("ffmpeg", "-i",
		util.AllConfig.Path.Videos+fileWithExt,
		"-vcodec", "h264", "-acodec", "aac",
		util.AllConfig.Path.Videos+filename+".mp4").Output()
	models.UpdateVideo(Db, "mp4_ready", filename)
	exec.Command("ffmpeg", "-i",
		util.AllConfig.Path.Videos+fileWithExt,
		util.AllConfig.Path.Videos+filename+".webm").Output()
	models.UpdateVideo(Db, "webm_ready", filename)
	exec.Command("ffmpeg", "-i",
		util.AllConfig.Path.Videos+fileWithExt,
		util.AllConfig.Path.Videos+filename+".m4a").Output()
	models.UpdateVideo(Db, "m4a_ready", filename)
	exec.Command("ffmpeg", "-i",
		util.AllConfig.Path.Videos+fileWithExt,
		util.AllConfig.Path.Videos+filename+".oga").Output()

	models.UpdateVideo(Db, "live", filename)
	os.Remove(util.AllConfig.Path.Videos + fileWithExt)
}
