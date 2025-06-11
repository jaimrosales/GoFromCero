package main

import "fmt"

func main() {
	//len(): # de elementos en el slice, el tamano del slice
	//cap(): # de elementos del array origen, a partir del indice donde se creo el slice, indica la capidad maxima que puede tener el slice

	animals := [5]string{":gorilla", ":dog", ":cat", ":bird", "elephant"}
	pets := animals[1:3]            //":dog",":cat"
	pets = append(pets, ":sheppar") // se agrega sheppar al slices pero se modifica el valor de animals en lan posicion 3

	//si agregas mas espacios por delante de elephant
	//para solucionar esto go crea un nuevo array, con referencia totalmente diferente al array original

	fmt.Println("animals: ", animals)
	fmt.Println("pets: ", pets)
	fmt.Println("tamano pets: ", len(pets))
	fmt.Println("capacidad pets: ", cap(pets))

	//tambien se puede crear un slices sin specificar el tamano
	mascotas := []string{":dog", ":cat"}
	fmt.Println("mascotas: ", mascotas)
	fmt.Println("tamano mascotas: ", len(mascotas))
	fmt.Println("capacidad mascotas: ", cap(mascotas))

	//generara un slice con una capacidad de 3 pero sin tamano
	maskotas := make([]string, 0, 3)
	fmt.Println("maskotas: ", maskotas)
	fmt.Println("tamano maskotas: ", len(maskotas))
	fmt.Println("capacidad maskotas: ", cap(maskotas))

	// los slices tienen valor cero, es nil el cual es el equivalente a null o nulo
}
