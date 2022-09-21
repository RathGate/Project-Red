package main

import (
	"projectRed/menu"
)

var player menu.Personnage = menu.P1

func main() {
	potions := menu.Item{
		Name:   "potions",
		Type:   "heal",
		Effect: 3,
	}
	pouet := menu.Item{
		Name:   "prouts",
		Type:   "heal",
		Effect: 3,
	}
	player.Init("Link", "Royal Knight", 1, 3, 1, map[*menu.Item]int{&potions: 3})
	player.Inventory.AddToInventory(&pouet, 5)
	player.AccessInventory()
	player.Inventory.RemoveFromInventory("proussts", 15)
	player.AccessInventory()
}
