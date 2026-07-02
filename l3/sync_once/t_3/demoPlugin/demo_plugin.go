package demoplugin

import (
	pl "syncplugin"
	"time"
)

type DemoPlugin struct{}

func (p *DemoPlugin) Execute() string {
	return "DemoPlugin executed successfully!"
}

func InitDemo() (pl.Plugin, error) {
	// Имитация длительной инициализации
	time.Sleep(1000 * time.Millisecond)
	return &DemoPlugin{}, nil
}
