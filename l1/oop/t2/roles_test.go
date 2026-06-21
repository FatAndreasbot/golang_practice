package users_test

import (
	"testing"
	"users"
)

func TestBasicUser(t *testing.T) {
	u := users.NewBasicUser("test")
	permissionName, _ := users.GetPermissionName(users.Read)

	hasPermission := u.HasPermission(permissionName)
	if !hasPermission {
		t.Error("basic user should have read permissions")
	}

	permissionName, _ = users.GetPermissionName(users.Edit)
	hasPermission = u.HasPermission(permissionName)
	if hasPermission {
		t.Error("basic user should not have edit permissions")
	}
}

func TestModerator(t *testing.T) {
	u := users.NewModerator("test")
	permissionName, _ := users.GetPermissionName(users.Read)

	hasPermission := u.HasPermission(permissionName)
	if !hasPermission {
		t.Error("moderator should have read permissions")
	}

	permissionName, _ = users.GetPermissionName(users.Edit)
	hasPermission = u.HasPermission(permissionName)
	if !hasPermission {
		t.Error("moderator should have edit permissions")
	}

	permissionName, _ = users.GetPermissionName(users.Delete)
	hasPermission = u.HasPermission(permissionName)
	if hasPermission {
		t.Error("moderator should not have delete permissions")
	}
}

func TestAdmin(t *testing.T) {
	u := users.NewAdmin("test")
	permissionName, _ := users.GetPermissionName(users.Read)

	hasPermission := u.HasPermission(permissionName)
	if !hasPermission {
		t.Error("admin should have read permissions")
	}

	permissionName, _ = users.GetPermissionName(users.Edit)
	hasPermission = u.HasPermission(permissionName)
	if !hasPermission {
		t.Error("admin should have edit permissions")
	}

	permissionName, _ = users.GetPermissionName(users.Delete)
	hasPermission = u.HasPermission(permissionName)
	if !hasPermission {
		t.Error("admin should have delete permissions")
	}
}
