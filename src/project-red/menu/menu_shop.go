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

var Shop *ShopKeeper

type ShopKeeper struct {
	NPC
}
type NPC struct {
	Name      string
	Class     string
	Inventory Inventory
}

func (shop *ShopKeeper) BuyMenu() {
	var firstBuy bool = true

	for {
		utils.ConsoleClear()
		keys := MapKeysToArr(&shop.Inventory)

		Box := box.New(box.Config{Px: 0, Py: 0, Type: "Double Single", Color: "Magenta", TitlePos: "Top"})
		Box.Print((P1.Name), utils.Format("O M E G A S H O P D U D E", "center", 48, []string{}))

		var answer int

		// SHOPKEEPER'S INVENTORY IS EMPTY:
		if len(shop.Inventory.Items) == 0 {
			fmt.Println(`"I have nothing left for now,`)
			fmt.Println(`please come back another time!"`)

			answer = GetInputInt(0, []int{0}, "")

			// SHOPKEEPER'S INVENTORY IS NOT EMPTY:
			// Prints shopkeeper's dialogue & player money
		} else {
			if firstBuy {
				utils.NPCLines(`"Here's what I have for you!"`, "magenta+b", 20)
				firstBuy = false
			} else {
				utils.NPCLines(`"Do you need something else ?"`, "magenta+b", 20)
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
			selectedItem, count := keys[answer-1], shop.Inventory.Items[keys[answer-1]]
			shop.SelectShopItem(selectedItem, count)
		}
	}
}

// ONCE AN ITEM IS SELECTED IN THE SHOP: (max = max amount of item in shop)
func (shop *ShopKeeper) SelectShopItem(item *Item, max int) {
	var count int = 1
	// If more than 1 item in shop, asks for amount needed:
	if max > 1 {
		utils.NPCLines(fmt.Sprintf(`"%v ? I have %v of them, %v ₽ each."`, item.Name, max, item.Price.Currency), "magenta+b", 20)
		utils.NPCLines(` do you need ?"`+"\n", "magenta+b", 20)
		fmt.Println(utils.Format("Money: %v ₽", "right", 50, []string{strconv.Itoa(P1.Inventory.Money)}))
		count = GetInputInt(max, []int{}, "")
	}

	// Dialogue for more than 1 item:
	if count > 1 {
		utils.NPCLines(fmt.Sprintf("So it'll be %v ₽ for those %v %vs, please.\n", (item.Price.Currency*count), count, item.Name), "magenta+b", 20)

		// Dialogue for 1 item:
	} else {
		utils.NPCLines(fmt.Sprintf("This %v will cost you %v ₽, please.\n", item.Name, item.Price.Currency), "magenta+b", 20)
	}

	// Prints money, user choices and asks for input:
	fmt.Println(utils.Format("Money: %v ₽", "right", 50, []string{strconv.Itoa(P1.Inventory.Money)}))
	fmt.Print("1 // Ok !")
	answer := GetInputInt(1, []int{}, "")

	if answer == 0 {
		return
	}

	shop.BuyItem(item, count)

}

// AFTER THE PLAYER HAS VALIDATED THEIR CHOICE:
func (shop *ShopKeeper) BuyItem(item *Item, count int) {

	// Check if the P1 has enough money to buy:
	if (item.Price.Currency * count) > P1.Inventory.Money {
		utils.NPCLines(`"Hey, don't buy if you can't pay !`+"\n", "magenta+b", 20)

		// Checks if the P1 has enough room in the bag to buy:
	} else if invFull, invCount := P1.Inventory.IsFull(); invFull || invCount+count > 10 {
		utils.NPCLines(`"It seems your bag is too heavy to buy this...`+"\n", "magenta+b", 20)

		// Buys the item:
	} else {
		// Removes item price from player money:
		P1.Inventory.Money -= item.Price.Currency * count
		// Adds the item to player's inventory and removes it from shopkeeper's inventory:
		P1.Inventory.AddToInventory(item, count)
		shop.Inventory.RemoveFromInventory(item, count)
		utils.UPrint((ansi.Color((utils.Format("๑๑๑ BOUGHT: %v %v(s) ๑๑๑\n", "center", 50, []string{strconv.Itoa(count), strings.ToUpper(item.Name)})), "blue+b")), 20)
		fmt.Println()
		utils.NPCLines(`It's always a pleasure doing business with you!"`+"\n", "magenta+b", 20)
	}
	time.Sleep(1500 * time.Millisecond)
}
