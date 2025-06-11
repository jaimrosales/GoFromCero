package main

import "fmt"

func main() {
	//declaracion clasica pero sus argumentos no van entre parentesis
	// for [declaracionVariable]; [condicion evaluar];[ejecucion posterior al finalizar ciclo]
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}

	//se puede declarar un for continuo, similar a un while(while no existe en go)
	//de esta forma no es necesario indicar la declaracion de inicio y posterior dentro del for
	// se puede declarar la variable fuera antes del for y la declaracion posterior dentro del bloque de evaluacion

	j := 0

	for j <= 5 {
		fmt.Println(j)
		j++
	}

	//delcaracion for forever
	// es util para hacer que procesos se procecen para siempre, como cuando se quiere usar un web socker o se trabaja con concurrencia
	// se puede utilizar break para salir del for
	k := 0
	for {
		if k == 6 {
			break
		}
		fmt.Println(k)
		k++
	}

	//FOR CON SLICES

	//se puede utilizar para iterar dentro de un slice o array de la siguiente manera
	food := []string{"🍕", "🍔", "🍎", "🌭"}

	for l, m := range food {
		fmt.Println("indice", l, "valor", m)
	}

	//tambien se puede usar un slice literal como sigue

	for n, o := range []string{"💻", "📱", "🧑‍💻", "🔬"} {
		fmt.Println("indice", n, "valor", o)
	}

	numbers := []uint8{2, 4, 6, 8}
	//el indice se omite para solo hacer caso al valor, y q (el valor) es una copia, no es exactamente el valor de numbers
	for _, q := range numbers {
		q *= 2
	}
	fmt.Println(numbers)

	// para editar el valor en el slice se utiliza el indice y en el bloque de codigo se usa el slices con su indice variable
	// go permite omitir en este caso el valor de range ya que trabajaras con los indices
	for i := range numbers {
		numbers[i] *= 2
	}
	fmt.Println(numbers)

	//FOR CON MAPS
	//se puede iterar en mapas
	comida := map[string]string{
		"pizza":     "🍕",
		"hamburger": "🍔",
		"apple":     "🍎",
		"hotdog":    "🌭",
	}

	//en caso de los mapas for range devolvera la llave y el contenido
	for key, value := range comida {
		fmt.Println("key: ", key, "valor: ", value)
	}

	//FOR SOBRE STRINGS
	//la siguiente manera muestra como iterar en un string
	// en el caso siguiente v es un tipo de dato rune, saldra su codigo unicode
	for i, v := range "RSTech" {
		fmt.Println("indice:", i, "value", v)
	}
	// para ver el valor se hace de la siguiente manera
	for i, v := range "RSTech" {
		fmt.Println("indice:", i, "value", string(v))
	}

}
