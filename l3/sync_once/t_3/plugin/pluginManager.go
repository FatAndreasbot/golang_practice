package syncplugin

import (
	"errors"
	"fmt"
	"sync"
)

var PluginDoesntExistsError error = errors.New("plugin is registered")

type PluginManager struct {
	plugins map[string]*pluginEntry
	mu      sync.RWMutex
}

func NewPluginManager() *PluginManager {
	return &PluginManager{
		plugins: make(map[string]*pluginEntry),
	}
}

func (pm *PluginManager) RegisterPlugin(name string, initFn func() (Plugin, error)) {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	pm.plugins[name] = &pluginEntry{
		initFn: initFn,
	}
}

// GetPlugin возвращает инициализированный плагин
func (pm *PluginManager) GetPlugin(name string) (Plugin, error) {
	// 1. Проверку существования плагина
	plugin_entryPtr, exists := pm.plugins[name]
	if !exists {
		return nil, PluginDoesntExistsError
	}

	// 2. Потокобезопасную однократную инициализацию
	var plugin Plugin
	var err error
	executed := false
	plugin_entryPtr.once.Do(func() {
		fmt.Printf("plugin %q was installed\n", name)
		plugin, err = plugin_entryPtr.initFn()
		executed = true
	})

	// 3. Обработку и кэширование ошибок
	if executed {
		plugin_entryPtr.pluginInstance = plugin
		plugin_entryPtr.instantionError = err
	}

	// 4. Возврат кэшированного результата
	return plugin_entryPtr.pluginInstance, plugin_entryPtr.instantionError
}
