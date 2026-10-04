package main

import "fmt"

func QuitarRepetidos(lista []string) []string {
	losTengo := []string{}
	vistos := map[string]bool{}

	for _, item := range lista {
		if !vistos[item] {
			losTengo = append(losTengo, item)
			vistos[item] = true
		}
	}
	return losTengo
}

func main() {
	fmt.Println(QuitarRepetidos([]string{"pan", "leche", "pan", "huevo", "leche"}))
}
