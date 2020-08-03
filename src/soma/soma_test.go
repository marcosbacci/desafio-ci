package main

import "testing"

func TestSoma(t *testing.T) {
	got := Soma(5, 5)
	want := 10

	if got != want {
		t.Errorf("Valor esperado %v, valor retornado %v", got, want)
	}
}
