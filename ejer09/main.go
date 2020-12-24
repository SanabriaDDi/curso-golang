// Mapas
package main

import "fmt"

func main() {
	paises := make(map[string]string) // Crea mapa con clave tipo string y valor tipo string
	fmt.Println(paises)

	paises["Mexico"] = "D.F."
	paises["Argentina"] = "Buenos Aires"
	//fmt.Println(paises)

	campeonato := map[string]int{ //Asignación directa usando llaves
		"Barcelona":    39,
		"Real Madrid":  38,
		"Chivas":       37,
		"Boca Juniors": 30}

	fmt.Println(campeonato)

	campeonato["River Plate"] = 25 // Asignacion a map
	campeonato["Chivas"] = 55      // Modifica valor por clave si lo encuentra, si no crea una nueva clave valor
	delete(campeonato, "Real Madrid")
	fmt.Println(campeonato)

	for equipo, puntaje := range campeonato {
		fmt.Printf("El equipo %s, tiene un puntaje de: %d\n", equipo, puntaje)
	}

	puntaje, existe := campeonato["Mineiro"]
	fmt.Printf("El puntaje capturado es %d, y el equipo existe %t \n", puntaje, existe)
}
