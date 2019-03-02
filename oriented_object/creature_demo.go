package main

import (
	"fmt"
)

type CreatureDemo struct {
	Name string
	Real bool
}

func (c CreatureDemo) Dump() {
	fmt.Println("name=", c.Name, "||real=", c.Real)
}

type FlyingCreatureDemo struct {
	CreatureDemo
	Wignum int
}

func (f FlyingCreatureDemo) Dump() {
	fmt.Println("name=", f.Name, "||real=", f.Real, "||wignum=", f.Wignum)
}

func main() {
	var fly FlyingCreatureDemo
	fly = FlyingCreatureDemo{
		CreatureDemo: CreatureDemo{
			Name: "flying",
			Real: true,
		},
		Wignum: 20,
	}
	fly.Dump()
}
