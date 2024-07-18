package main

import (
	"build/store"
	"fmt"
	currencyFmt "build/fmt"
)

func main(){
	fmt.Println("Something")
	fmt.Println("---------")
	
	
	// var product *store.Product = &store.Product{}


	// product := store.Product {
	// 	Name: "Danial",
	// 	Category: "Judo",
	// }

	// fmt.Println("Name: ", product.Name)
	// fmt.Println("Category" ,product.Category)	


	product := store.NewProduct("Samsung A71", "Phone", 150)

	fmt.Println(product.Price())
	
	product.SetPrice(300)
	
	fmt.Println(product.Price())


	fmt.Println(currencyFmt.ToCurrency(product.Price()))
}