package routers

import (
	"encoding/json"
	"net/http"

	"github.com/balboeng/sergeitter/db"
	"github.com/balboeng/sergeitter/models"
)

/*Register function to create register on db */
func Register(w http.ResponseWriter, r *http.Request) {
	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)
	_,existingUser,_ := db.CheckExistingUser(t.Email)
	_, status, errInsert := db.InsertRegister(t)
	if err != nil {
		http.Error(w, "Error in data "+err.Error(), 400)
		return
	}
	if len(t.Email) == 0{
		http.Error(w, "Email is required", 400)
		return
	}
	if len(t.Password) < 6{
		http.Error(w, "The password should be min 7 chars", 400)
		return
	}
	if existingUser {
		http.Error(w, "User already exist", 400)
		return
	}
	if errInsert != nil || !status {
		http.Error(w, "An error has ocurred"+ errInsert.Error(), 400)
		return
	}
	
	w.WriteHeader(http.StatusCreated)
}
