package main

import "fmt"

// las funciones se declaran con la palabra reservada func, y se llaman con su nombre seguido de sus argumentos
func main() {
	greet()
	saludar("Jaime", "Rosales")
	bienvenido("ricardo", "emmanuel", 1, 2, 3)
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
