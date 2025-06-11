package main

import "fmt"

func main() {

	var apple string = "apple" // se declara una variable con la palabra resevada var, el nombre de la variable  y el tipo de dato
	var banana string = "banana"
	var orange string = "orange"

	var ( // declaras de manera anidada
		coconut  string = "coconut" //
		pepino   string = "pepino"
		jitomate string = "jitomate"
	)

	var nuez, almendra, pepita string = "nuez", "almendra", "pepita" // declaras variables de un mismo tipo en una sola linea

	mango, limon, zanahoria := "mango", "limon", "zanahoria" // declaras la variable de manera dinamica, declaracion de variable corta, asigna el tipo de dato de manera automatica

	//apple =1 una vez asignado un tipo no se podra cambiar el tipo de dato para cada variable una vez asignado
	//apple := "manzana"   apple = manzana no se puede declarar por segunda vez una variable pero si se puede maciar el valor
	apple, toronja := "manzana", "toronja" // en este caso cambiaras el valor de apple y declaras asignas y tipeas la variable toronja

	fmt.Println(apple, banana, orange, coconut, pepino, jitomate, nuez, almendra, pepita, mango, limon, zanahoria, toronja) // se utiliza la variable, se imprime en terminal

}

//*toda variable no utilizada se marcara con errores
