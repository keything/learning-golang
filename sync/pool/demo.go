package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var (
		p *sync.Pool
		a *int
	)
	p = &sync.Pool{
		New: func() interface{} {
			fmt.Println("new int")
			return new(int)
		},
	}

	a = p.Get().(*int)
	fmt.Println("first get=", *a)
	*a = 1
	p.Put(a)

	a = p.Get().(*int)
	fmt.Println("second get=", *a)
	*a = 3
	p.Put(a)
	a = p.Get().(*int)
	fmt.Println("third get=", *a)
	a = p.Get().(*int)
	fmt.Println("(need new) four get=", *a)
	a = p.Get().(*int)
	fmt.Println("(need new) five get=", *a)
	a = p.Get().(*int)
	fmt.Println("(need new) six get=", *a)

	runtime.GC() //手动调用GC
	a = p.Get().(*int)
	fmt.Println("after manual gc get=", *a)
}
