package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (h *Handler) GiveFriendKey(w http.ResponseWriter, r *http.Request) {
	var requestBody map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&requestBody)

	if err != nil {
		WriteBadRequestErrorResponse(&w)
		return
	}

	//user verification
	sid, _ := strconv.Atoi(requestBody["sid"].(string))
	pin, _ := strconv.Atoi(requestBody["pin"].(string))
	rid, _ := strconv.Atoi(requestBody["rid"].(string))

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

	//give friendkey if the user has sufficient (1) caps to give...
	err = h.UserCache.GiveFriendKey(sid, rid)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)

		s := Status{
			Status:  http.StatusUnauthorized,
			Message: "Insufficient Caps",
		}
		//return copy of user for client records
		json.NewEncoder(w).Encode(&s)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
