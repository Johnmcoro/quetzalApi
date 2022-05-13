package service

type User struct {
}

type UserService interface {
	GetUser()
}

type userService struct {
}

func (us userService) GetUser() {

}

func NewUserService() *userService {
	return &userService{}
}
