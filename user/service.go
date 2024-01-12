package user

type (
	Service interface {

		// SignIn using Apple
		SignIn() (User, error)

		// GetInvitations to tournaments for a given user
		GetInvitations() ([]Invite, error)
	}

	MockService struct {
		Invitations []Invite
	}
)

func NewService() Service {
	return &MockService{}
}

func (service MockService) SignIn() (User, error) {
	return User{}, nil
}

func (service MockService) GetInvitations() ([]Invite, error) {
	return nil, nil
}
