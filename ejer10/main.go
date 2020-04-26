// Estructuras (POO)
package main

import (
	"fmt"
	"time"
	//"time"

	us "./user" // Importa paquete user.go. NOTA: ./user/user.go no funciona, se asume que el nombre de la carpeta es el nombre del archivo *.go
)

/* type usuario struct { // Define estructura
	Id        int
	Nombre    string
	FechaAlta time.Time
	Status    bool
} */

type pepe struct {
	us.Usuario
}

func main() {
	/* User := new(usuario) // Crea objeto
	User.Id = 10
	User.Nombre = "Pablo"
	fmt.Println(User) */

	u := new(pepe)
	u.AltaUsuario(1, "Diego", time.Now(), true)
	fmt.Println(u.Usuario)
}
