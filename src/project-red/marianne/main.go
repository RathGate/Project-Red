package main

import (
	"fmt"
	"projectRed/menu"
)

var Player menu.Character = menu.P1
var ShopKeeper menu.ShopKeeper
var GameEnded bool = false

func main() {
	ShopKeeper.Name = "ShopKeeper"
	ShopKeeper.Inventory.Money = 100
	potions2 := menu.Item{
		Name:   "Potion",
		Type:   "heal",
		Effect: map[string]interface{}{"type": "attack", "time": 0, "damage": 3, "element": "fire"},
		Price:  10,
	}
	moncul := map[string]interface{}{
		"type": "status",
		"time": 3, "damage": 3,
		"element": "poison",
	}
	poisonPot := menu.Item{
		Name:   "Poison Potion",
		Type:   "spell",
		Effect: moncul,
		Price:  1000,
	}
	menu.Shop = &ShopKeeper
	ShopKeeper.Inventory.Items = map[*menu.Item]int{&potions2: 1, &poisonPot: 1}
	Player.Init("Link", "Royal Knight", 1, 10, 7, map[*menu.Item]int{&potions2: 10, &poisonPot: 1})
	Player.AccessInventory()
	ShopKeeper.AccessInventory()
	fmt.Print("------\n")
	ShopKeeper.BuyMenu()
	fmt.Print("------\n")
	Player.AccessInventory()
	ShopKeeper.AccessInventory()
}
