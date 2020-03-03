package models

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type User struct {
	Id        int    `json:"id"`
	Email     string `json:"email"`
	Fans      int    `json:"fans"`
	Videos    int    `json:"videos"`
	Flavor    string `json:"flavor"`
	CreatedAt int64  `json:"created_at"`
}

const USER_SELECT = "SELECT id, fans, videos, email, UNIX_TIMESTAMP(created_at) as createdat from users"

func (u *User) Encode() string {
	b, _ := json.Marshal(u)
	s := string(b)
	sEnc := base64.StdEncoding.EncodeToString([]byte(s))
	return sEnc
}

func DecodeUser(s string) *User {
	var user User
	decoded, _ := base64.StdEncoding.DecodeString(s)
	err := json.Unmarshal([]byte(decoded), &user)
	if err != nil {
		return nil
	}
	return &user
}

func SelectUsers(db *sqlx.DB) ([]User, string) {
	users := []User{}
	sql := fmt.Sprintf("%s order by created_at desc", USER_SELECT)
	err := db.Select(&users, sql)
	s := ""
	if err != nil {
		s = err.Error()
	}

	return users, s
}
func SelectUser(db *sqlx.DB, id int) (*User, string) {
	user := User{}
	sql := fmt.Sprintf("%s where id=:id", USER_SELECT)
	rows, err := db.NamedQuery(sql, map[string]interface{}{"id": id})
	if err != nil {
		return &user, err.Error()
	} else {
		if rows.Next() {
			rows.StructScan(&user)
		}
	}

	return &user, ""
}

func IncrementUserCount(db *sqlx.DB, field string, id int) string {
	_, err := db.NamedExec(fmt.Sprintf("UPDATE users set %s=%s+1 where id=:id", field, field),
		map[string]interface{}{"id": id})
	if err != nil {
		return err.Error()
	}
	return ""
}
func UpdateUser(db *sqlx.DB, phrase, email string) string {
	_, err := db.NamedExec("UPDATE users set phrase=SHA1(:phrase) where email=:email",
		map[string]interface{}{"phrase": phrase, "email": email})
	if err != nil {
		return err.Error()
	}
	return ""
}
