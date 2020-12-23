package main

import "fmt"

func main() {

	//Ciclo básico
	/*for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}*/

	//Ciclo infinito:
	/*numero := 0
	for {
		fmt.Println("Continuo")
		fmt.Println("Ingrese el número secreto")
		fmt.Scanln(&numero)
		if numero == 100 {
			break // Rompe el ciclo
		}
	}*/

	/*var i = 0
	for i < 10 {
		fmt.Printf("\n Valor de i: %d", i)
		if i == 5 {
			fmt.Printf(" multiplicamos por 2 \n")
			i = i * 2
			continue	//Omite lo siguiente a la iteración
		}
		fmt.Printf("	Pasó por aquí \n")
		i++

	}*/

	var i int = 0

RUTINA: //Etiqueta
	fmt.Println("Antes del for")
	for i < 10 {
		if i == 4 {
			fmt.Println("voy a RUTINA sumando 2 a ", i, ": ")
			i = i + 2
			goto RUTINA // Regresa a la etiqueta RUTINA, diferente a continue
		}
		fmt.Printf("Valor de i: %d\n", i)
		i++
	}

}
