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
	"runtime"
	"strings"
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

// en la siguiente funcion recive un directorio, y un bool, y apoyandose del paquete dir, y la funcion info, obtiene los datos de la direccion(archivo)que
// se le ingresa, se llena el struct de file, la clase del archivo en poo
func getFile(dir fs.DirEntry, isHidden bool) (file, error) {
	info, err := dir.Info()
	if err != nil {
		return file{}, fmt.Errorf("dir.Info(): %v", err)
	}
	f := file{ //se llena la variable f con todos los datos que tiene cada archivo file con los datos de la dir, usando los paquetes dir e info
		name:             dir.Name(),
		isDir:            dir.IsDir(),
		isHidden:         isHidden,
		userName:         "jaimrosales", //se recivira del paquete final
		groupname:        "RStech",      //se recivira del paquete final
		size:             info.Size(),
		modificationTime: info.ModTime(),
		mode:             info.Mode().String(), // hace referencia a los permisos que tiene el usuario, ademas de si es un directorio o un link
	}
	setFile(&f) //se llama a la funcion set file y se le da acceso, con el operador de direccion a la direccion f, para que pueda modificar la struct de f

	return f, nil //se retuorna el archivo con todas sus propiedades
}

// la siguiente funcion agrega a la structura file, el tipo de archivo del directorio evaluado
func setFile(f *file) {

}

//funciones auxiliares a setFile

func isLink(f file) bool { //esta funcion evalua si el archivo a evaluar es un enlace
	//se verifica el modo del archivo evaluando el primer elemento del mismo, si es igual a L es un enlace,
	return strings.HasPrefix(strings.ToUpper(f.mode), "L") //esta funcion de strings permite verificar el prefijo(con que comienza) del strings que se le pasa, se convierte en mayuscula para que no haya diferencia por versiones

}
func isExec(f file) bool {
	if runtime.GOOS == Windows { //compara si el sistema operativo es windows, runtime.GOOS regresa el sistema operativo
		return strings.HasSuffix(f.name, exe) //compara si la extension del archivo el sufijo, es .exe para mandar el boleano
	}
	return strings.Contains(f.mode, "x") ///con contains verifica si el string contiene el string especificado
}
