package main

import (
	"fmt"
	"project-red/menu"
)

var currentMap *menu.Mapblock

func main() {
	fmt.Printf("\n")
	maps := menu.DeclareMaps()
	currentMap = maps[0]
	currentMap.PrintMap()
	for currentMap.Id != 9 {
		menu.ChangeMap(currentMap)
	}
}
