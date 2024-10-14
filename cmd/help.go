package cmd

import "fmt"

func cmdHelp() {
	helpText := `Usage: ./public-holidays-cli [OPTIONS]

Start the public holidays API.

Options:
  -h, --help      Show this help text and exit.
  -v, --version   Display the version of the application and exit.
  -j, --json      Output the version as JSON.

Examples:
  ./public-holidays-cli`
	fmt.Println(helpText)
}
