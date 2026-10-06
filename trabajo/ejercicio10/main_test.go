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
	{Nombre: "Cloro", Categoria: "limpieza", Precio: 12},
	{Nombre: "Jabón", Categoria: "limpieza", Precio: 8},
	{Nombre: "Café", Categoria: "comestibles", Precio: 25},
}

func TestHayMasCaroQue(t *testing.T) {
	casos := []struct {
		limite   float64
		esperado bool
	}{
		{28, true},
		{30, false},
		{4, true},
	}
	for _, c := range casos {
		obtenido := HayMasCaroQue(inventarioPrueba, c.limite)
		if obtenido != c.esperado {
			t.Errorf("HayMasCaroQue(%v) = %v, se esperaba %v", c.limite, obtenido, c.esperado)
		}
	}
}

func TestCategoriasUnicas(t *testing.T) {
	obtenido := CategoriasUnicas(inventarioPrueba)
	esperado := []string{"comestibles", "empaque", "limpieza"}
	if !reflect.DeepEqual(obtenido, esperado) {
		t.Errorf("CategoriasUnicas() = %v, se esperaba %v", obtenido, esperado)
	}
}

func TestCategoriasUnicasVacio(t *testing.T) {
	obtenido := CategoriasUnicas([]Producto{})
	if obtenido == nil || len(obtenido) != 0 {
		t.Errorf("CategoriasUnicas(vacío) = %#v, se esperaba una lista vacía (no nil)", obtenido)
	}
}

func TestPalabrasUnicas(t *testing.T) {
	obtenido := PalabrasUnicas([]string{"pan", "leche", "pan", "huevo", "café", "leche"})
	esperado := []string{"huevo", "café"}
	if !reflect.DeepEqual(obtenido, esperado) {
		t.Errorf("PalabrasUnicas() = %v, se esperaba %v", obtenido, esperado)
	}
}

func TestPalabrasUnicasNinguna(t *testing.T) {
	obtenido := PalabrasUnicas([]string{"pan", "leche", "pan", "leche"})
	if obtenido == nil || len(obtenido) != 0 {
		t.Errorf("PalabrasUnicas() = %#v, se esperaba una lista vacía (no nil)", obtenido)
	}
}

func TestDeCategorias(t *testing.T) {
	obtenido := DeCategorias(inventarioPrueba, []string{"limpieza", "empaque"})
	esperado := []Producto{
		{Nombre: "Tapas", Categoria: "empaque", Precio: 20},
		{Nombre: "Cajas", Categoria: "empaque", Precio: 5},
		{Nombre: "Cloro", Categoria: "limpieza", Precio: 12},
		{Nombre: "Jabón", Categoria: "limpieza", Precio: 8},
	}
	if !reflect.DeepEqual(obtenido, esperado) {
		t.Errorf("DeCategorias() = %v, se esperaba %v", obtenido, esperado)
	}
}

func TestDeCategoriasNinguna(t *testing.T) {
	obtenido := DeCategorias(inventarioPrueba, []string{"ropa"})
	if obtenido == nil || len(obtenido) != 0 {
		t.Errorf("DeCategorias(ropa) = %#v, se esperaba una lista vacía (no nil)", obtenido)
	}
}
