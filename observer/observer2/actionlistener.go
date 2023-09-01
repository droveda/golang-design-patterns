package observer2

type ActionListener interface {
	ActionExecuted(componentName string)
}
