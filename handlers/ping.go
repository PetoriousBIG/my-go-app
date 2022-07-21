package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Ping struct {
	l *log.Logger
}

func NewPing(l *log.Logger) *Ping {
	return &Ping{l}
}

func (p *Ping) Get(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("[DEBUG] Ping!")

	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		p.l.Println("[ERROR] pinging")
		http.Error(rw, "Oops", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "Pong %s", d)

}
