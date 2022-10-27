package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

func main() {
	var i1 Item
	fmt.Println(i1)
	fmt.Printf("%#v\n", i1)
	i2 := Item{0, 10}
	i2.Move(100, 200)
	fmt.Printf("i2 move %#v\n", i2)

	p1 := Player{"Jack", []Key{}, Item{500, 300}}

	fmt.Printf("p1 :%#v\n", p1)
	fmt.Printf("p1 :%#v\n", p1.X)
	fmt.Printf("p1 :%#v\n", p1.Item.X)

	p1.Move(10, 20)

	fmt.Printf("p1 :%#v\n", p1)
	// 一个type实现了哪个接口，他就是哪种类型
	ms := []mover{
		&p1,
		&i1,
	}
	moveAll(ms, 0, 0)

	fmt.Printf("i1: %#v\n", i1)
	fmt.Printf("p1: %#v\n", p1)
	var k Key = 10
	fmt.Println(k)
	p1.FoundKey(Jade)
	fmt.Println(p1)
	fmt.Println(p1.Keys)
}

func (p *Player) FoundKey(k Key) error {
	if k < Jade || k >= invalidKey {
		return fmt.Errorf("invalid key %#v", k)
	}
	// 如果Keys没有这个key，我们把这个key添加进去
	// if !containsKey(p.Keys,k){
	// 	p.Keys = append(p.Keys, k)
	// }
	if !slices.Contains(p.Keys, k) {
		p.Keys = append(p.Keys, k)
	}
	return nil
}

func containsKey(keys []Key, k Key) bool {
	for _, v := range keys {
		if v == k {
			return true
		}
	}
	return false
}

// Implement fmt.Stringer interface
func (k Key) String() string {
	switch k {
	case Jade:
		return "Jade"
	case Copper:
		return "Copper"
	case Crystal:
		return "Crystal"
	}
	// return fmt.Sprintf("<Key %v>", k)
	// 递归调用
	// 他会首先调用fmt将int 转换
	return fmt.Sprintf("<Key %d>", k)
}

// go's version enum
const (
	Jade Key = iota + 1
	Copper
	Crystal
	invalidKey // internal
)

type Key byte

func moveAll(ms []mover, x, y int) {
	for _, m := range ms {
		m.Move(x, y)
	}
}

// 在go中。含有单个方法的interface我们通常称为xxxer
type mover interface {
	Move(x, y int)
}

// Add a Keys a slice of Key
type Player struct {
	Name string
	Keys []Key
	Item
}

// if i want mutate Item , use pointer receiver
func (i *Item) Move(x, y int) {
	i.X = x
	i.Y = y
}

// 一般栈中的变量都是在方法用完，自动回收的
// 这里我们返回了一个以后还要用的变量
// Go compiler 会对这个做"逃逸分析"，将我们的变量分配到堆中
func NewItem(x, y int) (*Item, error) {
	if x < 0 || x > maxX || y < 0 || y > maxY {
		return nil, fmt.Errorf("%d,%d out of bound %d,%d", x, y, maxX, maxY)
	}
	i := Item{x, y}
	// i := ItemErrorf()
	// 	X: x,
	// 	Y: y,
	// }
	return &i, nil
}

// zero vs missing value
// 是我们人为设置的零值还是他没有值展示出来的默认值

const (
	maxX = 1000
	maxY = 600
)

type Item struct {
	X int
	Y int
}
