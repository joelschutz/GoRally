package mechanics

import (
	"math/rand"

	"github.com/joelschutz/gorally/models"
)

type DriverProjections struct {
	segmentIndex                 uint32
	targetSpeedsBySegment        []float64
	targetBreakingZonesBySegment []float64
}

func CalcDriverAcceleration(t models.Track, r models.Driver, v models.Vehicle, s models.Segmnent, vs VehicleState) float64 {
	paceNotes := fetchPaceNotes(t, r, vs)
	projections := calcDriverProjections(paceNotes, r, v)

	return 0
}

func calcDriverProjections(paceNotes []models.Segmnent, r models.Driver, v models.Vehicle) (p DriverProjections) {
	nextTargetSpeed := calcEstimateMaxSpeed(r, v, paceNotes[len(paceNotes)-1])
	p.targetSpeedsBySegment[len(paceNotes)-1] = nextTargetSpeed
	for i := len(paceNotes) - 2; i >= 0; i-- {
		segment := paceNotes[i]
		maxSpeed := calcEstimateMaxSpeed(r, v, segment)
		maxTorque := calcEstimateMaxTorque(r, v, segment)
		breakingZone := calcEstimateBreakingZone(maxSpeed, nextTargetSpeed, maxTorque)

		for {
			if breakingZone <= segment.Length {
				break
			}
			maxSpeed -= maxSpeed * (0.2 / float64(r.DrivingStyle[3]+1))
			breakingZone = calcEstimateBreakingZone(maxSpeed, nextTargetSpeed, maxTorque)
		}
		nextTargetSpeed = maxSpeed
		p.targetSpeedsBySegment[i] = maxSpeed
		p.targetBreakingZonesBySegment[i] = breakingZone
	}
	return p
}

func calcEstimateMaxSpeed(r models.Driver, v models.Vehicle, s models.Segmnent) float64 {
	realMS := CalcMaxSegmentSpeed(s, v)
	estimateError := ((.1 * (float64(r.DrivingStyle[0] + 1))) / (float64(r.TerrainSkills[s.Terrain]) + 1.)) * rand.Float64()
	if rand.Intn(2) == 0 {
		estimateError *= -1
	}
	return estimateError*realMS + realMS
}

func calcEstimateMaxTorque(r models.Driver, v models.Vehicle, s models.Segmnent) float64 {
	realMS := CalcMaxSegmentTorque(s, v)
	estimateError := ((.1 * (float64(r.DrivingStyle[2] + 1))) / (float64(r.VehicleSkills[v.DriveTrain]) + 1.)) * rand.Float64()
	if rand.Intn(2) == 0 {
		estimateError *= -1
	}
	return estimateError*realMS + realMS
}

func fetchPaceNotes(t models.Track, r models.Driver, vs VehicleState) (s []models.Segmnent) {
	noteCount := (r.DrivingStyle[1] + 4) / 2
	for i, segment := range t.Segments {
		if i > int(noteCount) {
			break
		}
		s = append(s, segment)
	}
	return s
}

func calcBreakingZone(initialSpeed, targetSpeed, maxTorque float64) (travelDistance float64) {
	currentSpeed := initialSpeed
	deltaTime := (initialSpeed - targetSpeed) / maxTorque

	for i := 0; i < int(deltaTime); i++ {
		travelDistance += currentSpeed - maxTorque
		currentSpeed -= maxTorque
	}

	speedDelta := currentSpeed - targetSpeed
	travelDistance += currentSpeed - speedDelta

	return travelDistance
}

func calcEstimateBreakingZone(r models.Driver, v models.Vehicle, s1Speed, s2Speed, s1Torque float64) (travelDistance float64) {
	realMS := calcBreakingZone(s1Speed, s2Speed, s1Torque)
	estimateError := ((.05) / (float64(r.DrivingStyle[4] + 1))) * rand.Float64()
	if rand.Intn(2) == 0 {
		estimateError *= -1
	}
	return estimateError*realMS + realMS
}
