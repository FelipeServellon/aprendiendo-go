package main

import "fmt"

func soloFrutas(listaDeSuper []string) []string {
	frutas := make([]string, 0, len(listaDeSuper))
	for i := 0; i < len(listaDeSuper); i++ {
		if listaDeSuper[i] == "Manzanas" || listaDeSuper[i] == "Bananos" || listaDeSuper[i] == "Naranjas" || listaDeSuper[i] == "Fresas" || listaDeSuper[i] == "Uvas" || listaDeSuper[i] == "Piña" || listaDeSuper[i] == "Sandía" || listaDeSuper[i] == "Mango" {
			frutas = append(frutas, listaDeSuper[i])
		}
	}
	return frutas
}

func soloPares(numeros []int) []int {
	numero := make([]int, 0, len(numeros))
	for i := 0; i < len(numeros); i++ {
		if numeros[i]%2 == 0 {
			numero = append(numero, numeros[i])
		}
	}
	return numero
}

func primerosPares(n int) []int {
	pares := make([]int, 0, n)
	for i := 1; i <= n; i++ {
		pares = append(pares, i*2)
	}
	return pares
}

func main() {
	resultado := soloFrutas([]string{"Manzanas", "Bananos", "Naranjas", "Fresas", "Uvas", "Piña", "Sandía", "Mango", "Pan", "Arroz", "Frijol", "Aceite", "Azúcar", "Sal", "Café", "Tomate", "Pollo", "Leche", "Huevos", "Queso", "Mantequilla", "Yogur", "Jamón", "Pasta", "Cebolla", "Papas", "Zanahorias", "Lechuga", "Ajo", "Cilantro", "Papel higiénico", "Detergente", "Jabón", "Shampoo", "Servilletas", "Bolsas de basura"})
	fmt.Println(resultado)

	fmt.Println(soloPares([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))

	fmt.Println(primerosPares(20))
}
