package menu

import (
	"fmt"
)

type Inventory struct {
	Items    map[*Item]int
	Money    int
	Capacity int
}

// ADDS [count] [item] TO THE INVENTORY
func (inventory *Inventory) AddToInventory(item *Item, count int) {

	// If inventory is full or too many items to put in the bag:
	invStatus, invCount := inventory.IsFull()
	if invStatus || invCount+count > 10 {
		fmt.Println("My bag can't bear all this for now...")
		return
	}

	// If item already in bag, adds [count] to its own count.
	for i := range inventory.Items {
		if item.Name == i.Name {
			inventory.Items[i] += count
			return
		}
	}
	// Else: adds item as a new item in bag.
	inventory.Items[item] = count
}

func (inv *Inventory) RemoveFromInventory(item *Item, count int) bool {
	item = RetrieveItemByName(item.Name, *inv)

	// Deletes [count] of [item] from the inventory.
	inv.Items[item] -= count

	// If 0 or less items in the inventory, deletes the full item:
	if inv.Items[item] <= 0 {
		delete(inv.Items, item)
		return true
	}
	return false
}

// CHECKS IF INVENTORY IS FULL:
// Returns 2 values: bool (true if full, false otherwise)
// and int (the number of items in the inventory)
func (inventory Inventory) IsFull() (bool, int) {
	count := 0
	for _, number := range inventory.Items {
		count += number
	}
	return (count >= 10), count
}

func (inventory *Inventory) UpgradeInventorySlot() {
	inventory.Capacity += 10
}
