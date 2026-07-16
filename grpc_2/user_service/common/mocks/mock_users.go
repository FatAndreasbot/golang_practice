package mocks

import (
	"proto/user_service"
	"user_service/data/models"
	data "user_service/data/stores/user"
)

func FillMockUsers(store data.UserStore) {
	testUser := models.NewUser("user", "user", user_service.UserRole_USER_ROLE_USER)
	testAdmin := models.NewUser("admin", "admin", user_service.UserRole_USER_ROLE_ADMIN)
	testBroker := models.NewUser("broker", "broker", user_service.UserRole_USER_ROLE_BROKER)

	store.AddUser(testUser)
	store.AddUser(testAdmin)
	store.AddUser(testBroker)
}
