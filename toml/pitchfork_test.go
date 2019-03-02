package main

import (
	"testing"
	"time"
)

func TestNewConfig(t *testing.T) {
	var (
		conf MyConfig
		err  error
	)

	conf, err = NewConfig("pitchfork.toml")
	if err != nil {
		t.Errorf("error should be nil.real = ", err)
	}
	if conf.Zookeeper.PitchforkRoot != "/pitchfork" {
		t.Errorf("error should be /pitchfork. real = ", conf.Zookeeper.PitchforkRoot)
	}

	if conf.Store.StoreCheckInterval.Seconds() != (time.Second * 5).Seconds() {
		t.Errorf("error should be 5s. real = ", conf.Store.StoreCheckInterval)
	}
}
