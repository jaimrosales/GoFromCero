package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func main() {
	fmt.Println(sum(0.4, -2.3, 21.6))
	fmt.Println(includes([]string{"a", "b", "c", "d"}, "d"))
	fmt.Println(includes([]string{"a", "b", "c"}, "d"))
	fmt.Println(includes([]int{1, 2, 3, 4, 5, 6}, 3))
}

//una forma de a la funcion generica de anadir todos los tipos de datos genericos se puede hacer mediante interfaces, el tema se ve mas a detalle en POO

type Number interface {
	~int | ~float64 | ~float32 | ~uint
}

//sin embargo el paquete constrains contiene interfaces que contienen distintos tipos de datos, integer: enteros con signo y sin signo etc
// se puede importar go mod init un nombre
//go mod tidy
//go get golang.org/x/exp/constraints el go mod init debe estar en donde el main

func sum[T constraints.Integer | constraints.Float](nums ...T) T {
	var total T
	for _, num := range nums {
		total += num
	}
	return total
}

// constrain comparable
//este constrains viene junto con go, y permite hacer restricciones de tipo cuando solo se quieren hacer comparaciones es decir se usan los operadores += o !=
//no se define el tipo de dato, ya que solo se comparara algo
// se usa el tipo comparable lo cual permitira ingresar cualquier tipo de dato, pero no se podra sumar o algo diferente a un operador de igualdad o diferencia

func includes[T comparable](list []T, value T) bool {
	for _, item := range list {
		if item == value {
			return true
		}
	}
	return false
}
