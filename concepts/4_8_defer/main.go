package main

import (
	"fmt"
	"os" //este paquete permite interactuar con el sistema operativo
)

//La funcion defer nos permite diferir o retrasar o aplasar la ejecucion de una funcion, hasta que la funcion donde fue llamado el defer retorne

func main() {
	defer fmt.Println(4) //se aplasa hasta el momento en que la funcion retorne para el caso en el que main retorne es decir al final
	fmt.Println(1)
	fmt.Println(2)

	defer fmt.Println(3) // el orden de funciones diferidas seran de la ultima a la primera

	// en el siguiente segmento se crea num y se difiere la impresion de num con valor 4, luego se cambia a 20 y se imprime num, como resultado
	//se imprimira 10 y posterior a eso se imprimira 4 debido a que cuando se difirio num tenia un valor de 4,
	//creo que se puede usar para guardas los valores actuales, como es el caso de el tiempo en el que se quedo del controlador de valvulas(un problema menos?)
	num := 4
	defer fmt.Println("imprime ", num)
	num = 10
	fmt.Println(num)

	//en efectivo es util para limpiar recursos, cerrar conecciones de red, o cerrar controladores de base de datos,
	//el puntero del archivo se guarda en file y el error en err, posterior a eso se manejo el error para salir en caso de no haber sido creado
	// posterior a eso se escribe en el archivo usando write, funcion la cual devuelve el numero de bytes y el error
	//debido a que para el caso no interesa la cantidad de bytes se usa _ para ignorar el valor, y err para recibir el error y manejarlo cuando se produsca alguno
	file, err := os.Create("test.txt") // nos permitecrear un archivo en el sistema operativo
	if err != nil {
		fmt.Println(err)
		return
	}
	//con la siguiente expresion difer, se difiere la funcion de cierre, de esta manera no es necesario escribir la funcion de cierre cada que
	//se manejen errores o se finalice el programa, para no dejar el archivo abierto
	defer file.Close()

	_, err = file.Write([]byte("hello RStechers"))
	if err != nil {
		/*file.Close()*/
		fmt.Println(err)
		return
	}
	/*file.Close()*/

}
