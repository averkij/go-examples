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
	Model string
}

type Instrument struct {
	Sound string
}

//MakeSound is more like extension method in .NET
func (inst *Instrument) MakeSound() {
	fmt.Println(inst.Sound)
}
