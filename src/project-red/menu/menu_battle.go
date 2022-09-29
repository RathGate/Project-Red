package menu

import (
	"fmt"
	"projectRed/utils"
	"strconv"
	"strings"
	"time"

	. "github.com/logrusorgru/aurora"
	"github.com/mgutz/ansi"
)

var TurnEnded bool
var DelayedAction map[string]interface{}

// GLOBAL FUNCTION FOR A BATTLE:
func TrainingFight(player *Character, enemy *Enemy) {
	var turn int = 1
	// While neither player nor enemy is dead:
	for ; !player.IsDead() && !enemy.IsDead(); turn++ {
		// Registers the player's action:
		player.RegisterPlayerAction(turn, enemy)

		// LAUNCHES THE TURN:
		BattleIntroduction(*player, *enemy)

		fmt.Println(ansi.Color(strings.Repeat("-", 50), "magenta+b"))
		fmt.Println(utils.Format("TURN %v", "center", 50, []string{strconv.Itoa(turn)}))
		fmt.Println(ansi.Color(strings.Repeat("-", 50), "magenta+b"))
		fmt.Print("\n")

		PrintBattleInfo(*player, *enemy)

		// PLAYER FASTER THAN ENEMY:
		if enemy.Stats.Initiative <= player.Stats.Initiative || DelayedAction["type"] == "run" {
			if outcome := PlayerTurn(turn, player, enemy); outcome == -1 {
				return
			} else if outcome == 1 {
				time.Sleep(500 * time.Millisecond)
				continue
			}
			PrintBattleInfo(*player, *enemy)
			time.Sleep(time.Second)
			if outcome := enemy.EnemyTurn(turn, player); outcome == 1 {
				time.Sleep(500 * time.Millisecond)
				continue
			}
			// ENEMY FASTER THAN PLAYER:
		} else {
			if outcome := enemy.EnemyTurn(turn, player); outcome == 1 {
				time.Sleep(500 * time.Millisecond)
				continue
			}
			PrintBattleInfo(*player, *enemy)
			time.Sleep(time.Second)

			if outcome := PlayerTurn(turn, player, enemy); outcome == -1 {
				return
			} else if outcome == 1 {
				time.Sleep(500 * time.Millisecond)
				continue
			}
		}
		// BATTLE ENDED:
		_ = GetInputInt(0, []int{}, "next")
	}
	GetBattleResults(turn, player, enemy)
	Goblin.Init("Goblin", 40, 40, 5, nil)
	player.Stats.Revert()
}

