
package models

import "github.com/jmoiron/sqlx"
import "fmt"

type Video struct {
	Id        int   `json:"id"` 
	CreatedAt int64 `json:"created_at"`
}

const VIDEO_SELECT = "SELECT id, UNIX_TIMESTAMP(created_at) as createdat from videos"

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
func InsertVideo(db *sqlx.DB) string {
	_, err := db.NamedExec("INSERT INTO videos (col) values (:col)",
		map[string]interface{}{"": ""})
	if err != nil {
		return err.Error()
	}
	return ""
}
