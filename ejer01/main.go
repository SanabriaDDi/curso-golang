package main

import (
	"fmt"
	"github.com/gookit/color"
)

func main() {
	//fmt.Println("Hola Mundo")

	var tamanio = 35;
	c := color.RGB(128,64,0) // color café
	for i := 0; i<tamanio; i++  {
		for j := 0; j < tamanio-i; j++ {
			fmt.Print(" ")
		}
		for k := 0; k < i*2-1; k++ {
			color.Info.Print("x")
		}
		fmt.Print("\n")

	}
	for l := 0; l < 4; l++ {

		for n := 0; n < tamanio-2; n++ {
			fmt.Print(" ")
		}
		c.Println("|||")
	}
}
