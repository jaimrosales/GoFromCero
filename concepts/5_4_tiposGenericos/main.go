package main

import (
	"fmt"
)

//tipos genericos
//Los tipos se pueden parametrizar o configurar utilizando parametros de tipos
//util para utilizar tipos de forma generica

// Suponiendo que se tiene una base de datos relacional, donde se almacenan productos con ID,
type Product struct {
	ID          uint
	description string
	Price       float64
}

// suponiendo que despues se decide utilizar tambien una base de datos documental, el ID ya no seria de tipo entero si no de tipo string
// por lo que tendrias que tener dos versiones diferentes de estructuras una con el ID tipo int y otra tipo string
// ademas que durante el codigo estaras intercambiando entre una y otra
type ProductV2 struct {
	ID          string
	description string
	Price       float64
}

//esto genera que pareciera que se trata de dos tipos de productos cuando solo se tiene un solo tipo de producto, se tiene demasiado codigo
//para este tipo de casos es cuando se utiliza tipos genericos, de manera que dentro de la estructura se configuren generalidades y poder tener una sola
//estructura sin importar que tipo de base de datos se utilice

//Para realizar esto se genera la estructura de manera normal y se agregan con corchetes los parametros de tipos posterior al nombre de manera similar a las funciones
//El concepto es similar a funciones, se agregan que tipos son validos para la funcion con "|" y el tipo del ID se define como T
// de esta manera no importa que tipo de ID (int o string) tenga la base de datos, se podra interactuar con el con una sola estructura
//sin embargo si se tiene que indicar el tipo que se manejara

type ProductV3[T uint | string] struct {
	ID          T
	description string
	Price       float64
}

func main() {
	product1 := Product{ID: 1, description: "zapatos", Price: 30}
	product2 := ProductV2{ID: "2abcd29", description: "zapatos", Price: 40}
	product3 := ProductV3[uint]{ID: 1, description: "zapatos", Price: 30} //es lo mismo que el producto 1 pero se indica el tipo a manejar
	product4 := ProductV3[string]{ID: "2abcd29", description: "zapatos", Price: 40}
	fmt.Println(product1)
	fmt.Println(product2)
	fmt.Println(product3)
	fmt.Println(product4)
}

//La recomendacion para saber cuando usar genericos, es basicamente cuando te veas forzado a duplicar el codigo, solo porque el tipo de cambio
//o porque se quiere dar soporte a varios tipos de codigo
