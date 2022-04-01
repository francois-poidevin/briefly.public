//go:build mage
// +build mage

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var (
	Default = Build
	goFiles = getGoFiles()
)

// Runs go mod download and then installs the binary.
func Build() {
	banner := figure.NewFigure("Briefly public", "", true)
	banner.Print()

	fmt.Println("")

	color.Red("# Pre Build ------------------------------------------------------------")
	mg.SerialDeps(Go.Format, Go.Tidy, Go.Deps)

	fmt.Println("")
	color.Red("# Build Artefact ----------------------------------------------------------------")
	mg.Deps(Bin.BrieflyPublic)
}

type Go mg.Namespace

// Tidy add/remove depenedencies.
func (Go) Tidy() error {
	color.Cyan("## Tidy go modules")
	return sh.RunV("go", "mod", "tidy", "-v")
}

// Deps install dependency tools.
func (Go) Deps() error {
	color.Cyan("## Vendoring dependencies")
	return sh.RunV("go", "mod", "vendor")
}

// Format runs gofmt on everything
func (Go) Format() error {
	color.Cyan("## Format everything")
	args := []string{"-w"}
	args = append(args, goFiles...)
	return sh.RunV("gofumpt", args...)
}

type Bin mg.Namespace

func (Bin) BrieflyPublic() error {
	return goBuild("briefly.public")
}

func goBuild(packageName string) error {
	color.Green(" > Building %s\n", packageName)

	return sh.RunV("go", "build", "-o", fmt.Sprintf("bin/%s", packageName))
}

func getGoFiles() []string {
	var goFiles []string

	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, "vendor/") {
			return filepath.SkipDir
		}

		if !strings.HasSuffix(path, ".go") {
			return nil
		}

		goFiles = append(goFiles, path)
		return nil
	})

	return goFiles
}
