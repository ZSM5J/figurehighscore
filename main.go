package main

import (
	"github.com/figurehighscore/config"
	"github.com/gorilla/mux"
	"github.com/figurehighscore/handler"
	"net/http"
	"github.com/rs/cors"
)

func init() {

	config.LoadConfiguration("config.json")
	// Register routes
	r := mux.NewRouter()
	r.HandleFunc("/api/figures/new", handler.NewFigureHandler).Methods("POST")
	r.HandleFunc("/api/results/new", handler.NewResultHandler).Methods("POST")
	r.HandleFunc("/api/login", handler.LoginHandler).Methods("POST")

	r.HandleFunc("/api/figures", handler.GetFigureListHandler).Methods("GET")
	r.HandleFunc("/api/results", handler.GetResultListHandler).Methods("GET")
	r.HandleFunc("/api/results/last25", handler.GetLastResultHandler).Methods("GET")
	r.HandleFunc("/api/players", handler.GetPlayerListHandler).Methods("GET")

	r.HandleFunc("/api/figures/{id}", handler.GetResultByFigureHandler).Methods("GET")
	r.HandleFunc("/api/figures/trash/{id}", handler.GetTrashResultByFigureHandler).Methods("GET")
	r.HandleFunc("/api/players/{token}", handler.GetResultByPlayerHandler).Methods("GET")
	r.HandleFunc("/api/max/{id}", handler.GetMaxResultHandler).Methods("GET")


	r.HandleFunc("/api/results/{id}", handler.DeleteResultHandler).Methods("DELETE")
	r.HandleFunc("/api/results/trash/{id}", handler.TrashResultHandler).Methods("PUT")
	r.HandleFunc("/api/results/undotrash/{id}", handler.UndoTrashResultHandler).Methods("PUT")
	r.HandleFunc("/api/figures/{id}", handler.DeleteFigureHandler).Methods("POST")
	r.HandleFunc("/api/players/{token}", handler.DeletePlayerHandler).Methods("DELETE")

	// Start HTTP server
	http.Handle("/", cors.AllowAll().Handler(r))
}


