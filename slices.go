package main

import "fmt"

func slices() {
	arreglo := [5]int{1, 28, 8, 4, 6}
	fmt.Println("Arreglo ", arreglo)

	slice := arreglo[:5]
	fmt.Println("Slice ", slice)
}
