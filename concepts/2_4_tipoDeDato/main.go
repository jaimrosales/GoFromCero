package main

import "fmt" // cuando se quieran saber las funciones de un paquete de go puedes buscarlo en go.dev /doc /stantard packege

// verbs: los verbos son indicadores de posicion

func main() {
	//bool, string,numeric
	var a bool = true
	var b string = "RSTech"
	var c byte = 255
	var d rune = 'a' //rune permite usar el codigo unicode
	var e float32 = 123.45
	var f float64 = 12439.199342

	// no es posible realizar operaciones entre distintos tipos de datos, debido a que go es un lenguaje estaticamente tipado
	// para hacerlo se puede hacer algo llamado casting

	var g uint16 = 2550
	h := uint16(c) + g

	_ = uint32(c) //el operador _ permite tener una variable sin usarla

	// printf, en el primer argumento se selecciona el formato y porsterior las variables a las que se les quiere dar formato
	fmt.Printf("Tipo: %T, Valor: %v,Tipo: %T, Valor: %v,Tipo: %T, Valor: %v,Tipo: %T, Valor: %v\n", a, a, b, b, c, c, d, d) //%T imprimira el tipo de dato de la variable y %v el valor por defecto de la variable:
	fmt.Printf(",Tipo: %T, Valor: %v,Tipo: %T, Valor: %v,Tipo: %T, Valor: %v", e, e, f, f, h, h)

	// el valor 0 de string es un valor vacio  ""
	// valor 0 de cualquier numerico es 0, 0.0
	// valor 0 de bolean es false

}
