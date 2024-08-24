package main

import (
	"fmt"
	// "reflect"
	// "strings"
	"regexp"
)

func main(){
	// product := "Kayak"

	// fmt.Println("Product:", product)

	// fmt.Println(strings.Contains(product, "ya"))
	// fmt.Println(strings.ContainsAny(product, "abc"))
	// fmt.Println(strings.ContainsRune(product, 'K'))
	// fmt.Println(strings.EqualFold(product, "KAYAK"))
	// fmt.Println(strings.HasPrefix(product, "Ka"))
	// fmt.Println(strings.HasSuffix(product, "yak"))
	
	// description := "A boat for one person"
	// // text := "Some text"
	// match, err := regexp.MatchString("[A-z]oat", description)
	// // Любой символ от A до z за которым следует символ oat
	// if err == nil{
	// 	fmt.Println("Match:", match)
	// } else {
	// 	fmt.Println("Error:",  err)
	// }

	pattern, compileErr := regexp.Compile("[A-z]oat")

	description := "A boat for one person"
	question := "Is that a goat?"
	preference := "I like oats"

	if compileErr == nil {
		fmt.Println("Description:", pattern.MatchString(description))
		fmt.Println("Question:", pattern.MatchString(question))
		fmt.Println("Preference:", pattern.MatchString(preference))
	} else {
		fmt.Println("Error:", compileErr)
	}
}