package handlers

import (
	bottle "bottled/bottles/bottle"
	"bottled/points"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func (h *Handler) CreateBottleHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&requestBody)

	if err != nil {
		WriteBadRequestErrorResponse(&w)
		return
	}

	//user verification
	sid, _ := strconv.Atoi(requestBody["sid"].(string))
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

	//check if sufficient hearts...

	err = h.HeartCache.Harm(sid, 4)

	err = nil
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

	enabled := false
	var long, lat float64

	if requestBody["enabled"].(string) == "true" {
		enabled = true
		lat, _ = strconv.ParseFloat(requestBody["lat"].(string), 64)
		long, _ = strconv.ParseFloat(requestBody["long"].(string), 64)
	}

	//bottle := h.BottleManager.CreateBottle(senderID, requestBody["message"].(string), requestBody["bottleType"].(string), bottleLives, p)

	bottle := bottle.Bottle{
		SenderID: sid,
		Tag:      requestBody["tag"].(string),
		Message:  requestBody["message"].(string),
		Lives:    3,
		Age:      (h.BottleCache).BottlesMade,
		Point: points.Point{
			Enabled: enabled,
			Lat:     lat,
			Long:    long,
			Age:     h.BottleCache.BottlesMade,
		},
	}

	//h.BottleCache.DB.AddBottle(&bottle)

	bottle.Point.BottleID = bottle.BottleID
	//list of EVERY SINGLE BOTTLE
	h.BottleCache.AllBottles[bottle.BottleID] = &bottle

	if bottle.Point.Enabled {
		h.BottleCache.Bottles["local"][bottle.Tag][bottle.BottleID] = &bottle
		fmt.Printf("\n%v", h.BottleCache.Bottles["local"][bottle.Tag][bottle.BottleID])
	} else {
		h.BottleCache.Bottles["global"][bottle.Tag][bottle.BottleID] = &bottle
	}

	fmt.Println("IIIIIIIII")
	h.BottleCache.BottlesMade = h.BottleCache.BottlesMade + 1

	h.BottleCache.BottleSenders[bottle.BottleID] = bottle.SenderID

	h.UserCache.AddCaps(sid, 1)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(&bottle)
	// In the future we could report back on the status of our DB, or our cache
	// (e.g. Redis) by performing a simple PING, and include them in the response.

}
