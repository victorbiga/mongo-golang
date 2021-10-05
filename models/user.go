package models

import "gopkg.in/mgo.v2/bson"

type Order struct {
	Id          bson.ObjectId `json":"id" bson:" _id"`
	Exchange    string        `json:"exchange" bson: "exchange"`
	Ticker      string        `json:"ticker" bson: "ticker"`
	Price       float32       `json:"price" bson: "price"`
	Amount      float32       `json:"amount" bson: "amount"`
	OrderAction string        `json:"orderaction" bson: "orderaction"`
}
