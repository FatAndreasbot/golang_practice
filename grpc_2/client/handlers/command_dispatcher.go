package handlers

import (
	"fmt"
	"sync"
)

type CommandDispatcher struct {
	actions map[string]func(...string) (string, error)
	mu sync.RWMutex
}

func NewCommandDispatcher() *CommandDispatcher{
	return &CommandDispatcher{
		actions: map[string]func(...string) (string, error){},
		mu: sync.RWMutex{},
	}
}

func (h *CommandDispatcher) AddCommand(name string, action func(...string) (string, error) ) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.actions[name] = action
}

func (h *CommandDispatcher) ExecuteCommand(command string, params ...string) {
	h.mu.RLock()
	action, ok := h.actions[command]
	h.mu.RUnlock()
	if !ok {
		fmt.Printf("command %q wass not found\n", command)
		return
	}
	result, err := action(params...)
	if err != nil{
		fmt.Printf("error while executing %q,\n%v\n\n", command, err)
		return
	}
	fmt.Println(result)
}
