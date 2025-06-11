package main

import "fmt"

//similar al if no es necesario poner la expresion entre parentesis
// no es necesaria la palabra reservada break
// en este caso se quiere evaluar caracter y en case se ponen los valores deseados con sus bloques de codigo y el bloque general
// go tambien permite agrupar casos en caso de que el bloque a ejecutar sea similar, separando cada caso por una coma en la misma linea
func main() {
	caracter := "🧟"
	pet := "🐕‍🦺"
	canSearch := false

	switch caracter {

	case "🦸":
		fmt.Println("es un superheroe")
	case "🧞":
		fmt.Println("es un superheroe")
	case "🦹", "🧟":
		fmt.Println("es un supervillano")
	default:
		fmt.Println("es un personaje normal")
	}

	//Se permite no indicar la expresion a evaluar general para el switch si no una propia para cada caso

	switch {
	case pet == "🐕‍🦺" || pet == "🐕":
		fmt.Println("es un perro")
	case pet == "🐈" || pet == "🐈‍⬛":
		fmt.Println("es un gato")
	default:
		fmt.Println("es un mascota")

	}

	// GRACIAS A ESTA PROPIEDAD SE PUEDEN HACER MAPAS DE ESTADOS
	//recordando que switch evaluara de arriba hacia abajo cada uno de sus casos (estados)
	switch {
	case !canSearch:
		fmt.Println("no esta permitida la busqueda")
	case caracter == "🦸" || pet == "🐕‍🦺":
		fmt.Println("es un superheroe")
	case caracter == "🦹" || pet == "🐈‍⬛":
		fmt.Println("es un supervillian")

	default:
		fmt.Println("no es un verdadero super")
	}

}
