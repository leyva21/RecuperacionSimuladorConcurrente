package models

import (
	"math/rand"
	"github.com/faiface/pixel"
)

type Carro struct {
	ID    int
	P1    pixel.Vec
	P2    pixel.Vec
	Width float64
	T     int
}

func NewCarro(id int, p1, p2 pixel.Vec) *Carro {
	return &Carro{
		ID:    id,
		P1:    p1,
		P2:    p2,
		Width: 30,
		T:     rand.Intn(10) + 5,
	}
}

func Estacionar(p *Casilleros, c *Carro) int {
	var freeNum int
	<-p.gateFree

	for i := range p.slots {
		if p.slots[i].free {
			c.P1 = p.slots[i].p
			c.P2 = pixel.V(c.P1.X, c.P1.Y-60)
			freeNum = i
			break
		}
	}
	p.slots[freeNum].free = false
	p.gateFree <- true

	return freeNum
}

func Salida(i int, p *Casilleros, c *Carro) {
	<-p.gateFree
	p.slots[i].free = true
	p.Capacity++
	p.gateFree <- true
	c.P1 = pixel.V(40, 360)
	c.P2 = pixel.V(40, 300)
}

func Irse(c *Carro) {
	c.P1 = pixel.V(-30, 0)
	c.P2 = pixel.V(-30, 0)
	c.Width = 0
}
