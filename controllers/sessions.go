package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tjarratt/babble"
	e "jjaa.me/email"
	"jjaa.me/models"
	"jjaa.me/util"
)

var babbler = babble.NewBabbler()

func SessionsNew(c *gin.Context) {
	BeforeAll("", c)
	lastEmail, _ := c.Cookie("lastEmail")
	c.HTML(http.StatusOK, "sessions__new.tmpl", gin.H{
		"lastEmail": lastEmail,
		"flash":     "",
		"name":      "name",
	})

}
func SessionsCreate(c *gin.Context) {
	user := models.User{}
	host := util.AllConfig.Http.Host
	email := c.PostForm("email")
	c.SetCookie("lastEmail", email, 3600, "/", host, false, false)
	password := c.PostForm("password")
	flash := ""

	if !strings.Contains(email, "@") || !strings.Contains(email, ".") || len(email) < 7 {
		flash = "not valid email"
	} else {
		user.Email = email
		user.Flavor = "user"
		sql := fmt.Sprintf("SELECT id, email, flavor from users where email=:email and phrase=SHA1(:phrase)")
		rows, err := Db.NamedQuery(sql, map[string]interface{}{"email": email, "phrase": password})
		if err != nil {
			flash = err.Error()
		} else {
			if rows.Next() {
				rows.StructScan(&user)
				c.SetCookie("user", user.Encode(), 3600*24*365, "/", host, false, false)
			} else {
				babbler.Count = 4
				phrase := babbler.Babble()
				fmt.Println(phrase)
				m := map[string]interface{}{"email": email, "phrase": phrase, "flavor": "user"}
				_, err = Db.NamedExec(`INSERT INTO users (email, phrase, flavor) 
values (:email, SHA1(:phrase), :flavor)`, m)
				if err != nil {
					flash = "was not able to login, review your emails from us."
					models.UpdateUser(Db, phrase, email)
					go e.Send(email, "info@many.pw", "your jjaa.me info", phrase)
				} else {
					go e.Send(email, "info@many.pw", "welcome to jjaa.me", phrase)
					flash = "check your email for your pass phrase"
				}
			}
		}
	}
	c.SetCookie("flash", flash, 3600, "/", host, false, false)
	c.Redirect(http.StatusFound, "/")
	c.Abort()
}
func SessionsDestroy(c *gin.Context) {
	host := util.AllConfig.Http.Host
	c.SetCookie("user", "", 3600, "/", host, false, false)

	c.Redirect(http.StatusFound, "/")
	c.Abort()
}
