package menu

import "fmt"

var Shop *ShopKeeper

type ShopKeeper struct {
	Character
}

func (s *ShopKeeper) ShopMenu() {
	fmt.Println(s.Name)
	s.AccessInventory()
}
func (s *ShopKeeper) BuyMenu() {
	if len(s.Inventory.Items) == 0 {
		fmt.Println("I have nothing left for now!")
	} else {
		fmt.Println("Here's what I have for you!")
		index := 0
		var list []string
		for item, count := range s.Inventory.Items {
			fmt.Printf("n°%v || %v %v₽\n", index+1, item.Name, item.Price)
			_ = count
			index++
			list = append(list, item.Name)
		}
		fmt.Println(list)
		var answer int
		fmt.Printf("What will it be ? (type a number between 1 & %v, '0' to go back)\n", index)
		fmt.Scanf("%d\n", &answer)
		for answer < 0 || answer > index {
			fmt.Println("That's not a valid answer!")
			fmt.Scanf("%d\n", &answer)
		}
		if answer == 0 {
			return
		}
		BuyItem(&P1, s, list[answer-1], 1)
	}
}

func BuyItem(player *Character, shop *ShopKeeper, item string, count int) {
	player.Inventory.AddToInventory(&Item{Name: item}, count)
	shop.Inventory.RemoveFromInventory(item, count)
}
