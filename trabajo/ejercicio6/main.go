package main

import "fmt"

type Producto struct {
	Nombre    string
	Categoria string
	Precio    float64
}

// EJERCICIO 1 — structs
// Total suma el precio de todos los productos.
func Total(productos []Producto) float64 {
	GranTotal := 0.0
	for i := 0; i != len(productos); i++ {
		GranTotal = GranTotal + productos[i].Precio
	}
	return GranTotal
}

// EJERCICIO 2 — maps (crear y escribir)
// ContarPorCategoria devuelve cuántos productos hay en cada categoría.
// Ej: {"comestibles": 2, "empaque": 1}
func ContarPorCategoria(productos []Producto) map[string]int {
	existencias := make(map[string]int)
	for i := 0; i < len(productos); i++ {
		existencias[productos[i].Categoria] = existencias[productos[i].Categoria] + 1

	}
	return existencias
}

// EJERCICIO 3 — maps (leer)
// SoloFrutas devuelve únicamente los elementos de la lista que son frutas.
// Mismo resultado que soloFrutas de ejercicio3, pero sin la cadena de ||.
func SoloFrutas(lista []string) []string {
	frutas := map[string]bool{
		"Manzanas": true, "Uvas": true,
	}

	resultado := []string{}
	for i := 0; i < len(lista); i++ {
		if frutas[lista[i]] {
			resultado = append(resultado, lista[i])
		}
	}
	return resultado
}

// EJERCICIO 4 — structs + maps
// MasCaroPorCategoria devuelve, por cada categoría, su producto más caro.
func MasCaroPorCategoria(productos []Producto) map[string]Producto {

	return nil
}

func main() {
	inventario := []Producto{
		{Nombre: "Leche", Categoria: "comestibles", Precio: 10},
		{Nombre: "Tapas", Categoria: "empaque", Precio: 20},
		{Nombre: "Azúcar", Categoria: "comestibles", Precio: 30},
	}

	fmt.Println(Total(inventario))
	fmt.Println(ContarPorCategoria(inventario))
	fmt.Println(SoloFrutas([]string{"Manzanas", "Pan", "Uvas", "Jabón"}))
	fmt.Println(MasCaroPorCategoria(inventario))
}
