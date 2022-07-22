package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
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

	err := response.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (pr *PingResponse) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)

	return e.Encode(pr)
}
