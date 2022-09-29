package menu

import (
	"fmt"
	"projectRed/utils"
)

type Equipment struct {
	Head  *Item
	Armor *Item
	Boots *Item
	Sword *Item
}

func (equipment *Equipment) Equip(item *Item) bool {
	currEquiped := equipment.GetPiece(item.Type)
	utils.UPrint(fmt.Sprintf("Do you want to equip %v?\n", item.Name), 20)
	if currEquiped != nil {
		utils.UPrint(fmt.Sprintf("[ %v is currently equiped ]\n", currEquiped.Name), 20)
		fmt.Println()
	}
	utils.UPrint("1 // Yes !", 20)
	answer := GetInputInt(1, []int{}, "")

	if answer == 0 {
		return false
	}

	// Applies and removes the buffs
	// P1.ApplyBuff(currEquiped, item)
	fmt.Println("jpp")
	// Swaps the items:
	p := P1.Equipment.Swap(item)
	if p != nil {
		P1.Inventory.AddToInventory(p, 1)
	}

	return true
}

func (equipment *Equipment) Swap(item *Item) *Item {
	println("Mais wtf")
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
	return nil
}
func (equipment *Equipment) GetPiece(part string) *Item {
	if part == "head" {
		return equipment.Head
	} else if part == "armor" {
		return equipment.Armor
	} else if part == "boots" {
		return equipment.Boots
	} else {
		return equipment.Sword
	}

}

func (player *Character) ApplyBuff(previous, current *Item) {
	// REMOVE BUFFS FROM PREVIOUS ARMOR:
	if previous != nil {
		player.Stats.Max_hp -= previous.Effect["HP"].(int)
		player.Stats.Atk += previous.Effect["Atk"].(int)
	}
	if current != nil {
		player.Stats.Max_hp += previous.Effect["HP"].(int)
		player.Stats.Atk += previous.Effect["Atk"].(int)
	}
}

func EquipmentToMap() map[string]*Item {
	var newmap map[string]*Item = map[string]*Item{
		"Head":  P1.Equipment.Head,
		"Armor": P1.Equipment.Armor,
		"Boots": P1.Equipment.Boots,
		"Sword": P1.Equipment.Boots,
	}
	return newmap
}
