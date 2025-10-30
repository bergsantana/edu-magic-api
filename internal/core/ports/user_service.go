package ports

type UserService interface {
	SignUp(name, email, password string) error
	Login(email, password string) (string, error)
}
