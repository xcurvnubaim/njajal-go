// controller/user.controller.go
package controller

import (
	"net/http"
	"njajal-go/service"

	"github.com/go-chi/render"
)

// UserController is the exported interface
type UserController interface {
	Register(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
}

// UserControllerImpl is the implementation of UserController
type UserControllerImpl struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserControllerImpl {
	return &UserControllerImpl{
		userService: userService,
	}
}

type ErrorResponse struct {
	Message string `json:"message"`
}

// Register is the implementation of the Register method
func (uc *UserControllerImpl) Register(w http.ResponseWriter, r *http.Request) {
	type RegisterRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var request RegisterRequest
	err := render.Decode(r, &request)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, ErrorResponse{Message: err.Error()})
		return
	}
	err = uc.userService.Register(r.Context(), request.Email, request.Password)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, ErrorResponse{Message: err.Error()})
		return
	}

	type JsonResponse struct {
		Message string `json:"message"`
	}
	response := JsonResponse{
		Message: "Register Success!",
	}
	render.JSON(w, r, response)
}

func (uc *UserControllerImpl) Login(w http.ResponseWriter, r *http.Request) {
	type LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var request LoginRequest
	err := render.Decode(r, &request)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, ErrorResponse{Message: err.Error()})
		return
	}
	token, err := uc.userService.Login(r.Context(), request.Email, request.Password)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, ErrorResponse{Message: err.Error()})
		return
	}
	type dataResponse struct {
		Email string `json:"email"`
		Token string `json:"token"`
	}
	type JsonResponse struct {
		Message string `json:"message"`
		Data    dataResponse
	}
	response := JsonResponse{
		Message: "Register Success!",
		Data: dataResponse{
			Email: request.Email,
			Token: token,
		},
	}
	render.JSON(w, r, response)
}
