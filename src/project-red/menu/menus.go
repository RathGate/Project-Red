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
	fmt.Println("\n0 // Quit")

	var answer int = AskUserInt(3, []int{})

	switch answer {
	case 1:
		P1.DisplayInfo()
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
		fmt.Printf("------ %v's INVENTORY: ------\n", strings.ToUpper(p.Name))
		fmt.Printf("%vMoney: %v₽\n\n", strings.Repeat(" ", 28-len(strconv.Itoa(p.Inventory.Money))), p.Inventory.Money)
		var index int
		if len(p.Inventory.Items) == 0 {
			fmt.Println("*** EMPTY ***")
		} else {
			arr := MapToArr(p.Inventory.Items)
			for index = 0; index < len(arr); index++ {
				fmt.Printf("%v // %v (x%v)\n", index+1, arr[index][0].(Item).Name, arr[index][1].(int))
			}
		}

		fmt.Println("\n0 // Quit")
		answer := AskUserInt(index, []int{})
		if answer == 0 {
			return
		} else {
			selectedItem, count := p.Inventory.SelectItem(answer - 1)
			selectedItem.ItemMenu(count, p.Inventory)
		}
	}
}

func (s *ShopKeeper) BuyMenu() {
	var firstBuy bool = true
	for {
		if firstBuy {
			fmt.Print("------------ SHOPKEEPER: -----------\n")
			fmt.Printf("%vMoney: %v₽\n\n", strings.Repeat(" ", 28-len(strconv.Itoa(P1.Inventory.Money))), P1.Inventory.Money)
		}
		var answer int

		// SHOPKEEPER'S INVENTORY IS EMPTY:
		if len(s.Inventory.Items) == 0 {
			fmt.Println(`"I have nothing left for now,`)
			fmt.Println(`please come back another time!"`)
			fmt.Println("\n0 // Quit")

			answer = AskUserInt(0, []int{0})

			// SHOPKEEPER'S INVENTORY IS NOT EMPTY:
		} else {
			if firstBuy {
				fmt.Print(`"Here's what I have for you!"` + "\n\n")
				firstBuy = false
			} else {
				fmt.Print(`"Do you need something else ?"` + "\n\n")
			}

			var index int

			// Prints each item to sell and asks for an input:
			arr := MapToArr(s.Inventory.Items)
			for index = 0; index < len(arr); index++ {
				fmt.Printf("%v // %v %v %v₽\n", index+1, arr[index][0].(Item).Name, strings.Repeat(" ", 25-len(arr[index][0].(Item).Name)), arr[index][0].(Item).Price)
			}

			fmt.Println("\n0 // Quit")
			answer = AskUserInt(index, []int{})
		}

		// PROCESSES THE USER'S INPUT:
		if answer == 0 {
			return
		} else {
			s.SelectShopItem(&P1, answer-1)
			// selectedItem.BuyItem(&P1, s, count)
		}
	}
}
