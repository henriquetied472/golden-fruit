package models

import (
	"golden-fruit/security"
	"math/rand"
	"time"
)

type Player struct {
	*Object
	Metadata struct {
		width, height int
	}
	
	Score int
}

type Coordinate struct {
	x, y int
}

var (
	dir_up = Coordinate{0, -1}
	dir_down = Coordinate{0, 1}
	dir_right = Coordinate{1, 0}
	dir_left = Coordinate{-1, 0}
	dir_nil = Coordinate{0, 0}
	Dir_Up = dir_up
	Dir_Down = dir_down
	Dir_Right = dir_right
	Dir_Left = dir_left
	Dir_Nil = dir_nil
)

func NewRandomP(w, h int) *Player {
	rand.Seed(time.Now().UnixNano())
	size := 20
	return &Player{
		Object: &Object{
			Size: size,
			X: float64(rand.Intn(w / size)),
			Y: float64(rand.Intn(h / size)),
			Color: "#13ed13",
		},
		Metadata: struct{width, height int}{
			width: w,
			height: h,
		},
		Score: 0,
	}
}

func (p *Player) Move(direction Coordinate) {
	switch direction {
	case dir_up:
		tmp := p.Object.addCoordinate(dir_up)
		if !(tmp.y < 0) {
			p.X = float64(tmp.x)
			p.Y = float64(tmp.y)
		}
	case dir_down:
		tmp := p.Object.addCoordinate(dir_down)
		if !(float64(tmp.y) > float64((p.Metadata.height / p.Size) - 1)) {
			p.X = float64(tmp.x)
			p.Y = float64(tmp.y)
		}
	case dir_right:
		tmp := p.Object.addCoordinate(dir_right)
		if !(float64(tmp.x) > float64((p.Metadata.width / p.Size) - 1)) {
			p.X = float64(tmp.x)
			p.Y = float64(tmp.y)
		}
	case dir_left:
		tmp := p.Object.addCoordinate(dir_left)
		if !(tmp.x < 0) {
			p.X = float64(tmp.x)
			p.Y = float64(tmp.y)
		}
	case dir_nil:
		tmp := p.Object.addCoordinate(dir_nil)
		p.X = float64(tmp.x)
		p.Y = float64(tmp.y)
	default:
		security.AlertAdulteration("Speed incresed anormaly")
	}
}

func (p *Player) UpdateColision(a *Apple, e *Enemy) {
	if p.areColiding(a.Object) {
		p.Score++
		a.Resurge(p.Metadata.width, p.Metadata.height)
	} else if p.areColiding(*e.Object) {
		if p.Score != 0{
			p.Score--
			e.Resurge(p.Metadata.width, p.Metadata.height)
		}
	}
}

func (p *Player) UpdateWallColision() {
	if p.X < 0 {
		p.X = 0
	} else if p.X > float64(p.Metadata.width - p.Size) {
		p.X = float64(p.Metadata.width - p.Size) 
	}

	if p.Y < 0 {
		p.Y = 0
	} else if p.Y > float64(p.Metadata.height - p.Size) {
		p.Y = float64(p.Metadata.height - p.Size) 
	}
}

func (p *Player) GetWidth() int { return p.Metadata.width }
func (p *Player) GetHeight() int { return p.Metadata.height }