package user

type LoginInput struct {
	Login    string
	Password string
}

type LoginPayload struct {
	Token string
}

type User struct {
	ID       string
	Login    string
	Password string
}
