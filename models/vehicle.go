package models

type VehicleState struct {
	Speed, Acceleration, Damage, Fuel float64
	Location                          uint64 // Segment index
}

type Vehicle struct {
	Name         string       `json:"name"`
	Manufacturer string       `json:"manufacturer"`
	Class        Class        `json:"cls"`
	DriveTrain   DriveTrain   `json:"driveTrain"`
	VehicleStats VehicleStats `json:"vehicleStats"`
}

type DriveTrain uint8

const (
	FWD DriveTrain = iota
	RWD
	AWD
)

// Represents vehicle stats
// [0]Weight
// [1]Power
// [2]Torque
// [3]BreakTorque
// [4]Gears
type VehicleStats [5]uint32
