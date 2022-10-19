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

	_ = json.NewDecoder(r.Body).Decode(&u)
	u.Name = u.Name + " ganteng"

	_ = json.NewEncoder(w).Encode(u)

	w.WriteHeader(http.StatusOK)
}
