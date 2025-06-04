package main

import "fmt"

func main() {
	//Operadores Comparacion: >, <, ==, !=, >=, <=
	//ya los conoces, usados para comparacion de intervalos
	//son operadores de expresion y no de asignacion

	fmt.Println((4 + 5) > 6)
	//Operadores Logicos: &&, ||
	//ya los conoces usados para logica boleana y compuertas logicas
	var age uint = 33
	fmt.Println("es adulto?  ", age >= 18 && age <= 60)
	fmt.Println("es nino o anciano?  ", age <= 18 || age >= 60)

	//operadores logico unario: !
	// es el operador de negacion
	fmt.Println(!(4 == 4))
}
