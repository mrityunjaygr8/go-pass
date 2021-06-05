package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mrityunjaygr8/go-pass/users"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type App struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (a App) migrate() {
	a.DB.AutoMigrate(&users.User{})
}

func (a *App) Initialize(username, password, dbname string) {
	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s", username, password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("An error has occurred")
	}
	a.DB = db
	a.migrate()

	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) initializeRoutes() {
	// a.Router.HandleFunc("/users/", a.)
	a.Router.HandleFunc("/api/users/", a.getAllUsers).Methods("GET")
	a.Router.HandleFunc("/api/users/", a.createUser).Methods("POST")
	a.Router.HandleFunc("/api/users/{id:[0-9]+}/", a.fetchUser).Methods("GET")
	a.Router.HandleFunc("/api/users/{id:[0-9]+}/", a.updateUser).Methods("PUT")
	a.Router.HandleFunc("/api/users/{id:[0-9]+}/", a.deleteUser).Methods("DELETE")
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}
