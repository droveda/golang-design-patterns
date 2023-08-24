package factory

import "fmt"

type Musket struct {
	Gun
}

func NewMusket() IGun {
	return &Musket{
		Gun{
			name:  "Musket gun",
			power: 1,
		},
	}
}

func (g *Musket) MakeSound() {
	fmt.Println("My sound is ppppp")
}
