package main

import (
	"projectRed/menu"
)

var Player menu.Character = menu.P1
var ShopKeeper = menu.ShopDude
var GameEnded bool = false

func main() {
	Player.Init()
	for !GameEnded {
		menu.OpenMenu()
	}
}
