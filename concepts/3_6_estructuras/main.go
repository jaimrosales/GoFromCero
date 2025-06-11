package main

import "fmt"

// las estructuras permiten almacenar una coleccion de campos similar a las clases en otros lenguajes
//declaracion de una structura   type [name] struct {se definen los campos de la estructura}
// por lo general no es regla las estructuras se declaran a nivel packete para tenerlos disponibles en el paquete

type Person struct {
	name        string
	age         uint8
	haschildren bool
}

// para crear instancias de una persona es
//[nombre]  := [structName]{se define cada valor}

func main() {
	alexis := Person{
		name:        "alexis",
		age:         42,
		haschildren: true,
	}

	fmt.Printf("%v \n", alexis)

	ximena := Person{"ximena", 19, false} //de esta manera no es necesario especificar el dato pero si seguir el orden
	fmt.Printf("%v \n", ximena)

	ruben := Person{age: 20}
	fmt.Printf("%v \n", ruben)
	fmt.Printf("%v \n", alexis.name) // con el operador . indicas el valor en especifico

	//tambien se pueden hacer punteros a estructuras
	hannia := &ximena
	//de esta manera se pueden modificar valores desde un puntero
	(*hannia).age = 20

	fmt.Printf("%v \n", hannia)

	//o bien
	hannia.age = 19
	fmt.Printf("%v \n", hannia)

}
