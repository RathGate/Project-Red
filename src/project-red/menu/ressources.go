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
	Price: Price{
		Currency: 10,
	},
}
var FairyBottle = Item{
	Name:        "Fairy Bottle",
	Description: `"A small light dances inside the bottle..."` + "\n" + `"What could be the use of it ?"`,
	Category:    "item",
	Type:        "unknown",
	BattleUse:   false,
}
var PoisonPotion = Item{
	Name:        "Poison Potion",
	Description: `"Uhh... This one seems a little bit doubtful... Are you sure\nyou wanna try it out ?"`,
	Category:    "item",
	Type:        "spell",
	BattleUse:   true,
	Effect:      map[string]interface{}{"type": "status", "time": 0, "damage": 3, "element": "poison"},
	Price: Price{
		Currency: 15,
	},
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
	Price: Price{
		Currency: 25,
	},
	BattleUse: false,
	Effect:    map[string]interface{}{"type": "skill", "learn": Fireball},
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
	Price: Price{
		Currency: 120,
	},
	BattleUse: false,
}
var OrganiserGuide = Item{
	Name:        "Organiser Guide",
	Description: `"You can learn a lot of things reading books...\nThis one will teach how to organise your bag.\nMax inventory space: +10."`,
	Category:    "item",
	Type:        "book",
	BattleUse:   false,
	Effect:      map[string]interface{}{"type": "expand"},
	Price: Price{
		Currency: 30,
	},
}
var ShopDude = ShopKeeper{
	NPC{
		Name:  "Oméga ShopDude",
		Class: "Shop",
		Inventory: Inventory{
			Items: map[*Item]int{&FireBook: 1, &TrucExtraCheros: 1, &Potion: 3, &OrganiserGuide: 2, &CrowFeather: 3, &BoarLeather: 3, &WolfHide: 3, &TrollSkin: 3},
		},
	},
}

// CRAFTABLES:
var WolfHide = Item{
	Name:        "Wolf Hide",
	Description: `"A beast pelt that can be used to craft armors."`,
	Category:    "item",
	Type:        "craftable",
	Effect:      map[string]interface{}{"type": "expand"},
	Price: Price{
		Currency: 4,
	},
	BattleUse: false,
}
var TrollSkin = Item{
	Name:        "Troll Skin",
	Description: `"A beast skin that can be used to craft armors."`,
	Category:    "item",
	Type:        "craftable",
	Effect:      map[string]interface{}{"type": "expand"},
	Price: Price{
		Currency: 7,
	},
	BattleUse: false,
}
var BoarLeather = Item{
	Name:        "Boar Leather",
	Description: `"A beast pelt that can be used to craft armors."`,
	Category:    "item",
	Type:        "craftable",
	Effect:      map[string]interface{}{"type": "expand"},
	Price: Price{
		Currency: 3,
	},
	BattleUse: false,
}
var CrowFeather = Item{
	Name:        "Crow Feather",
	Description: `"A feather said to have magical properties.` + "\n" + `Can be used to craft armors."`,
	Category:    "item",
	Type:        "craftable",
	Effect:      map[string]interface{}{"type": "expand"},
	Price: Price{
		Currency: 1,
	},
	BattleUse: false,
}

var MainMenu_Opt []string = []string{"Display Info", "Inventory", "Shop", "BlackSmith", "Training Battle", "Who Are They ?"}

var CapWild = Item{
	Name:        "Cap of the Wild",
	Description: `"This armor was apparently crafted for a hero who travels the wilds.` + "\n" + `Strangely enough, it's just your size."`,
	Category:    "equipment",
	Type:        "head",
	Price:       Price{Currency: 0, Items: map[*Item]int{&WolfHide: 1, &BoarLeather: 1}},
	Effect:      map[string]interface{}{"HP": 5, "Atk": 0},
}
var CapKnight = Item{
	Name:        "Cap of the Knight",
	Description: `"This armor was apparently crafted for a hero who travels the wilds.` + "\n" + `Strangely enough, it's just your size."`,
	Category:    "equipment",
	Type:        "head",
	Price:       Price{Currency: 0, Items: map[*Item]int{&WolfHide: 1, &BoarLeather: 1}},
	Effect:      map[string]interface{}{"HP": 0, "Atk": 5},
}
var TunicWild = Item{
	Name:        "Tunic of the Wild",
	Description: `"This armor was apparently crafted for a hero who travels the wilds.` + "\n" + `Strangely enough, it's just your size."`,
	Category:    "equipment",
	Type:        "armor",
	Price:       Price{Currency: 0, Items: map[*Item]int{&WolfHide: 2, &TrollSkin: 1}},
	Effect:      map[string]interface{}{"HP": 5, "Atk": 0},
}
var TrousersWild = Item{
	Name:        "Trousers of the Wild",
	Description: `"This armor was apparently crafted for a hero who travels the wilds.` + "\n" + `Strangely enough, it's just your size."`,
	Category:    "equipment",
	Type:        "boots",
	Price:       Price{Currency: 0, Items: map[*Item]int{&WolfHide: 2, &BoarLeather: 1}},
	Effect:      map[string]interface{}{"HP": 5, "Atk": 0},
}

var MasterSword = Item{
	Name:        "Master Sword",
	Description: `"A legendary weapon said to have saved the kingdom"` + "\n" + `"a long time ago.`,
	Category:    "equipment",
	Type:        "sword",
	Price:       Price{Currency: 0, Items: map[*Item]int{&CrowFeather: 2}},
	Effect:      map[string]interface{}{"HP": 0, "Atk": 10},
}

var SmithDude = BlackSmith{
	NPC{
		Name:  "Alpha SmithDude",
		Class: "Forge",
		Inventory: Inventory{
			Items: map[*Item]int{&CapWild: 1, &TunicWild: 1, &TrousersWild: 1, &MasterSword: 1},
		},
	},
}
