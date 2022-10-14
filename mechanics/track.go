package mechanics

import (
	"fmt"
	"math"

	"github.com/joelschutz/gorally/models"
)

type TrackState struct {
	s            models.Segmnent
	paceNotes    []models.Segmnent
	maxSpeed     float64
	maxTorque    float64
	distanceLeft float64
}

func CalcTrackTime(t models.Track, r models.Driver, v models.Vehicle) (tr models.TrackResult) {
	vs := VehicleState{}
	ts := TrackState{}
	segmentsTime := []float64{}
	for i, segment := range t.Segments {
		ts.distanceLeft = float64(segment.Length)
		ts.maxSpeed = CalcMaxSegmentSpeed(segment, v)
		ts.maxTorque = CalcMaxSegmentTorque(segment, v)
		ts.paceNotes = fetchPaceNotes(t, r, uint64(i))
		projections := CalcDriverProjections(r, v, vs, ts)
		for {
			segmentsTime[i]++
			acc := CalcDriverAcceleration(projections, vs, ts)
			if vs.Speed <= ts.maxSpeed && acc <= ts.maxTorque { // Check if car is too fast to make the segment
				speed := vs.Speed + acc
				distanceTraveled := speed
				vs.Speed = speed
				vs.Location = uint64(distanceTraveled)
				if distanceTraveled >= ts.distanceLeft {
					pTime := ts.distanceLeft / distanceTraveled
					segmentsTime[i] -= 1 - pTime
					break
				}
				ts.distanceLeft -= distanceTraveled
			} else {
				segmentsTime[i] += 5 // Adds 5s penalty
				vs.Speed = 0         // Stop vehicle
				vs.Damage += 1       // Apply Damage
			}
		}
	}
	tr.TotalTime = sumSegmentTimes(segmentsTime)
	tr.TimeBySegment = segmentsTime
	return tr
}

func sumSegmentTimes(times []float64) (sum float64) {
	for _, v := range times {
		sum += v
	}
	return sum
}

func fetchPaceNotes(t models.Track, r models.Driver, segmentIndex uint64) (s []models.Segmnent) {
	noteCount := FetchPaceNotesCount(r)
	for i := segmentIndex; i < noteCount+segmentIndex; i++ {
		s = append(s, t.Segments[i])
	}
	return s
}

func CalcMaxSegmentTorque(s models.Segmnent, v models.Vehicle) float64 {
	if s.Direction == models.Straight {
		return CalcMaxStraightTorque(s, v)
	}
	return CalcMaxCornerTorque(s, v)
}

func CalcMaxStraightTorque(s models.Segmnent, v models.Vehicle) float64 {
	terrainDrag := calcTerrainDrag(s.Terrain) // Admendional
	vehicleMass := float64(v.VehicleStats[0]) // kg
	gravity := 3.14                           // m/s2
	return terrainDrag * vehicleMass * gravity
}

func CalcMaxCornerTorque(s models.Segmnent, v models.Vehicle) float64 {
	terrainDrag := calcTerrainDrag(s.Terrain) // Admendional
	vehicleMass := float64(v.VehicleStats[0]) // kg
	gravity := 3.14                           // m/s2
	steringAngle := calcSteringAngle(s)       //radian
	return terrainDrag * vehicleMass * gravity * math.Cos(steringAngle)
	// return terrainDrag * vehicleMass * gravity
}

func CalcMaxSegmentSpeed(s models.Segmnent, v models.Vehicle) float64 {
	if s.Direction == models.Straight {
		return calcMaxStraightSpeed(s, v)
	}
	return calcMaxCornerSpeed(s, v)
}

func calcMaxStraightSpeed(s models.Segmnent, v models.Vehicle) float64 {
	power := float64(v.VehicleStats[1]) // W
	airDrag := 0.33                     // Admendional
	airDensity := 1.25                  // kg/m3
	crossArea := 1.                     // m2
	return math.Cbrt(2. * power / (airDrag * airDensity * crossArea))
}

func calcMaxCornerSpeed(s models.Segmnent, v models.Vehicle) float64 {
	terrainDrag := calcTerrainDrag(s.Terrain)  // Admendional
	gravity := 3.14                            // m/s2
	vehicleMass := float64(v.VehicleStats[0])  // kg
	cornerRadius := calcCornerRadius(s.Corner) // m
	slopeAngle := 0.                           // radian
	return math.Sqrt((vehicleMass*terrainDrag*gravity*math.Sin(slopeAngle))+(vehicleMass*gravity*math.Cos(slopeAngle))) * cornerRadius / vehicleMass
}

func calcTerrainDrag(t models.Terrain) float64 {
	return 0.6 / float64(t)
}

func calcCornerRadius(c models.CornerLevel) float64 {
	return 100 / float64(c)
}

func calcSteringAngle(s models.Segmnent) float64 {
	cornerPerimeter := float64(s.Length)       // m
	cornerRadius := calcCornerRadius(s.Corner) // m
	teta := (cornerPerimeter) / (cornerRadius) // radian
	fmt.Println("teta", (teta/2)*(180/math.Pi))
	fmt.Println("tetaRad", (teta / 2))
	return teta / 2
}
