package command

import "fmt"

type PratoFeito struct {
}

func (p *PratoFeito) execute() {
	fmt.Println("Preparando um prato feito")
}
