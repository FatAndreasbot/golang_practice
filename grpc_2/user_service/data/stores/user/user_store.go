package data

import (
	"errors"
	"t1"
	"user_service/data/models"

	"github.com/google/uuid"
)

type UserMemStore struct {
	store *t1.Cache[uuid.UUID, *models.User]
	usernames *t1.Cache[string, uuid.UUID]
}

func NewUserMemStore() *UserMemStore{
	return &UserMemStore{
		store: t1.NewCache[uuid.UUID, *models.User](),
	}
}

func (s *UserMemStore) AddUser(user *models.User) error {
	_, exists := s.usernames.Get(user.Username)
	if exists {
		return errors.New("username is not unique")
	}

	s.store.Set(user.ID, user)
	return nil
}

func (s *UserMemStore) GetByID(id uuid.UUID) (*models.User, error){
	if user, ok := s.store.Get(id); ok {
		return user, nil
	}
	return nil, errors.New("user was not found")
}

func (s *UserMemStore) GetByUsername(username string) (*models.User, error){
	return nil, errors.New("not implemented")
}
