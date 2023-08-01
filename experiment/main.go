package main

import "fmt"

type Uncle interface {
	Go(n string) (string, error)
	Run(speed int) (int, error)
}

type Me interface {
	Do(w string) (string, error)
}

type Aunty interface {
	Hell(u Uncle) (int, error)
}

type familyTree struct{}

func main() {

	ft := familyTree{}
	say(ft)
}

func say(u Uncle) {
	u.Go("ddd")
	u.Run(12)
}

func (familyTree) Go(n string) (string, error) {
	fmt.Println("Ding dong ....", n)
	return "go", nil
}

func (familyTree) Run(s int) (int, error) {
	fmt.Println("run rast mate", s)
	return 0, nil
}
