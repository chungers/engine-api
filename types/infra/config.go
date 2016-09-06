package infra

import (
	"encoding/json"
)

// Config defines a SWIM config
type Config struct {
	Groups []Group `json:",omitempty"`
}

// Group defines a scale group
type Group struct {
	Name       string           `json:",omitempty"`
	Driver     *string          `json:",omitempty"`
	Properties *json.RawMessage `json:",omitempty"`
}
