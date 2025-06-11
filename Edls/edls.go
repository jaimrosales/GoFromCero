package main

import "time"

//files types
const ( //constantes a utilizar en el edls
	fileRegular int = iota
	fileDirectory
	fileExecutable
	fileCompress
	fileImage
	fileLink
)

//file extension: constantes a utilizar en el edls
const (
	exe = ".exe"
	deb = ".deb"
	zip = ".zip"
	gz  = ".gz"
	tar = ".tar"
	rar = ".rar"
	png = ".png"
	jpg = ".jpg"
	gif = ".gif"
)

//la siguiente estructura contendra los diferentes campos de un archivo necesarios para renderizar la informacion
type file struct {
	name             string
	fileType         int
	isDir            bool
	isHidden         bool
	userName         string
	groupname        string
	size             int64
	modificationTime time.Time
	mode             string
}

//en la siguiente escritura se coontendra el estilos de renderizado
type styleFileType struct {
	incon  string
	color  string
	symbol string
}

var mapStylesByFileType = map[int]styleFileType{
	fileRegular:    {},
	fileDirectory:  {},
	fileExecutable: {},
	fileCompress:   {},
	fileImage:      {},
	fileLink:       {},
}
