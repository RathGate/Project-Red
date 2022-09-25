package menu

import (
	"fmt"
	"strconv"
	"strings"
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

type Inventory struct {
	Items map[*Item]int
	Money int
}

func (inv *Inventory) UseItem(item *Item, target *Character) {
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
		inv.Items[RetrieveItemByName(item.Name, *inv)]--
		// Deletes Item if its number hits 0 in Inventory:
		fmt.Printf("Used: %v - Remains: %v", item.Name, inv.Items[RetrieveItemByName(item.Name, *inv)])
		if inv.Items[RetrieveItemByName(item.Name, *inv)] == 0 {
			delete(inv.Items, RetrieveItemByName(item.Name, *inv))
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
		inv.RemoveFromInventory(item, 1)

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
		inv.RemoveFromInventory(item, 1)
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

func (inv *Inventory) RemoveFromInventory(item *Item, count int) {
	item = RetrieveItemByName(item.Name, *inv)

	inv.Items[item] -= count

	if inv.Items[item] <= 0 {
		delete(inv.Items, item)
		return
	}
}

func (inv Inventory) IsFull() (bool, int) {
	count := 0
	for _, number := range inv.Items {
		count += number
	}
	return (count >= 10), count
}

func (s *ShopKeeper) SelectShopItem(item *Item, max int) {
	var count int = 1
	if max > 1 {
		fmt.Printf(`"%v ? I have %v of them, %v₽ each.`, item.Name, max, item.Price)
		fmt.Print("\n" + `How many do you need ?"` + "\n")
		fmt.Printf("%vMoney: %v₽\n\n", strings.Repeat(" ", 28-len(strconv.Itoa(P1.Inventory.Money))), P1.Inventory.Money)
		count = AskUserInt(max, []int{})
	}
	if count > 1 {
		fmt.Printf("So it'll be %v₽ for those %v %vs, please.\n", item.Price*count, count, item.Name)
	} else {
		fmt.Printf("This %v will cost you %v₽, please.\n", item.Name, item.Price)
	}
	fmt.Printf("%vMoney: %v₽\n\n", strings.Repeat(" ", 28-len(strconv.Itoa(P1.Inventory.Money))), P1.Inventory.Money)

	fmt.Print("1 // Ok !")
	answer := AskUserInt(1, []int{})

	if answer == 0 {
		return
	}
	s.BuyItem(&P1, item, count)

}
func (i Item) ItemMenu(count int, inv Inventory) {
	fmt.Printf("You selected: %v (x%v)\n", i.Name, count)

	validAns := []int{0, 2}

	if i.Type == "heal" || i.Type == "book" {
		validAns = append(validAns, 1)
		fmt.Println("1 // Use")
	}
	fmt.Println("2 // Discard")
	answer := AskUserInt(0, validAns)
	switch answer {
	case 1:
		inv.UseItem(&i, &P1)
	case 2:
		inv.DiscardItem(&i, count)
	default:
		return
	}
}
