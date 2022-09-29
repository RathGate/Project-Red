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

func (player *Character) RegisterPlayerAction(turn int, enemy *Enemy) {

	TurnEnded = false
	DelayedAction = map[string]interface{}{}
	utils.ConsoleClear()
	BattleIntroduction(*player, *enemy)
	for !TurnEnded {
		fmt.Print(utils.Format(strings.Repeat("- ", 16), "center", 50, []string{}) + "\n")
		utils.UPrint(ansi.Color(utils.Format(">>> WHAT SHOULD I DO NOW... <<<", "center", 50, []string{}), "white+b")+"\n", 20)
		fmt.Print(utils.Format(strings.Repeat("- ", 16), "center", 50, []string{}) + "\n")
		fmt.Println(Magenta(utils.Format("Next Turn: %v", "right", 50, []string{strconv.Itoa(turn)}) + "\n"))

		PrintBattleInfo(*player, *enemy)

		utils.PrintMenuOpt([]string{"Attack", "Skills", "Items", "Run Away"})

		answer := GetInputInt(0, []int{1, 2, 3, 4}, "battle")
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
			max := 0
			for i, skill := range player.Skills {
				fmt.Printf("%v // %v\n", i+1, skill.Name)
				max++
			}
			input := GetInputInt(max, []int{}, "")
			if input == 0 {
				continue
			}
			fmt.Println(player.Skills[input-1].Name + "\n")
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
			input := GetInputInt(1, []int{}, "")
			if input == 1 {
				DelayedAction["type"] = "run"
				TurnEnded = true
			}
			continue
		}
	}

}

func PlayerTurn(turn int, player *Character, enemy *Enemy) int {
	switch DelayedAction["type"] {
	case "attack":
		utils.UPrint(fmt.Sprintf("%v uses Attack against %v !\n", player.Name, enemy.Name), 20)
		time.Sleep(250 * time.Millisecond)
		inflicted := player.Stats.Atk
		if IsCritical(20) {
			utils.UPrint(ansi.Color("Critical hit !!\n", "red+b"), 20)
			time.Sleep(250 * time.Millisecond)
			inflicted *= 2
		}
		enemy.Stats.Curr_hp -= inflicted
		utils.UPrint(fmt.Sprintf("%v dmg inflicted to %v !\n", inflicted, enemy.Name), 20)
	case "skill":
	// skill := DelayedAction["skill"].(Skill)
	// 	if player.Stats.Curr_sp -
	case "item":
		item := DelayedAction["item"].(*Item)
		utils.UPrint(fmt.Sprintf("%v uses an item: %v.\n", player.Name, item.Name), 20)
		time.Sleep(1000 * time.Millisecond)
		player.Inventory.UseItem(item, enemy, "delayed")

	case "run":
		utils.UPrint(fmt.Sprintf("%v tries to run away !\n", player.Name), 20)
		time.Sleep(1000 * time.Millisecond)
		if IsCritical(1) {
			utils.UPrint(ansi.Color(player.Name+" managed to run away successfully !\n", "green+b"), 20)
			time.Sleep(1000 * time.Millisecond)
			return -1
		}
		utils.UPrint(player.Name+" didn't make it to the exit...\n\n", 20)
	}
	time.Sleep(250 * time.Millisecond)

	if player.IsDead() {
		utils.UPrint(ansi.Color(utils.Format(player.Name+" has fainted !\n", "center", 50, []string{}), "red+b"), 20)
		if !player.IsRevived() {
			return 1
		}
	} else if enemy.IsDead() {
		utils.UPrint(ansi.Color(utils.Format(enemy.Name+" has fainted !\n", "center", 50, []string{}), "white+b"), 20)
		return 1
	}
	return 0
}

func IsCritical(max int) bool {
	// rand.Seed(time.Now().UnixNano())

	// return 1+rand.Intn(max-1) == 1
	return false
}

func GetBattleResults(turn int, player *Character, enemy *Enemy) {
	won := player.Stats.Curr_hp >= 1
	time.Sleep(250 * time.Millisecond)
	if !won {
		fmt.Println("Lost !")
	} else {
		utils.UPrint(ansi.Color(fmt.Sprintf("%v WINS THE BATTLE !\n", strings.ToUpper(player.Name)), "cyan+b"), 20)
		time.Sleep(250 * time.Millisecond)
		utils.UPrint(fmt.Sprintf("%v gets %v ₽ and %v exp !\n", player.Name, enemy.Money, enemy.Exp), 20)
		time.Sleep(250 * time.Millisecond)
		// GET LEVEL
	}
	_ = GetInputInt(0, []int{}, "")
}
