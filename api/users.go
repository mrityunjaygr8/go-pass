package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mrityunjaygr8/go-pass/users"
)

func (a *App) getAllUsers(w http.ResponseWriter, r *http.Request) {
	var all_users []users.User
	all_users, err := users.ListUsers(a.DB)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, all_users)
}

func (a *App) createUser(w http.ResponseWriter, r *http.Request) {
	var u users.User
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&u); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := u.CreateUser(a.DB); err != nil {
		if err.Error() == users.USER_EXISTS {
			respondWithError(w, http.StatusBadRequest, err.Error())
		} else {
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusCreated, u)
}

func (a *App) fetchUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	u := users.User{ID: uint(id)}

	err = u.FetchUser(a.DB)
	if err != nil {
		if err.Error() == users.USER_NOT_EXISTS {
			respondWithError(w, http.StatusNotFound, err.Error())
		} else {
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	respondWithJSON(w, http.StatusOK, u)
}

func (a *App) updateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	u := users.User{ID: uint(id)}
	err = u.FetchUser(a.DB)

	if err != nil {
		if err.Error() == users.USER_NOT_EXISTS {
			respondWithError(w, http.StatusNotFound, err.Error())
		} else {
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	decoder := json.NewDecoder(r.Body)
	if err = decoder.Decode(&u); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = u.UpdateUser(a.DB)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	respondWithJSON(w, http.StatusCreated, u)
}

func (a *App) deleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	u := users.User{ID: uint(id)}
	err = u.FetchUser(a.DB)
	if err != nil {
		if err.Error() == users.USER_NOT_EXISTS {
			respondWithError(w, http.StatusNotFound, err.Error())
		} else {
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	err = u.DeleteUser(a.DB)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusNoContent, fmt.Sprintf("%d deleted successfully", id))

}
