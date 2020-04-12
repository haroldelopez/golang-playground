package main

import (
	"fmt"
	"strconv"
)

func arreglos() {
	var arreglo [5]int
	arreglo2 := [27]int{1, 28, 8, 4, 1, 29, 9, 12, 3, 5, 4, 345, 46, 84, 83, 5, 345, 634, 56, 3456, 345, 14, 35, 2345}
	fmt.Println(arreglo)

	for i := 0; i < len(arreglo2); i++ {
		fmt.Println(arreglo2[i])
	}
	fmt.Println("ANTES ES " + strconv.Itoa(arreglo2[15]))
	arreglo2[15] = arreglo2[15] * 2
	fmt.Println("DESPUES ES " + strconv.Itoa(arreglo2[15]))

	var matriz [3][3]int
	matriz[1][2] = 5
	fmt.Println(matriz)
}
