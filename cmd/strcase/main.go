package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime/debug"

	"github.com/veggiemonk/strcase"
)

const usageStrCase = `
command line to convert any string to:

* camelCase or lower camel case: 
  - example: "theQuickBrownFoxJumpsOverTheLazyDog", 'type myInternalType struct {}'
  - usage: Internal (private) variables, functions, methods, and types in Go

* PascalCase or upper camel case: 
  - example: "TheQuickBrownFoxJumpsOverTheLazyDog", 'type MyExportedType struct {}'
  - usage: Exported (public) variables, functions, methods, and types in Go

* snake_case: 
  - example: "the_quick_brown_fox_jumps_over_the_lazy_dog"
  - usage: naming convention in Python.

* dash-case or kebab-case: 
  - "the-quick-brown-fox-jumps-over-the-lazy-dog"
  - usage: naming convention in CSS, also used in HTML and kubernetes manifests.

Library usage:
  see https://github.com/veggiemonk/strcase

Usage of strcase:
  strcase [subcommand] [args...]
  strcase [subcommand] [-]	# read from stdin

  example usage:
	echo "foo-bar" | strcase camel  # fooBar
	strcase camel "foo-bar"  # fooBar

	echo "fooBar" | strcase kebab  # foo-bar
	strcase kebab "fooBar"  # foo-bar

	echo "foo-bar" | strcase pascal # FooBar
	strcase pascal "foo-bar"  # FooBar

	echo "fooBar" | strcase snake # foo_bar
	strcase snake "fooBar"  # foo_bar
`

func main() {
	// length:  1		 2		 3
	// index:  [0]		[1]	 	[2]		[3]
	// prompt: cmd     subcmd   args...
	l := len(os.Args)
	hasNoArgs := l == 1
	hasOneArg := l == 2
	hasMoreArgs := l >= 3

	if hasNoArgs {
		fmt.Printf("%s\n", usageStrCase)
		os.Exit(1)
	}

	if hasOneArg {
		checkHelp()
		checkVersion()
		stdin()
	}

	if hasMoreArgs {
		checkHelp()
		if os.Args[2] == "-" {
			stdin()
		}
		subcmd := os.Args[1]
		for _, arg := range os.Args[2:] {
			switch subcmd {
			case "camel":
				fmt.Println(strcase.Camel(arg))
			case "kebab":
				fmt.Println(strcase.Kebab(arg))
			case "pascal":
				fmt.Println(strcase.Pascal(arg))
			case "snake":
				fmt.Println(strcase.Snake(arg))
			default:
				_, _ = fmt.Fprintf(
					os.Stderr,
					"err: unknown subcommand: %s\n",
					subcmd,
				)
				os.Exit(1)
			}
		}
		return
	}
}

func stdin() {
	subcmd := os.Args[1]
	isHelp := subcmd == "help" || subcmd == "-h" || subcmd == "--help"
	if isHelp {
		fmt.Printf("%s\n", usageStrCase)
		os.Exit(0)
	}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		switch subcmd {
		case "camel":
			fmt.Println(strcase.Camel(line))
		case "kebab":
			fmt.Println(strcase.Kebab(line))
		case "pascal":
			fmt.Println(strcase.Pascal(line))
		case "snake":
			fmt.Println(strcase.Snake(line))
		default:
			_, _ = fmt.Fprintf(
				os.Stderr,
				"err: unknown subcommand: %s\n",
				subcmd,
			)
			os.Exit(1)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("error: %s\n", err)
		os.Exit(1)
	}
	return
}

func checkHelp() {
	if len(os.Args) > 1 {
		subcmd := os.Args[1]
		isHelp := subcmd == "help" || subcmd == "-h" || subcmd == "--help"

		if isHelp {
			fmt.Printf("%s\n", usageStrCase)
			os.Exit(0)
		}
	}
}

func checkVersion() {
	if len(os.Args) > 1 {
		subcmd := os.Args[1]
		isVersion := subcmd == "version" || subcmd == "-v" || subcmd == "--version"

		if isVersion {
			printVersion()
			os.Exit(0)
		}
	}
}

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func printVersion() {
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		_, _ = fmt.Fprintln(os.Stderr, "error reading build-info")
		os.Exit(1)
	}
	fmt.Printf("Build:\n%s\n", bi)
	fmt.Printf("Version: %s\n", version)
	fmt.Printf("Commit: %s\n", commit)
	fmt.Printf("Date: %s\n", date)
}
