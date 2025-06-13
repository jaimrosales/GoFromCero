package main

import "fmt"

// la funcion panic nos permite entrar en panico, es decir detener la ejecucion del programa,
//al ejecutarse el programa go mostrara todo los lugares de donde proviene el error, para rastrearlo es recomentable leerlo de abajo hacia arriba

//La funcion recover permite la recuperacion de un panico, es decir recuperar la ejecucion, se puede convinar con la funcion defer para ejecutarla
//al momento de que panico returne de la funcion como se muestra en el ejemplo antes de validate zero,  ademas de una funcion anonima autoejecutada
// el recover devuelve un valor, el cual se compara para en caso de que sea diferente de nil, es decir exista un panico, de esta manera se recuperara
//del panico imprimira lo que hay dentro del iff y continuara con el programa

func main() {
	division(100, 10)
	division(200, 25)
	division(27, 0)
	division(124, 8)
}

func division(dividend, divisor int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("me recupere del manico")
		}
	}()
	validateZero(divisor)
	fmt.Println(dividend / divisor)
}

func validateZero(divisor int) {
	if divisor == 0 {
		panic("no puedes dividir entre cero🤕") //dentro de la funcion panico se indica porque se esta en panico
	}
}
