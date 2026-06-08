package app

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/proudynyu/pkgrun/src/cmd"
	"github.com/proudynyu/pkgrun/src/cwd"
	"github.com/proudynyu/pkgrun/src/file"
	"github.com/proudynyu/pkgrun/src/ui"
)

func runProgram(manager string, script string) error {
	program := exec.Command(manager, "run", script)
	program.Stderr = os.Stderr
	program.Stdin = os.Stdin
	program.Stdout = os.Stdout
	return program.Run()
}

func Run() int {
	pkgs, err := cmd.IdentifyNodePackage()
	if err != nil {
		fmt.Printf("An Error has occurred: %s\n", err.Error())
		return 1
	}
	if len(pkgs) <= 0 {
		fmt.Println("No package manager was found for the current repository")
		fmt.Println("Leaving...")
		return 1
	}
	fmt.Printf("-> Package managers identified: %s\n", pkgs)

	pkgInstalled := &cmd.PackageInstalled{Pkg: ""}
	for _, pkg := range pkgs {
		if cwd.PackageIsInstalled(string(pkg)) {
			pkgInstalled.Pkg = pkg
			break
		}
	}

	if pkgInstalled.Pkg == "" {
		fmt.Println("No package manager was found for the current repository")
		return 1
	}

	fmt.Printf("-> Using package manager: %s\n", pkgInstalled.Pkg)
	json, err := cwd.FindPackageJson()
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		return 1
	}

	jsonFile, err := file.ReadPackageJson(json)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
		return 1
	}

	chosenCmd := ui.BuildInteractiveCmdChoose(jsonFile)
	if chosenCmd == "" {
		fmt.Println("No scripts found in package.json")
		return 1
	}

	err = runProgram(string(pkgInstalled.Pkg), chosenCmd)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to run script: %v\n", err)
		return 1
	}
	return 0
}
