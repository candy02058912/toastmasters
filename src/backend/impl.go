package svr

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/gorilla/sessions"
)

var bufferSize int = 1

type impl struct {
	waitTime time.Duration
	c        chan int
}

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

func (s *impl) h1Handler(w http.ResponseWriter, r *http.Request) {
	req := h1request{}
	schema.NewDecoder().Decode(&req, r.URL.Query())

	tmp := <-s.c

	log.Printf("h1, %v: %v", tmp, req)
	time.Sleep(s.waitTime)

	resp := h1response{
		Answer:    req.A + req.B,
		TimeStamp: time.Now().Unix(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

	s.c <- tmp + 1
}

func (s *impl) h2Handler(w http.ResponseWriter, r *http.Request) {
	req := h1request{}
	schema.NewDecoder().Decode(&req, r.URL.Query())

	resp := h2response{}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

}

// NewServer creates a server implementation for tutorial.
func NewServer(port string, waitTime time.Duration) http.Server {
	r := mux.NewRouter()

	i := impl{
		waitTime: waitTime,
		c:        make(chan int, bufferSize),
	}

	for idx := 0; idx < bufferSize; idx++ {
		i.c <- 0
	}

	r.HandleFunc("/h1", i.h1Handler).Methods("GET")
	r.HandleFunc("/h2", i.h2Handler).Methods("GET")

	return http.Server{
		Handler:      r,
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
}
