package main

import (
	"github.com/tfriedel6/canvas/glfwcanvas"
)

var tick int = 1
var fps int

func Update(wnd *glfwcanvas.Window) {
	//player colision
	if tick == 20 {
		e.RefreshTargetLocation(*p.Object)
		e.RunToTarget(p.GetWidth(), p.GetHeight())
		tick = 1
	} else {
		tick++
	}

	p.UpdateColision(a, e)
	fps = int(wnd.FPS())
}