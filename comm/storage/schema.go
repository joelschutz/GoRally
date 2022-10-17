package storage

import "gorm.io/gorm"

type Driver struct {
	gorm.Model
	Name          string
	Age           uint64
	TerrainSkills TerrainSkills `gorm:"embedded"`
	DrivingStyle  DrivingStyle  `gorm:"embedded"`
	VehicleSkills VehicleSkills `gorm:"embedded"`
	CompetitorID  uint
}
type TerrainSkills struct{ Tarmac, Mud, HeavyGravel, LightGravel, Sand, Snow uint32 }
type VehicleSkills struct{ FWD, RWD, AWD uint32 }
type DrivingStyle struct{ Recklessness, Communication, Aggressiveness, Adaptability, Reflexes uint32 }

type Vehicle struct {
	gorm.Model
	Name         string
	Manufacturer string
	Class        Class
	DriveTrain   DriveTrain
	VehicleStats VehicleStats `gorm:"embedded"`
	CompetitorID uint
}
type DriveTrain uint8
type Class uint8
type VehicleStats struct{ Weight, Power, Torque, BreakTorque, Gears uint32 }

type Track struct {
	gorm.Model
	Name     string
	Country  string
	Segments []Segmnent
}

type Segmnent struct {
	gorm.Model
	TrackID                    uint
	Length                     float64
	Direction, Corner, Terrain uint8
	Cut, Unseen, Jump          bool
}

type Event struct {
	gorm.Model
	Name        string
	Class       Class
	Tracks      []Track      `gorm:"many2many:event_tracks;"`
	Competitors []Competitor `gorm:"many2many:event_competitors;"`
}

type Competitor struct {
	gorm.Model
	Vehicle Vehicle
	Driver  Driver
}
