package http

import (
	"net/http"

	"github.com/johnmcoro/quetzalapi/internal/service"
)

type UserHandler interface {
	GetUsers(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	UserService service.UserService
}

func NewUserHandler(us service.UserService) UserHandler {
	return &userHandler{
		UserService: us,
	}
}
func (u *userHandler) GetUsers(w http.ResponseWriter, r *http.Request) {

}
