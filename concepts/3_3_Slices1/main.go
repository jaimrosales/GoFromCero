package main

import "fmt"

// Slices: son apuntadores a un array, no poseen datos

func main() {
	things := [7]string{":police_car", ":blue_car", ":yellow_car", ":race_car", ":fire_truck", ":blue_balloon", ":red_ballon"}

	cars := things[0:4] // al declarar y definir un slice es neceesario indicar el rango de indices del array al cual se quiere apuntar [ini:final]
	//tomando en cuenta que el rango es representativo a [) no incluye al valor final
	//si quieres tomar en cuenta desde el principio usas [:n] y si quieres tomar desde un numero hasta el final [n:]

	fmt.Println("things: ", things)
	fmt.Println("cars: ", cars) //":police_car", ":blue_car", ":yellow_car", ":race_car", ":fire_truck"

	ballon := things[5:7]

	fmt.Println("globos: ", ballon)

	fmt.Println("ballon[0]: ", ballon[0])

	cars[1] = "ambulance" // no solo cambias cars, si no tambien cambias things debido a que el slice apunta a un a direccion en particular

	fmt.Println("things: ", things)
	fmt.Println("cars: ", cars)

}
