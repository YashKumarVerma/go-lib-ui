package ui

import (
	"os"

	"github.com/fatih/color"
	"github.com/kyokomi/emoji"
)

// ContextPrint is used to display message with emoji context
func ContextPrint(emojiName string, message string) {
	var ending string
	if len(message) == 0 {
		ending = ""
	} else {
		ending = "\n"
	}

	emoji.Printf("[ :%s:] %s %s", emojiName, message, ending)
}

// PrintLink with dimmed formatting
func PrintLink(context string, link string) {
	styleCyan := color.New(color.FgCyan).SprintFunc()
	emoji.Printf("[ :link:] %s %s \n", context, styleCyan(link))
}

// CheckError and display error message
func CheckError(e error, message string, quit bool) {
	var emoji string

	if quit {
		emoji = "collision"
	} else {
		emoji = "fire"
	}

	if e != nil {
		ContextPrint(emoji, message)
		if quit {
			os.Exit(1)
		}
	}
}
