package main

import (
	"fmt"
	"runtime"

	"github.com/godesde0/ejercicios"
	"github.com/godesde0/files"
	"github.com/godesde0/iteraciones"
	"github.com/godesde0/teclado"
	"github.com/godesde0/variables"
)

func main() {
	fmt.Println("Hola Mundo")
	variables.MuestroEnteros()
	fmt.Println("---------")
	variables.RestoVariables()

	fmt.Println("---------INFORMACION DE SISTEMA OPERATIVO------")
	if os := runtime.GOOS; os == "Linux." || os == "OS X." {
		fmt.Println("Esto no es windows, es", os)
	} else {
		fmt.Println("Windos")
	}

	fmt.Println("---------USANDO SWITCH------")
	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Linux")
	case "darwin":
		fmt.Println("Darwin")
	default:
		fmt.Printf("%s \n", os)
	}

	fmt.Println("---------USANDO SWITCH------")
	numero, texto := ejercicios.Convertir("125")
	fmt.Println("Valor numero: ", numero)
	fmt.Println("Valor Texto: ", texto)

	fmt.Println("---------EJERCICIO 01------")
	estado, texto := ejercicios.Convertir("152,5")
	fmt.Println(estado)
	fmt.Println(texto)

	fmt.Println("---------TECLADO------")
	teclado.IngresoNumeros()

	iteraciones.Iterar()
	ejercicios.TablaMultiplicar()

	fmt.Println("---------TABLA DE MULTIPLICAR POR ARCHIVO------")
	fmt.Println(ejercicios.TabladeMultiplicar())

	fmt.Println("---------GRABANDO ARCHIVO EN FILES------")
	// files.GrabaTabla()
	files.SumaTabla()

	fmt.Println("---------LEO ARCHIVO------")
	files.LeoArchivo()
}
