package user

type (
	Service interface {
		SignIn() (User, error)
	}

	MockService struct{}
)

func NewService() Service {
	return &MockService{}
}

func (service MockService) SignIn() (User, error) {
	return User{}, nil
}
