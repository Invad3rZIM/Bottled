package handlers

import (
	bottle "bottled/bottles/bottle"
	"bottled/points"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func (h *Handler) GetBottlesHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&requestBody)

	if err != nil {
		WriteBadRequestErrorResponse(&w)
		return
	}

	//user verification
	uid, _ := strconv.Atoi(requestBody["uid"].(string))
	pin, _ := strconv.Atoi(requestBody["pin"].(string))

	if !h.UserCache.ValidPin(uid, pin) {

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

	//conver to ints here for
	amount, _ := strconv.Atoi(requestBody["amount"].(string))
	enabled := false
	var long, lat float64

	if requestBody["enabled"].(string) == "true" {
		enabled = true
		lat, _ = strconv.ParseFloat(requestBody["lat"].(string), 64)
		long, _ = strconv.ParseFloat(requestBody["long"].(string), 64)
	}

	p := points.Point{
		Enabled: enabled,
		Lat:     lat,
		Long:    long,
	}

	var bottles []*bottle.Bottle

	if enabled {
		bottles = h.BottleCache.GetLocalBottles(requestBody["tag"].(string), p, amount, 10)
	} else {
		bottles = h.BottleCache.GetGlobalBottles(requestBody["tag"].(string), p, amount)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// In the future we could report back on the status of our DB, or our cache
	// (e.g. Redis) by performing a simple PING, and include them in the response.

	//s := StatusMessage{
	//	Status: 200,
	//}

	fmt.Print(bottles)
	json.NewEncoder(w).Encode(&bottles)

	//io.WriteString(w, fmt.Sprintf("Total Bottles: %d \n Just created! %v", len(h.bottleManager.bottles), bottle))
}
