package router

import (
	"github.com/jsam1904/Proyecto1-Tracker-Api/handlers"
	"github.com/jsam1904/Proyecto1-Tracker-Api/middleware"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
)

func New() *chi.Mux {
	r := chi.NewRouter()

	r.Use(chimiddleware.Logger)
	r.Use(chimiddleware.Recoverer)
	r.Use(middleware.CORS)

	r.Get("/teams", handlers.GetTeams)
	r.Get("/teams/{id}", handlers.GetTeam)
	r.Post("/teams", handlers.CreateTeam)
	r.Put("/teams/{id}", handlers.UpdateTeam)
	r.Delete("/teams/{id}", handlers.DeleteTeam)

	return r
}
