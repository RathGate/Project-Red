package menu

import "fmt"

var P1 Personnage

type Item struct {
	Name   string
	Type   string
	Effect interface{}
}

type Inventory struct {
	Items map[*Item]int
}

type Personnage struct {
	Name      string
	Class     string
	Level     int
	Max_hp    int
	Curr_hp   int
	Inventory Inventory
}

func (p *Personnage) Init(name, class string, level, max_hp, curr_hp int, inventory map[*Item]int) {
	p.Name = name
	p.Class = class
	p.Level = level
	p.Max_hp = max_hp
	p.Curr_hp = curr_hp
	p.Inventory.Items = inventory
	P1 = *p
}

func (p Personnage) DisplayInfo() {
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

func (p Personnage) AccessInventory() {
	fmt.Printf("----- %v's INVENTORY: -----\n", p.Name)
	if len(p.Inventory.Items) == 0 {
		fmt.Println("*** EMPTY ***")
	} else {
		for item, count := range p.Inventory.Items {
			fmt.Printf("   %v: %v\n", item.Name, count)
		}
	}
}

func MainMenu(player *Personnage) {
	fmt.Println(">>> `D` to Display character information")
	fmt.Println(">>> `I` to display current Inventory")
	fmt.Println(">>> `P` to take a Potion")
	fmt.Println(">>> `Q` to Quit game")

	var answer string
	fmt.Scanf("%s\n", &answer)

	switch answer {
	case "D":
		fmt.Println("Display")
	case "I":
		player.AccessInventory()
	case "Q":
		fmt.Println("Quit")
	case "P":
		for item := range player.Inventory.Items {
			if item.Name == "potions" {
				player.Inventory.UseItem(item, player)
			}
		}
	}
}

func (inv *Inventory) UseItem(item *Item, P1 *Personnage) {
	fmt.Printf("Current HP: %v/%v\n", P1.Curr_hp, P1.Max_hp)
	if item.Type == "heal" && P1.Curr_hp < P1.Max_hp {
		if P1.Curr_hp+item.Effect.(int) > P1.Max_hp {
			P1.Curr_hp = P1.Max_hp
		} else {
			P1.Curr_hp += item.Effect.(int)
		}
		inv.Items[item]--
		fmt.Printf("Current HP: %v/%v\n", P1.Curr_hp, P1.Max_hp)
	}

	if inv.Items[item] == 0 {
		delete(inv.Items, item)
	}
}
func (inv *Inventory) AddToInventory(item *Item, count int) {
	for i := range inv.Items {
		if item.Name == i.Name {
			inv.Items[i] += count
			return
		}
	}
	inv.Items[item] = count
}

func (inv *Inventory) RemoveFromInventory(item string, count int) {
	for i := range inv.Items {
		if item == i.Name {
			inv.Items[i] -= count

			if inv.Items[i] <= 0 {
				fmt.Printf("No more %v in inventory...\n", i.Name)
				delete(inv.Items, i)
			}
			return
		}
	}
	fmt.Print("Something went wrong...\n")
}
