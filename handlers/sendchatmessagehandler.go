package handlers

import (
	"bottled/messages"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func (h *Handler) SendChatMessageHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody map[string]interface{}

	json.NewDecoder(r.Body).Decode(&requestBody)

	for i, v := range requestBody {
		fmt.Printf("\n%s %v", i, v)
	}
	//	if err != nil {
	//		err = errors.Trace(err)
	//		core.WriteBadRequestErrorResponse(w)
	//		return
	//	}

	//conver to ints here for
	text := requestBody["text"].(string)

	//user verification
	sid, _ := strconv.Atoi(requestBody["userID"].(string))
	//pin, _ := strconv.Atoi(requestBody["pin"].(string))

	rid, _ := strconv.Atoi(requestBody["receiverID"].(string))

	bid, _ := strconv.Atoi(requestBody["bottleID"].(string))

	chatNum, _ := strconv.Atoi(requestBody["num"].(string))
	//user verification

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

	//implement the canpayfor check on hearts tomorrow

	m := messages.NewMessage(sid, rid, bid, text, chatNum)

	h.MessageCache.ContinueConversation(m)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	//return copy of user for client records
	json.NewEncoder(w).Encode(&m)

}
