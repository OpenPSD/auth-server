package userstore

import (
	"fmt"

	"github.com/openpsd/auth-server/entities"
)

type Userstore interface {
	CreateUser(user *entities.User) (*entities.User, error)
	GetUser(username string) (*entities.User, error)
}

type MemUserStore struct {
	users map[string]*entities.User
}

func NewMemUserStore() *MemUserStore {
	return &MemUserStore{
		users: make(map[string]*entities.User),
	}
}

// GetUser by username
func (m *MemUserStore) GetUser(username string) (*entities.User, error) {
	if user, ok := m.users[username]; ok {
		return user, nil
	}
	return nil, fmt.Errorf("failed to get user %s", username)
}

func (m *MemUserStore) CreateUser(user *entities.User) (*entities.User, error) {
	m.users[user.Username] = user
	return user, nil
}
