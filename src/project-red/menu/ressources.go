package menu

// ITEMS:
var Potion = Item{
	Name:     "Potion",
	Category: "item",
	Type:     "heal",
	Effect:   map[string]interface{}{"damage": 3},
	Price:    50,
}
var PoisonPotion = Item{
	Name:     "Poison Potion",
	Category: "item",
	Type:     "spell",
	Effect:   map[string]interface{}{"type": "status", "time": 0, "damage": 3, "element": "poison"},
	Price:    50,
}
var Iceball = Item{
	Name:     "Iceball",
	Category: "skill",
	Type:     "spell",
	Effect:   map[string]interface{}{"type": "attack", "time": 0, "damage": 5, "element": "ice"},
	Price:    500,
}

var FireBook = Item{
	Name:     "SpellBook - Fireball",
	Category: "book",
	Price:    25,
}
var Fireball = Item{
	Name:     "Fireball",
	Category: "skill",
	Type:     "spell",
	Effect:   map[string]interface{}{"type": "attack", "time": 0, "damage": 5, "element": "fire"},
}

var Punch = Item{
	Name:     "Punch",
	Category: "skill",
	Type:     "attack",
	Effect:   map[string]interface{}{"damage": 5},
}
var TrucExtraCheros = Item{
	Name:  "Truc Super Chéros",
	Price: 120,
}
var ShopDude = ShopKeeper{
	Character: Character{
		Name:  "Oméga ShopDude",
		Class: "Shop",
		Inventory: Inventory{
			Items: map[*Item]int{&FireBook: 1, &TrucExtraCheros: 1, &Potion: 3},
			Money: 200,
		},
	},
}
