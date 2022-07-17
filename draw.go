package main

import (
	"fmt"

	"github.com/tfriedel6/canvas"
)

func Draw(cv *canvas.Canvas) {
	//background
	cv.SetFillStyle("#12345d")
	cv.FillRect(0, 0, float64(width), float64(height))

	//player
	cv.SetFillStyle(p.Color)
	cv.FillRect(p.X * float64(p.Size), p.Y * float64(p.Size), float64(p.Size), float64(p.Size))

	//enemy
	cv.SetFillStyle(e.Color)
	cv.FillRect(e.X * float64(e.Size), e.Y * float64(e.Size), float64(e.Size), float64(e.Size))

	//apple
	cv.SetFillStyle(a.Color)
	cv.FillRect(a.X * float64(a.Size), a.Y * float64(a.Size), float64(a.Size), float64(a.Size))

	//score
	cv.SetFillStyle("#ffffff")
	cv.SetFont("~/.goldenfruit/Roboto-Regular.ttf", 15)
	cv.FillText(fmt.Sprint("Score: ", p.Score), 20, 20)

	//fps
	if sfps {
		cv.FillText(fmt.Sprint(fps, " FPS"), float64(width) - 20, 20)
	}
}