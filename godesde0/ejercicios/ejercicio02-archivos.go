package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero int
var val int
var errr error
var texto string

func TabladeMultiplicar() string {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Ingrese número para generar tabla de multiplicar")
		if scanner.Scan() {
			numero, errr = strconv.Atoi(scanner.Text())
			if err != nil {
				// panic("Error al ingresar el registro ")
				continue
			} else {
				break
			}
		}
	}
	fmt.Println("-----Imprimiendo la tabla del ", numero)
	for i := 1; i <= 10; i++ {
		val = i * numero
		texto += fmt.Sprintf("%d X %d = %d\n", numero, i, val)
	}

	return texto

}
