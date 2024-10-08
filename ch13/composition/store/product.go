package store

type Product struct{
	Name, Category string
	price float64
}

func NewProduct(name, category string, price float64) *Product{
	return &Product{
		Name:     name,
		Category: category,
		price:    price,
	}
}

func (p *Product) Price(taxRate float64) float64{
	return p.price + (p.price * taxRate)
}

