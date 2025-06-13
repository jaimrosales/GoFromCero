package main

import "fmt"

//las funciones variaticas permitiran recibir un numero n de argumentos que podran ser procesados dentro de la funcion
//las funciones anonimas son aquellas funciones que no tienen nombre

func main() {
	fmt.Println(sum(2))
	fmt.Println(suma(2, 3))
	fmt.Println(sum(2, 3, 12))
	fmt.Println(suma(2, 3, 12, 1, 24))
	//las funciones anonimas pueden ser asignadas a alguna variable como enseguida
	greet := func() {

		fmt.Println("hello, is mee")
	}
	greet()

	//o tambien pueden ser autoejecutadas, para hacer esto al final de la ecuacion se pondran unos parentesis para autoejecutarla
	// de igual manera se pueden agregar parametros y los argumentos se escribiran entre los parentesis de al final de la funcion
	func(name string) {
		fmt.Println("hello, is mee")
	}("rstechers")
}

//FUNCIONES VARIATICAS

//para configurarle se recibira un argumento al cual su tipo de dato se antepondran 3 puntos como se muestran
//con los tres puntos lo que generas es convertir a la variable en un slice y para el bloque de codigo de la funcion se usan las formas de procesar slice
//par ael ejemplo se crea una variable donde se guardara el resultado de la sum y posterior mente se itera por  cada dato dentro del slice

func sum(nums ...int) int {
	var total int
	for _, num := range nums {
		total += num
	}

	return total
}

//O tambien se pueden utilizar retornos nombrados

func suma(nums ...int) (total int) {
	for _, num := range nums {
		total += num
	}
	return
}
