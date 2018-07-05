package main

import "fmt"

type Mutable interface {
	mutate(newValue string)
}

type Data struct {
	name string
}

func (d *Data) mutate(newValue string) {
	d.name = newValue
}

func mutator(mute Mutable) {
	mute.mutate("mutate")
}

func main() {
	d := Data{name: "fresh"}
	fmt.Println(d.name)
	mutator(&d)
	fmt.Println(d.name)
}
