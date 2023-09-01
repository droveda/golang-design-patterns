package observer2

type Publisher struct {
	Actions []ActionListener
}

func (p *Publisher) Subscribe(a ActionListener) {
	p.Actions = append(p.Actions, a)
}

func (p *Publisher) UnSubscribe(a ActionListener) {
	//TODO
}

func (p *Publisher) SomethingHappened() {
	p.notify()
}

func (p *Publisher) notify() {
	for _, action := range p.Actions {
		action.ActionExecuted("AA BB CC")
	}
}
