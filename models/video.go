package models

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Video struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Comments    int    `json:"comments"`
	Status      string
	UrlSafeName string
	CreatedAt   int64 `json:"created_at"`
}

const VIDEO_SELECT = "SELECT id, url_safe_name as urlsafename, title, comments, UNIX_TIMESTAMP(created_at) as createdat from videos"

func SelectVideos(db *sqlx.DB, userId int) ([]Video, string) {
	videos := []Video{}
	sql := fmt.Sprintf("%s where user_id = %d order by created_at desc limit 1000", VIDEO_SELECT, userId)
	err := db.Select(&videos, sql)
	s := ""
	if err != nil {
		s = err.Error()
	}

	return videos, s
}
func InsertVideo(db *sqlx.DB, title, safeName string, id int) string {
	_, err := db.NamedExec("INSERT INTO videos (title, url_safe_name, user_id, status) values (:title, :safe_name, :id, :status)",
		map[string]interface{}{"title": title, "id": id,
			"safe_name": safeName,
			"status":    "not_uploaded"})
	if err != nil {
		return err.Error()
	}
	return ""
}
