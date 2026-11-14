package main

import "fmt"

type Producto struct {
	Nombre string
	Precio float64
}

type Bodega struct {
	Nombre string
	Stock  map[string]int
}

type Categoria struct {
	Nombre string
	Padre  *Categoria
}

/*
EJERCICIO 1 — & y *
Reponer suma «extra» a la «cantidad» original, no a una copia.
Recibe la dirección de la «cantidad» (*int), no el número.
Ej: n := 5; Reponer(&n, 3) -> n queda en 8
Ej: n := 0; Reponer(&n, 4) -> n queda en 4
*/
func Reponer(cantidad *int, extra int) {
	*cantidad = *cantidad + extra
}

/*
EJERCICIO 2 — puntero a un struct
SubirPrecio suma «aumento» al «Precio» del producto original.
Ej: leche := Producto{"Leche", 10}; SubirPrecio(&leche, 2.5) -> leche.Precio queda en 12.5
*/
func SubirPrecio(p *Producto, aumento float64) {

	p.Precio = p.Precio + aumento

}

/*
EJERCICIO 3 — constructor
NewBodega devuelve una Bodega lista para usar:
con el «nombre» recibido y con su map «Stock» ya creado,
para que se pueda escribir en él sin que el programa se detenga.
Ej: b := NewBodega("Central"); b.Stock["Leche"] = 5 -> funciona
*/
func NewBodega(nombre string) Bodega {
	return Bodega{}
}

/*
EJERCICIO 4 — nil
NombrePadre devuelve el nombre de la categoría padre de «c».
Si «c» no tiene padre («Padre» es nil), devuelve "".
Ojo: leer un campo a través de un puntero nil detiene el programa.
Ej, con estas dos categorías:
lacteos := Categoria{Nombre: "Lácteos"}
quesos := Categoria{Nombre: "Quesos", Padre: &lacteos}
NombrePadre(quesos)  -> "Lácteos"
NombrePadre(lacteos) -> ""
*/
func NombrePadre(c Categoria) string {
	return ""
}

func main() {
	n := 5
	Reponer(&n, 3)
	fmt.Println("Cantidad:", n)

	leche := Producto{"Leche", 10}
	SubirPrecio(&leche, 2.5)
	fmt.Println("Precio de la leche:", leche.Precio)

	central := NewBodega("Central")
	fmt.Println("Bodega:", central.Nombre, central.Stock)

	lacteos := Categoria{Nombre: "Lácteos"}
	quesos := Categoria{Nombre: "Quesos", Padre: &lacteos}
	fmt.Println("Padre de Quesos:", NombrePadre(quesos))
	fmt.Println("Padre de Lácteos:", NombrePadre(lacteos))
}
