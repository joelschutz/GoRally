package models

type TrackState struct {
	Location  uint32
	Terrain   Terrain
	PaceNotes Segments
	// MaxTorque    float64
	// DistanceLeft float64
}

type Track struct {
	Name     string   `json:"name"`
	Country  string   `json:"country"`
	Segments Segments `json:"segments"`
}

type Segments []Segment

func (s Segments) GetNode(index int) Point2D {
	return s[index]
}
func (s Segments) Size() int {
	return len(s)
}
func (s Segments) Looped() bool {
	return false
}
func (s Segments) Width(node int) float64 {
	if node <= len(s)-1 {
		return s[node].Width
	}
	return -1
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

// Corner Levels will be calculated on the fly for by the angle between nodes
type CornerLevel uint8

// const (
// 	Flat CornerLevel = iota
// 	Level6
// 	Level5
// 	Level4
// 	Level3
// 	Level2
// 	Level1
// 	Square
// 	HairPin
// )

// No longer relevant
// type Direction uint8

// const (
// 	Straight Direction = iota
// 	Left
// 	Right
// )

type Segment struct {
	Length float64 `json:"length"`
	Width  float64 `json:"width"`
	// Direction         Direction   `json:"direction"`
	// Corner            CornerLevel `json:"corner"`
	Terrain           Terrain `json:"terrain"`
	x, y              float64
	Cut, Unseen, Jump bool
}

func (bp Segment) X() float64 {
	return bp.x
}
func (bp Segment) Y() float64 {
	return bp.y
}
func (bp Segment) SetX(x float64) {
	bp.x = x
}
func (bp Segment) SetY(y float64) {
	bp.y = y
}
