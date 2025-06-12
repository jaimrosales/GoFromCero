package main

import (
	"fmt"
	"strings"
)

func main() {
	result := sum(2, 3)
	fmt.Println(result)
	low, up := convert("RStech")
	fmt.Println(low, up)

}

// cuando se quiere returnar un valor de una funcion la estructura es la siguiente
// func [name](argumentos)[Datatype del dato a retornar]{[codigo]return[valor de return]}
func sum(a, b int) int {
	return a + b
}

// las funciones pueden retornar multiples valores
func convert(text string) (string, string) {
	lower := strings.ToLower(text)
	upper := strings.ToUpper(text)

	return lower, upper
}
