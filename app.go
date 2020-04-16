// Package main consists of the application, namely the rest service
package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

// Type Application comprises a reference to the router as well as the database that is used to persist data
type Application struct {
	Router *mux.Router
	DB     *sql.DB
}

// Initialize the database connection
func (app *Application) Initialize(user, password, dbname string, port int) {
	connectionString :=
		fmt.Sprintf("port=%d user=%s password=%s dbname=%s sslmode=disable", port, user, password, dbname)

	var err error
	app.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	app.Router = mux.NewRouter()
	app.initializeRoutes()
}

// Initialize the REST handler
func (app *Application) initializeRoutes() {
	app.Router.HandleFunc("/footballer", app.getFootballers).Methods("GET")
}

// Run the application
func (app *Application) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, app.Router))
}

func main() {
	app := Application{}
	app.Initialize("postgres", "admin", "postgres", 5432)
	app.Run(":8090")
}

func (app *Application) getFootballers(writer http.ResponseWriter, reader *http.Request) {
	var players, err = getFootballers(app.DB)
	if err != nil {
		respondWithError(writer, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(writer, http.StatusOK, players)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
