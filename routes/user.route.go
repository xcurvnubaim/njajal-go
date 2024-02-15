package routes

import (
	"fmt"
	"net/http"
	"njajal-go/config"
	"njajal-go/controller"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

func UserRoutes(router chi.Router, UserController controller.UserController) chi.Router {
	// Protected routes
	router.Group(func(r chi.Router) {
		tokenAuth := config.TokenAuth
		// Seek, verify and validate JWT tokens
		r.Use(jwtauth.Verifier(tokenAuth))

		// Handle valid / invalid tokens.
		r.Use(jwtauth.Authenticator(tokenAuth))

		r.Get("/admin", func(w http.ResponseWriter, r *http.Request) {
			_, claims, _ := jwtauth.FromContext(r.Context())
			w.Write([]byte(fmt.Sprintf("protected area. hi %v", claims["user_id"])))
		})
	})

	// Public routes
	router.Group(func(r chi.Router) {
		r.Post("/register", UserController.Register)
		r.Post("/login", UserController.Login)
	})
	return router
}
