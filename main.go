package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gomodule/redigo/redis"
)

const contentType = "Content-Type"
const jsonUTF8 = "application/json; charset=utf-8"
const name = "samplesvc"

// These variables are assigned values during the build process using the -ldflags="-X ..." linker option.
var version = "Not set"
var origin = "Not set"
var commit = "Not set"
var branch = "Not set"
var built = "Not set"
var redisHost = getEnv("REDIS_SERVICE_HOST", "localhost")
var redisPort = getEnv("REDIS_SERVICE_PORT", "6379")
var maxRedisConns = 50

type info struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Origin  string `json:"origin"`
	Commit  string `json:"commit"`
	Branch  string `json:"branch"`
	Built   string `json:"built"`
}

type sampleLink struct {
	CollectionExerciseID string `json:"collectionExerciseId"`
	SampleSummaryID      string `json:"sampleSummaryId"`
}

func main() {
	port, overridden := os.LookupEnv("PORT")
	if !overridden {
		port = ":8081"
	}

	redisPool := redis.NewPool(func() (redis.Conn, error) {
		c, err := redis.Dial("tcp", redisHost+":"+redisPort)

		if err != nil {
			return nil, err
		}

		return c, err
	}, maxRedisConns)

	defer redisPool.Close()

	http.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(contentType, jsonUTF8)
		json, _ := json.Marshal(info{Name: name, Version: version, Origin: origin, Commit: commit, Branch: branch, Built: built})
		fmt.Fprintf(w, "%s\n", json)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "OK\n")
	})

	http.HandleFunc("/samples/", func(w http.ResponseWriter, r *http.Request) {

		id := strings.TrimPrefix(r.URL.Path, "/samples/")
		id = strings.TrimSuffix(id, "/attributes")

		c := redisPool.Get()
		defer c.Close()

		value, err := redis.String(c.Do("GET", "sampleunit:"+id))

		if err != nil {
			w.WriteHeader(404)
			fmt.Println(err)
			message := fmt.Sprintf("Could not GET %s", id)
			fmt.Fprintf(w, "%s\n", message)
		} else {
			// have to swap ' for " as its a python object written and ' is not valid JSON
			w.Header().Set(contentType, jsonUTF8)
			value = strings.Replace(value, "'", "\"", -1)
			fmt.Fprintf(w, "%s\n", value)
		}

	})
	log.Fatal(http.ListenAndServe(port, nil))
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
