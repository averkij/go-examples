package main

import "fmt"

func main() {
	b := Instrument{Sound: "Brrring!"}
	a := Guitar{Model: "Ibanez", Instrument: b}

	//Call Instrument method
	a.MakeSound()

	fmt.Println(a.Instrument.Sound)
}

type Guitar struct {
	Instrument
	Tool
	Model string
}

type Instrument struct {
	Sound string
}

type Tool struct {
	Sound string
}

//MakeSound is more like extension method in .NET
func (inst *Instrument) MakeSound() {
	fmt.Println(inst.Sound)
}

//We can't have two methods with the same name for Guitar
// func (tool *Tool) MakeSound() {
// 	fmt.Println(tool.Sound)
// }
