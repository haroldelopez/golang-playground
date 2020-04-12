package main

import (
	"fmt"
	"strconv"
)

func ciclos() {
	for i := 0; i < 15; i++ {
		fmt.Println("Ciclo for en " + strconv.Itoa(i))
	}

	i := 0
	for i < 15 {
		fmt.Println("Ciclo while en " + strconv.Itoa(i))
		if i == 9 {
			break
		}
		i++
	}
	fmt.Println("Voy a terminar el programa")
}
