package main

import (
	"config"
	"fmt"
	"models"
)

func main() {

	db, err := config.GetMySQLDB()
	if err != nil {
		fmt.Println(err)
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		fmt.Println("Product List")
		products, err2 := productModel.FindAll()
		if err2 != nil {
			fmt.Println(err2)
		} else {
			fmt.Println("Products: ", len(products))
			for _, product := range products {
				fmt.Println("Id:", product.Id)
				fmt.Println("Name:", product.Name)
				fmt.Println("Price:", product.Price)
				fmt.Println("Quantity:", product.Quantity)
				fmt.Println("Status:", product.Status)
				fmt.Println("----------------------------")
			}
		}

	}

}