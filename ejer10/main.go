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

//Herencia de Usuario
type pepe struct {
	us.Usuario
}

func main() {
	/* User := new(usuario) // Crea objeto
	User.Id = 10
	User.Nombre = "Pablo"
	fmt.Println(User) */

	//u := new(pepe)
	u := pepe{us.Usuario{Id: 10, Nombre: "asd", FechaAlta: time.Now(), Status: true}}
	u.ImprimirUsuario()
	u.AltaUsuario(1, "Diego", time.Now(), true)
	fmt.Println(u.Usuario)
	fmt.Println("------------------")

	d := us.Usuario{Id: 10, Nombre: "Diego", FechaAlta: time.Now(), Status: true}
	d.ImprimirUsuario()
	fmt.Println("------------------")

	

}
