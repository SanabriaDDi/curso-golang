package main

import "log"

func main() {
	/*archivo := "prueba.txt"
	f, err := os.Open(archivo)

	defer f.Close()

	if err != nil {
		println("Error abriendo el archivo")
		os.Exit(1)
	}*/

	ejemploPanic()
}

func ejemploPanic() {
	defer func() {
		reco := recover() // Recover verifica el ultimo panic
		if reco != nil {
			//Ocurrio un panic
			log.Fatalf("ocurrió un error que generó panic \n %v", reco)
		}
	}()
	a := 1
	if a == 1 {
		panic("Se encontro el valor de 1")
	}
}
