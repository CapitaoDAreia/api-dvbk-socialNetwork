package usersController

import (
	"api-dvbk-socialNetwork/src/auth"
	"api-dvbk-socialNetwork/src/database"
	"api-dvbk-socialNetwork/src/repository"
	"api-dvbk-socialNetwork/src/responses"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func FollowUser(w http.ResponseWriter, r *http.Request) {
	followerID, err := auth.ExtractUserID(r)
	if err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}

	parameters := mux.Vars(r)
	followedID, err := strconv.ParseUint(parameters["userId"], 10, 64)
	if err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}

	if followedID == followerID {
		responses.FormatResponseToCustomError(w, 500, errors.New("Do you want to follow yourself? Pff! "))
		return
	}

	DB, err := database.ConnectWithDatabase()
	if err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}
	defer DB.Close()

	repository := repository.NewUserRepository(DB)

	if err := repository.Follow(followedID, followerID); err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}

	responses.FormatResponseToJSON(w, http.StatusNoContent, nil)
}
