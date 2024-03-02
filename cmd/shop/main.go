package main

import (
	"ecommerce/db"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"gorm.io/gorm"

	_ "github.com/lib/pq"
)

type handler struct {
	DB *gorm.DB
}

func New(db *gorm.DB) handler {
	return handler{db}
}

var h = New(db.Connect())
var r = mux.NewRouter()

func mainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello world")
}

func HandleRequests() {
	r.HandleFunc("/", mainPage)
	r.HandleFunc("/items", h.CreateItem).Methods("POST")
	r.HandleFunc("/items/all", h.GetAllItems).Methods("GET")
	r.Path("/items").Queries("id", "{id}").HandlerFunc(h.UpdateItemById).Methods("PUT")
	r.Path("/items").Queries("id", "{id}").HandlerFunc(h.DeleteItemById).Methods("DELETE")
	// r.HandleFunc("/signin", packages.Signin)
	// r.HandleFunc("/signup", packages.Signup)
	// r.HandleFunc("/items/{id}/rating", h.GetItemRating).Methods("GET")
	// r.HandleFunc("/items/{id}/rating", h.PostRating).Methods("POST")
	// r.HandleFunc("/items/{id}/rating", h.UpdateRating).Methods("PUT")

	// r.HandleFunc("/items/{id}/comment", h.GetItemComments).Methods("GET")
	// r.HandleFunc("/items/{id}/comment", h.PostComment).Methods("POST")
	// r.HandleFunc("/items/{id}/comment", h.UpdateComment).Methods("PUT")
	// r.HandleFunc("/items/{id}/comment", h.DeleteComment).Methods("DELETE")

	// r.HandleFunc("/cart", h.GetCartItems).Methods("GET")

	// r.HandleFunc("/orders", h.GetUserOrders).Methods("GET")

	// r.HandleFunc("/addresses", h.PostAddress).Methods("POST")
	// r.HandleFunc("/addresses", h.UpdateAddress).Methods("PUT")
	// r.HandleFunc("/addresses", h.DeleteAddress).Methods("DELETE")

	// r.HandleFunc("/comments", h.GetUserComments).Methods("GET")
	fmt.Println("server is started")
	http.ListenAndServe(":7000", r)
}

func main() {
	HandleRequests()
}
