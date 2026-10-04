package main

import "fmt"

func main() {
	//fmt.Println("Congratulations!")

	//Declaración de variables
	var prueba1 int
	var prueba2 int = 20
	var prueba3 = 30
	prueba4 := 40

	fmt.Println(prueba1, prueba2, prueba3, prueba4)

	// Declaración de constantes
	const pruebaConst int = 1
	const pruebaConst2 = 2

	//fmt.Println(pruebaConst, pruebaConst2)

	//Printf
	nombre := "Platzi"
	cursos := 500
	fmt.Printf("%s tiene más de %d cursos\n", nombre, cursos)
}
