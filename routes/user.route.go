package routes

import (
	"njajal-go/controller"

	"github.com/go-chi/chi/v5"
)

func UserRoutes(router chi.Router, UserController controller.UserController) chi.Router {
	router.Post("/register", UserController.Register)
	router.Post("/login", UserController.Login)
	return router
}
