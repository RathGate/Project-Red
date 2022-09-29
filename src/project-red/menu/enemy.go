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
	enemy.Exp = 15
	enemy.Money = 40
}

func TrainingFight(player *Character, enemy *Enemy) {
	var turn int = 1
	for ; turn < 5 && !player.IsDead() && !enemy.IsDead(); turn++ {
		player.RegisterPlayerAction(turn, enemy)
		utils.ConsoleClear()
		BattleIntroduction(*player, *enemy)
		for _, char := range strings.Repeat("-", 50) {
			fmt.Print(string(char))
			time.Sleep(15 * time.Millisecond)
		}
		fmt.Println("\n" + utils.Format("TURN %v", "center", 50, []string{strconv.Itoa(turn)}))
		for _, char := range strings.Repeat("-", 50) {
			fmt.Print(string(char))
			time.Sleep(15 * time.Millisecond)
		}
		fmt.Print("\n\n")
		PrintBattleInfo(*player, *enemy)

		if enemy.Stats.Initiative <= player.Stats.Initiative || DelayedAction["type"] == "run" {
			if outcome := PlayerTurn(turn, player, enemy); outcome == -1 {
				return
			} else if outcome == 1 {
				time.Sleep(1500 * time.Millisecond)
				_ = GetInputInt(0, []int{}, "")
				continue
			}
			PrintBattleInfo(*player, *enemy)
			time.Sleep(time.Second)
			if outcome := enemy.EnemyTurn(turn, player); outcome == 1 {
				time.Sleep(1500 * time.Millisecond)
				_ = GetInputInt(0, []int{}, "")
				continue
			}
		} else {
			if outcome := enemy.EnemyTurn(turn, player); outcome == 1 {
				time.Sleep(1500 * time.Millisecond)
				_ = GetInputInt(0, []int{}, "")
				continue
			}
			PrintBattleInfo(*player, *enemy)
			time.Sleep(time.Second)

			if outcome := PlayerTurn(turn, player, enemy); outcome == -1 {
				return
			} else if outcome == 1 {
				time.Sleep(1500 * time.Millisecond)
				_ = GetInputInt(0, []int{}, "")
				continue
			}
		}
		_ = GetInputInt(0, []int{}, "next")
	}
	GetBattleResults(turn, player, enemy)
}

func PrintBattleInfo(player Character, enemy Enemy) {
	e_name := utf8.RuneCountInString(fmt.Sprintf("%v", enemy.Name))
	e_hp := utf8.RuneCountInString(fmt.Sprintf("HP: %v/%v", enemy.Stats.Curr_hp, enemy.Stats.Max_hp))
	e_sp := utf8.RuneCountInString(fmt.Sprintf("SP: %v/%v", enemy.Stats.Curr_sp, enemy.Stats.Max_sp))
	fmt.Println(ansi.Color(fmt.Sprintf("%v", strings.ToUpper(enemy.Name)), "red+b") + ansi.Color(utils.Format("%v", "right", 50-e_name, []string{strings.ToUpper(player.Name)}), "blue+b"))
	fmt.Println(fmt.Sprintf("HP: %v/%v", enemy.Stats.Curr_hp, enemy.Stats.Max_hp) + utils.Format("HP: %v/%v", "right", 50-e_hp, []string{strconv.Itoa(player.Stats.Curr_hp), strconv.Itoa(player.Stats.Max_hp)}))
	fmt.Println(fmt.Sprintf("SP: %v/%v", enemy.Stats.Curr_sp, enemy.Stats.Max_sp) + utils.Format("SP: %v/%v", "right", 50-e_sp, []string{strconv.Itoa(player.Stats.Curr_sp), strconv.Itoa(player.Stats.Max_sp)}))

	fmt.Println()
}
func BattleIntroduction(player Character, enemy Enemy) {
	Box := box.New(box.Config{Px: 0, Py: 0, Type: "Double Single", Color: "Red", TitlePos: "Top"})
	Box.Print((P1.Name), utils.Format("A N  E N E M Y  A P P E A R S !!", "center", 48, []string{}))
	fmt.Println(utils.Format("An agressive %v ambushed you !", "center", 50, []string{enemy.Name}))

	if enemy.Loot != nil {
		fmt.Println(utils.Format("It seems to carry something interesting..."+"\n", "center", 50, []string{}))
	}
}
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

	utils.UPrint(fmt.Sprintf("%v inflicted %v dmg to %v !\n\n", enemy.Name, inflicted, player.Name), 20)

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

func (inventory *Inventory) GetBattleItems() []*Item {
	list := []*Item{}
	for _, item := range MapKeysToArr(inventory) {
		if item.BattleUse {
			list = append(list, item)
		}
	}
	return list
}

var Goblin = Enemy{}

func (enemy *Enemy) IsDead() bool {
	return enemy.Stats.Curr_hp <= 0
}
