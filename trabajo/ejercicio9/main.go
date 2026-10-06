package main

import "fmt"

type Producto struct {
	Nombre    string
	Categoria string
	Precio    float64
}

// EJERCICIO 1 — filtro que devuelve solo un campo
// NombresBaratos devuelve los NOMBRES de los productos cuyo precio
// es menor que el límite, en el mismo orden en que aparecen.
// Si no hay ninguno, devuelve una lista vacía (no nil).
// Ej: límite 15 -> ["Leche", "Cajas"]
func NombresBaratos(productos []Producto, limite float64) []string {
	masBaratos := []string{}

	for _, v := range productos {
		if v.Precio < limite {
			masBaratos = append(masBaratos, v.Nombre)
		}
	}
	return masBaratos
}

// EJERCICIO 2 — dos acumuladores a la vez
// PromedioDeCategoria devuelve el precio promedio de los productos
// de la categoría pedida (suma de precios / cantidad de productos).
// Si la categoría no tiene productos, devuelve 0.
// Ej: "empaque" -> (20 + 5) / 2 = 12.5f
func PromedioDeCategoria(productos []Producto, categoria string) float64 {
	sumaDePrecios := 0.0
	cantidadDeProductos := 0.0

	for i := 0; i < len(productos); i++ {
		if categoria == productos[i].Categoria {
			cantidadDeProductos = cantidadDeProductos + 1
			sumaDePrecios = sumaDePrecios + productos[i].Precio
		}
	}
	if cantidadDeProductos == 0 {
		return 0
	}
	return sumaDePrecios / cantidadDeProductos
}

// EJERCICIO 3 — map como contador
// Repetidos devuelve las palabras que aparecen MÁS de una vez.
// Cada una sale una sola vez, en el orden en que se repitió por primera vez.
// Si ninguna se repite, devuelve una lista vacía (no nil).
// Ej: ["pan","leche","pan","huevo","leche","pan"] -> ["pan","leche"]
func Repetidos(lista []string) []string {
	vistos := map[string]int{}
	repetidos := []string{}

	for _, item := range lista {
		vistos[item]++
		if vistos[item] == 2 {
			repetidos = append(repetidos, item)
		}
	}

	return repetidos
}

func main() {
	inventario := []Producto{
		{Nombre: "Leche", Categoria: "comestibles", Precio: 10},
		{Nombre: "Tapas", Categoria: "empaque", Precio: 20},
		{Nombre: "Azúcar", Categoria: "comestibles", Precio: 30},
		{Nombre: "Cajas", Categoria: "empaque", Precio: 5},
	}

	fmt.Println(NombresBaratos(inventario, 15))
	fmt.Println(PromedioDeCategoria(inventario, "empaque"))
	fmt.Println(Repetidos([]string{"pan", "leche", "pan", "huevo", "leche", "pan"}))
}
