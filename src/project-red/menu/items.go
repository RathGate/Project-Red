package menu

import (
	"fmt"
	"time"
)

type Item struct {
	Name        string
	Description string
	Category    string
	Type        string
	Price       int
	Effect      map[string]interface{}
}

func (inventory *Inventory) UseItem(item *Item, target *Character) {
	// HEALING ITEMS:
	if item.Type == "heal" {

		// If Player HP bar is already full:
		if P1.Curr_hp == P1.Max_hp {
			fmt.Println("There's no need to take this right now...")
			time.Sleep(2 * time.Second)
			fmt.Println()
			return
		}
		// Else: Take the item
		fmt.Printf("Current HP: %v/%v\n", P1.Curr_hp, P1.Max_hp)
		if P1.Curr_hp+item.Effect["damage"].(int) > P1.Max_hp {
			P1.Curr_hp = P1.Max_hp
		} else {
			P1.Curr_hp += item.Effect["damage"].(int)
		}
		fmt.Printf("Current HP: %v/%v\n", P1.Curr_hp, P1.Max_hp)
		inventory.Items[RetrieveItemByName(item.Name, *inventory)]--
		// Deletes Item if its number hits 0 in Inventory:
		fmt.Printf("Used: %v - Remains: %v", item.Name, inventory.Items[RetrieveItemByName(item.Name, *inventory)])
		if inventory.Items[RetrieveItemByName(item.Name, *inventory)] == 0 {
			delete(inventory.Items, RetrieveItemByName(item.Name, *inventory))
		}
		time.Sleep(2 * time.Second)
		fmt.Print("\n")
		return
	}

	// SPELLS OR SPELL-LIKE ITEMS:
	if item.Type == "spell" {
		switch item.Effect["type"] {
		case "status":
			for i := item.Effect["time"].(int); i > 0 && !target.IsDead(); i-- {
				time.Sleep(1 * time.Second)
				target.Curr_hp -= 1
				fmt.Printf("%v/%v\n", target.Curr_hp, target.Max_hp)
			}
		}
		inventory.RemoveFromInventory(item, 1)

		return
	}

	if item.Type == "book" {
		temp := item.Effect["learn"].(Item)
		for _, i := range P1.Skills {
			if i.Name == temp.Name {
				fmt.Println("You already know that skill...")
				return
			}
		}
		target.Skills = append(target.Skills, temp)
		inventory.RemoveFromInventory(item, 1)
		return
	}
}
