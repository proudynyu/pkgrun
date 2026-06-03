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

func runProgram(manager string, script string) {
	program := exec.Command(manager, "run", script)
	program.Stderr = os.Stderr
	program.Stdin = os.Stdin
	program.Stdout = os.Stdout
	program.Run()
}

func Run() int {
	pkgs, err := cmd.IdentifyNodePackage()
	if err != nil {
		fmt.Printf("An Error has occured: %s\n", err.Error())
		return 1
	}
	if len(pkgs) <= 0 {
		fmt.Println("No package manager was found for the current repository")
		fmt.Println("Leaving...")
		return 1
	}
	fmt.Printf("-> Package managers identify: %s\n", pkgs)

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
		fmt.Printf("error: %s", err.Error())
		return 1
	}
	jsonFile, err := file.ReadPackageJson(json)
	if err != nil {
		panic(err.Error())
	}
	chosen_cmd := ui.BuildInteractiveCmdChoose(jsonFile)

	runProgram(string(pkgInstalled.Pkg), chosen_cmd)
	return 0
}
