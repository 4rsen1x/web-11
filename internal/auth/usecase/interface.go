package usecase

type Provider interface {
	AddUser(username, password string) error
	GetUser(username string) (string, error)
}
