package model

type Draw struct {
	Red   int
	Green int
	Blue  int
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
