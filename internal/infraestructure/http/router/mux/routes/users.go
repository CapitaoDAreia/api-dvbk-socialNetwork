package routes

import (
	"api-dvbk-socialNetwork/internal/application/services"
	repository "api-dvbk-socialNetwork/internal/infraestructure/database/repositories"
	"api-dvbk-socialNetwork/internal/infraestructure/http/controllers/usersController"
	"database/sql"
	"net/http"
)

func ConfigRoutes(db *sql.DB) []Route {

	repository := repository.NewUserRepository(db)
	services := services.NewUsersServices(repository)
	controllers := usersController.NewUsersController(services)

	var userRoutes = []Route{
		{
			URI:        "/users",
			Method:     http.MethodPost,
			Controller: controllers.CreateUser,
			NeedAuth:   false,
		},

		{
			URI:        "/users",
			Method:     http.MethodGet,
			Controller: controllers.GetUsers,
			NeedAuth:   true,
		},

		{
			URI:        "/users/{userId}",
			Method:     http.MethodGet,
			Controller: controllers.GetUser,
			NeedAuth:   true,
		},

		{
			URI:        "/users/{userId}",
			Method:     http.MethodPut,
			Controller: controllers.UpdateUser,
			NeedAuth:   true,
		},

		{
			URI:        "/users/{userId}",
			Method:     http.MethodDelete,
			Controller: controllers.DeleteUser,
			NeedAuth:   true,
		},
		{
			URI:        "/users/{userId}/follow",
			Method:     http.MethodPost,
			Controller: controllers.FollowUser,
			NeedAuth:   true,
		},
		{
			URI:        "/users/{userId}/unfollow",
			Method:     http.MethodPost,
			Controller: controllers.UnFollowUser,
			NeedAuth:   true,
		},
		{
			URI:        "/users/{userId}/followers",
			Method:     http.MethodGet,
			Controller: controllers.GetFollowersOfAnUser,
			NeedAuth:   true,
		},
		{
			URI:        "/users/{userId}/following",
			Method:     http.MethodGet,
			Controller: controllers.GetWhoAnUserFollow,
			NeedAuth:   true,
		},
		{
			URI:        "/users/{userId}/update-password",
			Method:     http.MethodPost,
			Controller: controllers.UpdateUserPassword,
			NeedAuth:   true,
		},
	}

	return userRoutes
}