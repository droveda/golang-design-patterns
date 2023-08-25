package observer

type Observer interface {
	Update(name string)
	GetId() string
}
