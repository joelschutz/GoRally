package models

type Driver struct {
	Name          string        `json:"name"`
	Age           uint64        `json:"age"`
	TerrainSkills TerrainSkills `json:"terrainSkills"`
	DrivingStyle  DrivingStyle  `json:"drivingStyle"`
	VehicleSkills VehicleSkills `json:"vehicleSkills"`
}

// Represents the driver skill for each terrain
//
// [0]Tarmac
// [1]Mud
// [2]HeavyGravel
// [3]LightGravel
// [4]Sand
// [5]Snow
type TerrainSkills [6]uint32

// Represents the driver skill for each drivetrain
//
// [0]FWD
// [1]RWD
// [2]AWD
type VehicleSkills [3]uint32

//	Represents the driver style for each caracteristic
//
// [0]Recklessness - Limits the max speed estimate accurace
// [1]Communication - Extend the max number os pacenotes called
// [2]Aggressiveness - Limits the max torque estimate accurace
// [3]Adaptability - Limits the track estimate resolution
// [4]Reflexes - Limits the grip estimate resolution
type DrivingStyle [5]uint32
