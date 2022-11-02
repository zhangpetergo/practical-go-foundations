package main

import (
	"fmt"
	"sync"
	"time"
)

// 使用sync.Once完成幂等性
func main() {
	p := Payment{
		From:   "Jack",
		To:     "Marry",
		Amount: 123.45,
	}
	p.Process()
	p.Process()
}

func (p *Payment) Process() {
	// 有参数的使用匿名方法
	t := time.Now()
	//p.once.Do(p.process)
	p.once.Do(func() {
		p.process(t)
	})
}
func (p *Payment) process(t time.Time) {
	format := t.Format(time.RFC3339)
	fmt.Printf("%v %s->%.2f->%s\n", format, p.From, p.Amount, p.To)
}

type Payment struct {
	From   string
	To     string
	Amount float64
	once   sync.Once
}
