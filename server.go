package main

import (
	_ "fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/CrazyWearsPJs/hackr/models/user"
	"github.com/CrazyWearsPJs/hackr/repo"
	"github.com/codegangsta/negroni"
	"github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2"
	_ "gopkg.in/mgo.v2/bson"
)

var (
	redis_conn redis.Conn
	mongo_conn *mgo.Session
)

var (
	Users *repo.UserRepo
)

const (
	defaultPort = 3000
	hackrdb     = "app28354182"
)

func main() {

	mongo_conn := SetupMongo()
	redis_conn := SetupRedis()

	mux := SetupMux()

	defer mongo_conn.Close()
	defer redis_conn.Close()

	mongo_db := mongo_conn.DB(hackrdb)
	Users = &repo.UserRepo{Collection: mongo_db.C("users")}
	u, err := user.New([]byte("agonz056@ucr.edu"), []byte("123"))
	if err != nil {
		panic(err)
	}

	Users.Add(u)

	n := negroni.New(negroni.NewRecovery(), negroni.NewLogger(), negroni.NewStatic(http.Dir("app")))
	n.UseHandler(mux)

	port := os.Getenv("PORT")
	if port == "" {
		port = strconv.Itoa(defaultPort)
	}

	n.Run(":" + port)
}

func SetupRedis() redis.Conn {
	redis_uri := GetConnectionString("REDISTOGO_URL")

	redis_url, err := url.Parse(redis_uri)
	if err != nil {
		log.Fatalf("Unable to connectionString  %v, got error %v\n", redis_uri, err)
		os.Exit(1)
	}

	redis_sess, err := redis.Dial("tcp", redis_url.Host)
	if err != nil {
		log.Fatalf("Can't connect to redis, got error %v\n", err)
	}

	return redis_sess
}

func SetupMongo() *mgo.Session {
	mongo_uri := GetConnectionString("MONGOHQ_URL")

	mongo_sess, err := mgo.Dial(mongo_uri)
	if err != nil {
		log.Fatalf("Can't connect to mongo, got error %v\n", err)
	}
	return mongo_sess
}

func GetConnectionString(s string) string {
	uri := os.Getenv(s)
	if uri == "" {
		log.Fatalf("no connection string for %v provided\n", s)
	}
	return uri
}
