package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"jjaa.me/api"
	"jjaa.me/models"
)

func ApiVersion(c *gin.Context) {
	ap := api.ApiResponse{}
	ap.Version = "1.0.0"
	ap.Items = []interface{}{"test1", "test2"}
	ap.SentAt = time.Now().Unix()
	c.JSON(http.StatusOK, ap)

}
func ApiLatest(c *gin.Context) {
	ap := api.ApiResponse{}
	items, _ := models.SelectVideos(Db, 0)
	ap.Items = []interface{}{}
	for _, item := range items {
		ap.Items = append(ap.Items, item)
	}
	ap.SentAt = time.Now().Unix()
	c.JSON(http.StatusOK, ap)
}