func (player *Character) RegisterPlayerAction(turn int, enemy *Enemy) {
	TurnEnded = false
	DelayedAction = map[string]interface{}{}

	utils.ConsoleClear()
	BattleIntroduction(*player, *enemy)

	// WHILE PLAYER HAS NOT TAKEN A DECISION:
	for !TurnEnded {
		// BATTLE UI:
		fmt.Print(utils.Format(strings.Repeat("- ", 16), "center", 50, []string{}) + "\n")
		utils.UPrint(ansi.Color(utils.Format(">>> WHAT SHOULD I DO NOW... <<<", "center", 50, []string{}), "white+b")+"\n", 20)
		fmt.Print(utils.Format(strings.Repeat("- ", 16), "center", 50, []string{}) + "\n")
		fmt.Println(Magenta(utils.Format("Next Turn: %v", "right", 50, []string{strconv.Itoa(turn)}) + "\n"))
		PrintBattleInfo(*player, *enemy)
		utils.PrintMenuOpt([]string{"Attack", "Skills", "Items", "Run Away"})

		answer := GetInputInt(0, []int{1, 2, 3, 4}, "battle")

		// BASED ON USER'S CHOICE:
		switch answer {
		case 1:
			fmt.Println(ansi.Color(utils.Format("● ● ● ●| A T T A C K |● ● ● ●\n", "center", 50, []string{}), "yellow"))
			fmt.Printf("You're about to attack with %v atk.\n\n", player.Stats.Atk)
			fmt.Print("1 // Confirm")
			if confirm := GetInputInt(1, []int{}, ""); confirm == 1 {
				DelayedAction["type"] = "attack"
				TurnEnded = true
			}
			continue

		case 2:
			fmt.Println(ansi.Color(utils.Format("● ● ● ●| S K I L L S |● ● ● ●\n", "center", 50, []string{}), "yellow"))
			position := 0
			if len(player.Skills) == 0 {
				utils.UPrint(fmt.Sprintln(utils.Format("●●●● E M P T Y ●●●●", "center", 50, []string{})), 20)
			} else {
				for i, skill := range player.Skills {
					fmt.Printf("%v // %v\n", i+1, skill.Name)
					position++
				}
			}
			input := GetInputInt(position, []int{}, "")
			if input == 0 {
				continue
			}
			player.Skills[input-1].UseSkill(player, enemy, "battle")
			continue

		case 3:
			fmt.Println(ansi.Color(utils.Format("● ● ● ●| I T E M S |● ● ● ●\n", "center", 50, []string{}), "yellow"))
			position := 0
			battleItems := player.Inventory.GetBattleItems()
			if len(battleItems) == 0 {
				utils.UPrint(fmt.Sprintln(utils.Format("●●●● E M P T Y ●●●●", "center", 50, []string{})), 20)
			} else {
				for index, item := range battleItems {
					fmt.Printf("%v // %v (x%v)\n", index+1, item.Name, player.Inventory.Items[item])
					position++
				}
			}
			input := GetInputInt(position, []int{}, "")
			if input == 0 {
				continue
			}

			item, _ := RetrieveItemByName(battleItems[input-1].Name, player.Inventory)
			item.ItemMenu(1, &player.Inventory, "battle")
			continue

		default:
			fmt.Println(ansi.Color(utils.Format("● ● ● ●| R U N  A W A Y |● ● ● ●", "center", 50, []string{}), "yellow"))
			fmt.Println()
			utils.UPrint("This is not a reliable option...\n", 20)
			time.Sleep(250 * time.Millisecond)
			utils.UPrint("Are you sure you wanna try it out ?\n\n", 20)

			fmt.Print("1 // Let's try!")
			input := GetInputInt(1, []int{}, "")
			if input == 1 {
				DelayedAction["type"] = "run"
				TurnEnded = true
			}
			continue
		}
	}
}

// PLAYS THE PREVIOUSLY REGISTERED ACTION BEFORE OR AFTER ENEMY:
func PlayerTurn(turn int, player *Character, enemy *Enemy) int {

	switch DelayedAction["type"] {

	case "attack":
		utils.UPrint(fmt.Sprintf("%v uses Attack against %v !\n", player.Name, enemy.Name), 20)
		time.Sleep(250 * time.Millisecond)
		inflicted := player.Stats.Atk
		if utils.IsCritical(20) == 1 {
			utils.UPrint(ansi.Color("Critical hit !!\n", "red+b"), 20)
			time.Sleep(250 * time.Millisecond)
			inflicted *= 2
		}
		enemy.Stats.Curr_hp -= inflicted
		utils.UPrint(fmt.Sprintf("%v dmg inflicted to %v !\n\n", inflicted, enemy.Name), 20)

	case "skill":
		skill := DelayedAction["skill"].(Skill)
		utils.UPrint(fmt.Sprintf("%v uses the Skill %v.\n", player.Name, skill.Name), 20)
		time.Sleep(1000 * time.Millisecond)
		skill.UseSkill(player, enemy, "delayed")

	case "item":
		item := DelayedAction["item"].(*Item)
		utils.UPrint(fmt.Sprintf("%v uses an item: %v.\n", player.Name, item.Name), 20)
		time.Sleep(1000 * time.Millisecond)
		player.Inventory.UseItem(item, enemy, "delayed")

	case "run":
		utils.UPrint(fmt.Sprintf("%v tries to run away !\n", player.Name), 20)
		time.Sleep(1000 * time.Millisecond)
		if utils.IsCritical(5) == 1 {
			utils.UPrint(ansi.Color(player.Name+" managed to run away successfully !\n", "green+b"), 20)
			time.Sleep(1000 * time.Millisecond)
			return -1
		}
		utils.UPrint(player.Name+" didn't make it to the exit...\n\n", 20)
	}
	time.Sleep(250 * time.Millisecond)

	// CHECKS FOR USER'S AND ENEMY'S HP (1 means one is dead)
	if player.IsDead() {
		utils.UPrint(ansi.Color(utils.Format(player.Name+" has fainted !\n", "center", 50, []string{}), "red+b"), 20)
		fmt.Println()
		if !player.IsRevived() {
			return 1
		}
	} else if enemy.IsDead() {
		utils.UPrint(ansi.Color(utils.Format(enemy.Name+" has fainted !\n", "center", 50, []string{}), "white+b"), 20)
		fmt.Print()
		return 1
	}
	return 0
}

