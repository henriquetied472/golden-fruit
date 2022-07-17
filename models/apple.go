package models

import (
	"math/rand"
	"time"
)

type Apple struct {
	Object
}

func SpawnApple(w, h int) *Apple {
	rand.Seed(time.Now().UnixNano())
	size := 20
	return &Apple{
		Object: Object{
			Size: size,
			X: float64(rand.Intn(w / size)),
			Y: float64(rand.Intn(h / size)),
			Color: "#ed1313",
		},
	}
}

func (a *Apple) Resurge(w, h int) {
	size := a.Size
	a.Object = Object{
		Size: size,
		X: float64(rand.Intn(w / size)),
		Y: float64(rand.Intn(h / size)),
		Color: "#ed1313",
	}
}