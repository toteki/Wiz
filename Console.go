package wiz

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"github.com/howeyc/gopass" //TODO: dev seems to recommend to switching to package terminal
	"os"
)

//		Colorful console input and output. Also does passwords.

//		*	*	*	*	*	*	*	*	*	*	*	*	*	*	*	*
//		*	*	*	*	*	*	*	*	*	*	*	*	*	*	*	*

//		Exposed functions:
//			SilentPrompt(prompt string) string
//			Prompt(prompt string) string
//				Prompt console for input - silent hides it (like a password)

//			Red(items ...interface{})
//			Yellow(items ...interface{})
//			Green(items ...interface{})
//			Blue(items ...interface{})
//			Purple(items ...interface{})
//			White(items ...interface{})
//				All of these print in a similar (bold first arg, newline each after)

//			Print(items ...interface{})
//				Uses faint/grey text when available. Does not bold first or newline before each arg.

//		*	*	*	*	*	*	*	*	*	*	*	*	*	*	*	*
//		*	*	*	*	*	*	*	*	*	*	*	*	*	*	*	*

// Displays a prompt to the interface and returns what the user enters.
// Hides user input from the console. Use for password entry for example.
func SilentPrompt(prompt string) string {
	defer color.Unset()
	color.Set(color.FgWhite, color.Bold)
	fmt.Println(prompt)
	silentPassword, err := gopass.GetPasswd() // Silent
	if err != nil {
		panic(err)
	}
	return string(silentPassword)
}

// Displays a prompt to the interface and returns what the user enters
func Prompt(prompt string) string {
	defer color.Unset()
	color.Set(color.FgWhite, color.Bold)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(prompt)
	scanner.Scan()
	text := scanner.Text()
	return text
}

//Prints red. The first item passed is bolded. Each item gets a new line.
func Blue(items ...interface{}) {
	defer color.Unset()
	color.Set(color.FgBlue, color.Bold)
	first := true
	for _, element := range items {
		fmt.Println(element)
		if first {
			color.Unset()
			color.Set(color.FgBlue)
			first = false
		}
	}
}

//Prints red. The first item passed is bolded. Each item gets a new line.
func Red(items ...interface{}) {
	defer color.Unset()
	color.Set(color.FgRed, color.Bold)
	first := true
	for _, element := range items {
		fmt.Println(element)
		if first {
			color.Unset()
			color.Set(color.FgRed)
			first = false
		}
	}
}

//Prints green. The first item passed is bolded. Each item gets a new line.
func Green(items ...interface{}) {
	defer color.Unset()
	color.Set(color.FgGreen, color.Bold)
	first := true
	for _, element := range items {
		fmt.Println(element)
		if first {
			color.Unset()
			color.Set(color.FgGreen)
			first = false
		}
	}
}

//Prints yellow. The first item passed is bolded. Each item gets a new line.
func Yellow(items ...interface{}) {
	defer color.Unset()
	color.Set(color.FgYellow, color.Bold)
	first := true
	for _, element := range items {
		fmt.Println(element)
		if first {
			color.Unset()
			color.Set(color.FgYellow)
			first = false
		}
	}
}

//Prints magenta. The first item passed is bolded. Each item gets a new line.
func Purple(items ...interface{}) {
	defer color.Unset()
	color.Set(color.FgMagenta, color.Bold)
	first := true
	for _, element := range items {
		fmt.Println(element)
		if first {
			color.Unset()
			color.Set(color.FgMagenta)
			first = false
		}
	}
}

//Prints white. The first item passed is bolded. Each item gets a new line.
func White(items ...interface{}) {
	defer color.Unset()
	color.Set(color.FgWhite, color.Bold)
	first := true
	for _, element := range items {
		fmt.Println(element)
		if first {
			color.Unset()
			color.Set(color.FgWhite)
			first = false
		}
	}
}

//Prints grey ("Faint white"). Spaces, not newlines, between items passed.
//Windows computers seem to ignore 'faint' option, resulting in regular white.
func Print(items ...interface{}) {
	defer color.Unset()
	color.Set(color.FgWhite, color.Faint)
	for _, element := range items {
		fmt.Print(element, " ")
	}
	fmt.Print("\n")
}
