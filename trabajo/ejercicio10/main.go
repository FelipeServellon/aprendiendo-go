package main

import "fmt"

type Producto struct {
	Nombre    string
	Categoria string
	Precio    float64
}

// EJERCICIO 1 — salida anticipada
// HayMasCaroQue devuelve true si AL MENOS UN producto cuesta
// más que el límite (estrictamente mayor). Si ninguno, false.
// Ej: límite 28 -> true (Azúcar cuesta 30)
// Ej: límite 30 -> false
func HayMasCaroQue(productos []Producto, limite float64) bool {
	for _, suPrecio := range productos {
		if suPrecio.Precio > limite {
			return true
		}
	}
	return false
}

// EJERCICIO 2 — map de vistos + lista de salida
// CategoriasUnicas devuelve cada categoría una sola vez,
// en el orden en que aparece por primera vez.
// Si no hay productos, devuelve una lista vacía (no nil).
// Ej: -> ["comestibles", "empaque", "limpieza"]
func CategoriasUnicas(productos []Producto) []string {
	return nil
}

// EJERCICIO 3 — contar primero, filtrar después
// PalabrasUnicas devuelve las palabras que aparecen EXACTAMENTE una vez,
// en el orden en que aparecen en la lista.
// Si todas se repiten, devuelve una lista vacía (no nil).
// Ej: ["pan","leche","pan","huevo","café","leche"] -> ["huevo","café"]
func PalabrasUnicas(lista []string) []string {
	return nil
}

// EJERCICIO 4 — map como catálogo
// DeCategorias devuelve los productos cuya categoría está en la lista
// de categorías pedidas, en el mismo orden en que aparecen los productos.
// Si ninguno coincide, devuelve una lista vacía (no nil).
// Ej: ["limpieza","empaque"] -> Tapas, Cajas, Cloro, Jabón
func DeCategorias(productos []Producto, categorias []string) []Producto {
	return nil
}

func main() {
	inventario := []Producto{
		{Nombre: "Leche", Categoria: "comestibles", Precio: 10},
		{Nombre: "Tapas", Categoria: "empaque", Precio: 20},
		{Nombre: "Azúcar", Categoria: "comestibles", Precio: 30},
		{Nombre: "Cajas", Categoria: "empaque", Precio: 5},
		{Nombre: "Cloro", Categoria: "limpieza", Precio: 12},
		{Nombre: "Jabón", Categoria: "limpieza", Precio: 8},
		{Nombre: "Café", Categoria: "comestibles", Precio: 25},
	}

	fmt.Println(HayMasCaroQue(inventario, 28))
	fmt.Println(CategoriasUnicas(inventario))
	fmt.Println(PalabrasUnicas([]string{"pan", "leche", "pan", "huevo", "café", "leche"}))
	fmt.Println(DeCategorias(inventario, []string{"limpieza", "empaque"}))
}
