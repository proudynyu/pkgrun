package main

import (
	"fmt"
	"os/exec"

	"github.com/proudynyu/pkgrun/src/cmd"
	"github.com/proudynyu/pkgrun/src/cwd"
	"github.com/proudynyu/pkgrun/src/file"
	"github.com/proudynyu/pkgrun/src/ui"
)

func main() {
	pkgs, err := cmd.IdentifyNodePackage()
	if err != nil {
		fmt.Printf("An Error has occured: %s\n", err.Error())
		return
	}
	if len(pkgs) <= 0 {
		fmt.Println("No package manager was found for the current repository")
		fmt.Println("Leaving...")
		return
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
		return
	}
	fmt.Printf("-> Using package manager: %s\n", pkgInstalled.Pkg)


	json, err := cwd.FindPackageJson()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}
	jsonFile, err := file.ReadPackageJson(json)
	if err != nil {
		panic(err.Error())
	}
	chosen_cmd := ui.BuildInteractiveCmdChoose(jsonFile)
	exec.Command(string(pkgInstalled.Pkg), "run", chosen_cmd)
}
