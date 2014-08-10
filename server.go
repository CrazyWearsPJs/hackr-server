package main

import (
	"github.com/go-martini/martini"
	"gopkg.in/mgo.v2"
    //"gopkg.in/mgo.v2/bson"
    "os"
    "fmt"
)

func main() {
    uri:= os.Getenv("MONGOHQ_URL")
    if uri == "" {
        fmt.Println("no connection string provided")
        os.Exit(1)
    }

    sess, err := mgo.Dial(uri)
    if err != nil {
        fmt.Printf("Can't connect to mongo, got error %v\n", err)
        os.Exit(1)
    }

    defer sess.Close()

    m := martini.Classic()
    m.Get("/", helloHandler)
	m.Run()
}

func helloHandler() (int, string) {
	return 200, "Hello World"
}
