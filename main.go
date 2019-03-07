// endpoints.go
package main

import (
	"bottled/handlers"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"bottled/database"
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

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "your-password"
	dbname   = "calhounio_demo"
)

func main() {

	//dbUrl := os.Getenv(DATABASE_UR
	//	db, err := sql.Open("postgres", "kirk:@/localhost")

	db, err := sql.Open("postgres", "user=kirk password=testing123 dbname=bottled sslmode=disable")
	//	db, err := sql.Open("postgres", dbUrl)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	d := database.NewDatabaseConnection(db)

	fmt.Printf("\n\n%v", &d)

	r := mux.NewRouter()

	h := handlers.NewHandler(d)

	//launch launches the background goroutine loop functions
	go h.HeartCache.Launch()
	go h.BottleCache.Launch()

	//go h.MedicManager(10, 60)

	//users := make(map[string]*handlers.User)
	//	users["kirk"] = h.CreateUser("kirk")
	//	users["rob"] = h.CreateUser("rob")

	//u := handlers.NewUser("danny")
	//	h.Hurt(users["kirk"].GetUserID(), 12)
	//	h.Hurt(users["rob"].GetUserID(), 5)
	//	h.Hurt(users["rob"].GetUserID(), 5)

	r.HandleFunc("/alive", ServerAliveHandler)
	r.HandleFunc("/", Ping)

	r.HandleFunc("/bottle/create", h.CreateBottleHandler).Methods("POST")
	r.HandleFunc("/bottle/receive", h.GetBottlesHandler).Methods("GET")

	r.HandleFunc("/user/create", h.CreateUserHandler).Methods("POST")

	r.HandleFunc("/messages/init", h.SendFirstMessageHandler).Methods("POST")
	r.HandleFunc("/messages/send", h.SendChatMessageHandler).Methods("POST")

	r.HandleFunc("/befriend", h.GiveFriendKey).Methods("POST")

	r.HandleFunc("/messages/getnew", h.GetNewMessagesHandler).Methods("GET")

	err = http.ListenAndServe(":"+os.Getenv("PORT"), r)
	if err != nil {
		panic(err)
	}

	for {
	}
	//	for {
	//		for _, v := range users {
	//			fmt.Printf("%s hearts = %.2f\n", v.GetName(), v.CurrentLives())
	//		}
	//		time.Sleep(60000 * time.Millisecond)
	//	}*/
}
