package factory

import "fmt"

type Ak47 struct {
	Gun
}

func NewAk47() IGun {
	return &Ak47{
		Gun: Gun{
			name:  "Ak47 gun",
			power: 4,
		},
	}
}

func (g *Ak47) MakeSound() {
	fmt.Println("My sound is kkkkkk")
}
