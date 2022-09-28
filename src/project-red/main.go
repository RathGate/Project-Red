package main

import (
	"projectRed/menu"
)

var Player menu.Character = menu.P1
var ShopKeeper = menu.ShopDude
var GameEnded bool = false

func main() {
	// Utilisation du package menu:
	Player.Init()

	var Goblin = &menu.Goblin

	Goblin.InitEnemy("Goblin", 40, 40, 5)

	for !GameEnded {
		menu.OpenMenu()
	}
}
