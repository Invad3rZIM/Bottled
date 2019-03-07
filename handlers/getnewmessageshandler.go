package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func WriteBadRequestErrorResponse(e *http.ResponseWriter) {
	w := *e
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)

	s := Status{
		Status:  http.StatusBadRequest,
		Message: "Invalid JSON",
	}
	//return copy of user for client records
	json.NewEncoder(w).Encode(&s)
	return
}

func (h *Handler) GetNewMessagesHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&requestBody)

	if err != nil {
		WriteBadRequestErrorResponse(&w)
		return
	}

	//user verification
	uid, _ := strconv.Atoi(requestBody["userID"].(string))
	//pin, _ := strconv.Atoi(requestBody["pin"].(string))

	//pin logic to implement tomorrow
	//	if h.UserCache.Users[uid].Verify(pin) {

	//	}

	ms := h.MessageCache.GetNewMessages(uid)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	//return copy of user for client records
	json.NewEncoder(w).Encode(&ms)

}
