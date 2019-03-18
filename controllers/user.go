package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/swhite24/go-rest-tutorial/models"
)

type (
	// UserController represents the controller for operating on the User resource
	UserController struct{}
)

type User struct {
	//Id     bson.ObjectId `json:"id" bson:"_id"`
	firstName   string        `json:"firstName" bson:"firstName"`
	lastName   string        `json:"lastName" bson:"lastName"`
	message   string        `json:"message" bson:"message"`
	email   string        `json:"email" bson:"email"`
	service string        `json:"service" bson:"service"`
}

func NewUserController() *UserController {
	return &UserController{}
}

// GetUser retrieves an individual user resource
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Stub an example user
	u := User{
		firstName:   "Bob Smith",
		lastName: "male",
		message: "Test",
		email:"test email",
		service:"security"
		//Id:     p.ByName("id"),
	}

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(u)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", uj)
	fmt.Println(uj)
}

// CreateUser creates a new user resource
func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Stub a user to be populated from the body
	u := models.User{}

	// Populate the user data
	json.NewDecoder(r.Body).Decode(&u)

	// Add an Id
	u.Id = "foo"

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(u)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", uj)

}

// RemoveUser removes an existing user resource
func (uc UserController) RemoveUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// TODO: only write status for now
	w.WriteHeader(200)
}