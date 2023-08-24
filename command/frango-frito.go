package command

import "fmt"

type FrangoFrito struct {
	Frango Frango
}

func (f *FrangoFrito) execute() {
	fmt.Println("Preparando frango frito", f.Frango.Qtde)
}
