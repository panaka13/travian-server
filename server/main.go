package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/panaka13/travian-server/db"
	"github.com/panaka13/travian-server/handler"
)

var router *mux.Router

func myInit() {
	router = mux.NewRouter()
	router.HandleFunc("/user", handler.CreateUserHandler).Methods("POST")
	router.HandleFunc("/user/{userid}", handler.GetUserHandler).Methods("GET")
	router.HandleFunc("/village", handler.CreateVillageHandler).Methods("POST")
	router.HandleFunc("/viilage/{villageid}", handler.GetVillageHandler).Methods("GET")
	router.HandleFunc("/village/structure", handler.UpdateVillageStructure).Methods("PUT")
}

func main() {
	myInit()
	db.InitDb()
	fmt.Println("Hello World")
	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	srv.ListenAndServe()
}
