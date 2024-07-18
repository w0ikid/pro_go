// This is main package
package main

import (
	"fmt"
)

// PrintHello prints "Hello, world!" using the fmt.Println function
func PrintHello(){
	fmt.Println("Hello, world!")
}

// PrintNumber prints number using the fmt.Println function
func PrintNumber(number int) {
	fmt.Println(number)
}

func main(){
	PrintHello()
	for i := 0; i < 5; i++ {
		PrintNumber(i)
	}
}

