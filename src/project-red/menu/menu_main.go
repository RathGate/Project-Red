package menu

import (
	"fmt"
	"os"
	"time"
)

func OpenMenu() {
	fmt.Print("------ MAIN MENU ------\n\n")
	fmt.Println("1 // Display info")
	fmt.Println("2 // Inventory")
	fmt.Println("3 // Shop")

	var answer int = GetInputInt(3, []int{})

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

func (inventory *Inventory) DiscardItem(item *Item, count int) {
	answer := count

	if count > 1 {
		fmt.Printf("How many %v do you wanna throw away ? (max %v)\n", item.Name, count)
		answer = GetInputInt(count, []int{})
	}
	if answer == 0 {
		return
	}
	fmt.Printf("You're about to throw %v %v away.\n", answer, item.Name)
	fmt.Print("Are you sure ?\n" + "\n")
	fmt.Print("1 // Ok !")
	confirm := GetInputInt(1, []int{})

	if confirm == 1 {
		inventory.RemoveFromInventory(item, answer)
	}
}
