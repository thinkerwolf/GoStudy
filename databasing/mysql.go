package databasing

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

const (
	userName = "root"
	password = "123"
	ip       = "127.0.0.1"
	port     = "3306"
	dbName   = "test"
)

func MysqlDatabase() {
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	fmt.Println("Mysql url path ", path)

	db, err := sql.Open("mysql", path)
	checkErr(err)

	stmt, e1 := db.Prepare("select id, title from blog")
	checkErr(e1)

	rows, e2 := stmt.Query()
	checkErr(e2)
	fmt.Println(rows)

	var id int
	var title string
	for b := rows.Next(); b; b = rows.Next() {
		rows.Scan(&id, &title)
		fmt.Println("blog [id=", id, ",title=", title, "]")
	}

	db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
