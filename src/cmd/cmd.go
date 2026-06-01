package cmd

import "os"

type PackageManagerError struct {}
func (e *PackageManagerError) Error() string {
	return "Was not possible to identify the current package manager for this repository"
}

type PackageManager string
const (
	Yarn PackageManager = "yarn"
	Pnpm PackageManager = "pnpm"
	Bun  PackageManager = "bun"
	Npm  PackageManager = "npm"
)

type LockFile struct {
	Filename string
	Manager PackageManager
}
var LockFilesPriority = []LockFile{
 	{Filename: "bun.lock", Manager: Bun},
	{Filename: "pnpm-lock.json", Manager: Pnpm},
	{Filename: "yarn.lock", Manager: Yarn},
	{Filename: "package-lock.json", Manager: Npm},
}

func IdentifyNodePackage() ([]PackageManager, error) {
	pkgs := []PackageManager{}
	for _, entry := range LockFilesPriority {
		if _, err := os.Stat(entry.Filename); err == nil {
			pkgs = append(pkgs, entry.Manager)
		}
	}
	if len(pkgs) == 0 {
		return nil, &PackageManagerError{}
	}
	return pkgs, nil
}

type PackageInstalled struct {
	Pkg PackageManager
}
