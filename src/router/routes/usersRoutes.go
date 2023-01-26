package routes

import (
	"api/src/controllers"
	"net/http"
)

var usersRoutes = []Route{
	{
		URI:               "/users",
		Method:            http.MethodPost,
		Function:          controllers.CreateUser,
		hasAuthentication: false,
	},
	{
		URI:               "/users",
		Method:            http.MethodGet,
		Function:          controllers.FindUsers,
		hasAuthentication: true,
	},
	{
		URI:               "/users/{userId}",
		Method:            http.MethodGet,
		Function:          controllers.FindUserById,
		hasAuthentication: true,
	},
	{
		URI:               "/users/{userId}",
		Method:            http.MethodPut,
		Function:          controllers.UpdateUser,
		hasAuthentication: true,
	},
	{
		URI:               "/users/{userId}",
		Method:            http.MethodDelete,
		Function:          controllers.DeleteUser,
		hasAuthentication: true,
	},
	{
		URI:               "/users/{userId}/follower",
		Method:            http.MethodPost,
		Function:          controllers.UserFollower,
		hasAuthentication: true,
	},
	{
		URI:               "/users/{userId}/unfollower",
		Method:            http.MethodPost,
		Function:          controllers.UserUnfollower,
		hasAuthentication: true,
	},
	{
		URI:               "/users/{userId}/followers",
		Method:            http.MethodGet,
		Function:          controllers.FindUsersFollowers,
		hasAuthentication: true,
	},
	{
		URI:               "/users/{userId}/following",
		Method:            http.MethodGet,
		Function:          controllers.Following,
		hasAuthentication: true,
	},
}
