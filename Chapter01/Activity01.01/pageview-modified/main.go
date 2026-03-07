package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/redis/go-redis/v9"
)

var dbClient *redis.Client
var ctx = context.Background()

var key = "pv"

func init() {
	dbClient = redis.NewClient(&redis.Options{
		Addr: "db:6379",
	})
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Ping from %s", r.RemoteAddr)
	pageView, err := dbClient.Incr(ctx, key).Result()
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "Hello, you're visitor #%v.\n", pageView)
}
