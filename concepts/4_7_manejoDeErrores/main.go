package main

import (
	"errors"
	"fmt"
	"strconv"
) // la libreria nos permite transformar strings a otros tipos de datos

// a diferencia de otros lenguajes Go, trabaja con errores y no con excepciones
//esto devido a que las excepciones no siempre son controladas por el desarrollador
//con la idea de que los errores sean abordados en el momento en que se presentan
//motivo por el cual retornan multiples valores, usualmente como ultimo argumento de retorno las funciones devuelven un error para poder validar
// si la funcion produjo un error y asi podermo controlar

//string atoi convierte un string a un numero, si se le escribe una letra podria ocasionar un error, como retorno string.atoi tiene un int y un error
// por eso el int se guarda en num y el error en err, por convencion la variable donde se guardan los errores es err
// el valor cero de los errores es nil, por lo que cuando el valor de err sea diferente a nil, es decir exista un error, se aboradara como en el if
//para el caso se publicara el valor del error y se terminara el programa

// para crear un error, se declara como variable y se agrega al paquete de errors, con .new

var errNotfound = errors.New("Not found")

// mapa de la funcion search, le llega una llave tipo int y devuelve un emogi
var food = map[int]string{
	1: "🍕",
	2: "🍔",
}

// una buena opractica para cuando se tieenen muchas funciones anidadas es agregar un comentario al valor del error con fmt.errorf("dedondeviene el error, %v",err)
func main() {
	num, err := strconv.Atoi("10")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("valor: %v, Tipo: %T\n", num, num)

	// crear un error muy basico es como sigue
	/*encontrado, err := buscar("RSTech")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("valor: %v, Tipo: %T", encontrado, encontrado)
	*/

	//ejemplo de un error personal en accion prueba cambiando el valor del parametro dentro de search por diferentes valor o letras "1","2""34"
	// el primer if es un ejemplo de cuando se quiera tratar un error de manera particular,

	found, err := search("23")
	if errors.Is(err, errNotfound) { // errors.is es una funcion que nos permite comprobar si el error comparado es el buscado como == pero con errores al usar %w
		fmt.Println("pudimos controlar el error correctamente")
		return
	}
	if err != nil {
		fmt.Println("search(): ", err)
		return
	}
	fmt.Printf("valor: %v, Tipo: %T\n", found, found)

}

// la siguiente funcion  recive una llave de tipo string, devuelve un valor tipo string o un error

func buscar(llave string) (string, error) {

	return "", errNotfound
}

//funcion search, recibe una llave, la convertira en int y buscara en un mapa dicho valor
// se transforma de string a numero la variable key y se guarda dentro de num y si existe error o no se guarda en err
//posteriormente se evalua si existe error y se genera una respuesta en caso de que exista, piublica el error
// despues de la comprobacion, con la funcion find food busca si existe alguna comida con esa designacion o si genera un error
// se vuelve a controlar el error y finalmente se devuelve el emoji y se dice que no existe un error

func search(key string) (string, error) {
	num, err := strconv.Atoi(key)
	if err != nil {
		fmt.Println(err)
		return "", fmt.Errorf("strconv.Atoi():%v", err)
	}
	emoji, err := findFood(num)
	if err != nil {
		fmt.Println(err)
		return "", fmt.Errorf("findFood():%w", err) //%w es un verbo adicional de la funcion errorf para poder envolver el error dentro del error, permitiendo a error .is inspeccionar el error dentro del error
	}

	return emoji, nil
}

// find food recive una llave id de tipo entero y retornara un string y un error
//a la variable value le asignara el valor de food con la respectiva llave y si existe o no existe se le asignara a la variable ok

//cuando ok sea cero, es decir no exista, entrara al if, retornando un string vacio y el error que se ah creado en este caso not found
// en caso de que se encuentre retornara el valor del emogi y el error de la funcion sera nil lo que servira para la comprobacion durante el programa principal

func findFood(id int) (string, error) {
	value, ok := food[id]
	if !ok {
		return "", errNotfound
	}
	return value, nil
}
