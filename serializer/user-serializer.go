package serializer

import (

	"github.com/eliemugenzi/simply-hired/db/models"
)

type UserResponse struct {
	ID uint `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	Role string `json:"role"`
}

type UserSerializer struct {
	User models.User
}

type UsersSerializer struct {
	Users []models.User
}

func (serializer *UserSerializer) Response() UserResponse {
	return UserResponse {
		ID: serializer.User.ID,
		FirstName: serializer.User.FirstName,
		LastName: serializer.User.LastName,
		Email: serializer.User.Email,
		Role: serializer.User.Role,
	}
}


func (serializer *UsersSerializer) Response() []UserResponse {
	response := []UserResponse{}
	for _, user := range serializer.Users {
		userSerializer := UserSerializer{User: user}
		response = append(response, userSerializer.Response())
	}

	return response
}