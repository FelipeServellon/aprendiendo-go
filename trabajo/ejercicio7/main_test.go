package main

import (
	"reflect"
	"testing"
)

var inventarioPrueba = []Producto{
	{Nombre: "Leche", Categoria: "comestibles", Precio: 10},
	{Nombre: "Tapas", Categoria: "empaque", Precio: 20},
	{Nombre: "Azúcar", Categoria: "comestibles", Precio: 30},
	{Nombre: "Cajas", Categoria: "empaque", Precio: 5},
}

func TestSumarPorCategoria(t *testing.T) {
	obtenido := SumarPorCategoria(inventarioPrueba)
	esperado := map[string]float64{"comestibles": 40, "empaque": 25}
	if !reflect.DeepEqual(obtenido, esperado) {
		t.Errorf("SumarPorCategoria() = %v, se esperaba %v", obtenido, esperado)
	}
}

func TestSumarPorCategoriaListaVacia(t *testing.T) {
	obtenido := SumarPorCategoria([]Producto{})
	if len(obtenido) != 0 {
		t.Errorf("SumarPorCategoria([]) = %v, se esperaba un map vacío", obtenido)
	}
}

func TestFiltrarPorCategoria(t *testing.T) {
	obtenido := FiltrarPorCategoria(inventarioPrueba, "empaque")
	esperado := []Producto{
		{Nombre: "Tapas", Categoria: "empaque", Precio: 20},
		{Nombre: "Cajas", Categoria: "empaque", Precio: 5},
	}
	if !reflect.DeepEqual(obtenido, esperado) {
		t.Errorf("FiltrarPorCategoria(\"empaque\") = %v, se esperaba %v", obtenido, esperado)
	}
}

func TestFiltrarPorCategoriaSinCoincidencias(t *testing.T) {
	obtenido := FiltrarPorCategoria(inventarioPrueba, "limpieza")
	if len(obtenido) != 0 {
		t.Errorf("FiltrarPorCategoria(\"limpieza\") = %v, se esperaba una lista vacía", obtenido)
	}
}

func TestQuitarRepetidos(t *testing.T) {
	obtenido := QuitarRepetidos([]string{"pan", "leche", "pan", "huevo", "leche"})
	esperado := []string{"pan", "leche", "huevo"}
	if !reflect.DeepEqual(obtenido, esperado) {
		t.Errorf("QuitarRepetidos() = %v, se esperaba %v", obtenido, esperado)
	}
}

func TestQuitarRepetidosSinRepetidos(t *testing.T) {
	obtenido := QuitarRepetidos([]string{"pan", "leche"})
	esperado := []string{"pan", "leche"}
	if !reflect.DeepEqual(obtenido, esperado) {
		t.Errorf("QuitarRepetidos() = %v, se esperaba %v", obtenido, esperado)
	}
}
