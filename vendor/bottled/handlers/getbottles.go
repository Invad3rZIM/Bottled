package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
)

//START HERE TOMORROW NIGHT
func (bm *BottleManager) GetLocalBottles(bottleType string, p Point, amount int, idealMaxDistance float64) []*Bottle {
	bs := make([]*Bottle, amount)

	return bs
}

//START HERE TOMORROW NIGHT
func (bm *BottleManager) GetGlobalBottles(bottleType string, p Point, amount int) []*Bottle {
	bs := make([]*Bottle, amount)

	return bs
}

func (h *Handler) GetNearestBottlesHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody map[string]interface{}

	json.NewDecoder(r.Body).Decode(&requestBody)

	//	if err != nil {
	//		err = errors.Trace(err)
	//		core.WriteBadRequestErrorResponse(w)
	//		return
	//	}

	//conver to ints here for
	id, _ := strconv.Atoi(requestBody["userID"].(string))
	amount, _ := strconv.Atoi(requestBody["amount"].(string))
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
		userID:  id,
	}

	var bottles []*Bottle

	if enabled {
		bottles = h.bottleManager.GetLocalBottles(requestBody["bottleType"].(string), p, amount, 10)
	} else {
		bottles = h.bottleManager.GetGlobalBottles(requestBody["bottleType"].(string), p, amount)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// In the future we could report back on the status of our DB, or our cache
	// (e.g. Redis) by performing a simple PING, and include them in the response.

	//s := StatusMessage{
	//	Status: 200,
	//}
	json.NewEncoder(w).Encode(bottles)

	//io.WriteString(w, fmt.Sprintf("Total Bottles: %d \n Just created! %v", len(h.bottleManager.bottles), bottle))
}
