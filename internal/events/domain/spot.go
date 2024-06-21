package domain

import (
	"errors"

	"github.com/google/uuid"
)

var (
	formatSpot                         = "XY (X is a letter and Y is a number)"
	ErrSpotInvalidSpotNumber           = errors.New("Spot number is invalid")
	ErrSpotInvalidSpotEmpty            = errors.New("Spot number is empty")
	ErrSpotInvalidSpotFormat           = errors.New("Spot number is not in a valid format: " + formatSpot)
	ErrSpotInvalidSpotFormatFirstChar  = errors.New("Spot first character is not on a valid format: " + formatSpot)
	ErrSpotInvalidSpotFormatSecondChar = errors.New("Spot second character is not on a valid format: " + formatSpot)
	ErrSpotNotFound                    = errors.New("Spot not found")
	ErrSpotAlreadyReserved             = errors.New("Spot already reserved")
)

type SpotStatus string

const (
	SpotStatusAvailable SpotStatus = "available"
	SpotStatusSold      SpotStatus = "Sold"
)

type Spot struct {
	ID       string
	EventID  string
	Name     string
	Status   SpotStatus
	TicketID string
}

func NewSpot(event *Event, name string) (*Spot, error) {
	spot := &Spot{
		ID:      uuid.New().String(),
		EventID: event.ID,
		Name:    name,
		Status:  SpotStatusAvailable,
	}

	// v := spot.Validate()
	// if v != nil {
	// 	return nil, v
	// }

	if err := spot.Validate(); err != nil {
		return nil, err
	}

	return spot, nil
}

func (s Spot) Validate() error {
	if len(s.Name) == 0 {
		return ErrSpotInvalidSpotEmpty
	}
	if len(s.Name) < 2 {
		return ErrSpotInvalidSpotFormat
	}
	if s.Name[0] < 'A' || s.Name[0] > 'Z' {
		return ErrSpotInvalidSpotFormatFirstChar
	}

	if s.Name[1] < '0' || s.Name[1] > '9' {
		return ErrSpotInvalidSpotFormatSecondChar
	}
	return nil
}

func (s *Spot) Reserve(ticketID string) error {
	if s.Status == SpotStatusSold {
		return ErrSpotAlreadyReserved
	}
	s.Status = SpotStatusSold
	s.TicketID = ticketID
	return nil
}
