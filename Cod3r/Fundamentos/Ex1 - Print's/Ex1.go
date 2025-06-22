package main

import "fmt"

func main() {
	fmt.Println("Hello, World!");
	fmt.Print("Primeiro programa!\n");
	fmt.Printf("Outro programa em %s!!!", "Go\n")
	x := 10
	y := 20.2412232
	s := "Olá, Mundo!"
	
	xs := fmt.Sprint(x)
	fmt.Println("O valor de x é: " + xs)
	fmt.Println("O valor de x é: ", x)

	fmt.Printf("\nO valor de Y é %.2f", y)
	fmt.Printf("\nMostrar tipo string %s", s)

	a := 1
	b := 1.9999
	c := false
	d := "Teste"
	fmt.Printf("\nValores: %d, %.2f, %f, %t, %s", a, b, b, c, d)
	fmt.Printf("\nValores: %v, %v, %v, %v", a, b, c, d)

}