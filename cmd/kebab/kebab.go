package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/veggiemonk/strcase"
)

const usageKebab = `
usage of kebab:
	echo "fooBar" | kebab  # foo-bar
	kebab "foo-bar"  # fooBar
`

func main() {
	// length:  1		 2		 3
	// index:  [0]		[1]	 	[2]
	// prompt: cmd     args...
	l := len(os.Args)
	hasNoArgs := l == 1
	hasOneArg := l == 2
	hasMoreArgs := l >= 3
	if hasNoArgs {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Println(strcase.Kebab(line))
		}
		if err := scanner.Err(); err != nil {
			fmt.Printf("error: %s\n", err)
			os.Exit(1)
		}
		return
	}

	if hasOneArg {
		isHelp := os.Args[1] == "help" || os.Args[1] == "-h" || os.Args[1] == "--help"
		if isHelp {
			fmt.Printf("%s\n", usageKebab)
			os.Exit(0)
		}

		if os.Args[1] == "-" {
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				line := scanner.Text()
				fmt.Println(strcase.Kebab(line))
			}
			if err := scanner.Err(); err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "error: %s\n", err)
				os.Exit(1)
			}
			return
		}

		fmt.Println(strcase.Kebab(os.Args[1]))
		return
	}

	if hasMoreArgs {
		for _, arg := range os.Args[1:] {
			fmt.Println(strcase.Kebab(arg))
		}
	}
}
