package db

import(
	"database/sql"
	"github.com/go-sql-driver/mysql"
)

type Database struct {
	
}

func ConnectToDB() {
	db, err := sql.Open("mysql","user:pass@/gowldb")
	if err != nil {
		panic(err)
	}
	err = db.Ping()
}