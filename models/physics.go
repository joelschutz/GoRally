package models

// Graph Interface
type Graph2D interface {
	GetNodes() []GraphNode
}
type GraphNode interface {
	Links() []GraphLink
}
type GraphLink interface {
	To() GraphNode
	ToIndex() int
}

type BaseGraph2D struct {
	Nodes []GraphNode
}

func (bg BaseGraph2D) GetNodes() []GraphNode {
	return bg.Nodes
}

// Point Interfaces
type Point2D interface {
	X() float64
	Y() float64
	SetX(x float64)
	SetY(y float64)
}

type Point3D interface {
	Point2D
	Z() float64
	SetZ(z float64)
}

// Generic Point Implementation
type BasicPoint2D struct {
	x, y float64
}

func (bp BasicPoint2D) X() float64 {
	return bp.x
}
func (bp BasicPoint2D) Y() float64 {
	return bp.y
}
func (bp BasicPoint2D) SetX(x float64) {
	bp.x = x
}
func (bp BasicPoint2D) SetY(y float64) {
	bp.y = y
}

func NewBasicPoint2D(x, y float64) Point2D {
	return BasicPoint2D{x: x, y: y}
}

// Shape types from Point2D
type Square2D [4]Point2D
type Traingle2D [3]Point2D

func TriangleArea(t Traingle2D) float64 {
	return ((t[0].X() * (t[1].Y() - t[2].Y())) + (t[1].X() * (t[2].Y() - t[0].Y())) + (t[2].X() * (t[0].Y() - t[1].Y()))) / 2
}

type Line2D [2]Point2D

// Spline Interface
type Spline2D interface {
	GetNode(index int) Point2D
	Size() int
	Looped() bool
}

type BasicSpline2D struct {
	nodes  []Point2D
	looped bool
}

func (bs BasicSpline2D) GetNode(index int) Point2D {
	return bs.nodes[index]
}
func (bs BasicSpline2D) Size() int {
	return len(bs.nodes)
}
func (bs BasicSpline2D) Looped() bool {
	return bs.looped
}

func NewBasicSpline2D(looped bool, nodes ...Point2D) Spline2D {
	return BasicSpline2D{nodes: nodes, looped: looped}
}

type BorderedSpline2D interface {
	Spline2D
	Width(node int) float64 // -1 should represent a invalid node
}

type BasicBorderedSpline2D struct {
	nodes    []Point2D
	looped   bool
	widthMap []float64
}

func (bbs BasicBorderedSpline2D) GetNodes() []Point2D {
	return bbs.nodes
}
func (bbs BasicBorderedSpline2D) Size() int {
	return len(bbs.nodes)
}
func (bbs BasicBorderedSpline2D) Looped() bool {
	return bbs.looped
}
func (bbs BasicBorderedSpline2D) Width(node int) float64 {
	if node <= len(bbs.widthMap)-1 {
		return bbs.widthMap[node]
	}
	return -1
}

func NewBasicBorderedSpline2D(looped bool, nodes ...Point2D) Spline2D {
	return BasicSpline2D{nodes: nodes, looped: looped}
}
