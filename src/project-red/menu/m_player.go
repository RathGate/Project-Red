package menu

// type Personnage struct {
// 	Name      string
// 	Class     string
// 	Level     int
// 	Max_hp    int
// 	Curr_hp   int
// 	Inventory []Item
// }

// type Item struct {
// 	Name  string
// 	Count int
// }

// func (p Personnage) DisplayInfo() {
// 	fmt.Println("Name:", p.Name)
// 	fmt.Println("Class:", p.Class)
// 	fmt.Println("Level:", p.Level)
// 	fmt.Println("Max HP:", p.Max_hp)
// 	fmt.Println("HP:", p.Curr_hp)
// 	fmt.Println("Inventory:")
// 	if p.Inventory == nil {
// 		fmt.Println("*** EMPTY ***")
// 	} else {
// 		for _, item := range p.Inventory {
// 			fmt.Printf("  %v %v(s)\n", item.Count, item.Name)
// 		}
// 	}
// 	fmt.Print("\n")
// }
// func (p Personnage) AccessInventory() {
// 	fmt.Printf("----- %v's INVENTORY: -----\n", p.Name)
// 	if len(p.Inventory) == 0 {
// 		fmt.Println("Nothing to see here...")
// 	} else {
// 		for _, item := range p.Inventory {
// 			fmt.Printf("### %v %v(s)\n", item.Count, item.Name)
// 		}
// 	}
// 	fmt.Print("\n")
// }

// func MainMenu() {
// 	fmt.Println(">>> `D` to Display Character Information")
// 	fmt.Println(">>> `I` to Display Current Inventory")
// 	fmt.Println(">>> `Q` to Quit Game")

// 	var answer string
// 	fmt.Scanf("%s", &answer)

// 	switch answer {
// 	case "D":
// 		fmt.Println("Display")
// 	case "I":
// 		p1.AccessInventory()
// 	case "Q":
// 		fmt.Println("Quit")
// 	}
// }

// func Init(name, class string, level, max_hp, curr_hp int, inventory []Item) Personnage {
// 	return Personnage{
// 		Name:      name,
// 		Class:     class,
// 		Level:     level,
// 		Max_hp:    max_hp,
// 		Curr_hp:   curr_hp,
// 		Inventory: inventory,
// 	}
// }