// PLAYS THE ENEMY'S TURN:

func (enemy *Enemy) EnemyTurn(turn int, player *Character) int {
	utils.UPrint(fmt.Sprintf("%v uses Attack against %v !\n", enemy.Name, player.Name), 20)
	time.Sleep(250 * time.Millisecond)
	inflicted := enemy.Stats.Atk
	if turn%3 == 0 {
		utils.UPrint(ansi.Color("Critical hit !!\n", "red+b"), 20)
		time.Sleep(250 * time.Millisecond)
		inflicted *= 2
	}
	player.Stats.Curr_hp -= inflicted

	utils.UPrint(fmt.Sprintf("%v inflicted %v dmg to %v !\n", enemy.Name, inflicted, player.Name), 20)
	fmt.Println()

	if player.IsDead() {
		utils.UPrint(ansi.Color(utils.Format(player.Name+" has fainted !\n", "center", 50, []string{}), "red+b"), 20)
		fmt.Print()
		if !player.IsRevived() {
			return 1
		}
	} else if enemy.IsDead() {
		utils.UPrint(ansi.Color(utils.Format(enemy.Name+" has fainted !\n", "center", 50, []string{}), "white+b"), 20)
		fmt.Println()
		return 1
	}
	return 0
}

func GetBattleResults(turn int, player *Character, enemy *Enemy) {
	won := player.Stats.Curr_hp >= 1
	time.Sleep(250 * time.Millisecond)
	if !won {
		fmt.Println()
		utils.UPrint(ansi.Color(utils.Format("G A M E  O V E R\n", "center", 50, []string{strings.ToUpper(player.Name)}), "red+b"), 100)
	} else {
		fmt.Println()
		utils.UPrint(ansi.Color(utils.Format("%v WINS THE BATTLE !\n", "center", 50, []string{strings.ToUpper(player.Name)}), "cyan+b"), 20)
		fmt.Println()
		time.Sleep(250 * time.Millisecond)
		utils.UPrint(fmt.Sprintf("%v gets %v ₽ and %v exp !\n", player.Name, enemy.Money, enemy.Exp), 20)
		time.Sleep(250 * time.Millisecond)

		_ = GetInputInt(0, []int{}, "next")

		player.Inventory.Money += enemy.Money
		player.BattleReward(enemy.Exp, enemy.Loot)

	}
	_ = GetInputInt(0, []int{}, "")
}
func (player *Character) BattleReward(xp int, item *Item) {

	// XP :::
	player.Stats.Exp += xp
	for player.Stats.Exp > player.Stats.Max_exp {
		player.LevelUp()
	}

	if item != nil {
		utils.UPrint("It seems the opponent carried something with them...\n", 20)
		if full, _ := player.Inventory.IsFull(); !full {
			time.Sleep(250 * time.Millisecond)
			player.Inventory.AddToInventory(item, 1)
			fmt.Println()
			utils.UPrint((ansi.Color((utils.Format("๑๑๑ OBTAINED: %v ๑๑๑\n", "center", 50, []string{item.Name})), "white+b")), 20)
		} else {
			utils.UPrint("But your inventory is full, too bad !", 20)
		}
		time.Sleep(250 * time.Millisecond)
	}
}
