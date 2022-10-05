package users

import (
	"encoding/json"
	"github.com/coba/model"
	"net/http"
)

func HandlerUsersAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var u model.Users

	json.NewDecoder(r.Body).Decode(&u)
	u.Name = u.Name + " ganteng"

	json.NewEncoder(w).Encode(u)

	w.WriteHeader(http.StatusOK)
}
