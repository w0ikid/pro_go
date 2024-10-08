package main

import (
	"fmt"
	"math/rand"
	"time"
)

type DispatchNotification struct{
	Customer string
	*Product
	Quantity int
}

var Customers = []string{"Alice", "Bob", "Charlie", "Dora"}

func DispatchOrders(channel chan DispatchNotification) {
	rand.Seed(time.Now().UTC().UnixNano())
	orderCount := rand.Intn(3) + 2
	fmt.Println("Order count:", orderCount)
	for i := 0; i < orderCount; i++ {
		channel <- DispatchNotification{
			Customer: Customers[rand.Intn(len(Customers)-1)],
			Product:  rand.Intn(10),
			Quantity: ProductList[rand.Intn(len(ProductList)-1)],
		}
	}
}