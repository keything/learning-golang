package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var (
		ip      *int
		age     int
		flagSet *flag.FlagSet
	)
	fmt.Println("args", flag.Args())
	flagSet = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	flagSet.Usage = func() {
		fmt.Printf("%s usage defined by user\n", os.Args[0])
		return
	}
	ip = flagSet.Int("ip", 1234, "help message for ip")

	flagSet.IntVar(&age, "age", 234, "help message for age")

	flagSet.Parse(os.Args[1:])
	fmt.Println("nargs=", flagSet.NArg()) //NArg:还没有被处理的字段的个数

	fmt.Println("nflag=", flagSet.NFlag()) //NFlag:已经被处理的字段个数

	fmt.Println("ip=", *ip, "||age=", age)
	fmt.Println("args=", flagSet.Args())
}
