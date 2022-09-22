package menu

import (
	"fmt"
	"os"
)

var P1 Character

type Character struct {
	Name      string
	Class     string
	Level     int
	Max_hp    int
	Curr_hp   int
	Inventory Inventory
}

func (p *Character) Init(name, class string, level, max_hp, curr_hp int, inventory map[*Item]int) {
	p.Name = name
	p.Class = class
	p.Level = level
	p.Max_hp = max_hp
	p.Curr_hp = curr_hp
	p.Inventory.Items = inventory
	p.Inventory.Money = 100
	P1 = *p
}

func (p Character) DisplayInfo() {
	fmt.Println("Name:", p.Name)
	fmt.Println("Class:", p.Class)
	fmt.Println("Level:", p.Level)
	fmt.Println("Max HP:", p.Max_hp)
	fmt.Println("HP:", p.Curr_hp)
	fmt.Println("Inventory:")
	if len(p.Inventory.Items) == 0 {
		fmt.Println("*** EMPTY ***")
	} else {
		for item, count := range p.Inventory.Items {
			fmt.Printf("   %v: %v\n", item.Name, count)
		}
	}
	fmt.Print("\n")
}

func (p Character) AccessInventory() {
	fmt.Printf("----- %v's INVENTORY: -----\n", p.Name)
	if len(p.Inventory.Items) == 0 {
		fmt.Println("*** EMPTY ***")
	} else {
		for item, count := range p.Inventory.Items {
			fmt.Printf("   %v: %v\n", item.Name, count)
		}
	}
}

func MainMenu() {
	fmt.Println(">>> `D` to Display character information")
	fmt.Println(">>> `I` to display current Inventory")
	fmt.Println(">>> `S` to Shop")
	fmt.Println(">>> `P` to take a Potion")
	fmt.Println(">>> `Q` to Quit game")

	var answer string
	fmt.Scanf("%s\n", &answer)

	switch answer {
	case "D":
		P1.DisplayInfo()
	case "I":
		P1.AccessInventory()
	case "S":
		Shop.BuyMenu()
	case "Q":
		os.Exit(0)
	case "P":
		for item := range P1.Inventory.Items {
			if item.Name == "Poison Potion" {
				P1.Inventory.UseItem(item, &P1)
			}
		}
	}
}

func (c *Character) IsDead() bool {
	return c.Curr_hp <= 0
}
