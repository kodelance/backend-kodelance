package user

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input UserInput) (User, error)
	LoginUser(input LoginInput) (User, error)
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
		return newUser, err
	}

	return newUser, nil
}

func (s *service) LoginUser(input LoginInput) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, errors.New("Email not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, errors.New("Password doesnt match")
	}

	return user, nil
}
