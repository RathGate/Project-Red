package menu

import (
	"fmt"
	"projectRed/utils"
	"strings"
	"time"

	"github.com/mgutz/ansi"
)

type Equipment struct {
	Head  *Item
	Armor *Item
	Boots *Item
	Sword *Item
}

func (equipment *Equipment) Equip(item *Item) *Item {
	var tmp *Item = nil
	switch item.Type {
	case "head":
		tmp = equipment.Head
		equipment.Head = item
		return tmp
	case "armor":
		tmp = equipment.Armor
		equipment.Armor = item
		return tmp
	case "boots":
		tmp = equipment.Boots
		equipment.Boots = item
		return tmp
	case "sword":
		tmp = equipment.Sword
		equipment.Sword = item
		return tmp
	}
	return tmp
}

func (player *Character) EquipItem(item *Item) bool {
	utils.UPrint(fmt.Sprintf("Do you want to wear %v ?", item.Name), 20)
	fmt.Print("\n\n")
	utils.UPrint("1 // Yes!", 20)

	if ans := GetInputInt(1, []int{}, ""); ans == 0 {
		return false
	}
	prev := player.Equipment.Equip(item)
	player.ApplyBuff(prev, item)
	player.Inventory.RemoveFromInventory(item, 1)
	if prev != nil {
		player.Inventory.AddToInventory(prev, 1)
	}

	utils.UPrint((ansi.Color((utils.Format("๑๑๑ EQUIPED: %v ๑๑๑", "center", 50, []string{strings.ToUpper(item.Name)})), "blue+b")), 20)

	time.Sleep(1000 * time.Millisecond)
	return true
}

func (player *Character) ApplyBuff(previous, current *Item) {
	// REMOVE BUFFS FROM PREVIOUS ARMOR:
	if previous != nil {
		player.Stats.Max_hp -= previous.Effect["HP"].(int)
		player.Stats.Atk -= previous.Effect["Atk"].(int)
	}
	if current != nil {
		player.Stats.Max_hp += current.Effect["HP"].(int)
		player.Stats.Atk += current.Effect["Atk"].(int)
	}
}

func EquipmentToArr() []*Item {
	return []*Item{P1.Equipment.Head, P1.Equipment.Armor, P1.Equipment.Boots, P1.Equipment.Boots}
}
