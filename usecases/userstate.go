package usecases

import (
	"fmt"

	"github.com/openpsd/auth-server/entities"
	"github.com/openpsd/auth-server/providers/userstore"
)

type Userstate struct {
	Userstore userstore.Userstore
}

func NewUserstate(userstore userstore.Userstore) *Userstate {
	return &Userstate{
		Userstore: userstore,
	}
}

// Login the user by validating the password
func (u *Userstate) Login(username string, password string) error {
	user, err := u.Userstore.GetUser(username)
	if err != nil {
		return err
	}
	if correctBcrypt([]byte(user.Hash), password) {
		user.IsLoggedIn = true
		return nil
	}
	return fmt.Errorf("invalid password %s", username)
}

// Logout the user
func (u *Userstate) Logout(username string) error {
	user, err := u.Userstore.GetUser(username)
	if err != nil {
		return err
	}
	user.IsLoggedIn = false
	return nil
}

// CreateUser with the given userstore
func (u *Userstate) CreateUser(username string, email string, password string) (*entities.User, error) {
	hash := hashBcrypt(password)
	user := entities.NewUser(username, email, hash)
	u.Userstore.CreateUser(user)
	return user, nil
}
