package main

import (
	"fmt"
	"strconv"
)

var numero int //Variable privada inicia con minúscula publica para otros paquetes inicia con letra mayúscula
var texto string
var status bool = true

func main() {
	var numero2, numero3, numero4 int
	numero2, numero3, numero4, texto, status := 2, 5, 67, "Este es mi texto", false

	var moneda int64 = 0

	numero2 = int(moneda)

	//texto = fmt.Sprintf("%d", moneda) // Convierte a texto

	texto = strconv.Itoa(int(moneda)) // Convierte a texto, posee más funciones
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
