package observer2

import "fmt"

type Button struct {
	Name string
}

func (b *Button) ActionExecuted(componentName string) {
	fmt.Println("Button action performed", b.Name, componentName)
}
