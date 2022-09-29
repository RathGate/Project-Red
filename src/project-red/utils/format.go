package utils

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/Delta456/box-cli-maker/v2"
	"github.com/mgutz/ansi"
)

// Removes all extra whitespaces before and after the
// given string:
func RemoveExtraSpace(str string) string {
	reg := regexp.MustCompile(`^ +| +$/g`)
	return reg.ReplaceAllString(str, "")
}

// Formats the given string according to given rules.
func Format(str string, typeof string, length int, args []string) string {

	for _, arg := range args {
		str = RemoveExtraSpace(strings.Replace(str, "%v", arg, 1))
	}

	// Counts the missing characters between the length of the string
	// and the given length.
	missing := length - utf8.RuneCountInString(str)
	if missing <= 0 {
		missing = 0
	}

	switch typeof {
	// "TITLE" becomes "--- TITLE ---"
	case "title":
		str = " " + str + " "
		str = strings.Repeat("-", missing/2) + str + strings.Repeat("-", missing/2+missing%2)
		// "TITLE" becomes "    TITLE    "
	case "center":
		str = " " + str + " "
		str = strings.Repeat(" ", missing/2) + str + strings.Repeat(" ", missing/2+missing%2)
		// "TITLE" becomes "        TITLE"
	case "right":
		str = strings.Repeat(" ", missing) + str
		// "TITLE" becomes "TITLE        "
	case "left":
		str = str + strings.Repeat(" ", missing)
	}
	return str
}

// Clears the console:
func ConsoleClear() {
	fmt.Print("\033[H\033[2J")
}

// Slowly prints the given string
// 2 chars per [speed]ms.
func UPrint(str string, speed int) {
	for i, char := range str {
		fmt.Print(string(char))
		if i%2 == 1 {
			time.Sleep(time.Duration(speed) * time.Millisecond)
		}
	}
	time.Sleep(200 * time.Millisecond)
}

// Blueprint for the colored boxes:
func PrintBox(title string, content string, color string) {
	ConsoleClear()
	Box := box.New(box.Config{Px: 0, Py: 0, Type: "Double Single", Color: color, TitlePos: "Top"})
	Box.Print((title), Format(content, "center", 48, []string{}))
	fmt.Println()
}

// Prints menu options from a given array.
func PrintMenuOpt(options []string) {
	for i, option := range options {
		UPrint(fmt.Sprintf("%v // %v\n", i+1, option), 20)
	}
}

// Prints the dialogue lines, the user can chose the color
// And the speed of the text.
func NPCLines(str string, color string, speed int) {
	UPrint(ansi.Color(Format(str, "center", 50, []string{})+"\n", color), speed)
}

// Returns success out of [max]
func IsCritical(max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max - 1)
}
