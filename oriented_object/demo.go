package main

import (
	"fmt"
	"time"
)

type Pet struct {
	name string
}

type Dog struct {
	Breed string
	Pet
	time.Duration
}

func (p *Pet) Speak() string {
	return fmt.Sprintf("my name is %v", p.name)
}

func (p *Pet) Name() string {
	return p.name
}

func (d *Dog) Speak() string {
	return fmt.Sprintf("%v and I am a %v", d.Pet.Speak(), d.Breed)
}
func (d *Dog) TimeD() {
	fmt.Println("time d ", d.Duration)
}

func main() {
	d := Dog{Pet: Pet{name: "spot"}, Breed: "pointer"}
	d.Duration = time.Second * 3
	fmt.Println(d.Name())
	fmt.Println(d.Speak())
	d.TimeD()
}
