package postsController

import (
	"api-dvbk-socialNetwork/internal/infraestructure/database"
	repository "api-dvbk-socialNetwork/internal/infraestructure/database/repositories"
	"api-dvbk-socialNetwork/internal/infraestructure/http/responses"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// --
func GetPost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postID, err := strconv.ParseUint(params["postId"], 10, 64)
	if err != nil {
		responses.FormatResponseToCustomError(w, http.StatusBadRequest, err)
		return
	}

	DB, err := database.ConnectWithDatabase()
	if err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}
	defer DB.Close()

	repository := repository.NewPostsRepository(DB)
	post, err := repository.SearchPost(postID)
	if err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}

	responses.FormatResponseToJSON(w, 200, post)
}
