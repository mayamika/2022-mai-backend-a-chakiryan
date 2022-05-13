package auth

type (
	LoginInput struct {
		Username string
		Password string
	}

	LoginPayload struct {
		Token string
	}
)

type (
	RegisterInput struct {
		Username string
		Name     string
		Surname  string
	}

	RegisterPayload LoginPayload
)
