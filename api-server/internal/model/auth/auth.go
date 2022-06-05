package auth

type LoginInput struct {
	Login    string
	Password string
}

type LoginPayload struct {
	Token string
}

type RegisterInput struct {
	Login    string
	Password string
	Name     string
	Surname  string
	Email    string
}

type RegisterPayload struct {
	Token string
}
