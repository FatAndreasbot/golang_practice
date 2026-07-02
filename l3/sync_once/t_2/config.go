package t2

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

type ConfigManager struct {
	config map[string]string
	once   sync.Once
	lock   sync.RWMutex
}

func (cm *ConfigManager) LoadConfig(configPath string) error {
	var ERR error
	cm.once.Do(func() {
		fileContent, err := os.ReadFile(configPath)
		if err != nil {
			ERR = err
			return
		}

		err = json.Unmarshal(fileContent, &cm.config)
		if err != nil {
			ERR = err
			return
		}
	})

	return ERR
}

func (cm *ConfigManager) Get(key string) string {
	cm.lock.RLock()
	defer cm.lock.RUnlock()
	val := cm.config[key]
	return val
}

func (cm *ConfigManager) PrintConfig() {
	cm.lock.RLock()
	defer cm.lock.RUnlock()
	configData, _ := json.Marshal(cm.config)
	configDataString := string(configData)
	fmt.Println(configDataString)

}
