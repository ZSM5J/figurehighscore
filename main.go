package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/rs/cors"
	"github.com/figurehighscore/handler"
)

func init() {


	// Register routes
	r := mux.NewRouter()
	r.HandleFunc("/api/figures/new", handler.NewFigureHandler).Methods("POST")
	r.HandleFunc("/api/results/new", handler.NewResultHandler).Methods("POST")

	r.HandleFunc("/api/figures", handler.GetFigureListHandler).Methods("GET")
	r.HandleFunc("/api/results", handler.GetResultListHandler).Methods("GET")
	r.HandleFunc("/api/players", handler.GetPlayerListHandler).Methods("GET")

	r.HandleFunc("/api/figures/{id}", handler.GetResultByFigureHandler).Methods("GET")
	r.HandleFunc("/api/players/{token}", handler.GetResultByPlayerHandler).Methods("GET")
	r.HandleFunc("/api/max/{id}", handler.GetMaxResultHandler).Methods("GET")

	r.HandleFunc("/api/results/{id}", handler.DeleteResultHandler).Methods("DELETE")
	r.HandleFunc("/api/results/trash/{id}", handler.TrashResultHandler).Methods("PUT")
	r.HandleFunc("/api/results/undotrash/{id}", handler.UndoTrashResultHandler).Methods("PUT")
	r.HandleFunc("/api/figures/{id}", handler.DeleteFigureHandler).Methods("DELETE")
	r.HandleFunc("/api/players/{token}", handler.DeletePlayerHandler).Methods("DELETE")

	// Start HTTP server
	http.Handle("/", cors.AllowAll().Handler(r))
}