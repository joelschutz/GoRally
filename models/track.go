package models

type Track struct {
	Name     string
	Segments []Segmnent
}

type Terrain uint8

const (
	Tarmac Terrain = iota
	Snow
	LightGravel
	HeavyGravel
)

type CornerLevel uint8

const (
	Flat CornerLevel = iota
	Level6
	Level5
	Level4
	Level3
	Level2
	Level1
	Square
	HairPin
)

type Direction uint8

const (
	Straight Direction = iota
	Left
	Right
)

type Segmnent struct {
	Length            uint32
	Direction         Direction
	Corner            CornerLevel
	Terrain           Terrain
	Cut, Unseen, Jump bool
}
