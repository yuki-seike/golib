package lib

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p *Person) Say() {
	fmt.Printf("My name is %s, I'm %d\n", p.Name, p.Age)
}
