package main

// los parametros de tipo permiten generar restricciones para trabajar con multiples tipos de datos ya sean funciones o tipos como estructuras
import "fmt"

//ve a tipos declarativos para entender la siguiente declaracion
type myintV2 int
type myintV3 int

var num1 myint = 2
var num2 myintV2 = 3
var num3 myintV3 = 4

func main() {
	fmt.Println(sum(2, 5, 7))
	fmt.Println(suma[float64](2.7, 5.4, 7.1))                      //al llamar una funcion con restricciones se especifica el tipo de dato
	fmt.Println(agregacion[myintV2]( /*num1,*/ num2 /*num3,*/, 8)) //solo correra un solo tipo de dato y su aproximacion indicada
	fmt.Println(agregacion(num1 /* num2 ,*/ /*num3,*/, 8))         // en caso de que no se agregue el tipo de constrain, go lo inferira del tipo del primer elemento
}

//funcion de ejemplo
func sum(nums ...int) int {
	var total int
	for _, num := range nums {
		total += num

	}
	return total
}

// para el siguiente caso no any no tendra uso debido a que el Tipo de dato any, no acepta operadores aritmeticos debido a que por ejemplo no se pueden sumar bool con strings
/*func suma(nums ...any) any {
	var total any
	for _, num := range nums {
		total += num
	}
	return total
}
*/
//para resolver lo anterior se hara uso de parametros de tipo es decir se restringiran los tipos para no aceptar bool o strings
//los parametros de tipo se denotan entre corchetes justo entre el nombre de la funcion y antes de los argumentos, pueden  ser usados
//en los parametros de la funcion como en el cuerpo o bloque de codigo de la funcion, cada parametro de tipo que se defina tiene que tener un tipo
// de restriccion definida, los parametros se definen con un nombre y luego se define el constrains
//por convencion se definen con letra mayuscula empezando por T, la t se transforma en el nuevo "tipo de dato a trabajar", es decir los tipos de datos aceptables
// que puede tomar

// existen tres tipos de constrains o
//1:constrains de tipo arbitrario
//2:constrains de tipo union de elementos
//3:constrains de tipo aproximacion de elementos

//el tipo arbitrario permite indicar con que tipo de elementos se va a trabajar
/*func suma[T int](nums ...T) T {
	var total T
	for _, num := range nums {
		total += num
	}
	return total
}*/

// El constrain de tipo union de elementos permite usar diferentes tipos de elementos o datos
// para indicarlo se separa cada tipo por un |
// cuando se manda a llamar la funcion se indicara la restriccion a utilizar dentro de los tipos disponibles dentro de la definicion
func suma[T int | float64](nums ...T) T {
	var total T
	for _, num := range nums {
		total += num
	}
	return total
}

//tipos declarativos
//los tipos declarativos son tipos de datos son datos creados por el programador con caracteristicas de los tipos de datos existentes
type myint int //se crea un nuevo tipo de dato de nombre"myint" a partir de un int
//sin embargo esto no significa que se pueden usar directamente como si fueran int al realizar una suma o si se quieren usar por ejemplo en la funcion suma pues no son int, son myint
// se puede usar si se agrega de nuevo con |,
//sin embargo si se crean numerosos tipos de datos con base en int, se tendria que agregar | enesima veces

//Tipo aproximacion de elementos
// para no tener que hacer eso se usa el constrains de tipo aproximacion de elementos, lo que permitira usar un tipo de dato en el caso int
// y  los tipos de datos derivados de int
//para indicar eso se usa ~ antes del tipo de dato, indicando aproximacion uy no igualdad

func agregacion[T ~int | float64](nums ...T) T {
	var total T
	for _, num := range nums {
		total += num
	}
	return total
}
