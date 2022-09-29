package utils

import (
	"fmt"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/Delta456/box-cli-maker/v2"
	"github.com/mgutz/ansi"
)

func RemoveExtraSpace(str string) string {
	reg := regexp.MustCompile(`^ +| +$/g`)
	return reg.ReplaceAllString(str, "")
}

func Format(str string, typeof string, length int, args []string) string {
	for _, arg := range args {
		str = RemoveExtraSpace(strings.Replace(str, "%v", arg, 1))
	}

	// var formatted []string

	// // FORMATTING:
	// for utf8.RuneCountInString(str) > length {
	// 	formatted = append(formatted, str[:length+1])
	// 	str = str[41:]
	// }
	// formatted = append(formatted, str)
	missing := length - utf8.RuneCountInString(str)
	if missing <= 0 {
		missing = 0
	}
	switch typeof {
	case "title":
		str = " " + str + " "
		str = strings.Repeat("-", missing/2) + str + strings.Repeat("-", missing/2+missing%2)
	case "center":
		str = " " + str + " "
		str = strings.Repeat(" ", missing/2) + str + strings.Repeat(" ", missing/2+missing%2)
	case "right":
		str = strings.Repeat(" ", missing) + str
	case "left":
		str = str + strings.Repeat(" ", missing)
	case "box-right":
	}
	return str
}
func ConsoleClear() {
	fmt.Print("\033[H\033[2J")
}

func UPrint(str string, speed int) {
	for i, char := range str {
		fmt.Print(string(char))
		if i%2 == 1 {
			time.Sleep(time.Duration(speed) * time.Millisecond)
		}
	}
	time.Sleep(200 * time.Millisecond)
}
func PrintBox(title string, content string, color string) {
	ConsoleClear()
	Box := box.New(box.Config{Px: 0, Py: 0, Type: "Double Single", Color: color, TitlePos: "Top"})
	Box.Print((title), Format(content, "center", 48, []string{}))
	fmt.Println()
}
func PrintMenuOpt(options []string) {
	for i, option := range options {
		UPrint(fmt.Sprintf("%v // %v\n", i+1, option), 20)
	}
}

func NPCLines(str string, color string, speed int) {
	UPrint(ansi.Color(Format(str, "center", 50, []string{})+"\n", color), speed)
}
