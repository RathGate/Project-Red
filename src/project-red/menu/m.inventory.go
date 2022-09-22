package menu

import (
	"fmt"
	"time"
)

type Item struct {
	Name   string
	Type   string
	Price  int
	Effect map[string]interface{}
}

type Inventory struct {
	Items map[*Item]int
	Money int
}

func (inv *Inventory) UseItem(item *Item, target *Character) {

	// HEALING ITEMS:
	if item.Type == "heal" {

		// If Player HP bar is already full:
		if P1.Curr_hp == P1.Max_hp {
			fmt.Println("No need to take this right now...")
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

		// Deletes Item if its number hits 0 in Inventory:
		inv.Items[item]--
		if inv.Items[item] == 0 {
			delete(inv.Items, item)
		}
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
		inv.Items[item]--
		if inv.Items[item] == 0 {
			delete(inv.Items, item)
		}
		return
	}
}
func (inv *Inventory) AddToInventory(item *Item, count int) {
	invStatus, invCount := inv.IsFull()
	if invStatus {
		fmt.Println("Inventory is full...")
		return
	}
	if invCount+count > 10 {
		fmt.Printf("Mh... The bag is too small to keep all this, so I'll take %v of them.\n", 10-invCount)
		count = 10 - invCount
	}
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
				delete(inv.Items, i)
			}
			return
		}
	}
	fmt.Print("Something went wrong...\n")
}

func (inv Inventory) IsFull() (bool, int) {
	count := 0
	for _, number := range inv.Items {
		count += number
	}
	return (count >= 10), count
}
