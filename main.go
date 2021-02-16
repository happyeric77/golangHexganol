package main

import (
	"project5/api"
	"strconv"
	"log"
	"os"
	"fmt"
	"net/http"
	"github.com/go-chi/chi"
	"project5/shortener"
	"project5/repository/redis"
	"project5/repository/mongo"
)

func main() {
	repo := chooseRepo()
	service := shortener.NewRedirectService(repo)
	handler := api.NewHandler(service)

	r := chi.NewRouter()
	r.Get("/{code}", handler.GET)
	r.Post("/", handler.POST)

	errs := make(chan error, 2)

	go func () {
		fmt.Printf("Listening to %s\n", httpPort())
		errs <- http.ListenAndServe(httpPort(), r)
	}()

	fmt.Printf("Terminated: %s\n", <-errs)	
}

func httpPort() string {
	port := ":7160"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	return fmt.Sprintf("%s", port)
}

func chooseRepo() shortener.RedirectRepository {
	switch os.Getenv("URL_DB") {
	case "redis":
		redisURL := "redis://127.0.0.1:6379" //Default redisURL
		if os.Getenv("REDIS_URL") != "" {
			redisURL = os.Getenv("REDIS_URL")
		}
		repo, err := redis.NewRedisRepository(redisURL)
		if err != nil {
			log.Fatal(err)
		}
		return repo
	case "mongo":
		mongoURL := "mongodb://127.0.0.1:27017/" // Default mongoURL
		if os.Getenv("MONGO_URL") != "" {
			mongoURL = os.Getenv("MONGO_URL")
		}
		mongoDB := "urlShortener" // Default mongoDB
		if os.Getenv("MONGO_DB") != "" {
			mongoDB = os.Getenv("MONGO_DB")
		}
		mongoTimeout := 5 // Default mongoTimeout (unit: second)
		if os.Getenv("MONGO_TIMEOUT") != "" {
			mongoTimeout, _ = strconv.Atoi(os.Getenv("MONGO_TIMEOUT"))
		}
		repo, err := mongo.NewMongoRepository(mongoURL, mongoDB, mongoTimeout)
		if err != nil {
			log.Fatal(err)
		}
		return repo
	}
	return nil
}