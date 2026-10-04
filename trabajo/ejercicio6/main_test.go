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

func TestTotal(t *testing.T) {
	obtenido := Total(inventarioPrueba)
	esperado := 65.0
	if obtenido != esperado {
		t.Errorf("Total() = %v, se esperaba %v", obtenido, esperado)
	}
}

func TestTotalListaVacia(t *testing.T) {
	if obtenido := Total([]Producto{}); obtenido != 0 {
		t.Errorf("Total([]) = %v, se esperaba 0", obtenido)
	}
}

func TestContarPorCategoria(t *testing.T) {
	obtenido := ContarPorCategoria(inventarioPrueba)
	esperado := map[string]int{"comestibles": 2, "empaque": 2}
	if !reflect.DeepEqual(obtenido, esperado) {
		t.Errorf("ContarPorCategoria() = %v, se esperaba %v", obtenido, esperado)
	}
}

func TestSoloFrutas(t *testing.T) {
	obtenido := SoloFrutas([]string{"Manzanas", "Pan", "Uvas", "Jabón", "Mango"})
	esperado := []string{"Manzanas", "Uvas", "Mango"}
	if !reflect.DeepEqual(obtenido, esperado) {
		t.Errorf("SoloFrutas() = %v, se esperaba %v", obtenido, esperado)
	}
}

func TestSoloFrutasSinFrutas(t *testing.T) {
	obtenido := SoloFrutas([]string{"Pan", "Jabón"})
	if len(obtenido) != 0 {
		t.Errorf("SoloFrutas() = %v, se esperaba una lista vacía", obtenido)
	}
}

func TestMasCaroPorCategoria(t *testing.T) {
	obtenido := MasCaroPorCategoria(inventarioPrueba)
	esperado := map[string]Producto{
		"comestibles": {Nombre: "Azúcar", Categoria: "comestibles", Precio: 30},
		"empaque":     {Nombre: "Tapas", Categoria: "empaque", Precio: 20},
	}
	if !reflect.DeepEqual(obtenido, esperado) {
		t.Errorf("MasCaroPorCategoria() = %v, se esperaba %v", obtenido, esperado)
	}
}
