package models

import "math"

type Track struct {
	Name     string
	Segments []Segmnent
}

type Terrain uint8

const (
	Tarmac Terrain = 1 + iota
	Mud
	HeavyGravel
	LightGravel
	Sand
	Snow
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
	Length            float64
	Direction         Direction
	Corner            CornerLevel
	Terrain           Terrain
	Cut, Unseen, Jump bool
}

func (s *Segmnent) IsLengthValid() bool {
	teoricalPerimeter := 2 * math.Pi * 100 / float64(s.Corner)
	return (s.Length <= teoricalPerimeter/2)
}
