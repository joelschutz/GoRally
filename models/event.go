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
