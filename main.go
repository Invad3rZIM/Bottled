// endpoints.go
package main

import (
	"bottled/handlers"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

func ServerAliveHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody map[string]interface{}
	// A very simple health check.

	//Grab and decode the request body from the submitted post body
	json.NewDecoder(r.Body).Decode(&requestBody)

	//	if err != nil {
	//		err = errors.Trace(err)
	//		core.WriteBadRequestErrorResponse(w)
	//		return
	//	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// In the future we could report back on the status of our DB, or our cache
	// (e.g. Redis) by performing a simple PING, and include them in the response.
	io.WriteString(w, requestBody["id"].(string))
}

func Ping(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Pong")
}

func main() {
	r := mux.NewRouter()

	h := handlers.NewHandler()
	go h.MedicManager(10, 60)

	users := make(map[string]*handlers.User)
	users["kirk"] = h.CreateUser("kirk")
	users["rob"] = h.CreateUser("rob")

	//u := handlers.NewUser("danny")
	h.Hurt(users["kirk"].GetUserID(), 12)
	h.Hurt(users["rob"].GetUserID(), 5)
	//	h.Hurt(users["rob"].GetUserID(), 5)

	r.HandleFunc("/alive", ServerAliveHandler)
	r.HandleFunc("/", Ping)
	r.HandleFunc("/bottle/create", h.CreateBottleHandler).Methods("POST")

	r.HandleFunc("/bottle/receive", h.CreateBottleHandler).Methods("GET")

	err := http.ListenAndServe(":"+os.Getenv("PORT"), r)
	if err != nil {
		panic(err)
	}
	for {
		for _, v := range users {
			fmt.Printf("%s hearts = %.2f\n", v.GetName(), v.CurrentLives())
		}
		time.Sleep(60000 * time.Millisecond)
	}
}
