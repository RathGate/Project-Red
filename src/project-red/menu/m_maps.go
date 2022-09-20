package menu

import (
	"fmt"

	"github.com/oleiade/reflections"
)

var Directions = []string{"Up", "Right", "Down", "Left"}

type Mapblock struct {
	Id                            int
	Discovered                    bool
	Data                          string
	Up, Right, Down, Left, Origin *Mapblock
}

func (m Mapblock) GetAvailableDir() []string {
	var result []string
	for _, dir := range Directions {
		val, _ := reflections.GetField(m, dir)
		if val.(*Mapblock) != nil {
			result = append(result, dir)
		}
	}
	return result
}

func (m Mapblock) PrintMap() {
	fmt.Printf("--- MAP: id_%v ---\nDiscovered: %v\n", m.Id, m.Discovered)
	for _, dir := range Directions {
		val, _ := reflections.GetField(m, dir)
		if val.(*Mapblock) != nil {
			fmt.Printf("%v: id_%v\n", dir, val.(*Mapblock).Id)
		}
	}
	fmt.Println()
}

func DeclareMaps() []*Mapblock {
	var map_0 = &Mapblock{
		Id: 0,
	}
	var map_1 = &Mapblock{
		Id:     1,
		Down:   map_0,
		Origin: map_0,
	}
	map_0.Up = map_1
	var map_2 = &Mapblock{
		Id:     2,
		Right:  map_1,
		Origin: map_1,
	}
	map_1.Left = map_2

	var map_3 = &Mapblock{
		Id:     3,
		Down:   map_1,
		Origin: map_1,
	}
	var map_4 = &Mapblock{
		Id:     4,
		Left:   map_1,
		Origin: map_1,
	}
	map_1.Right = map_4
	var map_5 = &Mapblock{
		Id:     5,
		Left:   map_3,
		Origin: map_3,
	}
	map_3.Right = map_5
	var map_6 = &Mapblock{
		Id:     6,
		Down:   map_5,
		Origin: map_5,
	}
	map_5.Up = map_6
	var map_7 = &Mapblock{
		Id:     7,
		Left:   map_4,
		Origin: map_4,
	}
	map_4.Right = map_7
	var map_8 = &Mapblock{
		Id:     8,
		Down:   map_7,
		Origin: map_7,
	}
	map_7.Up = map_8
	var map_9 = &Mapblock{
		Id:     9,
		Left:   map_8,
		Origin: map_8,
	}
	map_8.Right = map_9

	var maps = []*Mapblock{map_0, map_1, map_2, map_3, map_4, map_5, map_6, map_7, map_8, map_9}
	return maps
}

func (m Mapblock) PrintDirections() {
	fmt.Println("AVAILABLE DIRECTIONS:")
	fmt.Print("--")
	for _, dir := range m.GetAvailableDir() {
		fmt.Printf(" %v --", dir)
	}
	fmt.Print("\n\n")
}
func ChooseDirection(originMap *Mapblock) string {
	directions := originMap.GetAvailableDir()
	var answer string
	var is_valid bool = false
	originMap.PrintDirections()
	for !is_valid {
		fmt.Printf("%s", "Let's try... ")
		fmt.Scanf("%s\n", &answer)
		for _, dir := range directions {
			if answer == dir {
				is_valid = true
				return dir
			}
		}
		fmt.Printf("%s", "That's just a wall !\n")
	}
	fmt.Print("\n")
	return answer
}

func ChangeMap(originMap *Mapblock) {

	dir := ChooseDirection(originMap)
	val, _ := reflections.GetField(originMap, dir)
	*originMap = *val.(*Mapblock)
	if !originMap.Discovered {
		originMap.Discovered = true
	}
	fmt.Println()
	originMap.PrintMap()
}
