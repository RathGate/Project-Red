package menu

import (
	"fmt"
	"projectRed/utils"
	"strconv"
	"time"

	"github.com/mgutz/ansi"
)

type Item struct {
	Name        string
	Description string
	Category    string
	Type        string
	Price       Price
	BattleUse   bool
	Effect      map[string]interface{}
}

// Should have divided this one, not enough time...
func (inventory *Inventory) UseItem(item *Item, target *Enemy, environ string) bool {

	switch item.Type {

	case "heal":
		// CAN'T TAKE THE ITEM: RETURN
		if P1.Stats.Curr_hp == P1.Stats.Max_hp {
			utils.UPrint("There's no need to take this right now...\n", 20)
			return false
		}
		// ELSE: TAKE THE ITEM
		// If context = battle and action is recorded to be played later
		if environ == "battle" {
			DelayedAction = map[string]interface{}{"type": "item", "item": item}
			TurnEnded = true
			return true
		}
		// Else: actually uses the item
		prev := P1.Stats.Curr_hp

		if P1.Stats.Curr_hp+int((item.Effect["damage"].(float64))*float64(P1.Stats.Max_hp)) > P1.Stats.Max_hp {
			P1.Stats.Curr_hp = P1.Stats.Max_hp
		} else {
			P1.Stats.Curr_hp += int((item.Effect["damage"].(float64)) * float64(P1.Stats.Max_hp))
		}

		// LINES:
		// If in battle:
		if environ == "delayed" {
			utils.UPrint(fmt.Sprintf("%v restores %vHP of %v !\n", item.Name, int((item.Effect["damage"].(float64))*float64(P1.Stats.Max_hp)), P1.Name), 20)
			fmt.Println()
			return true
		}
		// Else:
		utils.UPrint((ansi.Color((utils.Format("๑๑๑ USED: %v ๑๑๑\n", "center", 50, []string{item.Name})), "yellow+b")), 20)
		fmt.Println()
		fmt.Printf("Previous HP: %v/%v\nCurrent HP: %v/%v\n", prev, P1.Stats.Max_hp, P1.Stats.Curr_hp, P1.Stats.Max_hp)

	case "poison":
		if environ == "battle" {
			DelayedAction = map[string]interface{}{"type": "item", "item": item}
			TurnEnded = true
			return true
		} else {
			utils.UPrint("Something is going wrong...\n", 100)
			fmt.Println()
			if environ == "" {
				utils.UPrint((ansi.Color((utils.Format("๑๑๑ USED: %v ๑๑๑\n", "center", 50, []string{item.Name})), "red+b")), 20)
				fmt.Println()
			}
			damage := int(float64(P1.Stats.Max_hp) * item.Effect["damage"].(float64))
			for i := item.Effect["time"].(int); i > 0; i-- {
				time.Sleep(1 * time.Second)
				P1.Stats.Curr_hp -= damage

				if P1.Stats.Curr_hp < 1 {
					P1.Stats.Curr_hp = 1
				}
				fmt.Printf("   → %v HP: %v/%v\n", P1.Name, P1.Stats.Curr_hp, P1.Stats.Max_hp)
				if P1.Stats.Curr_hp == 1 {
					break
				}
			}
			fmt.Println()
			utils.UPrint((ansi.Color((utils.Format("%v recovered from the poison !\n", "center", 50, []string{P1.Name})), "green+b")), 20)
		}

	case "book":
		// SPELLBOOK:
		if item.Effect["type"].(string) == "skill" {
			// Skill already learned:
			if _, learned := RetrieveSkillByName(item.Effect["learn"].(Skill).Name, &P1); learned {
				fmt.Printf("You already know that skill...\n\n")
				return false
				// Else: Learn the skill
			} else {
				utils.UPrint((ansi.Color((utils.Format("๑๑๑ USED: %v ๑๑๑\n", "center", 50, []string{item.Name})), "yellow+b")), 20)

				fmt.Printf("\nYou learned to use %v !\n", item.Effect["learn"].(Skill).Name)
				P1.Skills = append(P1.Skills, (item.Effect["learn"].(Skill)))
			}

			// INVENTORY BOOK
		} else if item.Effect["type"].(string) == "expand" {
			// Capacity at max:
			if inventory.Capacity >= 40 {
				fmt.Println("Hey look at you ! Do you really think your\nback could hold any more items?\n", inventory.Capacity)
				return false
			}
			inventory.UpgradeInventorySlot()
			utils.UPrint((ansi.Color((utils.Format("๑๑๑ USED: %v ๑๑๑\n", "center", 50, []string{item.Name})), "yellow+b")), 20)
			fmt.Printf("\nYour bag is bigger now ! It can hold up to %v items !\n", inventory.Capacity)
		}

	}
	// Removes the item from the inventory
	inventory.RemoveFromInventory(item, 1)

	_ = GetInputInt(0, []int{}, "")
	return true
}

func (item Item) DisplayDescription() {
	fmt.Println(item.Description)
	_ = GetInputInt(0, []int{}, "")
}

func (inventory *Inventory) DiscardItem(item *Item, count int) bool {
	answer := count

	if count > 1 {
		utils.UPrint(fmt.Sprintf("How many %v do you wanna throw away ? (max %v)\n", item.Name, count), 20)
		answer = GetInputInt(count, []int{}, "")
	}
	if answer == 0 {
		return false
	}
	fmt.Printf("You're about to throw %v %vaway.\n", answer, item.Name)
	fmt.Print("Are you sure ?\n" + "\n")
	fmt.Print("1 // Ok !")
	confirm := GetInputInt(1, []int{}, "")

	if confirm == 1 {
		utils.UPrint((ansi.Color((utils.Format("๑๑๑ DISCARDED: %v (x%v) ๑๑๑\n", "center", 50, []string{item.Name, strconv.Itoa(answer)})), "red+b")), 20)
		if inventory.RemoveFromInventory(item, answer) {
			utils.UPrint(fmt.Sprintf("\nThere's no more %v in the inventory...\n", item.Name), 20)
		} else {
			utils.UPrint(fmt.Sprintf("\n%v %v are still in the inventory !\n", strconv.Itoa(count-answer), item.Name), 20)
		}
		_ = GetInputInt(0, []int{}, "")
		return true
	}
	return false
}
