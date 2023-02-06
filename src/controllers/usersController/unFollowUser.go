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

func UnFollowUser(w http.ResponseWriter, r *http.Request) {
	followerID, err := auth.ExtractUserID(r)
	if err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}

	parameters := mux.Vars(r)

	followedID, err := strconv.ParseUint(parameters["userId"], 10, 64)
	if err != nil {
		responses.FormatResponseToCustomError(w, http.StatusBadRequest, err)
		return
	}

	if followedID == followerID {
		responses.FormatResponseToCustomError(w, http.StatusForbidden, errors.New("You are fated to follow yourself forever!"))
		return
	}

	DB, err := database.ConnectWithDatabase()
	if err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}

	repository := repository.NewUserRepository(DB)

	if err := repository.UnFollow(followedID, followerID); err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}

	responses.FormatResponseToJSON(w, http.StatusNoContent, nil)
}
