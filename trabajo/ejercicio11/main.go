package main

import (
	"errors"
	"fmt"
)

// Los errores que pueden devolver las funciones de este ejercicio.
var (
	ErrDivisionPorCero = errors.New("división por cero")
	ErrNoEncontrado    = errors.New("producto no encontrado")
	ErrSinStock        = errors.New("stock insuficiente")
)

/*
EJERCICIO 1 — devolver (valor, error)
Dividir devuelve a dividido entre b.
Si b es 0, no se puede dividir: devuelve 0 y ErrDivisionPorCero.
Si todo va bien, devuelve el resultado y nil (nil = "no hubo error").
Ej: Dividir(10, 2) -> 5, nil
Ej: Dividir(10, 0) -> 0, ErrDivisionPorCero
*/
func Dividir(a, b float64) (float64, error) {
	return 0, nil
}

/*
EJERCICIO 2 — el "ok" de un map
BuscarPrecio devuelve el precio del producto pedido.
Si el producto no está en el map, devuelve 0 y ErrNoEncontrado.
Ojo: "Muestra" SÍ existe y cuesta 0. Debe devolver 0 y nil, no un error.
Ej: "Leche"   -> 10, nil
Ej: "Muestra" -> 0, nil
Ej: "Pizza"   -> 0, ErrNoEncontrado
*/
func BuscarPrecio(precios map[string]float64, nombre string) (float64, error) {
	return 0, nil
}

/*
EJERCICIO 3 — devolver solo un error
Vender descuenta la cantidad del stock del producto.
Si el producto no está en el map, devuelve ErrNoEncontrado.
Si se piden más unidades de las que hay, devuelve ErrSinStock
y el stock NO cambia.
Si todo va bien, resta la cantidad y devuelve nil.
Ej: stock Leche 5, Vender 3 -> nil (queda 2)
Ej: stock Leche 5, Vender 9 -> ErrSinStock (sigue en 5)
*/
func Vender(stock map[string]int, nombre string, cantidad int) error {
	return nil
}

/*
EJERCICIO 4 — revisar el error de otra función
TotalPedido suma los precios de todos los productos del pedido,
usando BuscarPrecio para cada uno.
Si algún producto no existe, devuelve 0 y el error que dio BuscarPrecio.
Ej: ["Leche", "Café"]  -> 35, nil
Ej: ["Leche", "Pizza"] -> 0, ErrNoEncontrado
*/
func TotalPedido(precios map[string]float64, pedido []string) (float64, error) {
	return 0, nil
}

func main() {
	precios := map[string]float64{"Leche": 10, "Café": 25, "Muestra": 0}
	stock := map[string]int{"Leche": 5, "Café": 2}

	fmt.Println(Dividir(10, 2))
	fmt.Println(Dividir(10, 0))
	fmt.Println(BuscarPrecio(precios, "Muestra"))
	fmt.Println(BuscarPrecio(precios, "Pizza"))
	fmt.Println(Vender(stock, "Leche", 3), stock)
	fmt.Println(Vender(stock, "Café", 9), stock)
	fmt.Println(TotalPedido(precios, []string{"Leche", "Café"}))
	fmt.Println(TotalPedido(precios, []string{"Leche", "Pizza"}))
}
