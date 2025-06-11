package main

import "fmt"

const love = "Yaneli" // las constantes usualmente en go se ponen a nivel de packete y no a nivel de funcion para que sea pocible usarla en mas lugares

const ( // una buena practica es generar una seccion de constantes donde anides tus constantes y tener un codigo mas limpio
	siblings                    = 3
	hermano, hermanito, hermana = "ricardo", "Emiliano", "hannia"
)

const ( // lo que se esta usando es el indentificador iota para definir las variables dentro de un bloque de constantes
	Jan = iota + 1 //iota siempre valdra 0 por lo que al sumarle 1 el primer valor agregado sera 1
	Feb            //2
	Mar            //3
	Apr            //4
	May
	jun // todas las constantes tomaran un calor como si usaras un for para asignarles numeracion
)

func main() {
	const os string = "Linux" // Se declara con la palabra const, seguida del nombre el tipo y es necesario que se le asigne un valor, no se puede aseignar despues
	//	os = "linux" no permitira la asignacion

	const domain, engineer string = "RS.tech", "Jaime Rosales" // puedes aser mas de una asignacion por linea

	const place = "ixtlan del rio" //no se puede usar el operador de variable corta, sin embargo se puede inferir el tipo al remover el tipo de dato

	fmt.Println(os, domain, engineer, place, love, siblings, hermano, hermanito, hermana, May)

}

//las constantes no necesitan ser usadas para compilar un programa
