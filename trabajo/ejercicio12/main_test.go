package main

import "testing"

func TestReponer(t *testing.T) {
	casos := []struct {
		inicial, extra int
		esperado       int
	}{
		{5, 3, 8},
		{0, 4, 4},
		{7, 0, 7},
	}
	for _, c := range casos {
		n := c.inicial
		Reponer(&n, c.extra)
		if n != c.esperado {
			t.Errorf("con %d, Reponer(&n, %d) dejó n en %d; se esperaba %d", c.inicial, c.extra, n, c.esperado)
		}
	}
}

func TestSubirPrecio(t *testing.T) {
	casos := []struct {
		precio, aumento float64
		esperado        float64
	}{
		{10, 2.5, 12.5},
		{25, 5, 30},
		{0, 1, 1},
	}
	for _, c := range casos {
		p := Producto{"Leche", c.precio}
		SubirPrecio(&p, c.aumento)
		if p.Precio != c.esperado {
			t.Errorf("con precio %v, SubirPrecio(&p, %v) lo dejó en %v; se esperaba %v", c.precio, c.aumento, p.Precio, c.esperado)
		}
	}
}

func TestNewBodega(t *testing.T) {
	b := NewBodega("Central")
	if b.Nombre != "Central" {
		t.Errorf("NewBodega(%q).Nombre = %q; se esperaba %q", "Central", b.Nombre, "Central")
	}
	if b.Stock == nil {
		t.Fatalf("NewBodega(%q).Stock es nil: escribir en él detendría el programa", "Central")
	}
	b.Stock["Leche"] = 5
	if b.Stock["Leche"] != 5 {
		t.Errorf("tras escribir 5 en Stock[\"Leche\"] se leyó %d", b.Stock["Leche"])
	}
}

func TestNombrePadre(t *testing.T) {
	lacteos := Categoria{Nombre: "Lácteos"}
	quesos := Categoria{Nombre: "Quesos", Padre: &lacteos}
	frescos := Categoria{Nombre: "Frescos", Padre: &quesos}

	casos := []struct {
		c        Categoria
		esperado string
	}{
		{quesos, "Lácteos"},
		{frescos, "Quesos"},
		{lacteos, ""},
	}
	for _, c := range casos {
		obtenido := NombrePadre(c.c)
		if obtenido != c.esperado {
			t.Errorf("NombrePadre(%s) = %q; se esperaba %q", c.c.Nombre, obtenido, c.esperado)
		}
	}
}
