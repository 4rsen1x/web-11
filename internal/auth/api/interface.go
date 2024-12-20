package api

type Usecase interface {
	RegisterUser(username, password string) error
	AuthenticateUser(username, password string) (bool, error)
}
