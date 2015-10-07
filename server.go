package main

import (
	"net/http"

	"gopkg.in/mgo.v2"

	"github.com/RenatoAraujo/go-restful-application/controllers"
	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()

	// Get a UserController instance
	uc := controllers.NewUserController(getSession())

	r.GET("/user/:id", uc.GetUser)

	r.POST("/user", uc.CreateUser)

	r.DELETE("/user/:id", uc.RemoveUser)

	http.ListenAndServe("localhost:3000", r)
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost")

	if err != nil {
		panic(err)
	}

	return s
}
