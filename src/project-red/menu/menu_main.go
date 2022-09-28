package menu

import (
	"fmt"
	"os"
	"projectRed/utils"
	"time"

	"github.com/mgutz/ansi"
)

func OpenMenu() {
	utils.PrintBox(P1.Name, "M A I N  M E N U", "Green")
	var options int = 4
	if !Discovered {
		utils.PrintMenuOpt(MainMenu_Opt)
		options = 5
	} else {
		utils.PrintMenuOpt(MainMenu_Opt[:4])
	}

	var answer int = GetInputInt(options, []int{}, "")

	switch answer {
	case 1:
		// P1.DisplayInfo()
		PrintInfo(&P1)
	case 2:
		P1.AccessInventory()
	case 3:
		ShopDude.BuyMenu()
	case 4:
		TrainingFight(&P1, &Goblin)
	case 5:
		WhoAreThey()
	case 0:
		utils.UPrint(ansi.Color("Are you sure you wanna quit", "red+b"), 30)
		utils.UPrint(ansi.Color("...?\n\n", "red+b"), 100)
		utils.UPrint("1 // YES! ADIOS!", 20)
		answer2 := GetInputInt(1, []int{}, "")

		if answer2 == 1 {
			utils.UPrint(ansi.Color(utils.Format("♥ B I S O U S  B E B O U ♥\n", "center", 50, []string{}), "magenta+b"), 100)
			time.Sleep(time.Second)
			os.Exit(0)
		}

		fmt.Println()
	}
}

func WhoAreThey() {
	utils.PrintBox("Was it even tricky ?", "D I D  Y O U  F I N D  I T ?", "Yellow")
	utils.UPrint(fmt.Sprintln("We did ! "+ansi.Color("Steven Spielberg", "yellow+b")+" and "+ansi.Color("ABBA", "yellow+b")+" were hidden in\nthe powerpoint..."), 20)
	Discovered = true
	_ = GetInputInt(0, []int{}, "")
}

func (inventory *Inventory) DiscardItem(item *Item, count int) bool {
	answer := count

	if count > 1 {
		fmt.Printf("How many %v do you wanna throw away ? (max %v)\n", item.Name, count)
		answer = GetInputInt(count, []int{}, "")
	}
	if answer == 0 {
		return false
	}
	fmt.Printf("You're about to throw %v %v away.\n", answer, item.Name)
	fmt.Print("Are you sure ?\n" + "\n")
	fmt.Print("1 // Ok !")
	confirm := GetInputInt(1, []int{}, "")

	if confirm == 1 {
		if inventory.RemoveFromInventory(item, answer) {
			utils.UPrint(fmt.Sprintf("There's no more %v in the inventory...", item.Name), 5)
		} else {
			utils.UPrint(fmt.Sprintf("Successfully discarded %v %vs from the inventory.", answer, item.Name), 5)
		}
		_ = GetInputInt(0, []int{}, "")
		return true
	}
	return false
}
