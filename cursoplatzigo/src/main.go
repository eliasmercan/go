package main

import "fmt"

func main() {
	//Declaración de constantes
	const pi float64 = 3.14
	const pi2 = 3.1415

	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)

	//Declaracion de variables enteras
	base := 12
	var altura int = 14
	var area int
	fmt.Println(base, altura, area)

	//zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	//Area cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area cuadrado: ", areaCuadrado)

	//Operadores aritmeticos
	x := 10
	y := 50

	//suma
	result := x + y
	fmt.Println("Suma: ", result)

	//resta
	result = y - x
	fmt.Println("Resta: ", result)

	//multiplicacion
	result = x * y
	fmt.Println("Multiplicación: ", result)

	//multiplicacion
	result = y / x
	fmt.Println("División: ", result)
}
