package main

import (
	"fmt"
	"math"
)

func normalFunction(message string) {
	fmt.Println(message)
}

func tripleArgument(a int, b int, c string) {
	fmt.Println(a, b, c)
}

func returValue(a int) int {
	return a * 2
}

func dobleReturn(a int) (c, d int) {
	return a, a * 2
}

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

	//Division
	result = y / x
	fmt.Println("División: ", result)

	//Modulo
	result = y % x
	fmt.Println("Modulo: ", result)

	//Incremental
	x++
	fmt.Println("Incremental: ", x)

	//Decremental
	x--
	fmt.Println("Deceremental: ", x)

	// Recto calculo rectangulo, trapecio y de un circulo

	//Area rectangulo
	largo := 20
	ancho := 30

	areaRectangulo := largo * ancho
	fmt.Println("Area rectangulo: ", areaRectangulo)

	// Area trapecio

	base2 := 30

	areaTrapecio := ((base + base2) * altura) / 2
	fmt.Println("Area Trapecio: ", areaTrapecio)

	// Area Circulo

	var radio float64 = 30

	areaCirculo := (radio * radio) * math.Pi
	fmt.Println("Area Circulo: ", areaCirculo)

	//Uso de la función PrintF
	nombre := "Platzi"
	cursos := 500
	fmt.Printf("%s tiene más de %d cursos\n", nombre, cursos)
	//la directriz %v imprime por consola una variable si se desconoce el tipo de datos
	fmt.Printf("%v tiene más de %v cursos\n", nombre, cursos)

	//Uso de la función Sprintf
	message := fmt.Sprintf("%s tiene mas de un %d cursos", nombre, cursos)
	fmt.Println(message)

	//Tipo de datos para saber cual es el tipo de datos con %T
	fmt.Printf("message: %T\n", message)
	fmt.Printf("areaTrapecio: %T\n", areaTrapecio)

	//Funciones anonimas
	fmt.Println("--------Funciones anonimas--------")
	normalFunction("Hola Mundo")
	tripleArgument(1, 2, "hola")
	value := returValue(2)
	fmt.Println("value: ", value)

	value1, value2 := dobleReturn(2)
	fmt.Println("Doble retorno value: ", value1, value2)

	//Ciclo For
	fmt.Println("--------Ciclo For--------")
	for i := 0; i<=10; i++ {
		fmt.Println(i)
	}
}
