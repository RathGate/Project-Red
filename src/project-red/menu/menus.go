package menu

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func OpenMenu() {
	fmt.Print("------ MAIN MENU ------\n\n")
	fmt.Println("1 // Display info")
	fmt.Println("2 // Inventory")
	fmt.Println("3 // Shop")

	var answer int = AskUserInt(3, []int{})

	switch answer {
	case 1:
		// P1.DisplayInfo()
		PrintInfo(&P1)
	case 2:
		P1.AccessInventory()
	case 3:
		ShopDude.BuyMenu()
	case 0:
		fmt.Println("Bisou bébou <3")
		time.Sleep(time.Second)
		os.Exit(0)
		fmt.Println()
	}
}

func (p Character) AccessInventory() {
	for {
		keys := Sorted(&p.Inventory)
		fmt.Printf("------ %v's INVENTORY: ------\n", strings.ToUpper(p.Name))
		fmt.Printf("%vMoney: %v₽\n\n", strings.Repeat(" ", 28-len(strconv.Itoa(p.Inventory.Money))), p.Inventory.Money)
		var position = 0
		if len(p.Inventory.Items) == 0 {
			fmt.Println("*** EMPTY ***")
		} else {
			for index, element := range keys {
				fmt.Printf("%v // %v (x%v)\n", index+1, element.Name, p.Inventory.Items[element])
				position++
			}
		}

		answer := AskUserInt(position, []int{})
		if answer == 0 {
			return
		} else {
			selectedItem, count := keys[answer-1], p.Inventory.Items[keys[answer-1]]
			selectedItem.ItemMenu(count, p.Inventory)
		}
	}
}

func (s *ShopKeeper) BuyMenu() {
	var firstBuy bool = true
	for {
		keys := Sorted(&s.Inventory)
		if firstBuy {
			fmt.Print("------------ SHOPKEEPER: -----------\n\n")
		}
		var answer int

		// SHOPKEEPER'S INVENTORY IS EMPTY:
		if len(s.Inventory.Items) == 0 {
			fmt.Println(`"I have nothing left for now,`)
			fmt.Println(`please come back another time!"`)

			answer = AskUserInt(0, []int{0})

			// SHOPKEEPER'S INVENTORY IS NOT EMPTY:
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
			answer = AskUserInt(position, []int{})
		}

		// PROCESSES THE USER'S INPUT:
		if answer == 0 {
			return
		} else {
			selectedItem, count := keys[answer-1], s.Inventory.Items[keys[answer-1]]
			s.SelectShopItem(selectedItem, count)

			// s.SelectShopItem(&P1, answer-1)
			// selectedItem.BuyItem(&P1, s, count)
		}
	}
}

func (inv *Inventory) DiscardItem(i *Item, count int) {
	answer := count

	if count > 1 {
		fmt.Printf("How many %v do you wanna throw away ? (max %v)\n", i.Name, count)
		answer = AskUserInt(count, []int{})
	}
	if answer == 0 {
		return
	}
	fmt.Printf("You're about to throw %v %v away.\n", answer, i.Name)
	fmt.Print("Are you sure ?\n" + "\n")
	fmt.Print("1 // Ok !")
	confirm := AskUserInt(1, []int{})

	if confirm == 1 {
		inv.RemoveFromInventory(i, answer)
	}
}
