package handlers

import (
	"bottled/messages"
	"encoding/json"
	"net/http"
	"strconv"
)

func (h *Handler) SendFirstMessageHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody map[string]interface{}

	json.NewDecoder(r.Body).Decode(&requestBody)

	//	if err != nil {
	//		err = errors.Trace(err)
	//		core.WriteBadRequestErrorResponse(w)
	//		return
	//	}

	//conver to ints here for
	text := requestBody["text"].(string)

	//user verification
	sid, _ := strconv.Atoi(requestBody["userID"].(string))
	pin, _ := strconv.Atoi(requestBody["pin"].(string))

	if !h.UserCache.ValidPin(sid, pin) {

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)

		s := Status{
			Status:  http.StatusUnauthorized,
			Message: "Invalid Pin",
		}
		//return copy of user for client records
		json.NewEncoder(w).Encode(&s)
		return
	}

	bid, _ := strconv.Atoi(requestBody["bottleID"].(string))

	rid, ok := h.BottleCache.BottleSenders[bid]

	if !ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)

		s := Status{
			Status:  http.StatusUnauthorized,
			Message: "Invalid Bottle ID",
		}
		//return copy of user for client records
		json.NewEncoder(w).Encode(&s)
		return
	}

	//check if sufficient hearts...
	err := h.HeartCache.Harm(sid, 4)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)

		s := Status{
			Status:  http.StatusUnauthorized,
			Message: "Insufficient Hearts",
		}
		//return copy of user for client records
		json.NewEncoder(w).Encode(&s)
		return
	}

	m := messages.NewMessage(sid, rid, bid, text, 0)
	h.MessageCache.StartConversation(m)
	h.UserCache.AddCaps(sid, 1)
	h.BottleCache.AllBottles[bid].LoseLife()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	//return copy of user for client records
	json.NewEncoder(w).Encode(&m)

}

type Status struct {
	Status  int
	Message string
}
