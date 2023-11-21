package models

import "github.com/faiface/pixel"

type espacio struct {
	p    pixel.Vec
	free bool
}

func newEspacio(p pixel.Vec) espacio {
	return espacio{
		p:    p,
		free: true,
	}
}

func generarEspacio() []espacio {
	var espacios []espacio
	width := 40.0
	space := 10.0

	for x := 105.0; x <= 595; x += width + space {

		xM := (x + (x + width)) / 2
		p := pixel.V(xM, 445)

		espacios = append(espacios, newEspacio(p))
	}

	for x := 105.0; x <= 595; x += width + space {

		xM := (x + (x + width)) / 2
		p := pixel.V(xM, 15+80)

		espacios = append(espacios, newEspacio(p))
	}
	return espacios
}
