package storage

import (
	"context"
	"fmt"

	"github.com/joelschutz/gorally/comm/schema"
	"github.com/joelschutz/gorally/models"
)

type Storage interface {
	GetVehicle(ctx context.Context, index uint) (models.Vehicle, error)
	GetDriver(ctx context.Context, index uint) (models.Driver, error)
	GetTrack(ctx context.Context, index uint) (models.Track, error)
	GetEvent(ctx context.Context, index uint) (models.Event, error)
	GetCompetitor(ctx context.Context, index uint) (models.Competitor, error)
	AddVehicle(ctx context.Context, obj models.Vehicle) error
	AddDriver(ctx context.Context, obj models.Driver) error
	AddTrack(ctx context.Context, obj models.Track) error
	AddEvent(ctx context.Context, obj schema.EventSchema) error
	AddCompetitor(ctx context.Context, obj schema.CompetitorSchema) error
	ListVehicles(ctx context.Context) ([]models.Vehicle, error)
	ListDrivers(ctx context.Context) ([]models.Driver, error)
	ListTracks(ctx context.Context) ([]models.Track, error)
	ListEvents(ctx context.Context) ([]models.Event, error)
	ListCompetitors(ctx context.Context) ([]models.Competitor, error)
	UpdateVehicle(ctx context.Context, index uint, obj models.Vehicle) error
	UpdateDriver(ctx context.Context, index uint, obj models.Driver) error
	UpdateTrack(ctx context.Context, index uint, obj models.Track) error
	UpdateEvent(ctx context.Context, index uint, obj schema.EventSchema) error
	UpdateCompetitor(ctx context.Context, index uint, obj schema.CompetitorSchema) error
}

type MemoryDB struct {
	Vehicles    []models.Vehicle
	Drivers     []models.Driver
	Tracks      []models.Track
	Events      []models.Event
	Competitors []models.Competitor
}

func NewMemoryDB() *MemoryDB {
	return &MemoryDB{
		Vehicles:    []models.Vehicle{},
		Drivers:     []models.Driver{},
		Tracks:      []models.Track{},
		Events:      []models.Event{},
		Competitors: []models.Competitor{},
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

func (s *MemoryDB) GetEvent(ctx context.Context, index uint) (models.Event, error) {
	if len(s.Events) > int(index) {
		return s.Events[index], nil
	}
	return models.Event{}, fmt.Errorf("Index Out Of Range")
}

func (s *MemoryDB) GetCompetitor(ctx context.Context, index uint) (models.Competitor, error) {
	if len(s.Competitors) > int(index) {
		return s.Competitors[index], nil
	}
	return models.Competitor{}, fmt.Errorf("Index Out Of Range")
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

func (s *MemoryDB) ListEvents(ctx context.Context) ([]models.Event, error) {
	return s.Events, nil
}

func (s *MemoryDB) ListCompetitors(ctx context.Context) ([]models.Competitor, error) {
	return s.Competitors, nil
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

func (s *MemoryDB) UpdateEvent(ctx context.Context, index uint, obj schema.EventSchema) error {
	_, err := s.GetEvent(ctx, index)
	if err != nil {
		return err
	}
	s.Events[index] = models.Event{
		Name:        obj.Name,
		Class:       models.Class(obj.Class),
		Competitors: []models.Competitor{},
	}
	for _, v := range obj.Competitors {
		s.Events[index].Competitors = append(s.Events[index].Competitors, s.Competitors[v])
	}
	return nil
}

func (s *MemoryDB) UpdateCompetitor(ctx context.Context, index uint, obj schema.CompetitorSchema) error {
	_, err := s.GetCompetitor(ctx, index)
	if err != nil {
		return err
	}
	s.Competitors[index] = models.Competitor{
		Vehicle: s.Vehicles[obj.Vehicle],
		Driver:  s.Drivers[obj.Driver],
	}
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

func (s *MemoryDB) AddEvent(ctx context.Context, obj schema.EventSchema) error {
	e := models.Event{
		Name:        obj.Name,
		Class:       models.Class(obj.Class),
		Competitors: []models.Competitor{},
	}
	for _, v := range obj.Competitors {
		e.Competitors = append(e.Competitors, s.Competitors[v])
	}
	s.Events = append(s.Events, e)
	return nil
}

func (s *MemoryDB) AddCompetitor(ctx context.Context, obj schema.CompetitorSchema) error {
	s.Competitors = append(s.Competitors, models.Competitor{
		Vehicle: s.Vehicles[obj.Vehicle],
		Driver:  s.Drivers[obj.Driver],
	})
	return nil
}
