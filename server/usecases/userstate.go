package usecases

import (
	"fmt"

	"github.com/openpsd/auth-server/server/entities"
	"github.com/openpsd/auth-server/server/providers/oauth"
	"github.com/openpsd/auth-server/server/providers/userstore"
)

type Userstate struct {
	Userstore   userstore.Userstore
	Oauthclient oauth.Oauthclient
}

func NewUserstate(userstore userstore.Userstore, oauthclient oauth.Oauthclient) *Userstate {
	return &Userstate{
		Userstore:   userstore,
		Oauthclient: oauthclient,
	}
}

// Login the user by validating the password
func (u *Userstate) Login(username string, password string, challenge string) (string, error) {
	user, err := u.Userstore.GetUser(username)
	if err != nil {
		return "", err
	}
	if correctBcrypt([]byte(user.Hash), password) {
		user.IsLoggedIn = true
		acceptLoginRequest := entities.NewAcceptLoginRequest(username, true)
		redirectLink, err := u.Oauthclient.AcceptLoginRequest(challenge, acceptLoginRequest)
		return redirectLink, err
	}
	return "", fmt.Errorf("invalid password %s", username)
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
