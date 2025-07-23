package main

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	var a, b int
	fmt.Print("Ingresa el primer número: ")
	fmt.Scan(&a)
	fmt.Print("Ingresa el segundo número: ")
	fmt.Scan(&b)
	fmt.Printf("La suma es: %d\n", a+b)
}
