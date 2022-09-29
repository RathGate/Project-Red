package menu

import (
	"fmt"
	"projectRed/utils"
	"time"

	"github.com/mgutz/ansi"
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

func (inventory *Inventory) UseItem(item *Item, target *Enemy, environ string) bool {

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
			DelayedAction = map[string]interface{}{"type": "item", "item": item}
			TurnEnded = true
			return true
		}
		prev := P1.Stats.Curr_hp
		if P1.Stats.Curr_hp+int((item.Effect["damage"].(float64))*float64(P1.Stats.Max_hp)) > P1.Stats.Max_hp {
			P1.Stats.Curr_hp = P1.Stats.Max_hp
		} else {
			P1.Stats.Curr_hp += int((item.Effect["damage"].(float64)) * float64(P1.Stats.Max_hp))
		}
		if environ == "delayed" {
			utils.UPrint(fmt.Sprintf("%v restores %vHP of %v !\n", item.Name, int((item.Effect["damage"].(float64))*float64(P1.Stats.Max_hp)), P1.Name), 20)
			return true
		}
		utils.UPrint((ansi.Color((utils.Format("๑๑๑ USED: %v ๑๑๑\n", "center", 50, []string{item.Name})), "yellow+b")), 20)
		fmt.Println()
		fmt.Printf("Previous HP: %v/%v\nCurrent HP: %v/%v\n", prev, P1.Stats.Max_hp, P1.Stats.Curr_hp, P1.Stats.Max_hp)

	case "spell":
		switch item.Effect["type"] {
		case "status":
			for i := item.Effect["time"].(int); i > 0 && !target.IsDead(); i-- {
				time.Sleep(1 * time.Second)
				target.Stats.Curr_hp -= 1
				fmt.Printf("%v/%v\n", target.Stats.Curr_hp, target.Stats.Max_hp)
			}
		}
		utils.UPrint((ansi.Color((utils.Format("๑๑๑ SELECTED: %v ๑๑๑", "center", 50, []string{item.Name})), "yellow+b")), 20)

	case "book":

		if item.Effect["type"].(string) == "skill" {
			if _, learned := RetrieveSkillByName(item.Effect["learn"].(Skill).Name, &P1); learned {
				fmt.Printf("You already know that skill...\n\n")
				return false
			} else {
				utils.UPrint((ansi.Color((utils.Format("๑๑๑ USED: %v ๑๑๑\n", "center", 50, []string{item.Name})), "yellow+b")), 20)

				fmt.Printf("\nYou learned to use %v !\n", item.Effect["learn"].(Skill).Name)
				P1.Skills = append(P1.Skills, (item.Effect["learn"].(Skill)))
			}

		} else if item.Effect["type"].(string) == "expand" {
			if inventory.Capacity >= 40 {
				fmt.Println("Hey look at you ! Do you really think your\nback could hold any more items?\n", inventory.Capacity)
				return false
			}
			inventory.UpgradeInventorySlot()
			utils.UPrint((ansi.Color((utils.Format("๑๑๑ USED: %v ๑๑๑\n", "center", 50, []string{item.Name})), "yellow+b")), 20)
			fmt.Printf("\nYour bag is bigger now ! It can hold up to %v items !\n", inventory.Capacity)
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
