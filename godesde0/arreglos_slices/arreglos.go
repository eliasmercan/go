package arreglos_slices

import "fmt"

// var tabla [10]int

var tabla [10]int = [10]int{10, 0, 59}

func MuestroArreglos() {
	tabla[7] = 33
	tabla[2] = 54
	fmt.Println(tabla)
}
