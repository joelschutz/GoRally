package mechanics

import (
	"math"

	"github.com/joelschutz/gorally/models"
)

func CalcTrackTime(t models.Track, r models.Driver, v models.Vehicle) (tr models.TrackResult) {
	g := GetBasicSplineGradient(t.Segments, 0)
	tan := math.Atan2(-g.Y(), g.X())
	vs := models.VehicleState{Location: t.Segments.GetNode(0), Tangent: tan}
	ts := models.TrackState{}
	dp := DriverProjections{}

	for i, segment := range t.Segments {
		ts.Location = uint32(i)
		ts.PaceNotes = fetchPaceNotes(t, r, uint64(i))
		ts.Terrain = segment.Terrain
		projections := AppendDriverProjections(dp, r, v, vs, ts, 1)
		// Old Code Below
		for {
			tr.TimeBySegment[i]++
			tr.VStateBySecond = append(tr.VStateBySecond, vs)
			tr.TStateBySecond = append(tr.TStateBySecond, ts)
			acc := CalcDriverAcceleration(projections, vs, ts)
			if vs.Speed <= ts.MaxSpeed && acc <= ts.MaxTorque { // Check if car is too fast to make the segment
				speed := vs.Speed + acc
				distanceTraveled := speed
				vs.Speed = speed
				vs.Location = uint64(distanceTraveled)
				if distanceTraveled >= ts.DistanceLeft {
					pTime := ts.DistanceLeft / distanceTraveled
					tr.TimeBySegment[i] -= 1 - pTime
					break
				}
				ts.DistanceLeft -= distanceTraveled
			} else {
				tr.TimeBySegment[i] += 5 // Adds 5s penalty
				vs.Speed = 0             // Stop vehicle
				vs.Damage += 1           // Apply Damage
			}
		}
	}
	tr.TotalTime = sumSegmentTimes(tr.TimeBySegment)
	return tr
}

func sumSegmentTimes(times []float64) (sum float64) {
	for _, v := range times {
		sum += v
	}
	return sum
}

func fetchPaceNotes(t models.Track, r models.Driver, segmentIndex uint64) (s []models.Segment) {
	noteCount := FetchPaceNotesCount(r)
	for i := segmentIndex; i < noteCount+segmentIndex; i++ {
		s = append(s, t.Segments[i])
	}
	return s
}

func CalcTerrainStaticFiction(t models.Terrain) float64 {
	return 0.6 / float64(t)
}

func CalcTerrainKineticFiction(t models.Terrain) float64 {
	return 0.4 / float64(t)
}

func calcCornerRadius(c models.CornerLevel) float64 {
	return 100 / float64(c)
}
