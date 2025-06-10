package main

import "fmt"

//en Go los arrays son de tamano fijo, es decir no se pueden anadir mas elementos al array
func main() {

	var flags [3]string // declaracion
	flags[0] = ":flags" //asignacion de valores
	flags[1] = ":black_flag"
	flags[2] = ":white_flag"

	fmt.Println(flags)
	//segunda manera de declarar y asignar
	bandera := [3]string{":white_flag", ":black_flag", ":flags"}
	fmt.Println(bandera)
	//si no se conoce el valor exacto de arrays se puede usar la siguiente manera
	boat := [...]string{"lancha", "canoa", "yate"}
	// sin embargo no se permite anadir o modificar el tamano del array despues de su declaracion
	fmt.Println(boat)
}
