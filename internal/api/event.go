package api

// EventsContainerDTO represents the exchange format to wrap yokanban events.
type EventsContainerDTO struct {
	Event EventDTO `json:"event"`
}

// EventDTO represents the exchange format to create yokanban events.
type EventDTO struct {
	Type            string      `json:"type"`
	Events          interface{} `json:"events"` // no generics available at the moment :(
	SoftwareVersion string      `json:"softwareVersion"`
}
