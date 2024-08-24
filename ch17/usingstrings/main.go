package main

import "fmt"

func getProductName(index int) (name string, err error) {
	if len(Products) > index {
		name = fmt.Sprintf("Name of product: %v", Products[index].Name)
	} else {
		err = fmt.Errorf("Error for index: %v", index)
	}
	return
}

func main(){
	// fmt.Println("Product:", Kayak.Name, "Price:", Kayak.Price)
	// fmt.Print("Product:", Kayak.Name, "Price:", Kayak.Price, "\n")
	// fmt.Printf("Product: %v Price: $%4.2f", Kayak.Name, Kayak.Price)
	name, _ := getProductName(1)
	fmt.Println(name)

	_, err := getProductName(10)
	fmt.Println(err.Error())
}

