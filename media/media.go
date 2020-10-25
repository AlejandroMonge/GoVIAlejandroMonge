package media

import "fmt"

type Multimedia interface {
	Mostrar()
}

type Imagen struct {
	Titulo  string
	Formato string
	Canales string
}

type Audio struct {
	Titulo   string
	Formato  string
	Duracion string
}

type Video struct {
	Titulo  string
	Formato string
	Frames  string
}

type ContenidoWeb struct {
	Contenido []Multimedia
}

func (cw *ContenidoWeb) Mostrar() {

	for _, c := range cw.Contenido {
		fmt.Println("***************************************************")
		c.Mostrar()
	}

}

func (i *Imagen) Mostrar() {

	fmt.Println("Titulo: ", i.Titulo)
	fmt.Println("Formato: ", i.Formato)
	fmt.Println("Canales: ", i.Canales)

}

func (a *Audio) Mostrar() {

	fmt.Println("Titulo: ", a.Titulo)
	fmt.Println("Formato: ", a.Formato)
	fmt.Println("Duracion: ", a.Duracion)

}

func (v *Video) Mostrar() {

	fmt.Println("Titulo: ", v.Titulo)
	fmt.Println("Formato: ", v.Formato)
	fmt.Println("Frames: ", v.Frames)

}
