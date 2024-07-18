package store

type Charge struct {
	*Product
	Capacity int
	FastCharge bool
}

func NewCharge(name string, price float64, capacity int, fastcharge bool) *Charge {
	return &Charge{
		Product:    NewProduct(name, "Phone", 20),
		Capacity:   capacity,
		FastCharge: fastcharge,
}
