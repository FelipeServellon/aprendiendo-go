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

func TestNombresBaratos(t *testing.T) {
	obtenido := NombresBaratos(inventarioPrueba, 15)
	esperado := []string{"Leche", "Cajas"}
	if !reflect.DeepEqual(obtenido, esperado) {
		t.Errorf("NombresBaratos(15) = %v, se esperaba %v", obtenido, esperado)
	}
}

func TestNombresBaratosNinguno(t *testing.T) {
	obtenido := NombresBaratos(inventarioPrueba, 1)
	if obtenido == nil || len(obtenido) != 0 {
		t.Errorf("NombresBaratos(1) = %#v, se esperaba una lista vacía (no nil)", obtenido)
	}
}

func TestPromedioDeCategoria(t *testing.T) {
	casos := []struct {
		categoria string
		esperado  float64
	}{
		{"empaque", 12.5},
		{"comestibles", 20},
		{"limpieza", 0},
	}
	for _, c := range casos {
		obtenido := PromedioDeCategoria(inventarioPrueba, c.categoria)
		if obtenido != c.esperado {
			t.Errorf("PromedioDeCategoria(%q) = %v, se esperaba %v", c.categoria, obtenido, c.esperado)
		}
	}
}

func TestRepetidos(t *testing.T) {
	obtenido := Repetidos([]string{"pan", "leche", "pan", "huevo", "leche", "pan"})
	esperado := []string{"pan", "leche"}
	if !reflect.DeepEqual(obtenido, esperado) {
		t.Errorf("Repetidos() = %v, se esperaba %v", obtenido, esperado)
	}
}

func TestRepetidosNinguno(t *testing.T) {
	obtenido := Repetidos([]string{"pan", "leche", "huevo"})
	if obtenido == nil || len(obtenido) != 0 {
		t.Errorf("Repetidos() = %#v, se esperaba una lista vacía (no nil)", obtenido)
	}
}
