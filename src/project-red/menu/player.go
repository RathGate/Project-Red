package menu

import (
	"fmt"
	"projectRed/utils"
	"strings"
	"time"

	"github.com/mgutz/ansi"
)

var P1 Character

type Character struct {
	Name      string
	Class     string
	Skills    []Skill
	Inventory Inventory
	Stats     Stats
}

type Stats struct {
	Level      int
	Exp        int
	Max_exp    int
	Max_hp     int
	Curr_hp    int
	Max_sp     int
	Curr_sp    int
	Initiative int
	Atk        int
}

func (player *Character) Init(name, class string, atk int, items map[*Item]int, money int, skills []Skill) {
	// BASE CHARACTERISTICS:
	player.Name = name
	player.Class = class
	player.Skills = append(player.Skills, skills...)
	player.Inventory.Items = items
	player.Inventory.Capacity = 10
	player.Inventory.Money = 100

	// STATS:
	player.Stats.Level = 1
	player.Stats.Exp = 0
	player.Stats.Max_exp = 30
	switch class {
	case "Human":
		player.Stats.Max_hp = 100
	case "Elf":
		player.Stats.Max_hp = 80
	case "Dwarf":
		player.Stats.Max_hp = 120
	}
	player.Stats.Curr_hp = player.Stats.Max_hp / 2
	player.Stats.Max_sp, player.Stats.Curr_sp = 50, 50
	player.Stats.Initiative = 10
	player.Stats.Atk = atk
}

func (player *Character) CharacterCreation() {
	utils.ConsoleClear()
	// utils.UPrint(ansi.Color(`"Hello, stranger...`+"\n", "yellow"), 60)
	// time.Sleep(500 * time.Millisecond)
	// utils.UPrint(ansi.Color(`"Not a lot of lost souls come wandering down here...`+"\n", "yellow"), 60)

	// time.Sleep(500 * time.Millisecond)
	// utils.UPrint(ansi.Color("Tell me... ", "yellow"), 60)

	// time.Sleep(300 * time.Millisecond)
	// utils.UPrint(ansi.Color(`What's your name ?"`+"\n", "yellow"), 60)

	// playerName := GetInputStr("name")
	// utils.UPrint(ansi.Color(`"My eyes can't see anymore...`, "yellow"), 60)
	// time.Sleep(500 * time.Millisecond)
	// utils.UPrint(ansi.Color(" Are you\n", "yellow"), 60)
	// time.Sleep(200 * time.Millisecond)
	// utils.UPrint(ansi.Color(`a 'Human' ? `, "yellow"), 60)
	// time.Sleep(200 * time.Millisecond)
	// utils.UPrint(ansi.Color(`An 'Elf' ? `, "yellow"), 60)
	// time.Sleep(200 * time.Millisecond)
	// utils.UPrint(ansi.Color(`Or maybe a 'Dwarf'?"`+"\n", "yellow"), 60)

	// playerClass := GetInputStr("class")
	playerName := "moncul"
	playerClass := "Elf"

	player.Init(playerName, playerClass, 10, map[*Item]int{&Potion: 3, &FairyBottle: 1}, 100, []Skill{Punch})
	P1 = *player

	utils.UPrint(ansi.Color(`"I see. Take this. `, "yellow"), 60)
	time.Sleep(500 * time.Millisecond)
	utils.UPrint(ansi.Color(`Trust me, it'll be useful.`+"\n", "yellow"), 60)
	time.Sleep(500 * time.Millisecond)
	utils.UPrint(ansi.Color(`Don't be shy and take it ! `, "yellow"), 80)
	time.Sleep(500 * time.Millisecond)
	utils.UPrint(ansi.Color(`My old bones`+"\n"+`won't have any use of it anyway."`+"\n", "yellow"), 60)
	time.Sleep(500 * time.Millisecond)
	_ = GetInputInt(0, []int{}, "next")

	utils.UPrint((ansi.Color((utils.Format("๑๑๑ RECEIVED: Fairy Bottle ๑๑๑\n", "center", 50, []string{})), "white+b")), 20)

	time.Sleep(1000 * time.Millisecond)
	fmt.Println()
	utils.UPrint(ansi.Color(`Good Luck on your journey, young one !"`+"\n", "yellow"), 60)
	_ = GetInputInt(0, []int{}, "next")

	time.Sleep(1000 * time.Millisecond)
}

func (char Character) DisplayInfo() {
	fmt.Printf("----- %v'S INFO -----\n", strings.ToUpper(char.Name))
	fmt.Printf("HP: %v/%v\n", char.Stats.Curr_hp, char.Stats.Max_hp)
	fmt.Println("CLASS:", char.Class)
	fmt.Println("LEVEL:", char.Stats.Level)

	_ = GetInputInt(0, []int{0}, "")
}

func (char *Character) IsDead() bool {
	return char.Stats.Curr_hp <= 0
}

func (char *Character) IsRevived() bool {

	item, hasItem := RetrieveItemByName("Fairy Bottle", char.Inventory)
	// IF NO FAIRY IN INVENTORY: GAME OVER
	if !hasItem {
		return false
	}

	// ELSE:

	// Revive message:
	time.Sleep(1500 * time.Millisecond)
	_ = GetInputInt(0, []int{}, "")
	time.Sleep(1500 * time.Millisecond)
	utils.UPrint(ansi.Color(`"Do you really think it's time for a nap...?"`+"\n", "green+b"), 100)
	time.Sleep(250 * time.Millisecond)
	utils.UPrint(ansi.Color((utils.Format(`"Come on, wake up %v ! Wake up !"`, "", 50, []string{char.Name})+"\n\n"), "green+b"), 40)
	time.Sleep(time.Second)

	// Replenishes life:
	char.Stats.Curr_hp = char.Stats.Max_hp / 2
	char.Inventory.RemoveFromInventory(item, 1)

	utils.UPrint(fmt.Sprintf("%v has been revived by the %v !\n\n", char.Name, ansi.Color(item.Name, "green")), 40)
	time.Sleep(time.Second)
	return true
}
