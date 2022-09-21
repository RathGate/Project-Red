package main

import (
	"fmt"
	"project-red/menu"
)

var currentMap *menu.Mapblock

func main() {
	maps := menu.DeclareMaps()

	currentMap = maps[0]
	previousMap := currentMap
	previousMap.Id += 100
	fmt.Println(currentMap.Id, previousMap.Id)
	// menu.ChangeMap(currentMap, previousMap, maps[3])
	// fmt.Println(currentMap, maps)

	type Person struct {
		Name string
		Age  int
	}

	alice1 := Person{"Alice", 30}
	alice2 := alice1
	fmt.Println(alice1 == alice2)   // => true, they have the same field values
	fmt.Println(&alice1 == &alice2) // => false, they have different addresses

	alice2.Age += 10
	fmt.Println(alice1.Age, alice2.Age)

}
