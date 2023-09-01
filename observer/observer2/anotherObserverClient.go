package observer2

func Observer2Example() {
	p := Publisher{}

	b := Button{
		Name: "Buton",
	}

	t := TextArea{
		Name: "Text",
	}

	p.Subscribe(&b)
	p.Subscribe(&t)

	p.SomethingHappened()

}
