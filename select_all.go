package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type City struct {
	Id         int
	Name       string
	Population int
}

func main() {
	db, err := sql.Open("mysql", "root:Mylaptop56@tcp(localhost:3306)/test")
	fmt.Println("DB connection start")

	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}
	res, err := db.Query("SELECT * from cities")
	defer res.Close()

	if err != nil {
		log.Fatal(err)
	}
	for res.Next() {
		var city City
		err := res.Scan(&city.Id, &city.Name, &city.Population)

		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%v\n", city)
	}
}
