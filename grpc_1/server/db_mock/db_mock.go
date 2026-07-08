package dbmock

import (
	"fmt"
	"t1"
	"time"
)

type Role uint8

const (
	USER Role = iota
	ADMIN
)

type User struct {
	ID        int64
	Name      string
	Password  string
	Email     string
	Role      Role
	CreatedAt time.Time
	UpdatedAt time.Time
}

type DBMock struct {
	store  t1.Cache[int64, User]
	nextID int64
}

func NewDBMock() *DBMock {
	return &DBMock{
		store:  t1.NewCache[int64, User](),
		nextID: 1,
	}
}

func (db *DBMock) Create(u *User) int64 {
	db.store.Set(db.nextID, *u)
	u.ID = db.nextID
	db.nextID++

	return u.ID
}

func (db *DBMock) Retrieve(id int64) (*User, error) {
	user, exists := db.store.Get(id)
	if !exists {
		return nil, fmt.Errorf("user with id %d was not found", id)
	}
	return user, nil
}

func (db *DBMock) Update(userID int64, user User) {
	db.store.Set(userID, user)
}

func (db *DBMock) Delete(userID int64) {
	db.store.Delete(userID)
}
