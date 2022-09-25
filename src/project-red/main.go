package main

import (
	"projectRed/menu"
)

var Player menu.Character = menu.P1
var ShopKeeper = menu.ShopDude
var GameEnded bool = false

func main() {
	Player.Init()
	// ShopKeeper.DisplayInfo()
	for !GameEnded {
		menu.OpenMenu()
	}
}
