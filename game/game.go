package main

import (
	"fmt"
	"math"
	"slices"
	"sort"
)

func main() {
	i1 := Item{X: 1, Y: 1}
	fmt.Println(i1)
	fmt.Printf("%#v\n", i1)

	i1.Move(100, 200)
	fmt.Printf("%#v\n", i1)

	i2 := Item{1, 2}
	i3 := Item{Y: 10}

	ms := []mover{
		&i1,
		&i2,
		&i3,
	}
	moveAll(ms, 0, 0)
	for _, m := range ms {
		fmt.Println(m)
	}

	k := Jade
	fmt.Println("k: ", k)

    p1 := Player{
        Name: "Baga",
        Item: Item{500, 400},
    }
	p2 := Player{
		Name: "Joe",
		Item: Item{200, 300},
	}

	players := []Player{p1, p2}
	sortByDistance(players, 100, 200)
	fmt.Println(players)
/*     p1.FoundKey(Jade)
    fmt.Println(p1.Keys)
    p1.FoundKey(Jade)
    fmt.Println(p1.Keys) */

}

type Item struct {
	X int
	Y int
}

const (
	maxX = 1000
	maxY = 600
)

type Key byte

const (
	Jade Key = iota + 1
	Copper
	Crystal
    invalidKey 
)

func (k Key) String() string {
	switch k {
	case Jade:
		return "jade"
	case Copper:
		return "copper"
	case Crystal:
		return "crystal"
	}
	return fmt.Sprintf("<Key %d>", k)
}

func moveAll(ms []mover, x, y int) {
	for _, m := range ms {
		m.Move(x, y)
	}
}

type mover interface {
	Move(x, y int)
}

type Player struct {
    Name string
    Item 
    Keys []Key
}

func (p *Player) distanceFromPoint(x, y int) float64 {
	return math.Sqrt(float64((p.X - x)*(p.X -x) + (p.Y - y)* (p.Y - y))) 
}

func sortByDistance(players []Player, x, y int) {
	sort.Slice(players, func(i, j int) bool {
		return players[i].distanceFromPoint(x, y) > players[j].distanceFromPoint(x, y)
	})
}

func (p *Player) FoundKey(k Key) error {
    if k < Jade || k >= invalidKey {
        return fmt.Errorf("invalid key: %#v", k)
    }
    if !slices.Contains(p.Keys, k) {
        p.Keys = append(p.Keys, k)
    }
    
    return nil
}

func (i *Item) Move(x, y int) {
	i.X = x
	i.Y = y
}

func NewItem(x, y int) (*Item, error) {
	if x < 0 || x > maxX || y < 0 || y > maxY {
		return nil, fmt.Errorf("%d/%d out of bounds %d/%d", x, y, maxX, maxY)
	}
	i := Item{
		X: x,
		Y: y,
	}
	return &i, nil
}
