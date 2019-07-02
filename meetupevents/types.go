package meetupevents

type meetups struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	EventDate   string `json:"event_date"`
	Link        string `json:"link"`
	LimitRSVP   int64  `json:"rsvp_limit"`
	YesRSVP     int64  `json:"yes_rsvp_count"`
}
