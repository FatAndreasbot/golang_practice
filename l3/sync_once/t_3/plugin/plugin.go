package syncplugin

import "sync"

type Plugin interface {
	Execute() string
}

type pluginEntry struct {
	initFn          func() (Plugin, error)
	pluginInstance  Plugin
	instantionError error
	once            sync.Once
}
