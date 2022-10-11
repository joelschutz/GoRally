package mechanics

import (
	"fmt"
	"math"
	"time"

	"github.com/joelschutz/gorally/models"
)

type AppliedAcceleration float64

type VehicleState struct {
	Speed, Acceleration, Damage, Fuel float64
	Location                          uint64 // Segment index
}

func calcFinalVelocity(v models.Vehicle, vs VehicleState, aa AppliedAcceleration, tickTime time.Duration) (float64, error) {
	accSum := calcFinalAcceleratio(v, vs, aa)

	speed := vs.Speed + (accSum * tickTime.Seconds())
	distance := uint32(speed * tickTime.Seconds())

	if checkValidAppliedAcceleration(aa, v, distance) {
		return speed, nil
	}
	return 0, fmt.Errorf("Invalid AppliedAcceleration")
}

func calcFinalAcceleratio(v models.Vehicle, vs VehicleState, aa AppliedAcceleration) float64 {
	return vs.Acceleration + float64(aa) - calcAirDragForce(v, vs)
}

func calcAirDragForce(v models.Vehicle, vs VehicleState) float64 {
	airDrag := 0.33            // Admendional
	airDensity := 1.25         // kg/m3
	crossArea := 1.            // m2
	speed := float64(vs.Speed) // m/s
	return 0.5 * airDrag * airDensity * crossArea * math.Pow(speed, 2)
}

func calcMaxDeceleration(v models.Vehicle, distance uint32) float64 {
	torque := float64(v.VehicleStats.BreakTorque) // Nm
	vehicleMass := float64(v.VehicleStats.Weight) // kg
	d := float64(distance)                        // m
	return -(torque / (vehicleMass * d))
}

func calcMaxAccleration(v models.Vehicle, distance uint32) float64 {
	torque := float64(v.VehicleStats.Torque)      // Nm
	vehicleMass := float64(v.VehicleStats.Weight) // kg
	d := float64(distance)                        // m
	return torque / (vehicleMass * d)
}

// Checks

func checkValidAppliedAcceleration(aa AppliedAcceleration, v models.Vehicle, distance uint32) bool {
	maxAa := 0.
	if aa > 0 {
		maxAa = calcMaxAccleration(v, distance)
	}
	if aa < 0 {
		maxAa = calcMaxDeceleration(v, distance)
	}
	if aa > AppliedAcceleration(maxAa) {
		return false
	}
	return true
}
