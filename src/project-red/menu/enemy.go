package menu

// import (
// 	"fmt"
// 	"time"
// )

// type Enemy struct {
// 	Name  string
// 	Stats Stats
// 	Loot  *Item
// }

// func (enemy *Enemy) IsDead() bool {
// 	return enemy.Stats.Curr_hp <= 0
// }
// func (enemy *Enemy) InitEnemy(name string, curr_hp, max_hp, akt int) {
// 	enemy.Name = name
// 	enemy.Stats.Curr_hp = curr_hp
// 	enemy.Stats.Max_hp = max_hp
// 	enemy.Stats.Atk = akt
// 	enemy.Loot = &PoisonPotion
// 	enemy.Stats.Initiative = 10
// }
// func (player *Character) GetBattleInput() {
// 	var turnEnded bool

// 	for !turnEnded {
// 		fmt.Printf(">>> What should I do now... <<<\n\n")
// 		fmt.Println("1 // Attack")
// 		fmt.Println("2 // Skills")
// 		fmt.Println("3 // Items")
// 		fmt.Println("4 // Pass")
// 		answer := GetInputInt(0, []int{1, 2, 3, 4}, "battle")
// 		if answer == 3 {
// 			fmt.Printf("   ### ITEMS ###\n\n")
// 			position := 0
// 			battleItems := player.Inventory.GetBattleItems()
// 			if len(battleItems) == 0 {
// 				fmt.Println("*** EMPTY ***")
// 			} else {
// 				for index, item := range battleItems {
// 					fmt.Printf("%v // %v (x%v)\n", index+1, item.Name, player.Inventory.Items[item])
// 					position++
// 				}
// 			}
// 			input := GetInputInt(position, []int{}, "")
// 			if input == 0 {
// 				continue
// 			}
// 			return
// 		} else if answer == 2 {
// 			fmt.Printf("   ### SKILLS ###" + "\n\n")
// 			max := 0
// 			for i, skill := range player.Skills {
// 				fmt.Printf("%v // %v\n", i+1, skill.Name)
// 				max++
// 			}
// 			input := GetInputInt(max, []int{}, "")
// 			if input == 0 {
// 				continue
// 			}
// 			fmt.Println(player.Skills[input-1].Name + "\n")
// 		}

// 	}
// 	return
// }

// func TrainingFight(player *Character, enemy *Enemy) {
// 	BattleIntroduction(*player, *enemy)
// 	for turn := 0; !player.IsDead() && !enemy.IsDead(); turn++ {
// 		playerAtk := 4
// 		enemyAtk := enemy.EnemyPattern(player)
// 		fmt.Printf("--------- TURN %v --------\n\n", turn)

// 		if enemy.Stats.Initiative < player.Stats.Initiative {
// 			fmt.Println(playerAtk, enemyAtk)
// 		} else {
// 			fmt.Println(enemyAtk, playerAtk)
// 		}
// 	}
// }

// func BattleIntroduction(player Character, enemy Enemy) {
// 	fmt.Print("-------- AN ENEMY APPEARS !! --------\n" + "\n")
// 	fmt.Printf("An agressive %v ambushed you ! ", enemy.Name)
// 	if enemy.Loot != nil {
// 		fmt.Print("They\nseem to carry something interesting...\n\n")
// 	}
// 	time.Sleep(1000 * time.Millisecond)
// }
// func (enemy *Enemy) EnemyPattern(player *Character) int {
// 	// for i := 1; i < 3; i++ {
// 	// 	damage := enemy.Akt
// 	// 	if i%3 == 0 {
// 	// 		damage = enemy.Akt * 2
// 	// 		fmt.Print("Critical hit ! ")
// 	// 	}
// 	// 	player.Curr_hp -= damage
// 	// 	fmt.Printf("%v deals %vHP damage to %v !\n", enemy.Name, damage, player.Name)
// 	// 	fmt.Println(player.Curr_hp)
// 	// }
// 	return 69
// }

// func (inventory *Inventory) GetBattleItems() []*Item {
// 	list := []*Item{}
// 	for _, item := range MapKeysToArr(inventory) {
// 		if item.BattleUse {
// 			list = append(list, item)
// 		}
// 	}
// 	return list
// }

// var Goblin = Enemy{}
