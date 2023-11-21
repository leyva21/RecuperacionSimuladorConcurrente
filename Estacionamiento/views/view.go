package views

import (
	"Estacionamiento/models"
	"Estacionamiento/scenes"
	"sync"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

// Dibujar el contorno
func Estacionamiento(win *pixelgl.Window) {
	imd := imdraw.New(nil)
	imd.Color = colornames.Red

	p1 := pixel.V(100, 460)
	p2 := pixel.V(100, 20)
	p3 := pixel.V(600, 20)
	p4 := pixel.V(600, 460)
	p5 := pixel.V(100, 300)
	p6 := pixel.V(100, 180)

	imd.Push(p1)
	imd.Push(p5)
	imd.Line(1)

	imd.Push(p6)
	imd.Push(p2)
	imd.Line(1)

	imd.Push(p2)
	imd.Push(p3)
	imd.Line(1)

	imd.Push(p3)
	imd.Push(p4)
	imd.Line(1)

	imd.Push(p4)
	imd.Push(p1)
	imd.Line(1)
	imd.Draw(win)

}

func run() {
	cfg := pixelgl.WindowConfig{
		Bounds: pixel.R(0, 0, 640, 480),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	gateFree := make(chan bool, 1)
	mu := sync.Mutex{}

	p := models.NewParking(20, mu, gateFree)

	gateFree <- true

	i := 1

	for !win.Closed() {

		win.Clear(colornames.White)

		Estacionamiento(win)

		if i < 101 {
			go scenes.Run(i, p, win)

		}

		for _, c := range p.Cars {
			imd := imdraw.New(nil)
			imd.Color = colornames.Red

			imd.Push(c.P1)
			imd.Push(c.P2)
			imd.Line(c.Width)
			imd.Draw(win)
		}

		win.Update()

		i++

		time.Sleep(800 * time.Millisecond)

	}
}

func Show() {
	pixelgl.Run(run)
}
