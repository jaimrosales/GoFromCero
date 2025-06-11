package main

import "fmt"

func main() {
	//operadores aritmeticos:  (),*,/,%,+,-
	var a = 2 + 3 // se debe tener en cuenta la gerarquia de operaciones  () --> -
	fmt.Println(a)

	//operadores de asignacion:   =, +=,-=,*=,/+,%=
	// a la variable se le asginara o restara un valor adicional
	var b int = 5 //asignara 5 a b
	b += 2        // le sumara 2 a b y le asignara a b el valor
	b -= 4        // lo mismo pero con resta
	// sucedera lo mismo con modulo, multiplicacion y divicion
	fmt.Println(b)
	//declaracion post-incremento y post-decremento:  ++,--
	//(no son una expresion sino una declaracion)
	var c int = 6

	// fmt.Println(c++)			no es permitido por go
	// c = c++ *2				al no ser una operacion sino una asignacion no se puede usar como valor
	c++ // es necesario que primero se le asigne para posteriormente usarse
	fmt.Println(c)

}
