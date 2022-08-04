package handlers

import (
	"log"
	"net/http"

	"github.com/PetoriousBIG/my-go-app/util"
)

type Ping struct {
	l *log.Logger
}

type PingResponse struct {
	Response string `json:"response"`
}

func NewPing(l *log.Logger) *Ping {
	return &Ping{l}
}

func (p *Ping) Get(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("[DEBUG] Ping!")

	response := PingResponse{"Pong!"}

	err := util.ToJSON(response, rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
