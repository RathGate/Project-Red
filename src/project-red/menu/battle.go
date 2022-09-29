package menu

import (
	"fmt"
	"projectRed/utils"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/Delta456/box-cli-maker/v2"
	"github.com/mgutz/ansi"
)

func BattleIntroduction(player Character, enemy Enemy) {
	utils.ConsoleClear()
	Box := box.New(box.Config{Px: 0, Py: 0, Type: "Double Single", Color: "Red", TitlePos: "Top"})
	Box.Print((P1.Name), utils.Format("A N  E N E M Y  A P P E A R S !!", "center", 48, []string{}))
	fmt.Println(utils.Format("An agressive %v ambushed you !", "center", 50, []string{enemy.Name}))

	if enemy.Loot != nil {
		fmt.Println(utils.Format("It seems to carry something interesting..."+"\n", "center", 50, []string{}))
	}
}

func PrintBattleInfo(player Character, enemy Enemy) {
	e_name := utf8.RuneCountInString(fmt.Sprintf("%v", enemy.Name))
	e_hp := utf8.RuneCountInString(fmt.Sprintf("HP: %v/%v", enemy.Stats.Curr_hp, enemy.Stats.Max_hp))
	fmt.Println(ansi.Color(fmt.Sprintf("%v", strings.ToUpper(enemy.Name)), "red+b") + ansi.Color(utils.Format("%v", "right", 50-e_name, []string{strings.ToUpper(player.Name)}), "blue+b"))
	fmt.Println(fmt.Sprintf("HP: %v/%v", enemy.Stats.Curr_hp, enemy.Stats.Max_hp) + utils.Format("HP: %v/%v", "right", 50-e_hp, []string{strconv.Itoa(player.Stats.Curr_hp), strconv.Itoa(player.Stats.Max_hp)}))
	fmt.Println(utils.Format("SP: %v/%v", "right", 50, []string{strconv.Itoa(player.Stats.Curr_sp), strconv.Itoa(player.Stats.Max_sp)}))

	fmt.Println()
}

func (inventory *Inventory) GetBattleItems() []*Item {
	list := []*Item{}
	for _, item := range MapKeysToArr(inventory) {
		if item.BattleUse {
			list = append(list, item)
		}
	}
	return list
}

func (player *Character) LevelUp() {
	player.Stats.Exp = player.Stats.Max_exp - player.Stats.Exp
	player.Stats.Max_exp *= int(float64(player.Stats.Max_exp) * 1.3)
	player.Stats.Level++

	utils.UPrint(ansi.Color(utils.Format("●●●● LEVEL UP !! ●●●●\n", "center", 50, []string{}), "white+b"), 20)
	fmt.Println()
	time.Sleep(250 * time.Millisecond)
	utils.UPrint(ansi.Color(fmt.Sprintf("%v is now level %v !\n", player.Name, player.Stats.Level), "cyan"), 20)
	time.Sleep(250 * time.Millisecond)
	_ = GetInputInt(0, []int{}, "")
}
