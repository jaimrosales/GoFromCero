package main

//flags: nos permiten indicar alguna accion o filtro dentro del comando como --all --help --decorate son ejemplos, para poder escribirlos desde
//		desde la pantalla usamos el flag.parse que nos permitira ingresar el parse

//argument: indican la direccion, por lo general se ingresan despues de los flags, son los documents/blebleble/user/ ble/ pictures,
//			para poder interactuar con ellos en el codigo se utiliza el flag.arg despues del flag.parse

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
)

func main() {

	//filter flags   	flags de filtrado de salida
	flagPattern := flag.String("p", "", "filter by pattern")           // nos permite seleccionar el patron de archivo
	flagAll := flag.Bool("a", false, "all files including hide files") // nos mostrara todos los archivos
	flagNumberRecords := flag.Int("n", 0, "number of archivos")        // indica el numero maximo de archivos que se desean revisar en el filtrado

	//order flags  		flags que permitiran organizar la salida
	//hasOrderbytime := flag.Bool("t", false, "sort by time, oldest first")
	//hasOrderbySize := flag.Bool("s", false, "sort by file size, smallest first")
	//hasOrderReverse := flag.Bool("r", false, "reverse order while sorting")

	flag.Parse()        //un flag es con el que podremos meter comandos -h o --h esta parte nos permite ver mas que la direccion de memoria
	path := flag.Arg(0) // flag.arg nos regresa ela rgumento que se ah ingresa, en este programa se le asigna el string a la variable path, nos regresara el primer argumento la ruta
	if path == "" {
		path = "."
	}
	dirs, err := os.ReadDir(path) //Se lee el directorio enviado por el usuario, y se asigna a dirs, en caso de que haya un error se asigna a err
	if err != nil {               //se controla el error por medio de un panico
		panic(err)
	}

	fs := []file{} //este slice es donde se guardaran todos los archivos dentro de la direccion que se mostraran al llamarse la funcion edls
	for _, dir := range dirs {
		f, err := getFile(dir, false)
		if err != nil { //se controla el error por medio de un panico
			panic(err)
		}

		fs = append(fs, f)

	}

	fmt.Println(fs)
	fmt.Println("patron de filtro ", *flagPattern)
	fmt.Println("all", *flagAll)
	fmt.Println("flagNumberRecords: ", *flagNumberRecords)

}
func getFile(dir fs.DirEntry, isHidden bool) (file, error) {

	return file{}, nil

}
