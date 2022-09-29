package menu

var Discovered bool

// ITEMS:
var Potion = Item{
	Name:        "Potion",
	Description: `"A magic item that eases pain and heals wounds in the blink\nof an eye... Restores 30HP to its user."`,
	Category:    "item",
	Type:        "heal",
	BattleUse:   true,
	Effect:      map[string]interface{}{"damage": 0.33},
	Price:       50,
}
var PoisonPotion = Item{
	Name:        "Poison Potion",
	Description: `"Uhh... This one seems a little bit doubtful... Are you sure\nyou wanna try it out ?"`,
	Category:    "item",
	Type:        "spell",
	BattleUse:   true,
	Effect:      map[string]interface{}{"type": "status", "time": 0, "damage": 3, "element": "poison"},
	Price:       50,
}
var Iceball = Skill{
	Name:        "Iceball",
	Description: `"Throws an iceball to an enemy`,
	Type:        "attack",
	Damage:      10,
	CostType:    "SP",
	CostInt:     10,
}

var FireBook = Item{
	Name:        "SpellBook - Fireball",
	Description: `"You can learn a lot of things reading books...\nThis one will teach you the skill Iceball."`,
	Category:    "item",
	Type:        "book",
	Price:       25,
	Effect:      map[string]interface{}{"type": "skill", "learn": Fireball},
}

var Fireball = Skill{
	Name:        "Fireball",
	Description: `"Throws a fireball to an enemy.`,
	Type:        "attack",
	Damage:      8,
	CostType:    "SP",
	CostInt:     10,
}

var Punch = Skill{
	Name:        "Punch",
	Description: `"Basically throws a punch in your enemy's face."`,
	Type:        "attack",
	Damage:      8,
	CostType:    "SP",
	CostInt:     1,
}
var TrucExtraCheros = Item{
	Name:        "Truc Super Chéros",
	Description: `"I don't know what it is, but one thing is sure:\nit's expensive."`,
	Price:       120,
}
var OrganiserGuide = Item{
	Name:        "Organiser Guide",
	Description: `"You can learn a lot of things reading books...\nThis one will teach how to organise your bag.\nMax inventory space: +10."`,
	Category:    "item",
	Type:        "book",
	Effect:      map[string]interface{}{"type": "expand"},
}
var ShopDude = ShopKeeper{
	NPC{
		Name:  "Oméga ShopDude",
		Class: "Shop",
		Inventory: Inventory{
			Items: map[*Item]int{&FireBook: 1, &TrucExtraCheros: 1, &Potion: 3, &OrganiserGuide: 5},
			Money: 200,
		},
	},
}

var MainMenu_Opt []string = []string{"Display Info", "Inventory", "Shop", "Training Battle", "Who Are They ?"}
