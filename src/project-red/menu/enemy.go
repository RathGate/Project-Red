package menu

type Enemy struct {
	Name  string
	Stats Stats
	Loot  *Item
	Money int
	Exp   int
}

func (enemy *Enemy) InitEnemy(name string, curr_hp, max_hp, akt int) {
	enemy.Name = name
	enemy.Stats.Curr_hp = curr_hp
	enemy.Stats.Max_hp = max_hp
	enemy.Stats.Atk = akt
	enemy.Loot = &PoisonPotion
	enemy.Stats.Initiative = 10
	enemy.Exp = 25
	enemy.Money = 40
}

var Goblin = Enemy{}

func (enemy *Enemy) IsDead() bool {
	return enemy.Stats.Curr_hp <= 0
}
