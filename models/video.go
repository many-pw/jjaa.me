
package models

import "github.com/jmoiron/sqlx"
import "fmt"

type Video struct {
	Id        int   `json:"id"` 
	CreatedAt int64 `json:"created_at"`
}

func SelectVideos(db *sqlx.DB) ([]Video, string) {
	videos := []Video{}
	sql := fmt.Sprintf("SELECT id, UNIX_TIMESTAMP(created_at) as createdat from videos order by created_at desc")
	err := db.Select(&videos, sql)
	s := ""
	if err != nil {
		s = err.Error()
	}

	return videos, s
}
