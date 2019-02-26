package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type BottleManager struct {
	bottles map[int]*Bottle
}

func NewBottleManager() *BottleManager {
	bm := BottleManager{
		bottles: make(map[int]*Bottle),
	}

	return &bm
}

func (bm *BottleManager) AcceptBottle(id int) {
	b := bm.bottles[id]
	b.LoseLife(1) //accepting a bottle costs 1 life
}

func (bm *BottleManager) CreateBottle(senderID int, message string, bottleType string, lives int, point Point) *Bottle {
	b := NewBottle(senderID, message, bottleType, lives)
	b.AddLocation(point)

	bm.bottles[b.bottleID] = b

	fmt.Println(b)

	return b
}

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
	io.WriteString(w, fmt.Sprintf("Total Bottles: %d \n Just created! %v", len(h.bottleManager.bottles), bottle))
}
