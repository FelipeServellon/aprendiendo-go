package main

import "fmt"

func imprimirMensaje(mensaje string) {
	fmt.Println(mensaje)
}
func imprimirTripleValor(a, b, c int) (int, int, int) {
	return a, b, c
}
func retornarValor(b int) int {
	return b * 10
}
func retornarDobleValor(a int) (b, c int) {
	return a, a * 7
}
func main() {
	imprimirMensaje("Esto salió bien")
	x, y, z := imprimirTripleValor(1, 2, 3)
	fmt.Println(x, y, z)
	valor := retornarValor(1)
	fmt.Println(valor)
	valor1, valor2 := retornarDobleValor(3)
	fmt.Println("Valores 1 y 2:", valor1, valor2)
}
