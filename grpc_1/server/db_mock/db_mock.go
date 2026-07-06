package dbmock

import (
	"fmt"
	"t1"
)

type Role uint8

const (
	USER  Role = 1
	ADMIN Role = 2
)

type User struct {
	ID       int64
	Name     string
	Password string
	Role     Role
}

type DBMock struct {
	store  t1.Cache[int64, User]
	nextID int64
}

func newDBMock() DBMock {
	return DBMock{
		store:  t1.NewCache[int64, User](),
		nextID: 1,
	}
}

func (db *DBMock) Create(u *User) {
	db.store.Set(db.nextID, *u)
	u.ID = db.nextID
	db.nextID++
}

func (db *DBMock) Retrieve(id int64) (User, error) {
	user, exists := db.store.Get(id)
	if !exists {
		return User{}, fmt.Errorf("user with id %d was not found", id)
	}
	return user, nil
}

func (db *DBMock) Update(userID int64, user User) {
	db.store.Set(userID, user)
}

func (db *DBMock) Delete(userID int64) {
	db.store.Delete(userID)
}
