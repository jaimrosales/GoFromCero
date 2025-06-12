package main

import "fmt"

//las funciones tambien son tipos de datos lo que permite usarlas como parametors o retornos de funcion

func main() {
	nums := []int{2, 12, 23, 98, 21, 79}                           //slice de numeros aleatorios
	result := filter(nums, func(num int) bool { return num > 10 }) //se manda a llamar la funcion filter con callback como argumento con el filtro deseado
	fmt.Println(result)
}

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
