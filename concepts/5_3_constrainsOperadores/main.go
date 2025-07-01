package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func main() {
	fmt.Println(sum(0.4, -2.3, 21.6))
	fmt.Println(includes([]string{"a", "b", "c", "d"}, "d")) // [] indica el tipo de dato, se deja vacio para que go lo infiera compara si en la lista de strings existe  d
	fmt.Println(includes([]string{"a", "b", "c"}, "d"))      // compara si en la lista de strings existe  d
	fmt.Println(includes([]int{1, 2, 3, 4, 5, 6}, 3))        // compara si en la lista de numeros existe  3
	fmt.Println(filter([]int{2, 12, 23, 98, 21, 79}, lessThanTwenty))
	fmt.Println(filtro([]int{2, 12, 23, 98, 21, 79}, lessThanTwenty))
	fmt.Println(filtro([]int{2, 9, 20, 98, 21, 79}, menorqueVeinte))
}

//una forma de a la funcion generica de anadir todos los tipos de datos genericos se puede hacer mediante interfaces, el tema se ve mas a detalle en POO

type Number interface {
	~int | ~float64 | ~float32 | ~uint
}

//sin embargo el paquete constrains contiene interfaces que contienen distintos tipos de datos, integer: enteros con signo y sin signo etc
// se puede importar go mod init "un nombre para el mod"
//go mod tidy
//go get golang.org/x/exp/constraints el go mod init debe estar en donde el main
//una vez obtenido el paquete con go get, se agrega como cualquier libreria en import

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

// la funcion includes, manejara un constrains de tipo comparable, recibira una lista con el constrains generico comparable, y un valor con el constrains generico
//retorna un boleano
// el for itera por toda la lista y guarda el valor en item, para posterior evalue si el item es igual al valor, encaso de que si regresa true, con esto se evalua
// si un dato existe dentro de una lista

func includes[T comparable](list []T, value T) bool { //
	for _, item := range list {
		if item == value {
			return true
		}
	}
	return false
}

// tomando como ejemplo la funcion callback del 4_5, la funcion recive un slice de numero, y utilizando la funcion callback, filtraba dentro de un nuevo slice, los resultados con las especificaciones dadas
// para convertirla en funcion generica se agrega el parametro de tipo a la funcion entre corchetes antes de los parametros
// Utilizando el paquete constrains se indica el constrains general de numeros, para poder utilizar float e integer

// La funcion callback es el comparador, ingresa un numero y retorna un boleano identificando si se cumplio una condicion dada o no, en este caso ballback llama a leesThan Twnty
func filter[T constraints.Float | constraints.Integer](nums []T, callback func(T) bool) []T {
	result := make([]T, 0, len(nums))

	for _, num := range nums {
		if callback(num) {
			result = append(result, num)
		}
	}
	return result
}

func lessThanTwenty(num int) bool {
	return num < 20
}

// si se quiere comparar distintos tipos de datos no solo numericos si no tambien  string, se puede usar el constrains Ordered, que se encuentra dentro de constrans
// este tipo de dato permite los operadores <, <=, >=, >, con la funcion filtro
func filtro[T constraints.Ordered](nums []T, callback func(T) bool) []T {
	result := make([]T, 0, len(nums))

	for _, num := range nums {
		if callback(num) {
			result = append(result, num)
		}
	}
	return result
}
func menorqueVeinte(num int) bool {
	return num <= 20
}
