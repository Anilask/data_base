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
	Population string
}

func main() {
	db, err := sql.Open("mysql", "root:Mylaptop56@tcp(localhost:3306)/test")
	fmt.Println("Open DB")

	if err != nil {
		log.Fatal(err)
	}
	var myid int = 9
	res, err := db.Query("select *from cities where id =? ", myid)
	defer res.Close()

	if err != nil {
		log.Fatal(err)
	}
	if res.Next() {
		var city City
		err := res.Scan(&city.Id, &city.Name, &city.Population)

		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%v\n", city)

	} else {
		fmt.Println("no city found")
	}

}
