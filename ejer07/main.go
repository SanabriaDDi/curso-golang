// Funciones anonimas, es posible modificarlas en tiempo de ejecución

package main

import "fmt"

// Declara variable calculo 												var calculo
// De tipo funcion y recibe 2 datos de tipo entero 							func(int, int) int
// Devuelve un resultado de tipo entero										int
// Asigna una función indicando los nombre de las variables y tipos de dato	func(num1 int, num2 int) int
// Se deben respetar los datos de entrada del tipo de función al volver a definirla
var calculo func(int, int) int = func(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Printf("Sumo 5 + 7 = %d\n", calculo(5, 7))

	// Restamos
	calculo = func(num1 int, num2 int) int { // Reescribe la función
		return num1 - num2
	}
	fmt.Printf("Restamos 6 - 4 = %d\n", calculo(6, 4))

	// Dividimos
	calculo = func(num1 int, num2 int) int { // Reescribe la función
		return num1 / num2
	}
	fmt.Printf("Dividimos 12 / 3 = %d\n", calculo(12, 5))

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
// Crea función tabla que recibe un valor y retorna una función
func tabla(valor int) (retorno func() int) {
	numero := valor
	secuencia := 0

	if valor == 3 {

		retorno = func() int {

			secuencia--
			return numero * secuencia
		}
	} else {

		retorno = func() int {
			secuencia++
			return numero * secuencia
		}
	}

	return retorno
	/*return func() int {
		secuencia++
		return numero * secuencia
	}*/
}
