package services

import (
	"context"
	"fmt"

	"github.com/joelschutz/gorally/models"
)

type Storage interface {
	GetVehicle(ctx context.Context, index uint) (models.Vehicle, error)
	GetDriver(ctx context.Context, index uint) (models.Driver, error)
	GetTrack(ctx context.Context, index uint) (models.Track, error)
	AddVehicle(ctx context.Context, obj models.Vehicle) error
	AddDriver(ctx context.Context, obj models.Driver) error
	AddTrack(ctx context.Context, obj models.Track) error
	ListVehicles(ctx context.Context) ([]models.Vehicle, error)
	ListDrivers(ctx context.Context) ([]models.Driver, error)
	ListTracks(ctx context.Context) ([]models.Track, error)
	UpdateVehicle(ctx context.Context, index uint, obj models.Vehicle) error
	UpdateDriver(ctx context.Context, index uint, obj models.Driver) error
	UpdateTrack(ctx context.Context, index uint, obj models.Track) error
}

var db MemoryDB

type MemoryDB struct {
	Vehicles []models.Vehicle
	Drivers  []models.Driver
	Tracks   []models.Track
}

func NewMemoryDB() *MemoryDB {
	return &MemoryDB{
		Vehicles: []models.Vehicle{},
		Drivers:  []models.Driver{},
		Tracks:   []models.Track{},
	}
}

func (s *MemoryDB) GetVehicle(ctx context.Context, index uint) (models.Vehicle, error) {
	if len(s.Vehicles) > int(index) {
		return s.Vehicles[index], nil
	}
	return models.Vehicle{}, fmt.Errorf("Index Out Of Range")
}

func (s *MemoryDB) GetDriver(ctx context.Context, index uint) (models.Driver, error) {
	if len(s.Drivers) > int(index) {
		return s.Drivers[index], nil
	}
	return models.Driver{}, fmt.Errorf("Index Out Of Range")
}

func (s *MemoryDB) GetTrack(ctx context.Context, index uint) (models.Track, error) {
	if len(s.Tracks) > int(index) {
		return s.Tracks[index], nil
	}
	return models.Track{}, fmt.Errorf("Index Out Of Range")
}

func (s *MemoryDB) ListVehicles(ctx context.Context) ([]models.Vehicle, error) {
	return s.Vehicles, nil
}

func (s *MemoryDB) ListDrivers(ctx context.Context) ([]models.Driver, error) {
	return s.Drivers, nil
}

func (s *MemoryDB) ListTracks(ctx context.Context) ([]models.Track, error) {
	return s.Tracks, nil
}

func (s *MemoryDB) UpdateVehicle(ctx context.Context, index uint, obj models.Vehicle) error {
	_, err := s.GetVehicle(ctx, index)
	if err != nil {
		return err
	}
	s.Vehicles[index] = obj
	return nil
}

func (s *MemoryDB) UpdateDriver(ctx context.Context, index uint, obj models.Driver) error {
	_, err := s.GetDriver(ctx, index)
	if err != nil {
		return err
	}
	s.Drivers[index] = obj
	return nil
}

func (s *MemoryDB) UpdateTrack(ctx context.Context, index uint, obj models.Track) error {
	_, err := s.GetTrack(ctx, index)
	if err != nil {
		return err
	}
	s.Tracks[index] = obj
	return nil
}

func (s *MemoryDB) AddVehicle(ctx context.Context, obj models.Vehicle) error {
	s.Vehicles = append(s.Vehicles, obj)
	return nil
}

func (s *MemoryDB) AddDriver(ctx context.Context, obj models.Driver) error {
	s.Drivers = append(s.Drivers, obj)
	return nil
}

func (s *MemoryDB) AddTrack(ctx context.Context, obj models.Track) error {
	s.Tracks = append(s.Tracks, obj)
	return nil
}
