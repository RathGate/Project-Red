package main

import (
	"projectRed/menu"
)

var Goblin = &menu.Goblin
var Player menu.Character = menu.P1
var ShopKeeper = menu.ShopDude
var GameEnded bool = false

func main() {
	// Init. of the Goblin and the player
	Player.CharacterCreation()
	Goblin.Init("Goblin", 40, 40, 5, &menu.PoisonPotion)
	for !GameEnded {
		menu.OpenMenu()
	}
}
