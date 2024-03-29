package routes

import (
	"api/src/controllers"
	"net/http"
)

var postsRoutes = []Route{
	{
		URI:               "/posts",
		Method:            http.MethodPost,
		Function:          controllers.CreatePost,
		hasAuthentication: true,
	},
	{
		URI:               "/posts",
		Method:            http.MethodGet,
		Function:          controllers.FindAllPosts,
		hasAuthentication: true,
	},
	{
		URI:               "/posts/{postId}",
		Method:            http.MethodGet,
		Function:          controllers.FindPost,
		hasAuthentication: true,
	},
	{
		URI:               "/posts/{postId}",
		Method:            http.MethodPut,
		Function:          controllers.UpdatePost,
		hasAuthentication: true,
	},
	{
		URI:               "/posts/{postId}",
		Method:            http.MethodDelete,
		Function:          controllers.DeletePost,
		hasAuthentication: true,
	},
	{
		URI:               "/users/{userId}/posts",
		Method:            http.MethodGet,
		Function:          controllers.FindAllPostsFromUser,
		hasAuthentication: true,
	},
	{
		URI:               "/posts/{postId}/like",
		Method:            http.MethodPost,
		Function:          controllers.LikePost,
		hasAuthentication: true,
	},
	{
		URI:               "/posts/{postId}/unlike",
		Method:            http.MethodPost,
		Function:          controllers.UnlikePost,
		hasAuthentication: true,
	},
}
