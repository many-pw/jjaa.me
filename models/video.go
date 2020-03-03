package models

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Video struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Comments  int    `json:"comments"`
	Status    string
	CreatedAt int64 `json:"created_at"`
}

const VIDEO_SELECT = "SELECT id, title, comments, UNIX_TIMESTAMP(created_at) as createdat from videos"

func SelectVideos(db *sqlx.DB) ([]Video, string) {
	videos := []Video{}
	sql := fmt.Sprintf("%s order by created_at desc", VIDEO_SELECT)
	err := db.Select(&videos, sql)
	s := ""
	if err != nil {
		s = err.Error()
	}

	return videos, s
}
func InsertVideo(db *sqlx.DB, title string, id int) string {
	_, err := db.NamedExec("INSERT INTO videos (title, user_id, status) values (:title, :id, :status)",
		map[string]interface{}{"title": title, "id": id, "status": "not_uploaded"})
	if err != nil {
		return err.Error()
	}
	return ""
}
