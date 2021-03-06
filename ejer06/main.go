package main

import "fmt"

func main() {
	//fmt.Println(uno(5))
	//numero, estado := dos(1) // Se inicializan las variables en el orden que la función las retorna

	//fmt.Println(numero)
	//fmt.Println(estado)

	fmt.Printf(">>> Total de suma: %d <<<\n", suma(5, 46))
	fmt.Printf(">>> Total de suma: %d <<<\n", suma(2, 23, 45, 68))
	fmt.Printf(">>> Total de suma: %d <<<\n", suma(5))
	fmt.Printf(">>> Total de suma: %d <<<\n", suma(5, 46, 17, 25, 26, 98, 17))
}

func uno(numero int) (z int) { // Recibe un número entero y retorna un entero, puede retornar más de un tipo de dato
	z = numero * 2
	return
}

// Retorno de dos valores en una función
func dos(numero int) (int, bool) {
	if numero == 1 {
		return 5, false
	} else {
		return 10, true
	}

}

func suma(numeros ...int) (total int) { //Recibe una lista de parámetros

	for pos, numero := range numeros { // Range devuelve dos valores, se le pasa un vector
		// Primer valor es el número del elemento (_ El guin bajo aloja una variable que no se desea usar)
		total = total + numero
		fmt.Printf("\nItem: %d, valor: %d \n", pos, numero)
	}
	return
}
