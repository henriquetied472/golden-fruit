package main

import (
	"flag"
	"golden-fruit/models"
	"log"

	"github.com/tfriedel6/canvas/glfwcanvas"
)

const (
	width, height int = 800, 400
)

var p *models.Player
var e *models.Enemy
var a *models.Apple

var sfps bool

func init() {
	flag.BoolVar(&sfps, "fps", false, "Show FPS on the screen")
	p = models.NewRandomP(width, height)
	e = models.NewRandomE(width, height, *p.Object)
	a = models.SpawnApple(width, height)
}

func main() {
	wnd, cv, err := glfwcanvas.CreateWindow(width, height, "Golden Fruit")
	if err != nil {
		log.Fatal(err)
	}
	defer wnd.Window.Destroy()

	wnd.KeyDown = func(scancode int, rn rune, name string) {
		actualDir := models.Dir_Nil
		
		switch name {
		case "ArrowUp":
			actualDir = models.Dir_Up
		case "ArrowDown":
			actualDir = models.Dir_Down
		case "ArrowLeft":
			actualDir = models.Dir_Left
		case "ArrowRight":
			actualDir = models.Dir_Right
		}
		p.Move(actualDir)
	}

	wnd.MainLoop(func() {
		Update(wnd)
		Draw(cv)
	})
}
