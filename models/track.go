package models

type TrackState struct {
	Location     uint32
	PaceNotes    []Segmnent
	MaxSpeed     float64
	MaxTorque    float64
	DistanceLeft float64
}

type Track struct {
	Name     string     `json:"name"`
	Country  string     `json:"country"`
	Segments []Segmnent `json:"segments"`
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
	Length            float64     `json:"length"`
	Direction         Direction   `json:"direction"`
	Corner            CornerLevel `json:"corner"`
	Terrain           Terrain     `json:"terrain"`
	Cut, Unseen, Jump bool
}
