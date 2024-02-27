package clifunctions

import (
	"flag"
	"fmt"
	"github.com/Brian-Mashavakure/go-project-maker/pkg/osfunctions"
	"os"
)

// MAIN COMMAND LINE FUNCTION
func CliApp() {
	//creating make project command
	makeCmd := flag.NewFlagSet("make", flag.ExitOnError)

	//defining inputs for make command
	pathInput := makeCmd.String("path", "", "Path to where project directory will be stored in system")
	moduleInput := makeCmd.String("module", "", "String value for module url to be used for go mod init")

	//validation check after command input
	if len(os.Args) < 2 {
		fmt.Println("Expecting make sub command")
		os.Exit(1)
	}

	//parsing command arguments for input to main logic function
	makeCmd.Parse(os.Args[2:])

	//running the actual log
	//run switch statement checking second value in the command
	switch os.Args[1] {
	case "make":
		handleMake(makeCmd, pathInput, moduleInput)
	case " ":
		fmt.Print("Arguments required")
		os.Exit(1)
	default:
	}

}

// function to handle the make sub command
func handleMake(makeCmd *flag.FlagSet, path *string, module *string) {
	if *path == " " || *module == " " {
		fmt.Println("All arguments required for command to run")
		makeCmd.PrintDefaults()
		os.Exit(1)
	} else {
		osfunctions.MakeProjectDirectories(*path, *module)
	}

}
