package osfunctions

import (
	"fmt"
	"github.com/Brian-Mashavakure/go-project-maker/pkg/commandfunctions"
	"github.com/Brian-Mashavakure/go-project-maker/pkg/varsrepo"
	"log"
	"os"
	"path/filepath"
)

func createModFile(projectPath string, modulePath string) {
	err := commandfunctions.RunInitCommand(projectPath, modulePath)
	if err != nil {
		log.Fatal(err)
	}

	err2 := commandfunctions.RunModTidyCommand(projectPath)
	if err2 != nil {
		log.Fatal(err2)
	}

}

// function for manually getting project path and module url using fmt for testing i guess
// kept the function because i felt bad deleting it
func getProjectNameAndPath() (path string, module string) {
	fmt.Println("Enter project path and name (with a space inbetween): ")
	var projectName string
	var projectPath string

	_, err := fmt.Scan(&projectPath, &projectName)
	if err != nil {
		fmt.Println("Error ", err)
		return
	}

	fmt.Println("Now enter module path for go mod: ")
	var modulePath string

	_, moduleErr := fmt.Scan(&modulePath)
	if moduleErr != nil {
		fmt.Println("Module Error: ", err)
		return
	}

	finalPath := filepath.Join(projectPath + "\\" + projectName)
	return finalPath, modulePath

}

func MakeProjectDirectories(mainPath string, modulePath string) {

	//creating main directory
	_, err := os.Stat(mainPath)
	if os.IsNotExist(err) {
		err := os.Mkdir(mainPath, 0755)
		if err != nil {
			fmt.Println("Error ", err)
			log.Fatal(err)
		}
	}

	//create cmd directory
	cmdPath := filepath.Join(mainPath + "\\" + "cmd")
	_, cmdErr := os.Stat(cmdPath)
	if os.IsNotExist(cmdErr) {
		cmdErr2 := os.Mkdir(cmdPath, 0755)
		if cmdErr2 != nil {
			log.Fatal(cmdErr2)
		}
	}

	//create pkg directory
	pkgPath := filepath.Join(mainPath + "\\" + "pkg")
	_, pkgErr := os.Stat(pkgPath)
	if os.IsNotExist(pkgErr) {
		pkgErr2 := os.Mkdir(pkgPath, 0755)
		if pkgErr2 != nil {
			log.Fatal(pkgErr2)
		}
	}

	//create bin directory
	binPath := filepath.Join(mainPath + "\\" + "bin")
	_, binErr := os.Stat(binPath)
	if os.IsNotExist(binErr) {
		binErr2 := os.Mkdir(binPath, 0755)
		if binErr2 != nil {
			log.Fatal(binErr2)
		}
	}

	//create tests directory
	testsPath := filepath.Join(mainPath + "\\" + "tests")
	_, testsErr := os.Stat(testsPath)
	if os.IsNotExist(testsErr) {
		testsErr2 := os.Mkdir(testsPath, 0755)
		if testsErr2 != nil {
			log.Fatal(testsErr2)
		}
	}

	//create cmd main directory
	innerMainPath := filepath.Join(cmdPath + "\\" + "main")
	_, innerMainPathErr := os.Stat(innerMainPath)
	if os.IsNotExist(innerMainPathErr) {
		innerMainPathErr2 := os.Mkdir(innerMainPath, 0755)
		if innerMainPathErr2 != nil {
			log.Fatal(innerMainPathErr2)
		}
	}

	//write to main.go file
	mainGoPath := filepath.Join(innerMainPath + "\\" + "main.go")
	innerMainPathErr2 := os.WriteFile(mainGoPath, []byte(varsrepo.MainWriteString), 0660)
	if innerMainPathErr2 != nil {
		log.Fatal(innerMainPathErr2)
	}

	//create route handlers directory
	routesPath := filepath.Join(pkgPath + "\\" + "routehandlers")
	_, routesPathErr := os.Stat(routesPath)
	if os.IsNotExist(routesPathErr) {
		routesPathErr2 := os.Mkdir(routesPath, 0755)
		if routesPathErr2 != nil {
			log.Fatal(routesPathErr2)
		}
	}

	//write to route handlers file
	routeHandlerPath := filepath.Join(routesPath + "\\" + "routehandlers.go")
	routeHandlerPathErr2 := os.WriteFile(routeHandlerPath, []byte(varsrepo.HandlersWriteString), 0660)
	if routeHandlerPathErr2 != nil {
		log.Fatal(routeHandlerPathErr2)
	}

	//creating go.mod file
	createModFile(mainPath, modulePath)

}
