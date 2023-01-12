package schema

import "encoding/json"

type Payload struct {
	Action Action          `json:"action"`
	Data   json.RawMessage `json:"data"`
}

type Action struct {
	Target string `json:"target"`
	Method string `json:"method"`
	Index  uint   `json:"index"`
}

type EventSchema struct {
	Name        string `json:"name"`
	Class       uint8  `json:"calss"`
	Tracks      []uint `json:"tracks"`
	Competitors []uint `json:"competitors"`
}

type CompetitorSchema struct {
	Vehicle uint `json:"vehicle"`
	Driver  uint `json:"driver"`
}
