package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func (h *Handler) CreateBottleHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody map[string]interface{}

	json.NewDecoder(r.Body).Decode(&requestBody)

	//	if err != nil {
	//		err = errors.Trace(err)
	//		core.WriteBadRequestErrorResponse(w)
	//		return
	//	}

	//conver to ints here for
	senderID, _ := strconv.Atoi(requestBody["userID"].(string))
	bottleLives, _ := strconv.Atoi(requestBody["bottleLife"].(string))
	enabled := false
	var long, lat float64

	if requestBody["enabled"].(string) == "true" {
		enabled = true
		lat, _ = strconv.ParseFloat(requestBody["lat"].(string), 64)
		long, _ = strconv.ParseFloat(requestBody["long"].(string), 64)
	}

	p := Point{
		enabled: enabled,
		lat:     lat,
		long:    long,
		userID:  senderID,
	}

	bottle := h.bottleManager.CreateBottle(senderID, requestBody["message"].(string), requestBody["bottleType"].(string), bottleLives, p)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// In the future we could report back on the status of our DB, or our cache
	// (e.g. Redis) by performing a simple PING, and include them in the response.

	s := StatusMessage{
		Status: 200,
	}
	json.NewEncoder(w).Encode(s)

	io.WriteString(w, fmt.Sprintf("Total Bottles: %d \n Just created! %v", len(h.bottleManager.localBottles)+len(h.bottleManager.globalBottles), bottle))
}

func (bm *BottleManager) CreateBottle(senderID int, message string, bottleType string, lives int, point Point) *Bottle {
	b := NewBottle(senderID, message, bottleType, lives, bm.bottlesMade)

	if point.enabled {
		b.AddLocation(point)

		bm.localBottles[b.bottleID] = b
	} else {
		bm.globalBottles[b.bottleID] = b
	}

	bm.bottlesMade = bm.bottlesMade + 1

	return b
}
