package main

import (
	"fmt"

	"github.com/proudynyu/pkgrun/src/file"
	"github.com/proudynyu/pkgrun/src/cmd"
	"github.com/proudynyu/pkgrun/src/cwd"
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

	pkg_installed := &cmd.PackageInstalled{Pkg: ""}
	for _, pkg := range pkgs {
		if cwd.PackageIsInstalled(string(pkg)) {
			pkg_installed.Pkg = pkg
		}
	}

	if pkg_installed.Pkg == "" {
		fmt.Println("No package manager was found for the current repository")
		return
	}
	fmt.Printf("-> Using package manager: %s\n", pkg_installed.Pkg)


	json, err := cwd.FindPackageJson()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}
	json_file := file.ReadPackageJson(json)
	ui.BuildInteractiveCmdChoose(json_file)
}
