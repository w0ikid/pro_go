package main

import (
	"fmt"
)

func main(){
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Println(j)
		}
	}
}