package models

type Racer struct {
	Name  string
	Age   string
	Stats RacerStats
}

type RacerStats struct {
	DrivingSkills DrivingSkills
	DrivingStyle  DrivingStyle
	VehicleSkills VehicleSkills
}

type DrivingSkills struct {
	General, Snow, Tarmac, LightGravel, HeavyGravel int32
}

type DrivingStyle struct {
	Recklessness, Repairing, Dedication, Adaptability, Communication int32
}

type VehicleSkills struct {
	FWD, RWD, AWD int32
}
