package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"time"
)

type Config struct {
	Age        int
	Cats       []string
	Pi         float64
	Perfection []int
	DOB        time.Time
	D          duration `toml:"Duration"`
}

//https://golang.org/pkg/encoding/#TextUnmarshaler
type duration struct {
	time.Duration
}

func (d *duration) UnmarshalText(text []byte) error {
	var err error
	d.Duration, err = time.ParseDuration(string(text))
	return err
}

func main() {
	var conf Config
	tomlData := "Age = 25\nCats = [ \"Cauchy\", \"Plato\" ]\nPi = 3.14\nPerfection = [ 6, 28, 496, 8128 ]\nDOB = 1987-07-05T05:45:00Z\nDuration=\"8m30s\""
	if _, err := toml.Decode(tomlData, &conf); err != nil {
		fmt.Println("err=", err)
	}
	fmt.Println("succeed")
	fmt.Println("age=", conf.Age, "||cats=", conf.Cats, "||pi=", conf.Pi, "||DOB=", conf.DOB, "||duration=", conf.D)
	fmt.Println("total seconds=", conf.D.Seconds())
}
