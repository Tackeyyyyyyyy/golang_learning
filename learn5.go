package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)
func main() {

	fmt.Println("start")


	db, err := sql.Open("mysql", "root:@/name_list")
	if err != nil {
		panic(err.Error())
	}

	// 最後に切断
	defer db.Close()


	// 1件取得
	id := 2
	var name string
	err = db.QueryRow("SELECT first_name FROM names WHERE id = ? LIMIT 1", id).Scan(&name)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(id, name)


	// 全件取得
	rows, err := db.Query("SELECT id, first_name FROM names ORDER BY id")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close() // 遅延クローズ


	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(id, name)
	}

	fmt.Println("end")
}
