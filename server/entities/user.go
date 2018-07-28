package entities

// User struct
type User struct {
	Username          string
	Email             string
	Hash              []byte
	ExternalReference string
	IsLoggedIn        bool
}

func NewUser(username string, email string, hash []byte) *User {
	return &User{
		Username:   username,
		Email:      email,
		Hash:       hash,
		IsLoggedIn: false,
	}
}
