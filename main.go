// endpoints.go
package main

import (
	"bottled/handlers"
	"io"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func ServerAliveHandler(w http.ResponseWriter, r *http.Request) {
	//	var requestBody map[string]interface{}
	// A very simple health check.

	//Grab and decode the request body from the submitted post body
	//	err := json.NewDecoder(r.Body).Decode(&requestBody)

	//	if err != nil {
	//		err = errors.Trace(err)
	//		core.WriteBadRequestErrorResponse(w)
	//		return
	//	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// In the future we could report back on the status of our DB, or our cache
	// (e.g. Redis) by performing a simple PING, and include them in the response.
	io.WriteString(w, `this is a test`)
}

func main() {
	r := mux.NewRouter()

	h := handlers.NewHandler()
	go h.MedicManager(10)

	users := make(map[string]*handlers.User)
	users["kirk"] = h.CreateUser("kirk")
	users["rob"] = h.CreateUser("rob")

	h.Hurt(users["kirk"].GetUserID(), 12)

	//	h.Hurt(users["rob"].GetUserID(), 5)

	/*	for {
		for _, v := range users {
			fmt.Printf("%s hearts = %.2f\n", v.GetName(), v.CurrentLives())
		}
		time.Sleep(500 * time.Millisecond)
	}*/

	r.HandleFunc("/alive", ServerAliveHandler)

	err := http.ListenAndServe(":"+os.Getenv("PORT"), r)
	if err != nil {
		panic(err)
	}
}
