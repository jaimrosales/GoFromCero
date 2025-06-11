package main

import "fmt"

func main() {
	caracter := "🦸"
	// a diferencia de otros lenguajes la condicion del if no es necesario que este entre parentesis
	// se agrega entre llaves el codigo a reproducir cuando la condicion se cumple
	// para encadenar se puede usar else if como se muestra seguido de la siguiente condicion a evaluar
	// seguido de {bloque de codigo a ejecutar}
	// se puede usar un else{} como bloque de codigo generico

	if caracter == "🦸" {
		fmt.Println("es un superheroe")
	} else if caracter == "🦹" {
		fmt.Println("es un supervillano")
	} else {
		fmt.Println("es un personaje normal")
	}

	// se permite indicar que las variables esten unicamente al alcance de la estructura if
	// para declarar la se declara como normalmente despues del if seguido por un ; y la condicion
	if pet := "🐕‍🦺"; pet == "🐕‍🦺" {
		fmt.Println("es un perro")
	} else if pet == "🐈" {
		fmt.Println("es un gato")
	} else {
		fmt.Println("es un mascota")
	}
	//despues de aqui ya no se podra usar la variable

}
