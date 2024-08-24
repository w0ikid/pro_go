package main

import (
	"fmt"
	"time"
	// "time"
)

func CalcStoreTotal(data ProductData) {
	var storeTotal float64
	var channel chan float64 = make(chan float64, 2)
	for category, group := range data{
		go group.TotalPrice(category, channel)
	}
	fmt.Println("Starting to receive from channel")
	for i := 0; i < len(data); i++ {
		fmt.Println("-- channel read pending")
		fmt.Println("-- channel read pending", len(channel), "items in buffer, size", cap(channel))
		value := <- channel
		fmt.Println("-- channel read complete", value)
		storeTotal = storeTotal + value
		time.Sleep(time.Second)

	}
	fmt.Println("Total:", ToCurrency(storeTotal))
}
func (group ProductGroup) TotalPrice(category string, resultChannel chan float64){
	var total float64
	for _, p := range group {
		// fmt.Println(category, "product:", p.Name)
		total = total + p.Price
		time.Sleep(time.Microsecond * 100)
	}
	fmt.Println(category, "channel sending", ToCurrency(total))
	resultChannel <- total
	fmt.Println(category, "channel send complete")
}