// Arrays
package main

import "fmt"

//var tabla [10]int
//var matriz [5][7]int

func main() {
	//tabla[0] = 1
	//tabla[5] = 15

	/*tabla := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < len(tabla); i++ {
		fmt.Println(tabla[i])
	}*/
	//fmt.Println(tabla)
	//matriz[3][5] = 1
	//fmt.Println(matriz)

	//Slices vectores dinámicos
	//matriz := []int{1, 2, 3} //No se asigna longitud en un slice
	//fmt.Println(matriz)

	variante2()
	fmt.Println("-----------")
	variante3()
	fmt.Println("-----------")
	variante4()
}

func variante2() {
	elementos := [5]int{1, 2, 3, 4, 5} // vector
	porcion := elementos[3:]           //Crea slice porcion desde el elemento 3 hasta ek último
	fmt.Println(porcion)
}

func variante3() {
	elementos := make([]int, 5, 20) // Crea slice con make, pasa el array, el tamaño inicial, y el máximo que podrá tener en memoria
	fmt.Printf("Largo: %d, Capacidad: %d\n", len(elementos), cap(elementos))
}

func variante4() {
	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i) // La capacidad crece dinámicamente dependiendo de la longitud de los elemetos
	}
	fmt.Printf("Largo: %d, Capacidad: %d", len(nums), cap(nums))
}
