package models

import "sync"

type Casilleros struct {
	Capacity int
	Mu       sync.Mutex
	gateFree chan bool
	slots    []espacio
	Cars     []*Carro
}

func NewParking(c int, mu sync.Mutex, gateFree chan bool) *Casilleros {
	return &Casilleros{
		Capacity: c,
		Mu:       mu,
		gateFree: gateFree,
		slots:    generarEspacio(),
		Cars:     make([]*Carro, 0),
	}
}
