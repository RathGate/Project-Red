package menu

import (
	"fmt"
	"strings"
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

func (p *Character) Init() {
	// fmt.Println(`"Hello, stranger...`)
	// time.Sleep(1500 * time.Millisecond)
	// fmt.Println("Not a lot of lost souls comes wandering down here...")
	// time.Sleep(2000 * time.Millisecond)
	// fmt.Print("Tell me... ")
	// time.Sleep(1500 * time.Millisecond)
	// fmt.Println(`What's your name ?"`)
	// p.Name = GetInputStr("name")
	// fmt.Print(`"My eyes can't see anymore...`)
	// time.Sleep(1500 * time.Millisecond)
	// fmt.Print(" Are you\n")
	// fmt.Println(`a 'Human' ? An 'Elf' ? Or maybe a 'Dwarf'?"`)
	// p.Class = GetInputStr("class")
	p.Name = "Marianne"
	p.Class = "Human"
	p.Level = 1
	switch p.Class {
	case "Human":
		p.Max_hp = 100
	case "Elf":
		p.Max_hp = 80
	case "Dwarf":
		p.Max_hp = 120
	}
	p.Curr_hp = p.Max_hp / 2
	p.Inventory.Items = map[*Item]int{&PoisonPotion: 5, &Potion: 1}
	p.Inventory.Money = 100
	p.Skills = append(p.Skills, Punch)
	P1 = *p
}

func (p Character) DisplayInfo() {
	fmt.Printf("----- %v'S INFO -----\n", strings.ToUpper(p.Name))
	fmt.Printf("HP: %v/%v\n", p.Curr_hp, p.Max_hp)
	fmt.Println("CLASS:", p.Class)
	fmt.Println("LEVEL:", p.Level)

	_ = AskUserInt(0, []int{0})
}

func (c *Character) IsDead() bool {
	return c.Curr_hp <= 0
}
