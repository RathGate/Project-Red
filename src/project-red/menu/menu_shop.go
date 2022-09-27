package menu

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

var Shop *ShopKeeper

type ShopKeeper struct {
	Character
}

func (shop *ShopKeeper) BuyMenu() {
	var firstBuy bool = true

	for {
		keys := Sorted(&shop.Inventory)

		if firstBuy {
			fmt.Print("------------ SHOPKEEPER: -----------\n\n")
		}
		var answer int

		// SHOPKEEPER'S INVENTORY IS EMPTY:
		if len(shop.Inventory.Items) == 0 {
			fmt.Println(`"I have nothing left for now,`)
			fmt.Println(`please come back another time!"`)

			answer = GetInputInt(0, []int{0})

			// SHOPKEEPER'S INVENTORY IS NOT EMPTY:
			// Prints shopkeeper's dialogue & player money
		} else {
			if firstBuy {
				fmt.Print(`"Here's what I have for you!"` + "\n")
				firstBuy = false
			} else {
				fmt.Print(`"Do you need something else ?"` + "\n")
			}
			fmt.Printf("%vMoney: %v₽\n\n", strings.Repeat(" ", 27-len(strconv.Itoa(P1.Inventory.Money))), P1.Inventory.Money)

			var position int

			// Prints each item to sell and asks for an input:
			for index, element := range keys {
				fmt.Printf("%v // %v %v %v₽\n", index+1, element.Name, strings.Repeat(" ", 25-len(element.Name)), element.Price)
				position++
			}

			// Asks for user's input (an item or "quit")
			answer = GetInputInt(position, []int{})
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
		fmt.Printf(`"%v ? I have %v of them, %v₽ each.`, item.Name, max, item.Price)
		fmt.Print("\n" + `How many do you need ?"` + "\n")
		fmt.Printf("%vMoney: %v₽\n\n", strings.Repeat(" ", 28-len(strconv.Itoa(P1.Inventory.Money))), P1.Inventory.Money)
		count = GetInputInt(max, []int{})
	}

	// Dialogue for more than 1 item:
	if count > 1 {
		fmt.Printf("So it'll be %v₽ for those %v %vs, please.\n", item.Price*count, count, item.Name)

		// Dialogue for 1 item:
	} else {
		fmt.Printf("This %v will cost you %v₽, please.\n", item.Name, item.Price)
	}

	// Prints money, user choices and asks for input:
	fmt.Printf("%vMoney: %v₽\n\n", strings.Repeat(" ", 28-len(strconv.Itoa(P1.Inventory.Money))), P1.Inventory.Money)
	fmt.Print("1 // Ok !")
	answer := GetInputInt(1, []int{})

	if answer == 0 {
		return
	}

	shop.BuyItem(item, count)

}

// AFTER THE PLAYER HAS VALIDATED THEIR CHOICE:
func (shop *ShopKeeper) BuyItem(item *Item, count int) {

	// Check if the P1 has enough money to buy:
	if (item.Price * count) > P1.Inventory.Money {
		fmt.Println(`"Hey, don't buy if you can't pay !`)

		// Checks if the P1 has enough room in the bag to buy:
	} else if invFull, invCount := P1.Inventory.IsFull(); invFull || invCount+count > 10 {
		fmt.Println(`"It seems your bag is too heavy to buy this...`)

		// Buys the item:
	} else {
		// Removes item price from player money:
		P1.Inventory.Money -= item.Price * count
		// Adds the item to player's inventory and removes it from shopkeeper's inventory:
		P1.Inventory.AddToInventory(item, count)
		shop.Inventory.RemoveFromInventory(item, count)
		fmt.Printf("------ BOUGHT %v %v FROM %v ------\n\n", strings.ToUpper(item.Name), count, strings.ToUpper(shop.Name))
		fmt.Println(`It's always a pleasure doing business with you!"`)
	}
	time.Sleep(1000 * time.Millisecond)
}
