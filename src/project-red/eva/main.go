package main

import (
	"fmt"
	"os"
	//"time"
)

func main() {
	var c1 Character
	c1.init("evz", "human")
	c1.Menu()
}

func (c Character) Menu() {

	fmt.Println("C - Character Info")
	fmt.Println("I - Inventory")
	fmt.Println("Q - Quit")

	var selection string
	fmt.Scanf("%s\n", &selection)

	switch selection {
	case "C":
		c.DisplayInfo()
	case "I":
		c.AccessInventory()
	case "Q":
		c.Quit()
	}
}

func (c *Character) init(Name, Class string) {
	c.Name = Name
	c.Class = Class
	inventory := new(Inventory)
	inventory.init()
	c.Inventory = inventory
	characterstats := new(CharacterStats)
	characterstats.init()
	c.CharacterStats = characterstats
}

func (i *Inventory) init() {
	i.Items = make(map[*Item]int)
	i.Items[&Item{"Redbull"}] = 3
}

func (i Inventory) Display() {
	fmt.Println("----------")
	fmt.Println("Inventory")
	fmt.Println("----------")
	for item, quantity := range i.Items {
		fmt.Println(item.Name, ":", quantity)
	}
}

func (cs *CharacterStats) init() {
	cs.Level = 1
	cs.Hp = 50
	cs.MaxHp = 100
}

func (cs CharacterStats) Display() {
	fmt.Println("----------")
	fmt.Println("Character Stats")
	fmt.Println("----------")
	fmt.Println("Level:", cs.Level)
	fmt.Println("HP:", cs.Hp, "/", cs.MaxHp)
}

func (c Character) DisplayInfo() {
	fmt.Println("Name:", c.Name)
	fmt.Println("Class:", c.Class)
	c.CharacterStats.Display()
	fmt.Println("M - Back to menu")
	var btm string
	fmt.Scanf("%s\n", &btm)

	if btm == "M" {
		c.Menu()
	}
}

func (c Character) AccessInventory() {
	c.Inventory.Display()
	fmt.Println("M - Back to menu")
	//c.TakePot()

	var btm string
	fmt.Scanf("%s\n", &btm)

	if btm == "M" {
		c.Menu()
	}
}

func (c Character) Quit() {
	os.Exit(1)
}

/*func (c Character) Dead() {
	if c.CharacterStats.Hp <= 0 {
		var TryAgain string = "T"
		fmt.Scanln("%s\n", TryAgain)
		fmt.Println(c.Name, "has died, press T to try again.")
	}
}*/

type Inventory struct {
	Items map[*Item]int
}

type Item struct {
	Name string
}

type CharacterStats struct {
	Level int
	Hp    int
	MaxHp int
}

type Character struct {
	Name           string
	Class          string
	Inventory      *Inventory
	CharacterStats *CharacterStats
}
