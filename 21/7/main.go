// Description: 内部指针，&操作符不仅可以获得结构体的内存地址，还可以获得结构体内部字段的内存地址

package main

import "fmt"

type stats struct {
	level             int
	endurance, health int
}

func levelUp(s *stats) {
	s.level++
	s.endurance = 42 + (14 * s.level)
	s.health = 5 * s.endurance
}

type character struct {
	name string
	stats
}

func main() {
	player := character{name: "Matthias"}
	levelUp(&player.stats)

	fmt.Printf("%+v\n", player)
}
