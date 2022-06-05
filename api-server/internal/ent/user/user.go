// Code generated by entc, DO NOT EDIT.

package user

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldLogin holds the string denoting the login field in the database.
	FieldLogin = "login"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldSurname holds the string denoting the surname field in the database.
	FieldSurname = "surname"
	// FieldPasswordHash holds the string denoting the password_hash field in the database.
	FieldPasswordHash = "password_hash"
	// EdgeFriends holds the string denoting the friends edge name in mutations.
	EdgeFriends = "friends"
	// Table holds the table name of the user in the database.
	Table = "users"
	// FriendsTable is the table that holds the friends relation/edge. The primary key declared below.
	FriendsTable = "user_friends"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldLogin,
	FieldEmail,
	FieldName,
	FieldSurname,
	FieldPasswordHash,
}

var (
	// FriendsPrimaryKey and FriendsColumn2 are the table columns denoting the
	// primary key for the friends relation (M2M).
	FriendsPrimaryKey = []string{"user_id", "friend_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// LoginValidator is a validator for the "login" field. It is called by the builders before save.
	LoginValidator func(string) error
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// SurnameValidator is a validator for the "surname" field. It is called by the builders before save.
	SurnameValidator func(string) error
	// PasswordHashValidator is a validator for the "password_hash" field. It is called by the builders before save.
	PasswordHashValidator func(string) error
)