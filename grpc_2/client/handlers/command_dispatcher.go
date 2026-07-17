package handlers

import (
	"fmt"
	"sync"
)

type CommandDispatcher struct {
	actions map[string]func(...string) (string, error)
	helpMessages map[string]string
	mu sync.RWMutex
}

func NewCommandDispatcher() *CommandDispatcher{
	dispatcher := &CommandDispatcher{
		actions: make(map[string]func(...string) (string, error), 0),
		helpMessages: make(map[string]string, 0),
		mu: sync.RWMutex{},
	}
	dispatcher.AddCommand("help", "to print this message", func(params ...string) (string, error){
		var helpMessage string
		if len(params) == 0 {
			for command, message := range dispatcher.helpMessages{
				helpMessage = helpMessage + fmt.Sprintf("%q - %s\n", command, message)
			}
		} else {
			for _, command := range params {
				message, ok := dispatcher.helpMessages[command]
				if ok {
					helpMessage = helpMessage + fmt.Sprintf("%q - %s\n", command, message)
				} else {
					helpMessage = helpMessage + fmt.Sprintf("%q was not found", command)
				}
			}
		}
		return helpMessage, nil

	})

	return dispatcher
}

func (h *CommandDispatcher) AddCommand(name, helpString string, action func(...string) (string, error) ) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.actions[name] = action
	h.helpMessages[name] = helpString
}

func (h *CommandDispatcher) ExecuteCommand(command string, params ...string) {
	h.mu.RLock()
	action, ok := h.actions[command]
	h.mu.RUnlock()
	if !ok {
		h.ExecuteCommand("help")
		return
	}
	result, err := action(params...)
	if err != nil{
		fmt.Printf("error while executing %q,\n%v\n\n", command, err)
		return
	}
	fmt.Println(result)
}
