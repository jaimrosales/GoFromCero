package main

import "fmt"

//los mapas son estructuras de clave valor
// para crear un mapa se usa
//nombre := make(tipo de dato map[tipoDatoDeLlave]tipoDato de alamacenamiento)

func main() {
	music := make(map[string]string) //declaracion
	//Se asignan valores al mapa
	music["guitar"] = ":guitar"
	music["violin"] = ":violin"

	fmt.Println(music)

	tech := map[string]string{ //se declara y hacigna el mapa
		"computer": ":computer",
		"mouse":    ":mouse",
	}

	fmt.Println(tech)

	delete(tech, "computer")
	fmt.Println(tech)

	fmt.Println(music["guitarra"]) ///obtenemos el valor de la llave

	// cuando el mapa no encuentre el valor en el mismo mapa defvolvera el valor cero del tipo de dato
}
