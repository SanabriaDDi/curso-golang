package user

import (
	"fmt"
	"time"
)

type Usuario struct {
	Id        int
	Nombre    string
	FechaAlta time.Time
	Status    bool
}

// Crea puntero para usar la estructura (this *Usuario), *Usuario es un puntero

func (this *Usuario) AltaUsuario(id int, nombre string, fechaalta time.Time, status bool) {
	this.Id = id
	this.Nombre = nombre
	this.FechaAlta = fechaalta
	this.Status = status
}

func (this *Usuario) ImprimirUsuario() {
	fmt.Printf("Id: %d, Nombre: %s, Fecha de alta: %s, Status: %t\n", this.Id, this.Nombre, this.FechaAlta, this.Status)
}
