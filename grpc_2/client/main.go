package main

import (
	"bufio"
	"client/handlers"
	userservice_handlers "client/handlers/user_service"
	"fmt"
	"os"
	"strings"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)

	dispatcher := handlers.NewCommandDispatcher()
	userservice_handlers.AddUserServiceHandlers(dispatcher)



	for {
		if scanner.Scan() {
			var args []string
			input := strings.Split(scanner.Text(), " ")
			command := input[0]
			if len(input) > 1 {
				args = input[1:]
			}

			dispatcher.ExecuteCommand(command, args...)
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Error:", err)
		}
	}
}
