package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/fatih/color"
)

var (
	executableRegx *regexp.Regexp = regexp.MustCompile(`.*\.exe`)
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		printWarnings("No files to touch specified")
		return
	}

	for _, fileName := range args {
		if executableRegx.Match([]byte(fileName)) {
			msg := fmt.Sprintf(".exe could not be used as a file extension for %s", fileName)
			printWarnings(msg)
			continue
		}

		if _, err := os.Stat(fileName); err == nil {
			msg := fmt.Sprintf("File %s already exists.", fileName)
			printErrors(msg)
			continue
		}

		f, err := os.Create(fileName)
		if err != nil {
			panic(err)
		}

		defer f.Close()
	}

}

func printWarnings(msg string) {
	warn := color.New(color.FgBlack, color.BgRed, color.Bold)
	warnText := color.New(color.Bold)

	warn.Printf("WARN: ")
	warnText.Printf("%s\n", msg)
}

func printErrors(msg string) {
	err := color.New(color.FgBlack, color.BgRed, color.Bold)
	errText := color.New(color.Bold)

	err.Printf("ERROR: ")
	errText.Printf("%s\n", msg)
}
