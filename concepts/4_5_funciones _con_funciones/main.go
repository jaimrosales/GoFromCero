package main

import "fmt"

//las funciones tambien son tipos de datos lo que permite usarlas como parametors o retornos de funcion

func main() {

	//FUNCIONES COMO PARAMETROS
	nums := []int{2, 12, 23, 98, 21, 79}                           //slice de numeros aleatorios
	result := filter(nums, func(num int) bool { return num > 40 }) //se manda a llamar la funcion filter con callback como argumento con el filtro deseado
	fmt.Println(result)

	//FUNCIONES COMO RETORNO DE FUNCIONES
	// se vuelven a llamar los parentesis cada vez que se quieran mandar los parametros de las funciones a retornar
	resultado := sum(2)(3)
	fmt.Println(resultado)

	//otro ejemplo podria ser
	plus10 := sum(10)      //donde plus10 se usa como una especie de memoria cache, aqui vale 10
	fmt.Println(plus10(2)) //aqui vale 12
	fmt.Println(plus10(4)) //aqui se vera 14 en la terminal
	fmt.Println(plus10(1)) //aqui se vera 11 en la terminal
	fmt.Println(plus10(5)) //aqui se vera 15 en la terminal

	// esta es otra manera de reutilizar funciones

}

//FUNCIONES COMO PARAMETROS
//funcion como parametro de otra funcion, replicando la funcion filter de javascript
//la funcion filter recivira como argumento el slice de numeros y una funcion llamada callback

//la funcion callback tiene por parametro un valor de tipo entero y returna un entero

//regresando a la funcion filter retorna un slice de tipo int
//en el bloque de codigo se crea resultado donde se almacenaran los numeros filtrados,para crear el slices resultado se usa make

//al make se le indica que sera un slice tipo entero []int, y al no saber el tamano que tendra, se establece la capacidad maxima
//la cual seria el tamano de el slices que se quiere filtrar en este caso nums usando el comando len de la manera 0,len(nums) se cierra el make

// luego se itera, por el slice que se esta recibiendo usando for := range utilizando la nomenclatura de 4_2_for

//dentro del for se preguntara que si la logica de callback(bool) se cumple, es decir loque sea que este queriendo filtrar
//se agregara el valor iterado de nums a result y results tomara el valor del nuevo

// al terminar de itarar por cada valor dentro de nums, returnara la variable resultado

func filter(nums []int, callback func(int) bool) []int {
	result := make([]int, 0, len(nums))

	for _, num := range nums {
		if callback(num) {
			result = append(result, num)
		}
	}
	return result
}

//FUNCIONES COMO RETORNO DE FUNCIONES
// se crea la funcion suma la cual solo recivira un elemento "a" de tipo entero y retornara una funcion con argumento entero que returnara un entero
// en la parte del retorno de la funcion sum se desarrolla la funcion a desarrollar en este caso, la funcion tendra un argumento "b" tipo enteron y retornara un enterp
//esta funcion de retorno a su vez retornara la sumatoria de a mas b

// este tipo de casos funciona para reutilizar logica

func sum(a int) func(int) int {

	return func(b int) int {
		return a + b
	}

}
