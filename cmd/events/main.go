package main

import (
	"database/sql"
	"net/http"

	"m2bops.com/go/internal/events/infra/repository"
	"m2bops.com/go/internal/events/infra/service"
	"m2bops.com/go/internal/events/usecase"

	httpHandler "m2bops.com/go/internal/events/infra/http"
)

func main() {
	// I CANT DO DATABASE FROM JSON FILE
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/test_db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	eventRepo, err := repository.NewMysqlEventRepository(db)
	if err != nil {
		panic(err)
	}

	partnerBaseURLs := map[int]string{
		1: "http://localhost:9080/api1",
		2: "http://localhost:9080/api2",
	}

	partnerFactory := service.NewPartnerFactory(partnerBaseURLs)

	listEventsUseCase := usecase.NewListEventsUseCase(eventRepo)
	getEventsUseCase := usecase.NewGetEventsUseCase(eventRepo)
	listSpotsUseCase := usecase.NewListSpotsUseCase(eventRepo)
	buyTicketsUseCase := usecase.NewBuyTicketsUseCase(eventRepo, partnerFactory)

	eventsHandler := httpHandler.NewEventHandler(
		listEventsUseCase,
		listSpotsUseCase,
		getEventsUseCase,
		buyTicketsUseCase,
	)

	r := http.NewServeMux()
	r.HandleFunc("GET /events", eventsHandler.ListsEvents)
	r.HandleFunc("GET /events/{eventID}", eventsHandler.GetEvents)
	r.HandleFunc("GET /events/{eventID}/spots", eventsHandler.ListSpots)
	r.HandleFunc("POST /checkout", eventsHandler.BuyTicket)

	http.ListenAndServe(":8080", r)

}
