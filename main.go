package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/victorbiga/mongo-golang/controllers"
	"gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewOrderController(getSession())
	r.GET("/order/:id", uc.GetOrder)
	r.POST("/order", uc.CreateOrder)
	r.DELETE("/order/:id", uc.DeleteOrder)
	http.ListenAndServe("localhost:9000", r)

}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost:27017")

	if err != nil {
		panic(err)
	}
	return s
}
