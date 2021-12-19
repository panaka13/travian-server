package main

import (
	"flag"
	"net/http"
	"os"
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
	router.HandleFunc("/user", handler.FindUserHandler).Methods("GET")
	router.HandleFunc("/village", handler.CreateVillageHandler).Methods("POST")
	router.HandleFunc("/village/{villageid}", handler.GetVillageHandler).Methods("GET")
	router.HandleFunc("/village/structure", handler.UpdateVillageStructure).Methods("PUT", "OPTIONS")
}

func main() {
	myInit()

	databaseType := flag.String("databasetype", "sqlite3", "Database type")

	flag.Parse()

	if databaseType == nil {
		*databaseType = "sqlite3"
	}

	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		databaseUrl = "gorm.db"
	}

	db.InitDb(*databaseType, databaseUrl)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	cert := os.Getenv("CERT") != ""

	srv := &http.Server{
		Handler:      router,
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	if cert {
		srv.ListenAndServeTLS("cert/server.crt", "cert/server.key")
	} else {
		srv.ListenAndServe()
	}
}
