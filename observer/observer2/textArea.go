package observer2

import "fmt"

type TextArea struct {
	Name string
}

func (t *TextArea) ActionExecuted(componentName string) {
	fmt.Println("Text are performed", t.Name, componentName)
}
