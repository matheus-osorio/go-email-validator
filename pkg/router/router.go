package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/matheus-osorio/go-email-validator/pkg/controller"
)

var router *mux.Router

func Setup() {
	router = mux.NewRouter()
	router.HandleFunc("/verify", controller.CheckDomainExistence).Methods("GET")
	router.HandleFunc("/check", controller.CheckOwnershipEmail).Methods("POST")
	router.HandleFunc("/check/{uuid}", controller.ReceiveOwnershipEmail).Methods("PUT")
}

func Start() {
	err := http.ListenAndServe(":8080", router)
	fmt.Println(err)
}
