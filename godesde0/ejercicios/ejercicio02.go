package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero1 int
var variable int
var err error

func TablaMultiplicar() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Ingrese número para generar tabla de multiplicar")
	if scanner.Scan() {
		numero1, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("Error al ingresar el registro ")

		}
	}
	fmt.Println("-----Imprimiendo la tabla del ", numero1)
	for i := 1; i <= 10; i++ {
		variable = i * numero1
		fmt.Printf("%d X %d = %d\n", numero1, i, variable)
	}

}
