package router

import (
	"socially/server/middleware"

	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/event", middleware.GetAllEvents).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/event", middleware.CreateEvent).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/event", middleware.DeleteEvent).Methods("DELETE", "OPTIONS")
	return router
}
