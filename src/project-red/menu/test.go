package menu

func Sorted(inv *Inventory) []*Item {
	keys := make([]*Item, 0)
	for k := range inv.Items {
		keys = append(keys, k)
	}
	return keys
}

// fmt.Println(menu.Sorted(&Player.Inventory))
// for _, element := range menu.Sorted(&Player.Inventory) {
// 	fmt.Println(Player.Inventory.Items[element])
// }
