package factory

type DesertEagle struct {
	Gun
}

func NewDeserEagle() IGun {
	return &DesertEagle{
		Gun: Gun{
			name:  "Desert Eagle",
			power: 7,
		},
	}
}
