package menu

import (
	"fmt"
	"projectRed/utils"
	"strconv"

	"github.com/mgutz/ansi"
)

// MENU DISPLAYING THE INVENTORY
func (char *Character) AccessInventory() {
	for {

		utils.PrintBox(P1.Name, "C H A R A C T E R  I N V E N T O R Y", "Blue")
		utils.Format("Money: %v ₽\n\n", "right", 50, []string{strconv.Itoa(char.Inventory.Money)})

		keys := MapKeysToArr(&char.Inventory)
		var position = len(keys)
		if len(char.Inventory.Items) == 0 {
			utils.UPrint(fmt.Sprintln(utils.Format("●●●● E M P T Y ●●●●", "center", 50, []string{})), 20)

		} else {
			for i, key := range keys {
				utils.UPrint(fmt.Sprintf("%v // %v\n", i+1, key.Name), 20)
			}
		}

		answer := GetInputInt(position, []int{}, "")

		if answer == 0 {
			return
		} else {
			selectedItem, count := keys[answer-1], char.Inventory.Items[keys[answer-1]]
			selectedItem.ItemMenu(count, &char.Inventory, "")
		}
	}
}

// ONCE AN ITEM IS SELECTED IN THE INVENTORY MENU:
func (item *Item) ItemMenu(count int, inventory *Inventory, environ string) {
	for {
		utils.UPrint((ansi.Color((utils.Format("๑๑๑ SELECTED: %v (x%v) ๑๑๑", "center", 50, []string{item.Name, strconv.Itoa(count)})), "green+b")), 20)
		fmt.Print("\n\n")

		validAns := []int{0, 2}

		if item.Type == "heal" || item.Type == "book" {
			validAns = append(validAns, 1)
			utils.UPrint("1 // Use\n", 20)
		}

		fmt.Println("2 // Description")

		if environ != "battle" {
			validAns = append(validAns, 3)
			utils.UPrint("3 // Discard\n", 20)
		}

		answer := GetInputInt(0, validAns, "")

		switch answer {
		case 1:
			if inventory.UseItem(item, &P1, "battle") {
				return
			}
		case 2:
			item.DisplayDescription()
		case 3:
			if inventory.DiscardItem(item, count) {
				return
			}
		default:
			return
		}
	}
}
