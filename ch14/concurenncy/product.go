package main

import "strconv"

type Product struct{
	Name, Category string
	Price float64
}

var ProductList = []*Product{
	{Name: "Kayak", Category: "Watersports", Price: 279},
	{Name: "Lifejacket", Category: "Watersports", Price: 49.95},
	{Name: "Soccer Ball", Category: "Soccer", Price: 19.50},
	{Name: "Corner Flags", Category: "Soccer", Price: 34.95},
	{Name: "Stadium", Category: "Soccer", Price: 79500},
	{Name: "Thinking Cap", Category: "Chess", Price: 16},
}

type ProductGroup []*Product

type ProductData = map[string]ProductGroup

var Products = make(ProductData)

func ToCurrency(val float64) string{
	return "$" + strconv.FormatFloat(val, 'f', 2, 64)
}

func init(){
	for _, p := range ProductList {
		if _, ok := Products[p.Category]; ok{
			Products[p.Category] = append(Products[p.Category], p)
		} else{
			Products[p.Category] = ProductGroup{p}
		}
	}
}