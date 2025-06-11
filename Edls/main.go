package main

import (
	"flag"
	"fmt"
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

	flag.Parse() //un flag es con el que podremos meter comandos -h o --h esta parte nos permite ver mas que la direccion de memoria

	fmt.Println("patron de filtro ", *flagPattern)
	fmt.Println("all", *flagAll)
	fmt.Println("flagNumberRecords: ", *flagNumberRecords)
	fmt.Println("hasOrderbytime: ", *hasOrderbytime)
	fmt.Println("hasOrderbySize: ", *hasOrderbySize)
	fmt.Println("hasOrderReverse: ", *hasOrderReverse)
}
