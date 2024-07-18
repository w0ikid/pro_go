// dlv version

package main

import (
	"fmt"
)

func main(){
	fmt.Println("Hello world")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		for j := 0; i <= 3; j++ {
			fmt.Println(j)
		}
	}

}