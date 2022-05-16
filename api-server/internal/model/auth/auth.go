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
}

type RegisterPayload struct {
	Token string
}
