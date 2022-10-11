package models

type Driver struct {
	Name  string
	Age   string
	Stats DriverStats
}

type DriverStats struct {
	TerrainSkills TerrainSkills
	DrivingStyle  DrivingStyle
	VehicleSkills VehicleSkills
}

//	Represents the driver skill for each terrain
// 	[0]Tarmac
//  [1]Mud
//  [2]HeavyGravel
//  [3]LightGravel
//  [4]Sand
//  [5]Snow
type TerrainSkills [6]int32

//	Represents the driver skill for each drivetrain
// 	[0]FWD
//  [1]RWD
//  [2]AWD
type VehicleSkills [3]int32

type DrivingStyle struct {
	Recklessness   uint64 // Limits the max speed estimate accurace
	Communication  uint64 // Extend the max number os pacenotes called
	Aggressiveness uint64 // Limits the max torque estimate accurace
	Adaptability   uint64 // Extend the breaking zone estimate accurace
	Repairing      int32
}
