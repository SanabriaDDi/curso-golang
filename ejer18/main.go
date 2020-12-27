package main

import "fmt"

/*Middlewares:
Son interceptores que permiten ejecutar instrucciones comunes a varias
funciones que reciben y devuelven los mismos tipos de variables*/

var result int
func main() {
	fmt.Println("Inicio")

	result = operacionesMidd(sumar)(2,3)
	fmt.Println(result)
	result = operacionesMidd(restar)(10,6)
	fmt.Println(result)
	result = operacionesMidd(multiplicar)(2,4)
	fmt.Println(result)
}

func sumar(a, b int) int {
	return a + b
}
func restar(a, b int) int {
	return a - b
}
func multiplicar(a, b int) int {
	return a * b
}

//Respeta la entrada y salida el mismo tipo de dato
func operacionesMidd(f func(int, int) int) func(int, int) int{
	return func(a,b int) int {
		fmt.Println("Inicio de operación")
		return f(a,b)
	}
}
