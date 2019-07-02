package router

import (
	"log"
	"net/http"

	"github.com/DiegoSantosWS/meetupgobh/meetupevents"
	"github.com/GuilhermeCaruso/bellt"
)

// Router route to return data of events
func Router() {
	router := bellt.NewRouter()

	router.HandleGroup("/v1",
		router.SubHandleFunc("/allevents", meetupevents.HandlerAllEvents, "GET"),
	)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Println("algum erro aconteceu", err)
	}
}
