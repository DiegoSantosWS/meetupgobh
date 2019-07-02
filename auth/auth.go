package auth

import (
	"github.com/eladmica/go-meetup/meetup"
)

// Authetication faz autenticação com meetup
func Authetication() *meetup.Client {
	client := meetup.NewClient(nil)
	client.Authentication = meetup.NewKeyAuth("<YOUR_API_KEY>")
	return client
}
