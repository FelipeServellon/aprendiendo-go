package main

import "fmt"

type product struct {
	name     string
	category string
	price    float64
}

func FindAll(item []product, category string) []product {
	result := make([]product, 0, len(item))
	for i := 0; i != len(item); i++ {
		if item[i].category == category {
			result = append(result, item[i])
		}
	}
	return result

}

func main() {
	resultado3 := FindAll([]product{
		{name: "Leche", category: "comestibles", price: 10},
		{name: "Tapas", category: "empaque", price: 20},
		{name: "Azúcar", category: "comestibles", price: 30}}, "comestibles")
	fmt.Println(resultado3)
}
