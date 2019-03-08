package handlers

import (
	"bottled/hearts/heart"
	users "bottled/users/user"
	"encoding/json"
	"net/http"
)

/*
func (h *Handler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&requestBody)

	if err != nil {
		fmt.Println(err.Error())
		//	WriteBadRequestErrorResponse(&w)
		return
	}

	this is how i type my words like this i'm unsure what

	//conver to ints here for
	name := requestBody["name"].(string)

	user := users.NewUser(name, 0, 3, 0, 0, 0)
	h.UserCache.Users[user.UserID] = user

	//add user to database
	h.DB.AddUser(user)
	h.DB.AddHeart(heart.NewHeart(user))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	//return copy of user for client records
	json.NewEncoder(w).Encode(&user)
}*/

func (h *Handler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&requestBody)

	if err != nil {
		WriteBadRequestErrorResponse(&w)
		return
	}

	//conver to ints here for
	name := requestBody["name"].(string)

	user := users.NewUser(name, 0, 3, 0, 0, 0)
	h.UserCache.Users[user.UserID] = user

	//add user to database
	h.DB.AddUser(user)
	h.DB.AddHeart(heart.NewHeart(user))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	//return copy of user for client records
	json.NewEncoder(w).Encode(&user)
}
