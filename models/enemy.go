package models

import (
	"math/rand"
	"time"

	"github.com/jpierer/astar"
)

type Enemy struct {
	*Object
	
	Target Object
}

func NewRandomE(w, h int, trgt Object) *Enemy {
	rand.Seed(time.Now().UnixNano())
	size := 20
	return &Enemy{
		Object: &Object{
			Size: size,
			X: float64(rand.Intn(w / size)),
			Y: float64(rand.Intn(h / size)),
			Color: "#ed13ed",
		},
		Target: trgt,
	}
}

func (e *Enemy) RunToTarget(w, h int) {
	grid, _ := astar.New(astar.Config{
		GridWidth: w / e.Size,
		GridHeight: h / e.Size,
	})

	origin := astar.Node{
		X: int(e.X),
		Y: int(e.Y),
	}

	target := astar.Node{
		X: int(e.Target.X),
		Y: int(e.Target.Y),
	}

	nodes, _ := grid.FindPath(origin, target)

	e.X = float64(nodes[len(nodes) -1].X)
	e.Y = float64(nodes[len(nodes) -1].Y)
}

func (e *Enemy) RefreshTargetLocation(trgt Object) {
	e.Target = trgt
}

func (e *Enemy) Resurge(w, h int) {
	size := e.Size
	e.Object = &Object{
		Size: size,
		X: float64(rand.Intn(w / size)),
		Y: float64(rand.Intn(h / size)),
		Color: "#ed13ed",
	}
}