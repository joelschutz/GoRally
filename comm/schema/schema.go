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
