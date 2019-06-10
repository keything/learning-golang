package main

import (
	"github.com/BurntSushi/toml"
	"io/ioutil"

	"os"
	"time"
)

type MyConfig struct {
	Zookeeper Zookeeper
	Store     Store
}
type Zookeeper struct {
	Addrs         []string
	Timeout       myduration
	PitchforkRoot string
	StoreRoot     string
	VolumeRoot    string
}

type Store struct {
	StoreCheckInterval  myduration
	NeedleCheckInterval myduration
	RackCheckInterval   myduration
}
type myduration struct {
	time.Duration
}

//https://golang.org/pkg/encoding/#TextUnmarshaler
func (d *myduration) UnmarshalText(text []byte) error {
	var err error
	d.Duration, err = time.ParseDuration(string(text))
	return err
}

func NewConfig(path string) (c *MyConfig, err error) {
	var (
		openFile *os.File
		blob     []byte
	)
	if openFile, err = os.Open(path); err != nil {
		return nil, err
	}
	if blob, err = ioutil.ReadAll(openFile); err != nil {
		return nil, err
	}
	c = new(MyConfig)
	if _, err := toml.Decode(string(blob), c); err != nil {
		return nil, err
	}
	return c, nil

}
