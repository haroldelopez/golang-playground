package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

var y int16
var x int = 45

func main() {

	n, nerr := fmt.Println("1", 2)
	fmt.Println(n)
	fmt.Println(nerr)

	/*go say("Hola", 2*time.Second)
	go say("Cómo estás?", 1*time.Second)*/

	fmt.Printf("%T\n", y)

	//Con esto le digo al programa que espere que terminen de ejecutar las go routines para finalizar. S
	time.Sleep(4 * time.Second)
}

func say(texto string, segs time.Duration) {
	fmt.Println("He ingresado a say")
	time.Sleep(segs)
	fmt.Println("He finalizado say " + texto)
}
func TiempoTranscurrido() {
	//	ahora := time.N
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

func SolicitarInformacion() {
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
}
