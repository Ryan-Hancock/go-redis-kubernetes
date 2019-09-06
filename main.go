package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis"
)

//KEY redis key for operation.
const KEY = "kcount"

var (
	rClient *redis.Client
)

func errorHandler(w http.ResponseWriter, err string, httpCode int) {
	http.Error(
		w,
		err,
		httpCode,
	)
}

func handle(w http.ResponseWriter, r *http.Request) {
	rClient.Incr(KEY)

	val, err := rClient.Get(KEY).Result()
	if err != nil {
		log.Panicln(err)
		errorHandler(w, "error with redis client get", http.StatusBadRequest)
		return
	}

	log.Panicf("result: %s", val)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(val))
}

func main() {
	rClient = redis.NewClient(&redis.Options{Addr: os.Getenv("REDIS_URL")})
	pong, err := rClient.Ping().Result()
	if err != nil {
		fmt.Println("Client failed ping")
		panic(err)
	}

	fmt.Printf("Serving 0.0.0.0:8770 | Redis: %s", pong)
	http.HandleFunc("/inc", handle)
	http.ListenAndServe(":8770", nil)
}
