package peoplez

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	model "github.com/seggga/back2/internal/peoplez/model"
)

// HomeHandler sends homepage
func homeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, "Welcome to PEOPLEZ")
	}
}

// NewUserHandler adds a new user
func newUserHandler(stor model.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		decoder := json.NewDecoder(req.Body)
		var u model.User
		err := decoder.Decode(&u)
		if err != nil {
			err = fmt.Errorf("unable to parse request body %w", err)
			JSONError(w, err, http.StatusBadRequest)
			return
		}
		user := model.User{
			ID:      uuid.New(),
			Name:    u.Name,
			Surname: u.Surname,
		}

		err = stor.AddUser(user)
		if err != nil {
			err = fmt.Errorf("error creating new user %w", err)
			JSONError(w, err, http.StatusBadRequest)
			return
		}

		w.Header().Set("Application", "Peoplez")
		w.WriteHeader(http.StatusCreated)
	}
}

// NewUnionHandler adds a new user
func newUnionHandler(stor model.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		decoder := json.NewDecoder(req.Body)
		var u model.Union
		err := decoder.Decode(&u)
		if err != nil {
			err = fmt.Errorf("unable to parse request body %w", err)
			JSONError(w, err, http.StatusBadRequest)
			return
		}
		union := model.Union{
			ID:          uuid.New(),
			Name:        u.Name,
			Aim:         u.Aim,
			Contact:     u.Contact,
			Manager:     u.Manager,
			Description: u.Description,
		}

		err = stor.AddUnion(union)
		if err != nil {
			err = fmt.Errorf("error creating new union %w", err)
			JSONError(w, err, http.StatusBadRequest)
			return
		}

		w.Header().Set("Application", "Peoplez")
		w.WriteHeader(http.StatusCreated)
	}
}
