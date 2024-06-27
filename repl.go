package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			"help",
			"help menu",
			helpCommand,
		},
		"exit": {
			"exit",
			"exit the program",
			exitCommand,
		},
		"map": {
			"map",
			"Lists some location areas",
			mapCommand,
		},
	}
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("pokedex > ")
		scanner.Scan()
		cleaned := cleanInput(scanner.Text())
		if len(cleaned) == 0 {
			continue
		}
		commandName := cleaned[0]

		commandList := getCommands()
		command, ok := commandList[commandName]
		if !ok {
			fmt.Println("Invalid command")
			continue
		}

		command.callback(cfg)
	}

}

func cleanInput(inputString string) (splitWords []string) {
	lowered := strings.ToLower(inputString)
	splitWords = strings.Fields(lowered)
	return
}
