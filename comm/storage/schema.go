package storage

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/joelschutz/gorally/models"
	"gorm.io/gorm"
)

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
	Segments Segments
}

type Segment struct {
	Length                     float64
	Direction, Corner, Terrain uint8
	Cut, Unseen, Jump          bool
}

type Segments []Segment

// Scan scan value into Jsonb, implements sql.Scanner interface
func (j *Segments) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	return json.Unmarshal(bytes, &j)
}

// Value return json value, implement driver.Valuer interface
func (j Segments) Value() (driver.Value, error) {
	s, err := json.Marshal(j)
	if err != nil {
		return nil, err
	}
	return s, nil
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

func DriverToStorage(d models.Driver) Driver {
	return Driver{
		Name: d.Name,
		Age:  d.Age,
		TerrainSkills: TerrainSkills{
			Tarmac:      d.TerrainSkills[0],
			Mud:         d.TerrainSkills[1],
			HeavyGravel: d.TerrainSkills[2],
			LightGravel: d.TerrainSkills[3],
			Sand:        d.TerrainSkills[4],
			Snow:        d.TerrainSkills[5],
		},
		VehicleSkills: VehicleSkills{
			FWD: d.VehicleSkills[0],
			RWD: d.VehicleSkills[1],
			AWD: d.VehicleSkills[2],
		},
		DrivingStyle: DrivingStyle{
			Recklessness:   d.DrivingStyle[0],
			Communication:  d.DrivingStyle[1],
			Aggressiveness: d.DrivingStyle[2],
			Adaptability:   d.DrivingStyle[3],
			Reflexes:       d.DrivingStyle[4],
		},
	}
}

func DriverFromStorage(d Driver) models.Driver {
	return models.Driver{
		Name: d.Name,
		Age:  d.Age,
		TerrainSkills: models.TerrainSkills{
			d.TerrainSkills.Tarmac,
			d.TerrainSkills.Mud,
			d.TerrainSkills.HeavyGravel,
			d.TerrainSkills.LightGravel,
			d.TerrainSkills.Sand,
			d.TerrainSkills.Snow,
		},
		VehicleSkills: models.VehicleSkills{
			d.VehicleSkills.FWD,
			d.VehicleSkills.RWD,
			d.VehicleSkills.AWD,
		},
		DrivingStyle: models.DrivingStyle{
			d.DrivingStyle.Recklessness,
			d.DrivingStyle.Communication,
			d.DrivingStyle.Aggressiveness,
			d.DrivingStyle.Adaptability,
			d.DrivingStyle.Reflexes,
		},
	}
}
func VehicleToStorage(d models.Vehicle) Vehicle {
	return Vehicle{
		Name:         d.Name,
		Manufacturer: d.Manufacturer,
		Class:        Class(d.Class),
		DriveTrain:   DriveTrain(d.DriveTrain),
		VehicleStats: VehicleStats{
			Weight:      d.VehicleStats[0],
			Power:       d.VehicleStats[1],
			Torque:      d.VehicleStats[2],
			BreakTorque: d.VehicleStats[3],
			Gears:       d.VehicleStats[4],
		},
	}
}

func VehicleFromStorage(d Vehicle) models.Vehicle {
	return models.Vehicle{
		Name:         d.Name,
		Manufacturer: d.Manufacturer,
		Class:        models.Class(d.Class),
		DriveTrain:   models.DriveTrain(d.DriveTrain),
		VehicleStats: models.VehicleStats{
			d.VehicleStats.Weight,
			d.VehicleStats.Power,
			d.VehicleStats.Torque,
			d.VehicleStats.BreakTorque,
			d.VehicleStats.Gears,
		},
	}
}

func TrackToStorage(d models.Track) Track {
	t := Track{
		Name:     d.Name,
		Country:  d.Country,
		Segments: []Segment{},
	}

	for _, v := range d.Segments {
		t.Segments = append(t.Segments, Segment{
			Corner:    uint8(v.Corner),
			Terrain:   uint8(v.Terrain),
			Direction: uint8(v.Direction),
			Length:    v.Length,
			Cut:       v.Cut,
			Jump:      v.Jump,
			Unseen:    v.Unseen,
		})
	}
	return t
}

func TrackFromStorage(d Track) models.Track {
	t := models.Track{
		Name:     d.Name,
		Country:  d.Country,
		Segments: []models.Segment{},
	}

	for _, v := range d.Segments {
		t.Segments = append(t.Segments, models.Segment{
			Corner:    models.CornerLevel(v.Corner),
			Terrain:   models.Terrain(v.Terrain),
			Direction: models.Direction(v.Direction),
			Length:    v.Length,
			Cut:       v.Cut,
			Jump:      v.Jump,
			Unseen:    v.Unseen,
		})
	}
	return t
}

func EventFromStorage(d Event) models.Event {
	t := models.Event{
		Name:        d.Name,
		Class:       models.Class(d.Class),
		Tracks:      []models.Track{},
		Competitors: []models.Competitor{},
	}

	for _, v := range d.Tracks {
		t.Tracks = append(t.Tracks, TrackFromStorage(v))
	}

	for _, v := range d.Competitors {
		t.Competitors = append(t.Competitors, CompetitorFromStorage(v))
	}
	return t
}

func CompetitorFromStorage(d Competitor) models.Competitor {
	return models.Competitor{
		Vehicle: VehicleFromStorage(d.Vehicle),
		Driver:  DriverFromStorage(d.Driver),
	}
}
