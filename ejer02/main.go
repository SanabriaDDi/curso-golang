package main

import (
	"fmt"
	"strconv"
)

// Si no se asigna un valor a una variable se inicializa encaso de ser numerico con 0, de texto en cadena vacía y booleano en false
var numero int //Variable privada inicia con minúscula publica para otros paquetes inicia con letra mayúscula
var texto string
var status bool = true

func main() {
	var numero2, numero3, numero4 int
	// Creación de variables encadenada asignado el tipo de dato en función del valor daro
	numero2, numero3, numero4, texto, status := 2, 5, 67, "Este es mi texto", false

	var Moneda int64 = 0 // El nombre de la variable comienza en mayúscula cuando se requiere que sea
	//visible por otros paquetes, lo mismo aplica a los nombres de las funciones

	numero2 = int(Moneda) // Es necesario realizar la conversion de la variable

	//texto = fmt.Sprintf("%d", Moneda) // Convierte a texto, se debe especificar en el verbo que tipo de dato va a recibir

	texto = strconv.Itoa(int(Moneda)) // Convierte entero a texto, posee más funciones
	fmt.Println(numero2)
	fmt.Println(numero3)
	fmt.Println(numero4)
	fmt.Println(texto)
	fmt.Println(status)

	mostrarStatus()
}

func mostrarStatus() {
	fmt.Println(status)
}
