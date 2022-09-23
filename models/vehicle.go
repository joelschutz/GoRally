package models

type Vehicle struct {
	Name         string
	Manufacturer string
	VehicleStats VehicleStats
}

type DriveTrain uint8

const (
	FWD DriveTrain = iota
	RWD
	AWD
)

type VehicleStats struct {
	Weight, Power, Torque, Gears uint32
	Class                        Class
	DriveTrain                   DriveTrain
}
