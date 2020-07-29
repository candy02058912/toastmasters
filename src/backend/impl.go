package svr

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

var bufferSize int = 3

var waitTimeMap = map[string]string{
	"plain":      "3000ms",
	"chocolate":  "2500ms",
	"strawberry": "000ms",
}

type impl struct {
	waitTime  time.Duration
	toastType string
	l         sync.Mutex
	c         chan int
	count     int
}

func (s *impl) h1Handler(w http.ResponseWriter, r *http.Request) {

	s.l.Lock()
	s.count++
	s.l.Unlock()

	tmp := <-s.c

	if s.count > bufferSize {
		log.Fatal("request limit exceed.")
	}

	req := h1request{}
	schema.NewDecoder().Decode(&req, r.URL.Query())

	log.Printf("h1: %v", req)
	time.Sleep(s.waitTime)

	resp := h1response{
		Output:    s.toastType,
		TimeStamp: time.Now().Unix(),
	}

	s.c <- tmp + 1

	s.l.Lock()
	s.count--
	s.l.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

}

// NewServer creates a server implementation for tutorial.
func NewServer(port string, toastType string) http.Server {
	r := mux.NewRouter()

	wt, err := time.ParseDuration(waitTimeMap[toastType])
	if err != nil {
		log.Fatalf("error time duration format %s: %v", waitTimeMap[toastType], err)
	}

	i := impl{
		waitTime:  wt,
		toastType: toastType,
		count:     0,
		c:         make(chan int, bufferSize),
	}
	i.c <- 0

	r.HandleFunc("/h1", i.h1Handler).Methods("GET")

	return http.Server{
		Handler:      r,
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
}
