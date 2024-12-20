package usecase

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func (u *Usecase) RegisterUser(username, password string) error {
	existingPassword, err := u.p.GetUser(username)
	if err != nil {
		return err
	}
	if existingPassword != "" {
		return errors.New("user already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	return u.p.AddUser(username, string(hashedPassword))
}

func (u *Usecase) AuthenticateUser(username, password string) (bool, error) {
	hashedPassword, err := u.p.GetUser(username)
	if err != nil {
		return false, err
	}
	if hashedPassword == "" {
		return false, nil
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil, nil
}
