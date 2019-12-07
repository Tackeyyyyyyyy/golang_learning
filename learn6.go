package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	squirrel "github.com/Masterminds/squirrel"
)

func main() {

	fmt.Println("start")

	// クエリビルダを使う
	db, err := sql.Open("mysql", "root:@/name_list")
	if err != nil {
		panic(err.Error())
	}

	// 最後に切断
	defer db.Close()

	sql, _, err := squirrel.Select("id, last_name").From("names").ToSql()

	rows, err := db.Query(sql)
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(id, name)
	}

}
