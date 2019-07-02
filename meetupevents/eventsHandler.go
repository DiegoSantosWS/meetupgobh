package meetupevents

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/DiegoSantosWS/meetupgobh/auth"
	"github.com/DiegoSantosWS/meetupgobh/httputils"
	"github.com/DiegoSantosWS/meetupgobh/utils"

	"github.com/eladmica/go-meetup/meetup"
)

var err error

// HandlerAllEvents return list with all events
func HandlerAllEvents(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		err = errors.New("Method not allowed")
		log.Println(fmt.Sprintf("[meetupevents ] error %s", err))
		http.Error(w, err.Error(), http.StatusMethodNotAllowed)
		return
	}

	client := auth.Authetication()
	param := meetup.GetEventsParams{
		Page:   2,
		Scroll: "future_or_past",
		Desc:   false,
	}
	events, err := client.GetEvents("go-belo-horizonte", &param)
	if err != nil {
		log.Println(fmt.Sprintf("[main meetup go bh] ocorreu um erro. error [%s]", err))
	}

	mtup := []meetups{}
	for _, ev := range events {
		evs := meetups{
			Name:        ev.Name,
			Description: ev.Description,
			EventDate:   utils.DateConvertUnix(int64(ev.Time), "RFC3339"),
			Link:        ev.Link,
			LimitRSVP:   int64(ev.RSVPLimit),
			YesRSVP:     int64(ev.YesRSVPCount),
		}
		mtup = append(mtup, evs)
	}

	meetp, err := json.Marshal(mtup)
	if err != nil {
		http.Error(w, err.Error(), http.StatusPreconditionFailed)
		return
	}
	httputils.WriteJSON(w, meetp)
}
