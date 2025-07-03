package main

import (
	"time"

	"github.com/fatih/color"
)

// la siguiente constante servira para indicar si es de sistemas operativos windows
const Windows = "Windows"

// files types
const ( //constantes a utilizar en el edls
	fileRegular int = iota
	fileDirectory
	fileExecutable
	fileCompress
	fileImage
	fileLink
)

// file extension: constantes a utilizar en el edls
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

// la siguiente estructura contendra los diferentes campos de un archivo necesarios para renderizar la informacion
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

// en la siguiente escritura se coontendra el estilos de renderizado, PROPIEDADES DE CADA ESTILO
type styleFileType struct {
	icon   string
	color  color.Attribute
	symbol string
}

// en el siguiente mapa se define cad uno de los estilos haciendolos de styleFileType
var mapStylesByFileType = map[int]styleFileType{
	fileRegular:    {icon: "📄"},
	fileDirectory:  {icon: "📂", color: color.FgBlue, symbol: "/"},
	fileExecutable: {icon: "🚀", color: color.FgGreen, symbol: "*"},
	fileCompress:   {icon: "📦", color: color.FgRed},
	fileImage:      {icon: "📷", color: color.FgMagenta},
	fileLink:       {icon: "🔗", color: color.FgCyan},
}

var (
	blue    = color.New(color.FgBlue).Add(color.Bold).SprintFunc()
	green   = color.New(color.FgGreen).Add(color.Bold).SprintFunc()
	red     = color.New(color.FgRed).Add(color.Bold).SprintFunc()
	magenta = color.New(color.FgMagenta).Add(color.Bold).SprintFunc()
	cyan    = color.New(color.FgCyan).Add(color.Bold).SprintFunc()
	yellow  = color.New(color.FgYellow).SprintFunc()
)
