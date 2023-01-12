package mechanics

import (
	"math"
	"math/rand"

	"github.com/joelschutz/gorally/models"
)

type GripEstimate struct {
	ErrorAmount float64
	Random      bool
	VehicleMass float64
	Terrain     models.Terrain
	Gravity     float64
}

func (ge GripEstimate) X() float64 {
	errAmount := ge.ErrorAmount
	if ge.Random {
		errAmount *= rand.Float64()
		if rand.Intn(2) == 0 {
			errAmount *= -1
		}
	}
	return CalcTerrainKineticFiction(ge.Terrain) * ge.VehicleMass * ge.Gravity * errAmount
}
func (ge GripEstimate) Y() float64 {
	errAmount := ge.ErrorAmount
	if ge.Random {
		errAmount *= rand.Float64()
		if rand.Intn(2) == 0 {
			errAmount *= -1
		}
	}
	return CalcTerrainStaticFiction(ge.Terrain) * ge.VehicleMass * ge.Gravity * errAmount
}
func (ge GripEstimate) SetX(x float64) {
	return
}
func (ge GripEstimate) SetY(y float64) {
	return
}

type ProjectionNode struct {
	location     models.Point2D
	SegmentIndex int
	tangent      float64
	velocity     float64
	links        []models.GraphLink
}

func (pn ProjectionNode) Links() []models.GraphLink {
	return pn.links
}

type ProjectionLink struct {
	node          models.GraphNode
	nodeIndex     int
	tangentDelta  float64
	velocityDelta float64
}

func (pl ProjectionLink) To() models.GraphNode {
	return pl.node
}

func (pl ProjectionLink) ToIndex() int {
	return pl.nodeIndex
}

type DriverProjectionsGraph struct {
	Nodes []ProjectionNode
}

func (dpg DriverProjectionsGraph) GetNodes() []ProjectionNode {
	return dpg.Nodes
}

type DriverProjectionGenerations struct {
	Generations []DriverProjectionsGraph
}

type DriverProjections struct {
	segmentIndex          uint32
	projectionGenerations DriverProjectionGenerations
	generationsBySegment  [][]int
}

func CalcDriverAcceleration(p DriverProjections, vs models.VehicleState, ts models.TrackState) (acc float64) {
	if ts.DistanceLeft <= p.targetBreakingZonesBySegment[0] {
		acc = p.targetSpeedsBySegment[1] - vs.Speed
		if acc < (-1 * p.targetTorqueBySegment[0]) {
			acc = -1 * p.targetTorqueBySegment[0]
		}
	} else {
		acc = p.targetSpeedsBySegment[0] - vs.Speed
		if acc > p.targetTorqueBySegment[0] {
			acc = p.targetTorqueBySegment[0]
		}

	}
	return acc
}

func AppendDriverProjections(dp DriverProjections, r models.Driver, v models.Vehicle, vs models.VehicleState, ts models.TrackState, projectionLengthInSegments int) (p DriverProjections) {
	paceNotes := ts.PaceNotes
	p.segmentIndex = uint32(ts.Location)

	var lastGeneration DriverProjectionsGraph
	var generationCount int
	if p.segmentIndex > 0 {
		lastSegmentGenerations := dp.generationsBySegment[p.segmentIndex-1]
		generationCount = len(lastSegmentGenerations)
		lastGenerationIndex := lastSegmentGenerations[generationCount-1]
		lastGeneration = dp.projectionGenerations.Generations[lastGenerationIndex]
	} else {
		startNode := ProjectionNode{
			location:     vs.Location,
			SegmentIndex: -1,
			tangent:      vs.Tangent,
			velocity:     0,
		}
		lastGeneration = DriverProjectionsGraph{[]ProjectionNode{startNode}}
	}
	lastSegmentFound := p.segmentIndex

	for {
		segmentsFound := []int{int(lastSegmentFound)}
		var newGeneration DriverProjectionsGraph
		for _, node := range lastGeneration.GetNodes() {
			// Calc next generation
			/// Find Driver Grip Resolution(the amount of points to create)
			gripResolution := FetchGripResolution(r)
			// Estimate Grip
			gripEstimate := calcEstimateGrip(r, v, ts.Terrain)
			// Calc velocity vector
			velocityVector := models.NewBasicPoint2D(node.location.X(), node.location.Y()+node.velocity)
			// Plot each vector
			for teta := -math.Pi; teta < math.Pi; teta += (math.Pi / 2) / float64(gripResolution) {
				newPoint := models.NewBasicPoint2D(velocityVector.X()+math.Sin(teta)*float64(gripEstimate.X()), velocityVector.Y()+math.Cos(teta)*float64(gripEstimate.Y()))
				currentOffset := float64(p.segmentIndex)
				trackResolution := FetchTrackResolution(r)
				isOut := IsPointInSplineSectionBoarders(paceNotes, newPoint, trackResolution, currentOffset, currentOffset+trackResolution*5)
				if isOut {
					for i := 0; i < 4; i++ {
						currentOffset = trackResolution * 5 * float64(i)
						nextOffset := trackResolution * 5 * float64(i+1)
						if !IsPointAboveSection(paceNotes, newPoint, nextOffset) {
							// Point Out of the Track
							isOut = IsPointInSplineSectionBoarders(paceNotes, newPoint, trackResolution, currentOffset, nextOffset)
							break
						}
					}
				}
				newVelocity := DistanceBetweenTwoPoints(node.location, newPoint)

				if !isOut && (newVelocity > 0) {
					if uint32(currentOffset) != lastSegmentFound {
						lastSegmentFound = uint32(currentOffset)
						segmentsFound = append(segmentsFound, int(lastSegmentFound))
					}
					newNode := ProjectionNode{
						location: newPoint,
						tangent:  node.tangent + teta,
						velocity: newVelocity,
					}
					node.links = append(node.links, ProjectionLink{
						nodeIndex:     len(newGeneration.Nodes),
						velocityDelta: newVelocity - node.velocity,
						tangentDelta:  teta,
						node:          &newNode,
					})
					newGeneration.Nodes = append(newGeneration.Nodes, newNode)
				}
			}
		}
		if lastSegmentFound > p.segmentIndex+uint32(projectionLengthInSegments) {
			break
		}
		dp.projectionGenerations.Generations = append(dp.projectionGenerations.Generations, newGeneration)
		dp.generationsBySegment[generationCount] = append(dp.generationsBySegment[generationCount], segmentsFound...)
		generationCount++
		lastGeneration = newGeneration
	}

	return dp
}

func calcEstimateGrip(r models.Driver, v models.Vehicle, t models.Terrain) (g GripEstimate) {
	g.ErrorAmount = (.1 * (float64(r.DrivingStyle[2] + 1))) / (float64(r.VehicleSkills[v.DriveTrain]) + 1.)
	g.Terrain = t
	g.VehicleMass = float64(v.VehicleStats[0])
	g.Gravity = 3.14
	return g
}

func FetchPaceNotesCount(r models.Driver) uint64 {
	return (uint64(r.DrivingStyle[1]) + 4) / 2
}

func FetchGripResolution(r models.Driver) int {
	return int(uint64(r.DrivingStyle[4])/2 + 6)
}

func FetchTrackResolution(r models.Driver) float64 {
	return 0.1 / float64(r.DrivingStyle[3]+1)
}
