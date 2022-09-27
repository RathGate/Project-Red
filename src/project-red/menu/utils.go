package menu

// UTILITARY FUNCTIONS:

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"time"

	"golang.org/x/exp/slices"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// PUTS ALL KEYS OF A MAP INTO AN ARRAY
// Used to dodge the random map keys problem in the inventories:
func Sorted(inv *Inventory) []*Item {
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
func GetInputInt(max int, arr []int) int {
	var answer int
	stdin := bufio.NewReader(os.Stdin)
	fmt.Println("\n0 // Quit")
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
			fmt.Print("I'm afraid my old ears will have you repeat it again...\n")
			time.Sleep(1500 * time.Millisecond)
			continue
		}
		if inputType == "name" && !regexp.MustCompile(`^([a-zA-Z0-9]{2,10})$`).MatchString(answer) {
			DiscardBuffer(stdin)
			fmt.Println()
			fmt.Println(`"I'm afraid my old ears will have you repeat it again..."`)
			time.Sleep(1500 * time.Millisecond)
			fmt.Println("[Not a valid name! (2-10 characters, alphanum only)]")
			continue
		}
		if inputType == "class" && !slices.Contains([]string{"Human", "Elf", "Dwarf"}, answer) {
			fmt.Println()
			DiscardBuffer(stdin)
			fmt.Println(`"I'm afraid my old ears will have you repeat it again..."`)
			time.Sleep(1500 * time.Millisecond)
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
	fmt.Print("------ INIT RESULTS ------\n\n")
	fmt.Printf("NAME: %v\n", char.Name)
	fmt.Printf("CLASS: %v\n", char.Class)
	fmt.Printf("HP: %v/%v\n", char.Curr_hp, char.Max_hp)
	fmt.Printf("MONEY: %v₽\n", char.Inventory.Money)
	fmt.Printf("INVENTORY:\n")
	if len(char.Inventory.Items) == 0 {
		fmt.Println("    --- EMPTY ---")
	} else {
		for item, count := range char.Inventory.Items {
			fmt.Printf("    → %v (x%v)\n", item.Name, count)
		}
	}
	fmt.Println()
	fmt.Printf("SKILLS:\n")
	if len(char.Skills) == 0 {
		fmt.Println("    --- EMPTY ---")
	} else {
		for _, item := range char.Skills {
			fmt.Printf("    → %v\n", item.Name)
		}
	}
	fmt.Println()
}

// USED TO FIND AN ITEM IN AN INVENTORY BASED ON ITS NAME:
func RetrieveItemByName(name string, inventory Inventory) *Item {
	for item := range inventory.Items {
		if name == item.Name {
			return item
		}
	}
	return &Item{Name: "0emp0"}
}
