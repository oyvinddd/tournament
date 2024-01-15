package user

type (
	Service interface {

		// SignIn using Apple
		SignIn() (User, error)

		InviteUser() error

		// GetInvitations to tournaments for a given user
		GetInvitations() ([]Invite, error)
	}

	MockService struct {
		Invitations []Invite
	}
)

func (service MockService) InviteUser() error {
	//TODO implement me
	panic("implement me")
}

func NewService() Service {
	return &MockService{}
}

func (service MockService) SignIn() (User, error) {
	return User{}, nil
}

func (service MockService) GetInvitations() ([]Invite, error) {
	return nil, nil
}
