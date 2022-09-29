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
}

func (enemy *Enemy) InitEnemy(name string, curr_hp, max_hp, akt int) {
	enemy.Name = name
	enemy.Stats.Curr_hp = curr_hp
	enemy.Stats.Max_hp = max_hp
	enemy.Stats.Atk = akt
	enemy.Loot = &PoisonPotion
	enemy.Stats.Initiative = 10
}

func TrainingFight(player *Character, enemy *Enemy) {

	for turn := 1; turn < 5; turn++ {
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
			switch PlayerTurn(turn, player, enemy) {
			case -1:
				_ = GetInputInt(0, []int{}, "")
				return
			case 1:
				break
			}
			fmt.Println(enemy.Stats.Initiative)
		} else {
			fmt.Println(enemy.Stats.Initiative)
			PlayerTurn(turn, player, enemy)
		}
		_ = GetInputInt(0, []int{}, "next")
	}
	fmt.Println("hello?")
	_ = GetInputInt(0, []int{}, "")
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
func (enemy *Enemy) EnemyPattern(player *Character) int {
	// for i := 1; i < 3; i++ {
	// 	damage := enemy.Akt
	// 	if i%3 == 0 {
	// 		damage = enemy.Akt * 2
	// 		fmt.Print("Critical hit ! ")
	// 	}
	// 	player.Curr_hp -= damage
	// 	fmt.Printf("%v deals %vHP damage to %v !\n", enemy.Name, damage, player.Name)
	// 	fmt.Println(player.Curr_hp)
	// }
	return 69
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
