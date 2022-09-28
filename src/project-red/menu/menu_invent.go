package menu

import (
	"fmt"
	"strconv"
	"strings"
)

// MENU DISPLAYING THE INVENTORY
func (char *Character) AccessInventory() {
	for {
		keys := MapKeysToArr(&char.Inventory)
		fmt.Printf("------ %v's INVENTORY: ------\n", strings.ToUpper(char.Name))
		fmt.Printf("%vMoney: %v₽\n\n", strings.Repeat(" ", 28-len(strconv.Itoa(char.Inventory.Money))), char.Inventory.Money)
		var position = 0
		if len(char.Inventory.Items) == 0 {
			fmt.Println("*** EMPTY ***")
		} else {
			for index, element := range keys {
				fmt.Printf("%v // %v (x%v)\n", index+1, element.Name, char.Inventory.Items[element])
				position++
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
	fmt.Printf("You selected: %v (x%v)\n", item.Name, count)

	validAns := []int{0, 2}

	if item.Type == "heal" || item.Type == "book" {
		validAns = append(validAns, 1)
		fmt.Println("1 // Use")
	}
	fmt.Println("2 // Description")
	if environ != "battle" {
		validAns = append(validAns, 3)
		fmt.Println("3 // Discard")
	}

	answer := GetInputInt(0, validAns, "")
	switch answer {
	case 1:
		inventory.UseItem(item, &P1)
	case 2:
		item.DisplayDescription()
	case 3:
		inventory.DiscardItem(item, count)
	default:
		return
	}
}
