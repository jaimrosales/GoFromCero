package main

import "fmt"

func main() {
	PrintList("RStech", "gopheers", "goFromCero") //esta funcion no permite argtumentos difernetes a string debido a que es estaticamente tipado,
	PrintListint(1, 2, 3)                         // al realizar esto se tienen dos funciones diferentes que en base haceen lo mismo con diferente tipo de dato
	PrintListAny("RStech", "gopheers", "goFromCero")
	PrintListAny(1, 2, 3)
	PrintListAny(1, "RStech", 3)
}

//Ejemplo de una funcion no generica, es decir una funcion normal
func PrintList(list ...string) { //los 3 puntos indican una funcion variatica
	for _, item := range list {
		fmt.Println(item)
	}
}

//en caso de que se quiera imprimir una lista de un tipo de datos diferentes es necesarion duplicar la logica,
func PrintListint(list ...int) { //los 3 puntos indican una funcion variatica
	for _, item := range list {
		fmt.Println(item)
	}
}

//las funciones genericas permiten escribir funciones de manera generica que permiten trabajar con multiples de datos
//el primer caso es el tipo de funcion generica es el uso del tipode dato any
func PrintListAny(list ...any) {
	for _, item := range list {
		fmt.Println(item)
	}
}
