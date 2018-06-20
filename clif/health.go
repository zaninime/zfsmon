package clif

import "errors"

type Health int

const (
	Online Health = iota
	Degraded
	Faulted
	Offline
	Unavail
	Removed
)

var (
	ErrUnknownValue = errors.New("Unknown value")
)

func NewHealthFromCliOutput(s string) (*Health, error) {
	healthMap := map[string]Health{
		"ONLINE":   Online,
		"DEGRADED": Degraded,
		"FAULTED":  Faulted,
		"OFFLINE":  Offline,
		"UNAVAIL":  Unavail,
		"REMOVED":  Removed,
	}

	health, ok := healthMap[s]

	if !ok {
		return nil, ErrUnknownValue
	}

	return &health, nil
}
