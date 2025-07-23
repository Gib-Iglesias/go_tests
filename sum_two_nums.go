package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("Ingresa el primer número: ")
	fmt.Scan(&a)
	fmt.Print("Ingresa el segundo número: ")
	fmt.Scan(&b)
	fmt.Printf("La suma es: %d\n", a+b)
}
