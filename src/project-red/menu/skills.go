package menu

import (
	"fmt"
	"projectRed/utils"
	"time"

	"github.com/mgutz/ansi"
)

type Skill struct {
	Name        string
	Description string
	CostType    string
	CostInt     int
	Damage      int
	Type        string
}

func (skill Skill) UseSkill(player *Character, target *Enemy, environ string) bool {
	switch skill.Type {
	case "attack":
		if player.Stats.Curr_sp-skill.CostInt < 0 {
			utils.UPrint(fmt.Sprintf("Not enough SP to use %v...\n", skill.Name), 20)
			time.Sleep(time.Second)
			fmt.Println()
			return false
		}
		if environ == "battle" {
			DelayedAction = map[string]interface{}{"type": "skill", "skill": skill}
			TurnEnded = true
			return true
		}

		crit, damage := utils.IsCritical(20), skill.Damage
		if crit == 1 {
			damage *= 2
		}
		// EFFECT ON TARGET
		if (target.Stats.Curr_hp - damage) < 0 {
			target.Stats.Curr_hp = 0
		} else {
			target.Stats.Curr_hp -= damage
		}
		// DEPLETES USER SP
		player.Stats.Curr_sp -= skill.CostInt

		if environ == "delayed" {
			if utils.IsCritical(20) == 1 {
				utils.UPrint(ansi.Color("Critical hit !!\n", "red+b"), 20)
				time.Sleep(250 * time.Millisecond)
			}
			utils.UPrint(fmt.Sprintf("%v dmg inflicted to %v !\n\n", damage, target.Name), 20)
			return true
		}
	case "heal":
		fmt.Println("To do")
	}
	return true
}
