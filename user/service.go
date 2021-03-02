package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input Register) (User, error)
	Login(input Login) (User, error)
	Token(input TokenInput) (Token, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input Register) (User, error) {
	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	Password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.Password = string(Password)
	newUser, err := s.repository.Register(user)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}

func (s *service) Login(input Login) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}
	if user.ID == 0 {
		return user, errors.New("user tidak ditemukan")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s *service) Token(input TokenInput) (Token, error) {
	token := Token{}
	token.UserID = input.UserID
	token.Token = input.Token

	newToken, err := s.repository.UserToken(token)
	if err != nil {
		return newToken, err
	}
	return newToken, nil
}
