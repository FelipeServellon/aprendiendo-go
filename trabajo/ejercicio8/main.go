package main

import "fmt"

/* func QuitarRepetidos(lista []string) []string {
	visto := map[string]bool{}
	resultado := []string{}

	for _, s := range lista {
		if visto[s] {
			continue
		}
		visto[s] = true
		resultado = append(resultado, s)
	}
	return resultado
} */

func QuitarRepetidos(lista []string) []string {
	listaLimpia := map[string]bool{}

	for i:=0; i<len(lista);i++{
		if lista[i]
			continue
		}else{
			listaLimpia = append(listaLimpia,lista[i])
		}
	}

	return nil
}

func main() {
	fmt.Println(QuitarRepetidos([]string{"pan", "leche", "pan", "huevo", "leche"}))
}
