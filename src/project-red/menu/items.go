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
	BattleUse   bool
	Effect      map[string]interface{}
}

func (inventory *Inventory) UseItem(item *Item, target *Character, environ string) bool {

	switch item.Type {

	case "heal":
		if P1.Stats.Curr_hp == P1.Stats.Max_hp {
			fmt.Println("There's no need to take this right now...")
			time.Sleep(2 * time.Second)
			fmt.Println()
			return false
		}
		// Else: Take the item
		if environ == "battle" {
			DelayedAction = map[string]interface{}{"action": "item", "item": item, "target": target}
			TurnEnded = true
			return true
		}
		fmt.Printf("Previous HP: %v/%v | ", P1.Stats.Curr_hp, P1.Stats.Max_hp)
		if P1.Stats.Curr_hp+item.Effect["damage"].(int) > P1.Stats.Max_hp {
			P1.Stats.Curr_hp = P1.Stats.Max_hp
		} else {
			P1.Stats.Curr_hp += item.Effect["damage"].(int)
		}
		fmt.Printf("Current HP: %v/%v\n", P1.Stats.Curr_hp, P1.Stats.Max_hp)

	case "spell":
		switch item.Effect["type"] {
		case "status":
			for i := item.Effect["time"].(int); i > 0 && !target.IsDead(); i-- {
				time.Sleep(1 * time.Second)
				target.Stats.Curr_hp -= 1
				fmt.Printf("%v/%v\n", target.Stats.Curr_hp, target.Stats.Max_hp)
			}
		}

	case "book":

		if item.Effect["type"].(string) == "skill" {
			if _, learned := RetrieveSkillByName(item.Effect["learn"].(Skill).Name, target); learned {
				fmt.Println("You already know that skill...")
				return false
			} else {
				target.Skills = append(target.Skills, (item.Effect["learn"].(Skill)))
			}

		} else if item.Effect["type"].(string) == "expand" {
			if inventory.Capacity >= 40 {
				fmt.Println("Hey look at you ! Do you really think your\nback could hold any more items?\n", inventory.Capacity)
				return false
			}
			inventory.UpgradeInventorySlot()
			fmt.Printf("Your bag is bigger now ! It can hold up to %v items !\n", inventory.Capacity)
		}

	}
	inventory.RemoveFromInventory(item, 1)

	_ = GetInputInt(0, []int{}, "")
	return true
}

func (item Item) DisplayDescription() {
	fmt.Println(item.Description)
	_ = GetInputInt(0, []int{}, "")
}
