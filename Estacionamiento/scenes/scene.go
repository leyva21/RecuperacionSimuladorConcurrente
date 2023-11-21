package scenes

import (
	"Estacionamiento/models"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"time"
)

func carRoutine(c *models.Carro, p *models.Casilleros, win *pixelgl.Window) {
	p.Mu.Lock()
	p.Cars = append(p.Cars, c)

	for _, c := range p.Cars {
		imd := imdraw.New(nil)
		imd.Color = colornames.Red

		imd.Push(c.P1)
		imd.Push(c.P2)
		imd.Line(c.Width)
		imd.Draw(win)
	}

	win.Update()

	if p.Capacity > 0 {
		p.Capacity--
		i := models.Estacionar(p, c)
		p.Mu.Unlock()

		time.Sleep(time.Duration(c.T) * time.Second)

		models.Salida(i, p, c)

		time.Sleep(500 * time.Millisecond)

		models.Irse(c)
	} else {
		models.Irse(c)
		p.Mu.Unlock()
	}
}

func Run(i int, p *models.Casilleros, win *pixelgl.Window) {
	c := models.NewCarro(i, pixel.V(40, 180), pixel.V(40, 120))
	go carRoutine(c, p, win)
}
