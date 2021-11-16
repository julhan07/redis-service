package http

import (
	"ghitub/julhan07/redis-go/controllers"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
)

func RunApp(r *mux.Router, redis *redis.Client) *mux.Router {

	user := controllers.NewUserController(redis)

	r.HandleFunc("/user/list", user.GetAll).Methods("GET")

	return r
}
