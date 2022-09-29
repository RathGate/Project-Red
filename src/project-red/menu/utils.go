package menu

// UTILITARY FUNCTIONS:

import (
	"bufio"
	"fmt"
	"os"
	"projectRed/utils"
	"regexp"
	"strings"
	"time"

	"github.com/mgutz/ansi"
	"golang.org/x/exp/slices"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// PUTS ALL KEYS OF A MAP INTO AN ARRAY
// Used to dodge the random map keys problem in the inventories:
func MapKeysToArr(inv *Inventory) []*Item {
	keys := make([]*Item, 0)
	for k := range inv.Items {
		keys = append(keys, k)
	}
	return keys
}

// USED TO EMPTY BUFFER AND MAKE SUCCEEDING SCANS WORK:
func DiscardBuffer(r *bufio.Reader) {
	r.Discard(r.Buffered())
}

// GET A VALID NUMERIC INPUT FROM USER (max: value if input must be between 0 & max)
// (arr: if valid values are not a chain of numbers ([0, 2, 3] for exemple)
func GetInputInt(max int, arr []int, environ string) int {
	var answer int
	stdin := bufio.NewReader(os.Stdin)
	if environ == "" {
		utils.UPrint("\n0 // Back\n", 20)
	}
	if environ == "next" {
		utils.UPrint("\n0 // Next\n", 20)
	}
	for {
		fmt.Printf("   → ")
		if _, err := fmt.Fscanln(stdin, &answer); err != nil {
			DiscardBuffer(stdin)
			continue
		}

		if max != 0 && (answer < 0 || answer > max) {
			DiscardBuffer(stdin)
			continue
		}
		if len(arr) == 0 && max == 0 && !(answer == 0) {
			DiscardBuffer(stdin)
			continue
		}
		if len(arr) != 0 && !(slices.Contains(arr, answer)) {
			DiscardBuffer(stdin)
			continue
		}
		break
	}
	fmt.Print("\n")
	return answer
}

// GET A VALID STRING INPUT FROM USER:
// (inputType if to display the right errors, used for player.Name and player.Class)
func GetInputStr(inputType string) string {
	var answer string
	stdin := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("   → ")
		if _, err := fmt.Fscanln(stdin, &answer); err != nil {
			DiscardBuffer(stdin)
			fmt.Println()
			utils.UPrint(ansi.Color(`"I'm afraid my old ears will have you repeat it again..."`+"\n", "yellow"), 60)
			time.Sleep(500 * time.Millisecond)
			continue
		}
		if inputType == "name" && !regexp.MustCompile(`^([a-zA-Z]{2,10})$`).MatchString(answer) {
			DiscardBuffer(stdin)
			fmt.Println()
			utils.UPrint(ansi.Color(`"I'm afraid my old ears will have you repeat it again..."`+"\n", "yellow"), 60)
			time.Sleep(500 * time.Millisecond)
			fmt.Println("[Not a valid name! (2-10 characters, letters only)]")
			continue
		}
		if inputType == "class" && !slices.Contains([]string{"human", "elf", "dwarf"}, strings.ToLower(answer)) {
			fmt.Println()
			DiscardBuffer(stdin)
			utils.UPrint(ansi.Color(`"I'm afraid my old ears will have you repeat it again..."`+"\n", "yellow"), 60)
			time.Sleep(500 * time.Millisecond)
			fmt.Print("[Not a valid class! ('Human' / 'Elf' / 'Dwarf')]\n")
			continue
		}
		break
	}
	fmt.Print("\n")
	return cases.Title(language.English).String(answer)
}

// FONCTION DE TEST (pour print infos et inventaires)
func PrintInfo(char *Character) {
	utils.PrintBox(P1.Name, "C H A R A C T E R  I N F O", "Cyan")

	fmt.Println(ansi.Color("NAME: ", "cyan+B") + char.Name)
	fmt.Println(ansi.Color("CLASS: ", "cyan+B") + char.Class)
	fmt.Printf(ansi.Color("HP: ", "cyan+B")+"%v/%v\n", char.Stats.Curr_hp, char.Stats.Max_hp)
	fmt.Printf(ansi.Color("MONEY: ", "cyan+B")+"%v ₽\n", char.Inventory.Money)
	fmt.Printf(ansi.Color("INVENTORY: ", "cyan+B")+"%v slots\n", char.Inventory.Capacity)
	if len(char.Inventory.Items) == 0 {
		utils.UPrint(fmt.Sprintln(utils.Format("●●●● E M P T Y ●●●●", "center", 50, []string{})), 20)
	} else {
		for item, count := range char.Inventory.Items {
			utils.UPrint(fmt.Sprintf("    → %v (x%v)\n", item.Name, count), 15)
		}
	}
	fmt.Println()
	fmt.Println(ansi.Color("SKILLS: ", "cyan+B"))
	if len(char.Skills) == 0 {
		utils.UPrint(fmt.Sprintln(utils.Format("●●●● E M P T Y ●●●●", "center", 50, []string{})), 20)
	} else {
		for _, skill := range char.Skills {
			utils.UPrint(fmt.Sprintf("    → %v\n", skill.Name), 15)
		}
	}
	_ = GetInputInt(0, []int{}, "")
}

// USED TO FIND AN ITEM IN AN INVENTORY BASED ON ITS NAME:
func RetrieveItemByName(name string, inventory Inventory) (*Item, bool) {
	for item := range inventory.Items {
		if name == item.Name {
			return item, true
		}
	}
	return &Item{}, false
}

// USED TO FIND A SKILL IN AN INVENTORY BASED ON ITS NAME:
func RetrieveSkillByName(name string, player *Character) (Skill, bool) {
	var s Skill
	for _, skill := range player.Skills {
		if name == skill.Name {
			return skill, true
		}
	}
	return s, false
}
