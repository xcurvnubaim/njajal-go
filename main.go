// main.go
package main

import (
	"fmt"
	"net/http"
	"njajal-go/config"
	"njajal-go/controller"
	db "njajal-go/db/sqlc"
	"njajal-go/repository"
	"njajal-go/routes"
	"njajal-go/service"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func main() {
	var (
		conn *db.Queries = config.DbInit()
		userRepository = repository.NewUserRepository(conn)
		userService = service.NewUserService(userRepository)
		userController = controller.NewUserController(userService)
	)


	// Create a new router
	r := chi.NewRouter()

	// Use middleware.Logger
	r.Use(middleware.Logger)
	r.Use(middleware.CleanPath)

	// Set up middleware to handle JSON responses
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Get("/api", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API is working!"))
	})
	r.Get("Healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API is working!"))
	})
	// Define the user routes
	r.Mount("/api/user", routes.UserRoutes(r,userController))

	// Start the server
	fmt.Println("Initializing server at http://localhost:3000")
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
