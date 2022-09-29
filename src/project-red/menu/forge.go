package menu

import (
	"fmt"
	"projectRed/utils"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/Delta456/box-cli-maker/v2"
	"github.com/mgutz/ansi"
)

type Price struct {
	Currency int
	Items    map[*Item]int
}
type BlackSmith struct {
	NPC
}

func (forge *BlackSmith) BuyMenu() {
	var firstBuy bool = true

	for {
		utils.ConsoleClear()
		keys := MapKeysToArr(&forge.Inventory)

		Box := box.New(box.Config{Px: 0, Py: 0, Type: "Double Single", Color: "White", TitlePos: "Top"})
		Box.Print((P1.Name), utils.Format("A L P H A S M I T H G U Y", "center", 48, []string{}))

		var answer int

		// SHOPKEEPER'S INVENTORY IS EMPTY:
		if len(forge.Inventory.Items) == 0 {
			fmt.Println(`"I have nothing left for now,`)
			fmt.Println(`please come back another time!"`)

			answer = GetInputInt(0, []int{0}, "")

			// SHOPKEEPER'S INVENTORY IS NOT EMPTY:
			// Prints shopkeeper's dialogue & player money
		} else {
			if firstBuy {
				utils.NPCLines(`"Here's what I can craft for you!"`, "white+b", 20)
				firstBuy = false
			} else {
				utils.NPCLines(`"Do you need something else ?"`, "white+b", 20)
			}
			fmt.Println("\n" + utils.Format("Money: %v ₽", "right", 50, []string{strconv.Itoa(P1.Inventory.Money)}))

			var position int

			// Prints each item to sell and asks for an input:
			for index, element := range keys {
				length := utf8.RuneCountInString(fmt.Sprintf("%v // %v", index+1, element.Name))
				fmt.Println(fmt.Sprintf("%v // %v", index+1, element.Name) + utils.Format("%v ₽", "right", 50-length, []string{strconv.Itoa(element.Price.Currency)}))

				position++
			}

			// Asks for user's input (an item or "quit")
			answer = GetInputInt(position, []int{}, "")
		}

		// PROCESSES THE USER'S INPUT:
		if answer == 0 {
			return
		} else {
			selectedItem := keys[answer-1]
			forge.SelectSmithItem(selectedItem)
		}
	}
}

func (forge *BlackSmith) SelectSmithItem(item *Item) {
	utils.NPCLines(fmt.Sprintf(`"You want me to craft this: %v ?"`, item.Name), "white+b", 20)
	fmt.Println()
	utils.NPCLines(fmt.Sprintf(`"I need %v ₽ and these materials to start working:"`, item.Price.Currency), "white+b", 20)

	for key := range item.Price.Items {
		utils.UPrint(fmt.Sprintf("    → %v (x%v)\n", key.Name, item.Price.Items[key]), 20)
	}
	fmt.Println()
	utils.UPrint("1 // Ok !", 20)
	answer := GetInputInt(1, []int{}, "")

	if answer == 0 {
		return
	}

	forge.CraftItem(item)
}

func (forge *BlackSmith) CraftItem(item *Item) {
	if item.Price.Currency > P1.Inventory.Money {
		utils.NPCLines(`"Hey, don't buy if you can't pay !`+"\n", "white+b", 20)
	} else if !P1.CanCraft(item) {
		utils.NPCLines(`"You don't have enough materials to craft this.`, "white+b", 20)
	} else {
		// Deletes items from inventory:
		for i, count := range item.Price.Items {
			P1.Inventory.RemoveFromInventory(i, count)
		}
		P1.Inventory.AddToInventory(item, 1)
		P1.Inventory.Money -= item.Price.Currency

		utils.UPrint((ansi.Color((utils.Format("๑๑๑ CRAFTED: %v ๑๑๑\n", "center", 50, []string{strings.ToUpper(item.Name)})), "cyan+b")), 20)
		fmt.Println()
		utils.NPCLines(`"It's always a pleasure doing business with you!"`+"\n", "white+b", 20)
	}
	time.Sleep(1000 * time.Millisecond)
	_ = GetInputInt(0, []int{}, "")
}

func (player *Character) CanCraft(item *Item) bool {
	for material := range item.Price.Items {
		match, exists := RetrieveItemByName(material.Name, P1.Inventory)
		if !exists || P1.Inventory.Items[match] < item.Price.Items[material] {
			return false
		}
	}
	return true

}
