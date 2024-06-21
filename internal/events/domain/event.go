package domain

import (
	"errors"
	"time"
)

var (
	ErrEventNameRequired  = errors.New("Event name is required")
	ErrEventDateFuture    = errors.New("Event date needs to be in a future")
	ErrEventCapacityEmpty = errors.New("Event needs to have a capacity")
	ErrEventpriceEmpty    = errors.New("Event needs to have a price")
	ErrEventNotFound      = errors.New("Event not found")
)

type Rating string

const (
	RatingLivre Rating = "L"
	Rating10    Rating = "10"
	Rating12    Rating = "12"
	Rating14    Rating = "14"
	Rating16    Rating = "16"
	Rating18    Rating = "18"
)

type Event struct {
	ID           string
	Name         string
	Location     string
	Organization string
	Rating       Rating
	Date         time.Time
	ImageURL     string
	Capacity     int
	Price        float64
	PartnerID    int
	Spots        []Spot
	Tickets      []Ticket
}

func (e Event) Validate() error {
	if e.Name == "" {
		return ErrEventNameRequired
	}
	if e.Date.Before(time.Now()) {
		return ErrEventDateFuture
	}
	if e.Capacity <= 0 {
		return ErrEventCapacityEmpty
	}
	if e.Price < 0 {
		return ErrEventpriceEmpty
	}
	return nil
}

func (e *Event) AddSpot(name string) (*Spot, error) {
	spot, err := NewSpot(e, name)
	if err != nil {
		return nil, err
	}
	e.Spots = append(e.Spots, *spot)
	return spot, nil
}
