package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	var nombre string
	/*var edad int
	fmt.Printf("Introduzca su nombre\n")
	fmt.Scanf("%d\n",&nombre)
	fmt.Printf("Introduzca su apellido\n")
	fmt.Scanf("%d\n",&apellido)
	fmt.Printf("Introduzca su edad\n")
	fmt.Scanf("%d\n",&edad)
	fmt.Printf("¡Hola! "+nombre + " "+apellido +", su edad registrada es de %d",edad)*/
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Introduzca su nombre\n")
	var err error
	nombre, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Hola " + nombre)
	}

	//	ahora := time.Now
	//	futuro := time.Now
	bandera := false
	for bandera == false {
		fmt.Println("Tiempo transcurrido es ")
		/*	if time.Since(ahora) == time.Until(futuro) {
			bandera = true
		}*/
	}

	//fmt.Println("Voy a parar por " + (t.Sleep(30 * time.Second).).)
	time.Sleep(30 * time.Second)
	fmt.Println("Adiós")
}