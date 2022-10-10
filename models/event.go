package models

type Event struct {
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
	Vehicle Vehicle
	Racer   Racer
}

type TrackResult struct {
	TotalTime     float64
	TimeBySegment []float64
}

type EventResults struct {
	TotalTimes map[Competitor][]TrackResult
}

type Ranking map[Competitor]float64
