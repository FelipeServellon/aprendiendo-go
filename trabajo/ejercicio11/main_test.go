package main

import (
	"errors"
	"testing"
)

func preciosPrueba() map[string]float64 {
	return map[string]float64{"Leche": 10, "Café": 25, "Muestra": 0}
}

func TestDividir(t *testing.T) {
	casos := []struct {
		a, b        float64
		esperado    float64
		errEsperado error
	}{
		{10, 2, 5, nil},
		{9, 3, 3, nil},
		{10, 0, 0, ErrDivisionPorCero},
	}
	for _, c := range casos {
		obtenido, err := Dividir(c.a, c.b)
		if obtenido != c.esperado || !errors.Is(err, c.errEsperado) {
			t.Errorf("Dividir(%v, %v) = %v, %v; se esperaba %v, %v", c.a, c.b, obtenido, err, c.esperado, c.errEsperado)
		}
	}
}

func TestBuscarPrecio(t *testing.T) {
	casos := []struct {
		nombre      string
		esperado    float64
		errEsperado error
	}{
		{"Leche", 10, nil},
		{"Muestra", 0, nil},
		{"Pizza", 0, ErrNoEncontrado},
	}
	for _, c := range casos {
		obtenido, err := BuscarPrecio(preciosPrueba(), c.nombre)
		if obtenido != c.esperado || !errors.Is(err, c.errEsperado) {
			t.Errorf("BuscarPrecio(%q) = %v, %v; se esperaba %v, %v", c.nombre, obtenido, err, c.esperado, c.errEsperado)
		}
	}
}

func TestVender(t *testing.T) {
	casos := []struct {
		nombre        string
		cantidad      int
		errEsperado   error
		stockEsperado int
	}{
		{"Leche", 3, nil, 2},
		{"Leche", 5, nil, 0},
		{"Leche", 9, ErrSinStock, 5},
		{"Pizza", 1, ErrNoEncontrado, 0},
	}
	for _, c := range casos {
		stock := map[string]int{"Leche": 5}
		err := Vender(stock, c.nombre, c.cantidad)
		if !errors.Is(err, c.errEsperado) {
			t.Errorf("Vender(%q, %d) = %v; se esperaba %v", c.nombre, c.cantidad, err, c.errEsperado)
		}
		if stock[c.nombre] != c.stockEsperado {
			t.Errorf("tras Vender(%q, %d) el stock quedó en %d; se esperaba %d", c.nombre, c.cantidad, stock[c.nombre], c.stockEsperado)
		}
	}
}

func TestTotalPedido(t *testing.T) {
	casos := []struct {
		pedido      []string
		esperado    float64
		errEsperado error
	}{
		{[]string{"Leche", "Café"}, 35, nil},
		{[]string{"Leche", "Muestra", "Leche"}, 20, nil},
		{[]string{"Leche", "Pizza"}, 0, ErrNoEncontrado},
		{[]string{}, 0, nil},
	}
	for _, c := range casos {
		obtenido, err := TotalPedido(preciosPrueba(), c.pedido)
		if obtenido != c.esperado || !errors.Is(err, c.errEsperado) {
			t.Errorf("TotalPedido(%v) = %v, %v; se esperaba %v, %v", c.pedido, obtenido, err, c.esperado, c.errEsperado)
		}
	}
}
