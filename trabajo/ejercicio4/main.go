package main

import "fmt"

/*
	EJERCICIO #1:

Lo interpreté como una autorización de credenciales o como si la target era autizada o no.
*/
func Contains(names []string, target string) bool {
	for _, name := range names {
		if target == name {
			return true
		}
	}
	return false
}

/*
EJERCICIO #2:

Lo trabajé como una validación del target con el slice de nums, para que la respuesta fuera cuántas veces se repite.
*/
func CountOccurrences(nums []int, target int) int {
	contador := 0
	for i := 0; i < len(nums); i++ {
		if target == nums[i] {
			contador = contador + 1
		}
	}
	return contador
}

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
	/* Entrada ejercicio #1 */
	resultado1 := Contains([]string{"Jeanca", "David", "Tatsuya", "Felipe", "Ema"}, "Ema")
	fmt.Println(resultado1)

	/* Entrada ejercicio #2 */
	resultado2 := CountOccurrences([]int{1234, 5678, 91011, 1213, 1415}, 1213)
	fmt.Println(resultado2)

	/* Entrada ejercicio #3 */
	resultado3 := FindAll([]product{
		{name: "Leche", category: "comestibles", price: 10},
		{name: "Tapas", category: "empaque", price: 20},
		{name: "Azúcar", category: "comestibles", price: 30}}, "comestibles")
	fmt.Println(resultado3)
}
