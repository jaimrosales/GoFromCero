package main

import (
	"fmt"
	"strings" // este packete tiene distintas funcionalidades para trbaajar con strings
)

// las funciones se declaran con la palabra reservada func, y se llaman con su nombre seguido de sus argumentos
func main() {
	greet()
	saludar("Jaime", "Rosales")
	bienvenido("ricardo", "emmanuel", 1, 2, 3)

	name := "alexis"
	toUpperCase(name)
	fmt.Println(name)
	aMayusculas(&name)
	fmt.Println(name)
}

func greet() {
	fmt.Println("hello gophers")
}

func saludar(firstName string, lastname string) {
	fmt.Println("hello ", firstName, " ", lastname)
}

// si tienes tipos de datos similares en los primeros valores puedes omitir el tipo de dato
func bienvenido(primerNombre, apellido string, a, b, c int) {
	fmt.Println("hello ", primerNombre, " ", apellido, a, b, c)
}

func toUpperCase(text string) { //las funciones solo trabajaran con parametros por valor no modifican variables
	text = strings.ToUpper(text)

}

// para modificar los valores de las variables se usan funciones para pasarlos por referencia, punteros
func aMayusculas(text *string) {
	*text = strings.ToUpper(*text)
}
