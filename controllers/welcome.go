package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"jjaa.me/models"
	"jjaa.me/util"
)

var Db *sqlx.DB
var flash = ""
var user *models.User

func ValidAdminUser(c *gin.Context) bool {
	json, _ := c.Cookie("user")
	user = models.DecodeUser(json)
	if user == nil || user.Flavor != "admin" {
		SetFlash("you need to login", c)
		c.Redirect(http.StatusFound, "/sessions/new")
		c.Abort()
		return false
	}
	return true
}
func BeforeAll(flavor string, c *gin.Context) bool {
	gdpr, _ := c.Cookie("gdpr_ok")
	if gdpr == "" {
		c.HTML(http.StatusOK, "gdpr.tmpl", gin.H{
			"user":  user,
			"flash": flash,
		})
		return false
	}
	flash, _ = c.Cookie("flash")
	SetFlash("", c)

	json, _ := c.Cookie("user")
	user = models.DecodeUser(json)

	if flavor == "" {
		return true
	}

	if flavor == "user" {
		if user == nil {
			SetFlash("you need to login", c)
			c.Redirect(http.StatusFound, "/sessions/new")
			c.Abort()
			return false
		}
		return true
	}
	if user == nil || user.Flavor != "admin" {
		SetFlash("you need to login", c)
		c.Redirect(http.StatusFound, "/sessions/new")
		c.Abort()
		return false
	}
	return true
}

func SetFlash(s string, c *gin.Context) {
	host := util.AllConfig.Http.Host
	c.SetCookie("flash", s, 3600, "/", host, false, false)
}

func WelcomeIndex(c *gin.Context) {
	if !BeforeAll("", c) {
		return
	}

	if user == nil {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"user":  nil,
			"flash": flash,
		})
		return
	}

	user, _ = models.SelectUser(Db, user.Id)

	c.HTML(http.StatusOK, "homepage.tmpl", gin.H{
		"user":  user,
		"flash": flash,
	})
}
