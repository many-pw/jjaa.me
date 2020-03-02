package persist

import "fmt"
import "jjaa.me/util"
import "github.com/jmoiron/sqlx"
import _ "github.com/go-sql-driver/mysql"

func Connection() *sqlx.DB {
	url := fmt.Sprintf("%s:%s@(%s:%d)/%s", util.AllConfig.Db.User,
		util.AllConfig.Db.Password,
		util.AllConfig.Db.Host,
		util.AllConfig.Db.Port,
		util.AllConfig.Db.Database)
	db, err := sqlx.Connect("mysql", url)
	if err != nil {
		fmt.Println(err)
	}

	return db
}
