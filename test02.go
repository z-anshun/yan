package main

import (
	"fmt"
	"math/rand"
)

/*
角色有血量、蓝量、攻击力、技能四个属性

A角色：血量 100 蓝量 0 攻击力10（每次攻击减少对方10血量，初始蓝量为0）

​              技能 无消耗、每次攻击有50%概率触发连击（又一次攻击），每回合理论上可以无限连击

B角色：血量 300 蓝量 50 攻击力 20 （每次攻击回复 10 蓝量 ，初始蓝量为0）

​              技能  蓝量达到50自动触发、降低对方10%当前攻击力

回合制游戏，不考虑攻击速度

比试十次，分别由每回合都是A先手5次，和每回合都是B先手5次组成

请你计算出他们各自的胜率

### 题目要求

任何方式*/
type person struct {
	blood int
	blue  int
	att   float32
}

var p2 = person{
	blood: 300,
	blue:  0,
	att:   20,
}
var p1 = person{
	blood: 100,
	blue:  0,
	att:   10,
}
var count1, count2 int

func main() {
	for i := 0; i < 10; i++ {
		for {
			if i%2 == 0 {
				k1 := p1.start1()
				k2 := p2.start2()
				if k1 == 1 || k2 == 1 {
					break
				}
			}
			if i%2 == 1 {
				k1 := p1.start2()
				k2 := p2.start1()
				if k1 == 1 || k2 == 1 {
					break
				}
			}
		}

	}
	f1 := float32(float32( count1) /float32 (count2 + count1) )
	f2 := 1 - f1
	fmt.Printf("p1的胜率：%f,p2的胜率：%f", f1, f2)
}
func (p *person) start1() int {

	fmt.Println("1)攻击")
	p.blue += 10
	if p.blue > 50 {
		p2.att *= 0.9
	}
	p2.blood -= 20

	if p2.blood <= 0 {
		count1++
		p2.blood = 300
		p2.blue = 0
		p2.att = 20
		return 1
	}
	return 0
}
func (p *person) start2() int {
	for i := 0; i < 1; i++ {
		fmt.Println("2）攻击")
		p1.blood -= 10
		k := rand.Int() % 2
		fmt.Println(k)
		if k == 1 {
			i -= 1
		} //等于1就连击，就是再来一次
	}
	if p1.blood <= 0 {
		count2++
		p1.blood = 100
		p1.blue = 0
		p1.att = 10
		return 1
	}
	return 0
}
