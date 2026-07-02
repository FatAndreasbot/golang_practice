package t1_test

import (
	"t1"
	"testing"
)

func TestDatabase(t *testing.T) {
	db := t1.Database{}

	_ = db.GetConnection()
	_ = db.GetConnection()

}
