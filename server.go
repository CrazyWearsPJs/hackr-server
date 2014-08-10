package main

import (
	"github.com/go-martini/martini"
	//"gopkg.in/mgo.v2"
    //"gopkg.in/mgo.v2/bson"
    //"os"
    //"fmt"
)

func main() {
	m := martini.Classic()
	connect()
    m.Get("/", helloHandler)
	m.Run()
}

func helloHandler() (int, string) {
	return 200, "Hello World"
}

func connect() {
    if uri == ""{
        fmt.Println("no connection string provided")
        os.Exit(1)
    }

    sess, err := mgo.Dial(uri)
    if err != nil {
        fmt.Printf("Can't connect to mongo, go err %v\n", err)
        os.Exit(1)
    }
    defer sess.Close()
}

