package main

import (
	"composition/store"
	"fmt"
)

func main(){
	fmt.Println("TempleOs!")

	phone_samsung := store.NewProduct("SamsungA71", "Phone", 150)

	phone_xiaomi_special := &store.Product{
		Name:     "XiaomiSPECIALMODEL",
		Category: "Phone",
	}

	for _, p := range []*store.Product{phone_samsung, phone_xiaomi_special}{
		fmt.Println("Name: ", p.Name, "Category:", p.Category, "Price: ", p.Price(0.2))
	}

}