package anothercommand

type Button struct {
	Command
}

func (b *Button) Press() {
	b.Command.execute()
}
