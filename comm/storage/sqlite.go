package storage

import (
	"context"

	"github.com/joelschutz/gorally/comm/schema"
	"github.com/joelschutz/gorally/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SqliteDB struct {
	db *gorm.DB
}

func NewSqliteDB(fileName string) (*SqliteDB, error) {
	db, err := gorm.Open(sqlite.Open(fileName), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Migrate the schema
	db.AutoMigrate(&Driver{})
	db.AutoMigrate(&Track{})
	db.AutoMigrate(&Event{})
	db.AutoMigrate(&Vehicle{})
	db.AutoMigrate(&Competitor{})

	return &SqliteDB{db: db}, nil
}

func (s *SqliteDB) GetVehicle(ctx context.Context, index uint) (v models.Vehicle, err error) {
	o := Vehicle{}
	err = s.db.First(&o, index).Error
	return VehicleFromStorage(o), nil
}

func (s *SqliteDB) GetDriver(ctx context.Context, index uint) (d models.Driver, err error) {
	o := Driver{}
	err = s.db.First(&o, index).Error
	return DriverFromStorage(o), nil
}

func (s *SqliteDB) GetTrack(ctx context.Context, index uint) (t models.Track, err error) {
	o := Track{}
	err = s.db.First(&o, index).Error
	return TrackFromStorage(o), nil
}

func (s *SqliteDB) GetEvent(ctx context.Context, index uint) (e models.Event, err error) {
	o := Event{}
	err = s.db.First(&o, index).Error
	return EventFromStorage(o), nil
}

func (s *SqliteDB) GetCompetitor(ctx context.Context, index uint) (e models.Competitor, err error) {
	o := Competitor{}
	err = s.db.First(&o, index).Error
	return CompetitorFromStorage(o), nil
}

func (s *SqliteDB) ListVehicles(ctx context.Context) (e []models.Vehicle, err error) {
	o := []Vehicle{}
	err = s.db.Find(&o).Error
	for _, v := range o {
		e = append(e, VehicleFromStorage(v))
	}
	return
}

func (s *SqliteDB) ListDrivers(ctx context.Context) (e []models.Driver, err error) {
	o := []Driver{}
	err = s.db.Find(&o).Error
	for _, v := range o {
		e = append(e, DriverFromStorage(v))
	}
	return
}

func (s *SqliteDB) ListTracks(ctx context.Context) (e []models.Track, err error) {
	o := []Track{}
	err = s.db.Find(&o).Error
	for _, v := range o {
		e = append(e, TrackFromStorage(v))
	}
	return
}

func (s *SqliteDB) ListEvents(ctx context.Context) (e []models.Event, err error) {
	o := []Event{}
	err = s.db.Find(&o).Error
	for _, v := range o {
		e = append(e, EventFromStorage(v))
	}
	return
}

func (s *SqliteDB) ListCompetitors(ctx context.Context) (e []models.Competitor, err error) {
	o := []Competitor{}
	err = s.db.Find(&o).Error
	for _, v := range o {
		e = append(e, CompetitorFromStorage(v))
	}
	return
}

func (s *SqliteDB) UpdateVehicle(ctx context.Context, index uint, obj models.Vehicle) error {
	o := []Vehicle{}
	err := s.db.First(&o, index).Error
	if err != nil {
		return err
	}
	return s.db.Model(&o).Updates(VehicleToStorage(obj)).Error
}

func (s *SqliteDB) UpdateDriver(ctx context.Context, index uint, obj models.Driver) error {
	o := []Driver{}
	err := s.db.First(&o, index).Error
	if err != nil {
		return err
	}
	return s.db.Model(&o).Updates(DriverToStorage(obj)).Error
}

func (s *SqliteDB) UpdateTrack(ctx context.Context, index uint, obj models.Track) error {
	o := []Track{}
	err := s.db.First(&o, index).Error
	if err != nil {
		return err
	}
	return s.db.Model(&o).Updates(TrackToStorage(obj)).Error
}

func (s *SqliteDB) UpdateEvent(ctx context.Context, index uint, obj schema.EventSchema) error {
	o := Event{}
	err := s.db.First(&o, index).Error
	if err != nil {
		return err
	}
	i := Event{
		Name:  o.Name,
		Class: o.Class,
	}
	for _, v := range obj.Tracks {
		t := Track{}
		err := s.db.First(&t, v).Error
		if err != nil {
			return err
		}
		i.Tracks = append(i.Tracks, t)
	}
	for _, v := range obj.Competitors {
		t := Competitor{}
		err := s.db.First(&t, v).Error
		if err != nil {
			return err
		}
		i.Competitors = append(i.Competitors, t)
	}

	return s.db.Model(&o).Updates(i).Error
}

func (s *SqliteDB) UpdateCompetitor(ctx context.Context, index uint, obj schema.CompetitorSchema) error {
	o := Competitor{}
	err := s.db.First(&o, index).Error
	if err != nil {
		return err
	}

	t := Vehicle{}
	err = s.db.First(&t, obj.Vehicle).Error
	if err != nil {
		return err
	}
	d := Driver{}
	err = s.db.First(&d, obj.Driver).Error
	if err != nil {
		return err
	}
	i := Competitor{
		Vehicle: t,
		Driver:  d,
	}

	return s.db.Model(&o).Updates(i).Error
}

func (s *SqliteDB) AddVehicle(ctx context.Context, obj models.Vehicle) (err error) {
	o := VehicleToStorage(obj)
	return s.db.Create(&o).Error
}

func (s *SqliteDB) AddDriver(ctx context.Context, obj models.Driver) (err error) {
	o := DriverToStorage(obj)
	return s.db.Create(&o).Error
}

func (s *SqliteDB) AddTrack(ctx context.Context, obj models.Track) (err error) {
	o := TrackToStorage(obj)
	return s.db.Create(&o).Error
}

func (s *SqliteDB) AddEvent(ctx context.Context, obj schema.EventSchema) (err error) {
	o := Event{
		Name:  obj.Name,
		Class: Class(obj.Class),
	}
	for _, v := range obj.Tracks {
		t := Track{}
		err := s.db.First(&t, v).Error
		if err != nil {
			return err
		}
		o.Tracks = append(o.Tracks, t)
	}
	for _, v := range obj.Competitors {
		t := Competitor{}
		err := s.db.First(&t, v).Error
		if err != nil {
			return err
		}
		o.Competitors = append(o.Competitors, t)
	}

	return s.db.Create(&o).Error
}

func (s *SqliteDB) AddCompetitor(ctx context.Context, obj schema.CompetitorSchema) (err error) {
	t := Vehicle{}
	err = s.db.First(&t, obj.Vehicle).Error
	if err != nil {
		return err
	}
	d := Driver{}
	err = s.db.First(&d, obj.Driver).Error
	if err != nil {
		return err
	}
	o := Competitor{
		Vehicle: t,
		Driver:  d,
	}

	return s.db.Create(&o).Error
}
