package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

// RegisterRoutes registers telepathy-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	// this line is used by starport scaffolding # 1
	r.HandleFunc("/telepathy/user", createUserHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/telepathy/user", listUserHandler(cliCtx, "telepathy")).Methods("GET")
	r.HandleFunc("/telepathy/user/{key}", getUserHandler(cliCtx, "telepathy")).Methods("GET")
	r.HandleFunc("/telepathy/user", setUserHandler(cliCtx)).Methods("PUT")
	r.HandleFunc("/telepathy/user", deleteUserHandler(cliCtx)).Methods("DELETE")
	r.HandleFunc("/telepathy/user/{key}/follow", followUserHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/telepathy/user/{key}/unfollow", unfollowUserHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/telepathy/user/{key}/thoughts", listThoughtsByCreatorHandler(cliCtx, "telepathy")).Methods("GET")

	r.HandleFunc("/telepathy/comment", createCommentHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/telepathy/comment", listCommentHandler(cliCtx, "telepathy")).Methods("GET")
	r.HandleFunc("/telepathy/comment/{key}", getCommentHandler(cliCtx, "telepathy")).Methods("GET")
	r.HandleFunc("/telepathy/comment", setCommentHandler(cliCtx)).Methods("PUT")
	r.HandleFunc("/telepathy/comment", deleteCommentHandler(cliCtx)).Methods("DELETE")
	r.HandleFunc("/telepathy/comment/{key}/like", likeCommentHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/telepathy/comment/{key}/dislike", dislikeCommentHandler(cliCtx)).Methods("POST")

	r.HandleFunc("/telepathy/thought", createThoughtHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/telepathy/thought", listThoughtHandler(cliCtx, "telepathy")).Methods("GET")
	r.HandleFunc("/telepathy/thought/{key}", getThoughtHandler(cliCtx, "telepathy")).Methods("GET")
	r.HandleFunc("/telepathy/thought", setThoughtHandler(cliCtx)).Methods("PUT")
	r.HandleFunc("/telepathy/thought", deleteThoughtHandler(cliCtx)).Methods("DELETE")
	r.HandleFunc("telepathy/thought/{key}/like", likeThoughtHandler(cliCtx)).Methods("POST")
	r.HandleFunc("telepathy/thought/{key}/dislike", dislikeThoughtHandler(cliCtx)).Methods("POST")
	r.HandleFunc("telepathy/thought/{key}/comments", listCommentsForThoughtHandler(cliCtx, "telepathy")).Methods("GET")
}
