package t2_test

import (
	"t2"
	"testing"
)

func TestConfig(t *testing.T) {
	var cm t2.ConfigManager

	err := cm.LoadConfig("./mock_config_file.json")
	if err != nil {
		t.Error(err)
	}

	testValue := cm.Get("app_name")
	if testValue == "" {
		t.Error("testValue is empty")
	}
	err = cm.LoadConfig("asd")
	if err != nil {
		t.Error("the init relaunched")
	}
}
