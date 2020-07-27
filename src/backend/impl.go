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

var waitTimeMap = map[string]string{
	"plain":      "1000ms",
	"chocolate":  "2500ms",
	"strawberry": "3000ms",
}

type impl struct {
	waitTime  time.Duration
	toastType string
	c         chan int
}

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

func (s *impl) h1Handler(w http.ResponseWriter, r *http.Request) {
	// occupy some memory for no reason.
	var a [100000]int
	a[2] = 1

	req := h1request{}
	schema.NewDecoder().Decode(&req, r.URL.Query())

	tmp := <-s.c

	log.Printf("h1, %v: %v", tmp, req)
	time.Sleep(s.waitTime)

	resp := h1response{
		Output:    s.toastType,
		TimeStamp: time.Now().Unix(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

	s.c <- tmp + 1 + a[3]
}

func (s *impl) h2Handler(w http.ResponseWriter, r *http.Request) {
	req := h1request{}
	schema.NewDecoder().Decode(&req, r.URL.Query())

	resp := h2response{}

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
		c:         make(chan int, bufferSize),
	}

	for idx := 0; idx < bufferSize; idx++ {
		i.c <- idx
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
