package main

import "fmt"

type Producto struct {
	Nombre    string
	Categoria string
	Precio    float64
}

// EJERCICIO 1 — map acumulador con float64
// SumarPorCategoria devuelve cuánto dinero suma cada categoría.
// Ej: {"comestibles": 40, "empaque": 25}
func SumarPorCategoria(productos []Producto) map[string]float64 {
	total := make(map[string]float64)

	for i := 0; i < len(productos); i++ {
		total[productos[i].Categoria] = total[productos[i].Categoria] + productos[i].Precio
	}
	return total
}

// EJERCICIO 2 — filtro sobre structs
// FiltrarPorCategoria devuelve solo los productos de la categoría pedida.
// Si no hay ninguno, devuelve una lista vacía (no nil).
func FiltrarPorCategoria(productos []Producto, categoria string) []Producto {
	solicitados := []Producto{}

	for i := 0; i < len(productos); i++ {
		if productos[i].Categoria == categoria {
			solicitados = append(solicitados, productos[i])
		}
	}
	return solicitados
}

// EJERCICIO 3 — map como registro de vistos
// QuitarRepetidos devuelve la lista sin elementos duplicados,
// conservando el orden de la primera aparición.
// Ej: ["pan","leche","pan"] -> ["pan","leche"]
func QuitarRepetidos(lista []string) []string {
	listaLimpia := map[string]bool{}
	resultado := []string{}

	for i := 0; i < len(lista); i++ {
		if listaLimpia[lista[i]] == false {
			listaLimpia[lista[i]] = true
			resultado = append(resultado, lista[i])
		}
	}

	return resultado
}

func main() {
	inventario := []Producto{
		{Nombre: "Leche", Categoria: "comestibles", Precio: 10},
		{Nombre: "Tapas", Categoria: "empaque", Precio: 20},
		{Nombre: "Azúcar", Categoria: "comestibles", Precio: 30},
		{Nombre: "Cajas", Categoria: "empaque", Precio: 5},
	}

	fmt.Println(SumarPorCategoria(inventario))
	fmt.Println(FiltrarPorCategoria(inventario, "empaque"))
	fmt.Println(QuitarRepetidos([]string{"pan", "leche", "pan", "huevo", "leche"}))
}
