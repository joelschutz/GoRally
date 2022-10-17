package storage

import (
	"context"

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
	err = s.db.First(&v, index).Error
	return
}

func (s *SqliteDB) GetDriver(ctx context.Context, index uint) (d models.Driver, err error) {
	err = s.db.First(&d, index).Error
	return
}

func (s *SqliteDB) GetTrack(ctx context.Context, index uint) (t models.Track, err error) {
	err = s.db.First(&t, index).Error
	return
}

func (s *SqliteDB) GetEvent(ctx context.Context, index uint) (e models.Event, err error) {
	err = s.db.First(&e, index).Error
	return
}

func (s *SqliteDB) ListVehicles(ctx context.Context) (e []models.Vehicle, err error) {
	err = s.db.Find(&e).Error
	return
}

func (s *SqliteDB) ListDrivers(ctx context.Context) (e []models.Driver, err error) {
	err = s.db.Find(&e).Error
	return
}

func (s *SqliteDB) ListTracks(ctx context.Context) (e []models.Track, err error) {
	err = s.db.Find(&e).Error
	return
}

func (s *SqliteDB) ListEvents(ctx context.Context) (e []models.Event, err error) {
	err = s.db.Find(&e).Error
	return
}

func (s *SqliteDB) UpdateVehicle(ctx context.Context, index uint, obj models.Vehicle) error {
	o, err := s.GetVehicle(ctx, index)
	if err != nil {
		return err
	}
	return s.db.Model(&o).Updates(obj).Error
}

func (s *SqliteDB) UpdateDriver(ctx context.Context, index uint, obj models.Driver) error {
	o, err := s.GetDriver(ctx, index)
	if err != nil {
		return err
	}
	return s.db.Model(&o).Updates(obj).Error
}

func (s *SqliteDB) UpdateTrack(ctx context.Context, index uint, obj models.Track) error {
	o, err := s.GetTrack(ctx, index)
	if err != nil {
		return err
	}
	return s.db.Model(&o).Updates(obj).Error
}

func (s *SqliteDB) UpdateEvent(ctx context.Context, index uint, obj models.Event) error {
	o, err := s.GetEvent(ctx, index)
	if err != nil {
		return err
	}
	return s.db.Model(&o).Updates(obj).Error
}

func (s *SqliteDB) AddVehicle(ctx context.Context, obj models.Vehicle) (err error) {
	err = s.db.Create(&obj).Error
	return
}

func (s *SqliteDB) AddDriver(ctx context.Context, obj models.Driver) (err error) {
	err = s.db.Create(&obj).Error
	return
}

func (s *SqliteDB) AddTrack(ctx context.Context, obj models.Track) (err error) {
	err = s.db.Create(&obj).Error
	return
}

func (s *SqliteDB) AddEvent(ctx context.Context, obj models.Event) (err error) {
	err = s.db.Create(&obj).Error
	return
}
