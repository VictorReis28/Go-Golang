package main

import m "math"


func main() {
	const PI float64 = 3.1415 // Declaração de constante
	var raio = 3.2            // Tipo inferido pelo compilador (float64)

	area := PI * m.Pow(raio, 2)                // Declaração de variável com inferência de tipo
	println("A área da circunferência é:", area) // Imprime a área
	// Declaração de constantes e variáveis
	const (
		a = 1
		b = 2
		c = 3
	)

	var(
		d = 4
		e = 5
		f = 6
	)
	println("Constantes:", a, b, c) // Imprime constantes
	println("Variáveis:", d, e, f) // Imprime variáveis

	
	var g, h = true, false // Declaração de múltiplas variáveis
	println("Variáveis booleanas:", g, h) // Imprime variáveis booleanas

	i, j, k := 10, false, "oi!" // Declaração de múltiplas variáveis com inferência de tipo
	println("Variáveis inteiras:", i, j, k) // Imprime variáveis 
}