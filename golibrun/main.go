package main

import (
	"fmt"

	"github.com/yuki-seike/golib"
)

func main() {
	fmt.Printf("ok!\n")

	p := golib.Person{"hawk", 10}
	p.Say()
}
