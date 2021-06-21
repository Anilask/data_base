package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:Mylaptop56@tcp(localhost:3306)/test")

	defer db.Close()
	if err != nil {
		log.Fatal(err)

	}
	sql := "DELETE FROM cities WHERE id IN (1,2,3)"
	res, err := db.Exec(sql)

	if err != nil {
		panic(err.Error())
	}
	affeectedRows, err := res.RowsAffected()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The statement affected %d rows \n", affeectedRows)
}
