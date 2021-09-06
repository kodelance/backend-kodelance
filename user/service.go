package user

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input UserInput) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input UserInput) (User, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	if err != nil {
		fmt.Println("Error, password gagal di enkripsi")
	}
	user := User{
		Email:    input.Email,
		Password: string(passwordHash),
		Fullname: input.Fullname,
	}

	newUser, err := s.repository.Save(user)
	if err != nil {
		fmt.Println("Gagal membuat user baru")
		return newUser, err
	}

	return newUser, nil
}
