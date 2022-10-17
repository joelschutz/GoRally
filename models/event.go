package models

import (
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Name        string
	Class       Class
	Tracks      []Track
	Competitors []Competitor
}

type Class uint8

const (
	LeagueF Class = iota
	LeagueE
	LeagueD
	LeagueC
	LeagueB
	LeagueA
	LeagueS
)

type Competitor struct {
	gorm.Model
	Vehicle Vehicle
	Driver  Driver
}

type TrackResult struct {
	gorm.Model
	TotalTime      float64
	TimeBySegment  []float64
	VStateBySecond []VehicleState
	TStateBySecond []TrackState
}

type EventResults map[Competitor][]TrackResult

type Ranking map[Competitor]float64
