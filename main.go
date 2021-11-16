package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"ghitub/julhan07/redis-go/databases"
	router "ghitub/julhan07/redis-go/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	dbRedis := databases.RedisConnection()

	r.Use(mux.CORSMethodMiddleware(r))

	srv := &http.Server{
		Handler:      router.RunApp(r, dbRedis),
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("service running on port 8080")
	log.Fatal(srv.ListenAndServe())
}
