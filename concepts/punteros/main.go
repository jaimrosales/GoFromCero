package main

// Puntero: variables que almacenan la direccion en memoria de un valor

import "fmt"

func main() {
	var color string = ":red_square"
	// el  nos permitira obtener la direccion de memoria de una variable, para usarlo  solo se colorara antes del nombre de la variable

	fmt.Printf("tipo: %T, Valor %s, Diraccion: %v\n", color, color, &color)

	//la sintaxis para declarar un puntero es igual a la sintaxis de una variable normal, se le agrega un * antes del tipo de dato para indicar que se
	//esta creando un puntero a la direccion de otra variable
	//anteponiendo el * al nombre del puntero se desreferencia, se podria decir que * indica que se apunta a una referencia desreferenciacion = referencia de referencia

	var pointerColor *string //se crea el puntero
	pointerColor = &color    //se le asigna al puntero el valor de direccion
	fmt.Printf("Tipo: %T, Valor: %v, Desreferenciacion: %s \n", pointerColor, pointerColor, *pointerColor)

	//Se puede cambiar el valor usando el * (operador de referenciacion)
	*pointerColor = ":blue_square"
	fmt.Println(*pointerColor)

}
