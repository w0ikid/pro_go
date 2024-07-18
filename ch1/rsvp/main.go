package main

import (
	"fmt"
	"html/template"
)

type Rspv struct{
	Name, Email, Phone string
	WillAttend bool
}

var responses = make([]*Rspv, 0, 10)
var templates = make(map[string]*template.Template, 3)

func loadTemplates(){
	// TODO - load templates here
	templateNames := [5]string{"welcome", "form", "thanks", "sorry", "list"}
	for index, name := range templateNames {
		t, err := template.ParseFiles("layout.html", name + ".html")
		if err == nil {
			templates[name] = t
			fmt.Println("Loaded template", index, name)
		} else {
			panic(err)
		}
	}
}

func main(){
	fmt.Println("work")

	loadTemplates()
}