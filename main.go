package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		panic(fmt.Errorf("no command line args passed"))
	}
	/*
	 * Format of the 'args' array: [`<COMMAND_NAME_1> <ARG1> <ARG2> .. <ARG N>`, `<COMMAND_NAME_2> <ARG1> <ARG2> .. <ARG N>`]
	 *
	 * We loop through the list of commands passed in as input arguments and handle each one of them
	 */
	for _, arg := range args {
		handle(arg)
		//arg will have each command passed in the command line argumens
	}

}

/*
 * This method should parse each input and assigns it into different variables.
 *
 * The value of the function parameter "input" will be of the format - "`<COMMAND_NAME_1> <ARG1> <ARG2> .. <ARG N>".
 * We split the string and retrieve the input data appropriately into the variable inputListForOneCommand.
 *
 */
func handle(command string) {
	inputsForOneCommand := strings.Split(command, " ")
	fmt.Println(inputsForOneCommand)
	//TODO: implement the logic to handle each input
}
