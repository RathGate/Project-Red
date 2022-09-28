package menu

import (
	"fmt"
	"projectRed/utils"
	"strconv"
	"strings"

	. "github.com/logrusorgru/aurora"
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
		fmt.Println((Bold(utils.Format(">>> WHAT SHOULD I DO NOW... <<<", "center", 50, []string{}))))
		fmt.Print(utils.Format(strings.Repeat("- ", 16), "center", 50, []string{}) + "\n")
		fmt.Println(Magenta(utils.Format("Next Turn: %v", "right", 50, []string{strconv.Itoa(turn)}) + "\n"))

		fmt.Println("1 // Attack")
		fmt.Println("2 // Skills")
		fmt.Println("3 // Items")
		fmt.Println("4 // Pass")

		answer := GetInputInt(0, []int{1, 2, 3, 4}, "battle")
		switch answer {
		case 1:
			fmt.Println(utils.Format("● ● ● ●| A T T A C K |● ● ● ●", "center", 50, []string{}))

			input := GetInputInt(1, []int{}, "")
			if input == 1 {
				DelayedAction["type"] = "attack"
				TurnEnded = true
			}
			continue
		case 2:
			fmt.Println(utils.Format("● ● ● ●| S K I L L S |● ● ● ●", "center", 50, []string{}))
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
			fmt.Println(utils.Format("● ● ● ●| I T E M S |● ● ● ●", "center", 50, []string{}))
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

			RetrieveItemByName(battleItems[input-1].Name, player.Inventory).ItemMenu(1, &player.Inventory, "battle")
			continue
		default:
			fmt.Println(utils.Format("● ● ● ●| G U A R D |● ● ● ●", "center", 50, []string{}))
			input := GetInputInt(1, []int{}, "")
			if input == 1 {
				DelayedAction["type"] = "guard"
				TurnEnded = true
			}
			continue
		}
	}

}
