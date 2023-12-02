package model

type Draw struct {
	Red   int
	Green int
	Blue  int
}

func (d *Draw) Power() int {
	return d.Red * d.Green * d.Blue
}

type Game struct {
	Id    int
	Draws []Draw
}

type World struct {
	MaxRed   int
	MaxGreen int
	MaxBlue  int
}
