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
	"regexp" //paquete de expresiones regulares
	"runtime"
	"sort"
	"strings"
	"time"

	"golang.org/x/exp/constraints"
)

func main() {

	//filter flags   	flags de filtrado de salida
	flagPattern := flag.String("p", "", "filter by pattern")           // nos permite seleccionar el patron de archivo
	flagAll := flag.Bool("a", false, "all files including hide files") // nos mostrara todos los archivos
	flagNumberRecords := flag.Int("n", 0, "number of archivos")        // indica el numero maximo de archivos que se desean revisar en el filtrado

	//order flags  		flags que permitiran organizar la salida
	hasOrderbytime := flag.Bool("t", false, "sort by time, oldest first")
	hasOrderbySize := flag.Bool("s", false, "sort by file size, smallest first")
	hasOrderReverse := flag.Bool("r", false, "reverse order while sorting")

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
	//Se itera por cada uno de los archivos dentro del directorio para guardarlos en fs
	for _, dir := range dirs {
		isHidden := isHidden(dir.Name(), path)

		if isHidden && !*flagAll {
			continue
		}

		//antes de agregar el archivo a fs se filtra, y se filtrara solo cuando se escriba algo para filtrar
		if *flagPattern != "" {
			isMatched, err := regexp.MatchString("(?i)"+*flagPattern, dir.Name()) // la funcion matchstring permite evaluar si la primera expresion se encuentra en la segunda, "(?i)" nos sirve para indicar que es insensible a mayus concatenado
			if err != nil {
				panic(err)
			}
			if !isMatched {
				continue // continue, lo que hace es que saltara a la siguiente iteracion dentro del for, por lo que no llegara a la siguiente linea de codigo
			}
		}

		f, err := getFile(dir, isHidden)
		if err != nil { //se controla el error por medio de un panico
			panic(err)
		}

		fs = append(fs, f) //slice de archivos a mostrar

	}

	//ordenamiento alfabetico, default
	if !*hasOrderbySize || !*hasOrderbytime {
		orderByName(fs, *hasOrderReverse)
	}

	//Solo se podra ordenar por tamano en caso de que no se ordene por fecha
	if *hasOrderbySize && !*hasOrderbytime {
		orderBySize(fs, *hasOrderReverse)
	}
	//Solo se podra ordenar por fecha en caso de que no se ordene por tamano
	if *hasOrderbytime && !*hasOrderbySize {
		orderByTime(fs, *hasOrderReverse)
	}

	//para filtrar cuantos registros se quieren imprimir, justo antes de imprimir se setea el flag
	if *flagNumberRecords == 0 || *flagNumberRecords > len(fs) { //si el flag nunca se especifico o si se especifico un tamano mayor a la cantidad de registros dentro del directorio, el valor a mostrar sera la longitud de registros dentro del directorio
		*flagNumberRecords = len(fs)
		//Se pasa este valor a la funcion print list, para que tome en cuenta este valor como limite a iterar en la impresion
	}

	//IMPRESIONES

	//fmt.Println(fs)                   //imprime el slice de archivos
	printList(fs, *flagNumberRecords) //Imprime el Slice de archivos con el formato deseado
	//Flags
	fmt.Println("patron de filtro ", *flagPattern)
	fmt.Println("all", *flagAll)
	fmt.Println("flagNumberRecords: ", *flagNumberRecords)

}

func mySort[T constraints.Ordered](i, j T, isReverse bool) bool {
	if isReverse {
		return i > j
	}
	return i < j
}

func orderByName(files []file, isReverse bool) {
	sort.SliceStable(files, func(i int, j int) bool {
		return mySort(
			strings.ToLower(files[i].name),
			strings.ToLower(files[j].name),
			isReverse)

	})
}

func orderBySize(files []file, isReverse bool) {
	sort.SliceStable(files, func(i int, j int) bool {
		return mySort(
			files[i].size,
			files[j].size,
			isReverse)
	})
}

func orderByTime(files []file, isReverse bool) {
	sort.SliceStable(files, func(i int, j int) bool {
		return mySort(
			files[i].modificationTime.Unix(),
			files[j].modificationTime.Unix(),
			isReverse)

	})
}

// La siguiente es la funcion de salida
// esta funcion se encargara de imprimir cada estructura dentro del slice y utilisando la funcion print format se le dara formato a la impresion
// recive el slices de archivos y la cantidad de iteraciones deseadas
func printList(fs []file, nRecords int) {
	for _, file := range fs[0:nRecords] { // dentro de file guardara cada valor fs desde su posicion 0 hasta la cantidad especificada de registros deseados
		style := mapStylesByFileType[file.fileType] // se manda a llamar el mapa donde se tienen los datos de icono, color y simbolo, relacionados con cada tipo de archivo, recordando que cada tipo de archivo tiene un int asignado, y se le asigna a style

		fmt.Printf("%s  %s  %s  %10d  %s  %s %s %s \n", file.mode, file.userName, file.groupname, file.size, file.modificationTime.Format(time.DateTime), style.icon, file.name, style.symbol) // dentro del verbo, entre % y la letra indicadora de tipo, si se pone un numero se indica el espacio que se quiere usar para el verbo para dar un sentido de columnas
		// al file modification tiene muchos datos que no son realmente necesarios para la aplicacion, para no imprimir todo se usa la funcion format del paquete time y se le da el formato datetime
	}
}

// La siguiente funcion es para la organizacion y artructuracion de las archivos dentro del directorio
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
	switch {
	case isLink(*f):
		f.fileType = fileLink
	case f.isDir:
		f.fileType = fileDirectory
	case isExec(*f):
		f.fileType = fileExecutable
	case isCompress(*f):
		f.fileType = fileCompress
	case isImage(*f):
		f.fileType = fileImage
	default:
		f.fileType = fileRegular

	}
}

// funciones auxiliares a setFile
// esta funcion evalua si el archivo a evaluar es un enlace
func isLink(f file) bool {
	//se verifica el modo del archivo evaluando el primer elemento del mismo, si es igual a L es un enlace,
	return strings.HasPrefix(strings.ToUpper(f.mode), "L") //esta funcion de strings permite verificar el prefijo(con que comienza) del strings que se le pasa, se convierte en mayuscula para que no haya diferencia por versiones

}

func isExec(f file) bool {
	if runtime.GOOS == Windows { //compara si el sistema operativo es windows, runtime.GOOS regresa el sistema operativo
		return strings.HasSuffix(f.name, exe) //compara si la extension del archivo el sufijo, es .exe para mandar el boleano
	}
	return strings.Contains(f.mode, "x") ///con contains verifica si el string contiene el string especificado
}

// compara si el sufijo del nombre del archivo para identificar que tipo de archivo es y si es comprimido
func isCompress(f file) bool {
	return strings.HasSuffix(f.name, zip) ||
		strings.HasSuffix(f.name, gz) || //archivos comprimidos de unix
		strings.HasSuffix(f.name, rar) ||
		strings.HasSuffix(f.name, deb) //archivos comprimidos de linux, ubuntu o debian
}

// compara si el sufijo del nombre del archivo para identificar que tipo de archivo es y si es un tipo imagen
func isImage(f file) bool {
	return strings.HasSuffix(f.name, png) ||
		strings.HasSuffix(f.name, jpg) ||
		strings.HasSuffix(f.name, gif)
}

func isHidden(fileName, basePath string) bool {
	return strings.HasPrefix(fileName, ".")
}
