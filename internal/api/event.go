package api

// EventsContainerDTO represents the exchange format to wrap yokanban events.
type EventsContainerDTO struct {
	Event EventDTO `json:"event"`
}

// EventDTO represents the exchange format to create yokanban events.
type EventDTO struct {
	Type            string            `json:"type"`
	Events          *[]ColumnEventDTO `json:"events"`
	SoftwareVersion string            `json:"softwareVersion"`
}
