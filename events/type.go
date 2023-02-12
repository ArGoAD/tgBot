package events

type Fetcher interface {
	Fetch(limit int) ([]Event, error)
}

type Processor interface {
	Processor(e Event) error
}

type Type int

const (
	Unknown Type = iota
	message
)

type Event struct {
	Type Type
	Text string
}
