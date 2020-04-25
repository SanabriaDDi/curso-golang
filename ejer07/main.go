// Funciones anonimas, es posible modificarlas en tiempo de ejecución

package main

import "fmt"

var calculo func(int, int) int = func(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Printf("Sumo 5 + 7 = %d\n", calculo(7, 7))

	// Restamos
	calculo = func(num1 int, num2 int) int { // Reescribe la función
		return num1 - num2
	}
	fmt.Printf("Restamos 6 - 4 = %d\n", calculo(6, 4))

	// Dividimos
	calculo = func(num1 int, num2 int) int { // Reescribe la función
		return num1 / num2
	}
	fmt.Printf("Dividimos 12 / 3 = %d\n", calculo(12, 3))

	operaciones()

	//Closure
	tablaDel := 2
	MiTabla := tabla(tablaDel)
	for i := 0; i < 10; i++ {
		fmt.Println(MiTabla())
	}

	tablaDel = 3
	MiTabla = tabla(tablaDel)
	for i := 0; i < 10; i++ {
		fmt.Println(MiTabla())
	}
}

func operaciones() {
	resultado := func() int { // Función anonima dentro de función
		var a int = 23
		var b int = 23
		return a + b
	}

	fmt.Println(resultado())
}

// CLOSURES
func tabla(valor int) func() int {
	numero := valor
	secuencia := 0
	return func() int {
		secuencia++
		return numero * secuencia
	}
}
