package main

type Person struct {
	Name     string
	LastName string
	Age      int
	Height   float32
	Weight   float32
}

func (p *Person) SetName(name string) {
	p.Name = name
}

func UpdateLastName(p *Person, lastName string) {
	p.LastName = lastName
}
