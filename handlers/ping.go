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

func (p *Ping) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Ping")

	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Oops", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "Pong %s", d)

}
