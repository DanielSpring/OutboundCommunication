package models

import "gopkg.in/mgo.v2/bson"

type (
	// User represents the structure of our resource
	User struct {
		//Id     bson.ObjectId `json:"id" bson:"_id"`
		firstName   string        `json:"firstName" bson:"firstName"`
		lastName   string        `json:"lastName" bson:"lastName"`
		message   string        `json:"message" bson:"message"`
		email   string        `json:"email" bson:"email"`
		service string        `json:"service" bson:"service"
	}
)
