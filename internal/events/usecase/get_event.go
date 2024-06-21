package usecase

import "m2bops.com/go/internal/events/domain"

type GetEventInputDTO struct {
	ID string
}

type GetEventOutputDTO struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Location     string  `json:"location"`
	Organization string  `json:"organization"`
	Rating       string  `json:"rating"`
	Date         string  `json:"date"`
	Capacity     int     `json:"capacity"`
	Price        float64 `json:"price"`
	PartnerID    int     `json:"partner_id"`
}

type GetEventsUseCase struct {
	repo domain.EventRepository
}

func NewGetEventsUseCase(repo domain.EventRepository) *GetEventsUseCase {
	return &GetEventsUseCase{repo: repo}
}

func (uc *GetEventsUseCase) Execute(input GetEventInputDTO) (*GetEventOutputDTO, error) {

	event, err := uc.repo.FindEventByID(input.ID)
	if err != nil {
		return nil, err
	}

	return &GetEventOutputDTO{
		ID:           event.ID,
		Name:         event.Name,
		Location:     event.Location,
		Organization: event.Organization,
		Rating:       string(event.Rating),
		Date:         event.Date.Format("1234-12-12 12:12:12"),
		Capacity:     event.Capacity,
		Price:        event.Price,
		PartnerID:    event.PartnerID,
	}, nil
}
