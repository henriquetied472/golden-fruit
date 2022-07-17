package models

type Object struct {
	X, Y  float64
	Size  int
	Color string
}

func (obj *Object) addCoordinate(coord Coordinate) Coordinate {
	x := obj.X + float64(coord.x)
	y := obj.Y + float64(coord.y)
	return Coordinate{int(x), int(y)}
}

func (obj *Object) areColiding(o Object) bool {
	if obj.X == o.X && obj.Y == o.Y {
		return true
	}
	return false
}
