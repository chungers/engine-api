// +build experimental

package infra

import (
	"encoding/json"
)

// URLString is a URL string
type URLString string

// Config defines a SWIM config
type Config struct {
	// Groups is a map keyed by the name of the group
	Groups map[string]Group `json:",omitempty"`

	// Addresses is an association of url string to addresses
	Addresses map[URLString]Address `json:",omitempty"`
}

// Group defines a scale group
type Group struct {

	// Plugin is the name of the plugin (user/repo) that implements a GroupDriver interface.
	Plugin *string `json:",omitempty"`

	// Properties capture the free-form properties for the plugin GroupDriver
	Properties *json.RawMessage `json:",omitempty"`
}

// Addresse models an endpoint for the services running in a swarm
type Address struct {
	// Service is the swarm service name
	Service string `json:",omitempty"`

	// TODO - add L4, L7 load balancer configs
}
