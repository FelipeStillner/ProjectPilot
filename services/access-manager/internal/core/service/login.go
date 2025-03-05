package core

type LoginInput struct {
	Username string
	Password string
}

func (s *AccessService) Login(input LoginInput) (*string, error) {
	err := input.verify()
	if err != nil {
		return nil, err
	}

	// TODO

	token := "idk"

	return &token, nil
}

func (input LoginInput) verify() error {
	return nil
}
