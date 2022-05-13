package user

type User struct {
	ID       string
	Username string
	Name     string
	Surname  string
}

var (
	Anonymous = &User{}
)
