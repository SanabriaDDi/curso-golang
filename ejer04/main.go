package main

import (
	"bufio"
	"fmt"
	"os"
)

var numero1 int
var numero2 int
var resultado int
var leyenda string

func main() {
	fmt.Println("Ingrese número 1: ")
	//fmt.Scanf("%d", &numero1)
	fmt.Scanln(&numero1)

	fmt.Println("Ingrese número 2: ")
	//fmt.Scanf("%d", numero2)
	fmt.Scanln(&numero2)

	fmt.Println("Descripción: ")

	scanner := bufio.NewScanner(os.Stdin) // bufio detecta entradas por teclado, recibe por parametro el estandar input del sistema (os)
	if scanner.Scan() {                   // Comprueba que lo que se escaneo no dio error
		leyenda = scanner.Text()
	}
	resultado = numero2 + numero2
	fmt.Println(leyenda, resultado)

	//fmt.Printf("%d", numero1+numero2)

	// En windows no se ejecuta de manera correcta ya que en linux se usa un solo
	// byte (\n) y windows usa dos (\n \r) con lo que el segundo Scanf detecta como
	// entrada el segundo byte. Se soluciona cambiando por Scanln.
}
