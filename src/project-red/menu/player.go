package menu

import (
	"fmt"
	"strings"
	"time"
)

var P1 Character

type Character struct {
	Name      string
	Class     string
	Skills    []Item
	Level     int
	Max_hp    int
	Curr_hp   int
	Inventory Inventory
}

func (player *Character) Init() {
	fmt.Println(`"Hello, stranger...`)
	time.Sleep(1500 * time.Millisecond)
	fmt.Println("Not a lot of lost souls comes wandering down here...")
	time.Sleep(2000 * time.Millisecond)
	fmt.Print("Tell me... ")
	time.Sleep(1500 * time.Millisecond)
	fmt.Println(`What's your name ?"`)
	player.Name = GetInputStr("name")
	fmt.Print(`"My eyes can't see anymore...`)
	time.Sleep(1500 * time.Millisecond)
	fmt.Print(" Are you\n")
	fmt.Println(`a 'Human' ? An 'Elf' ? Or maybe a 'Dwarf'?"`)
	player.Class = GetInputStr("class")
	player.Level = 1
	switch player.Class {
	case "Human":
		player.Max_hp = 100
	case "Elf":
		player.Max_hp = 80
	case "Dwarf":
		player.Max_hp = 120
	}
	player.Curr_hp = player.Max_hp / 2
	player.Inventory.Items = map[*Item]int{&PoisonPotion: 5, &Potion: 1}
	player.Inventory.Money = 100
	player.Skills = append(player.Skills, Punch)
	P1 = *player
}

func (char Character) DisplayInfo() {
	fmt.Printf("----- %v'S INFO -----\n", strings.ToUpper(char.Name))
	fmt.Printf("HP: %v/%v\n", char.Curr_hp, char.Max_hp)
	fmt.Println("CLASS:", char.Class)
	fmt.Println("LEVEL:", char.Level)

	_ = GetInputInt(0, []int{0})
}

func (char *Character) IsDead() bool {
	return char.Curr_hp <= 0
}
