package main

import (
	"fmt"

	"./media"
)

func main() {

	cw := media.ContenidoWeb{}
	opc := ""
	titulo := ""
	formato := ""
	canales := ""
	duracion := ""
	frames := ""

	for true {
		fmt.Println("Agregar Imagen ......... 1")
		fmt.Println("Agregar Audio .......... 2")
		fmt.Println("Agregar Video .......... 3")
		fmt.Println("Mostrar Todo ........... 4")
		fmt.Println("Salir .................. 0")
		fmt.Print("Elige una opcion: ")
		fmt.Scan(&opc)
		if opc == "1" {

			fmt.Print("\nTitulo: ")
			fmt.Scan(&titulo)
			fmt.Print("\nFormato: ")
			fmt.Scan(&formato)
			fmt.Print("\nCanales: ")
			fmt.Scan(&canales)

			img := media.Imagen{titulo, formato, canales}

			cw.Contenido = append(cw.Contenido, &img)

		} else if opc == "2" {
			fmt.Print("\nTitulo: ")
			fmt.Scan(&titulo)
			fmt.Print("\nFormato: ")
			fmt.Scan(&formato)
			fmt.Print("\nDuracion: ")
			fmt.Scan(&duracion)

			vdo := media.Video{titulo, formato, duracion}

			cw.Contenido = append(cw.Contenido, &vdo)
		} else if opc == "3" {
			fmt.Print("\nTitulo: ")
			fmt.Scan(&titulo)
			fmt.Print("\nFormato: ")
			fmt.Scan(&formato)
			fmt.Print("\nDuracion: ")
			fmt.Scan(&frames)

			vdo := media.Audio{titulo, formato, frames}

			cw.Contenido = append(cw.Contenido, &vdo)
		} else if opc == "4" {
			cw.Mostrar()
		} else if opc == "0" {
			break
		}
	}
}
