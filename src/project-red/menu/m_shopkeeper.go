package menu

import (
	"fmt"
	"strings"
	"time"
)

var Shop *ShopKeeper

type ShopKeeper struct {
	Character
}

func (s *ShopKeeper) ShopMenu() {
	fmt.Println(s.Name)
	s.AccessInventory()
}

func (s *ShopKeeper) BuyItem(player *Character, item *Item, count int) {
	if (item.Price * count) > player.Inventory.Money {
		fmt.Println(`"Hey, don't buy if you can't pay !`)
	} else if invFull, invCount := player.Inventory.IsFull(); invFull || invCount+count > 10 {
		fmt.Println(`"It seems your bag is too heavy to buy this...`)
	} else {
		player.Inventory.AddToInventory(item, count)
		s.Inventory.RemoveFromInventory(item, count)
		fmt.Printf("------ BOUGHT %v %v FROM %v ------\n\n", strings.ToUpper(item.Name), count, strings.ToUpper(s.Name))
		fmt.Println(`It's always a pleasure doing business with you!"`)
	}
	time.Sleep(1000 * time.Millisecond)
	// player.Inventory.AddToInventory(item, count)
	// s.Inventory.RemoveFromInventory(item, count)
}
