package ejercicios

import (
	"strconv"
)

func Convertir(valor string) (int, string) {
	numero, err := strconv.Atoi(valor)
	if err != nil {
		return 0, "Hubo un error " + err.Error()
	}
	/*
		condicionar para validar si hay un error

	*/
	if numero > 100 {
		return numero, "Es Mayor a 100"
	} else {
		return numero, "Es menor a 100"
	}
}
