package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/Nebiyou-Daniel/go-api/internal/middleware"
	)

// Handler is the main handler for the API

func Handler(r *chi.Mux){
	// Global middleware
	r.Use(chimiddle.StripSlashes)

	// Routes
	r.Route("/accounts", func(router chi.Router){

		// Middleware for /accounts route
		router.Use(middleware.Authorization)

		router.Get("/coins", GetCoinBalance)
	})
}

