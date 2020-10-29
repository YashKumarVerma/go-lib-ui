// Package ui provides few easy to use api methods to write clean outputs with styling support along with emojis
package ui

import (
	"os"

	"github.com/fatih/color"
	"github.com/kyokomi/emoji"
)

// ContextPrint writes the output in the format of "[emoji] : message".
// a complete list of supported emojis available here : http://www.unicode.org/emoji/charts/emoji-list.html
func ContextPrint(emojiName string, message string) {
	var ending string
	if len(message) == 0 {
		ending = ""
	} else {
		ending = "\n"
	}

	emoji.Printf("[ :%s:] %s %s", emojiName, message, ending)
}

// PrintLink is used to style printing links to the console, dimmed, with a hyperlink icon
func PrintLink(context string, link string) {
	styleCyan := color.New(color.FgCyan).SprintFunc()
	emoji.Printf("[ :link:] %s %s \n", context, styleCyan(link))
}

// CheckError shortens your code. Simply pass the error variable into the function,
// the error message your want to display if error was not nil, and finally the
// third parameter describes if you want to halt the execution of the program if err.
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
