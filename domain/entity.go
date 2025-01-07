package domain

type Product struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Price float64 `json:"price"`
}

/*
OUTPUT EXAMPLE OF JSON
{
    "id": 1,
    "name": "Laptop",
    "price": 1500.0
}
*/