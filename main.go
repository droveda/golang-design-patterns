package main

import (
	"fmt"

	"github.com/droveda/design-patterns/builder"
	"github.com/droveda/design-patterns/command"
	"github.com/droveda/design-patterns/factory"
	_ "github.com/droveda/design-patterns/observer"
	"github.com/droveda/design-patterns/observer/observer2"
	"github.com/droveda/design-patterns/singleton"
	_ "github.com/droveda/design-patterns/strategy"
)

func main() {
	fmt.Println("Hello!")

	// factoryExample()
	// builderExample()
	// singletonExample()
	// someTests()

	// commandExample()
	// anothercommand.CommandClient()
	// bridge.BridgeExample()
	// strategy.DoStrategyExample()
	// observer.ObserverExample()
	observer2.Observer2Example()
}

func factoryExample() {
	musket, _ := factory.GetGun(factory.MusketType)
	println(musket.GetName(), " - power:", musket.GetPower())
	musket.MakeSound()

	ak47, _ := factory.GetGun(factory.Ak47Type)
	println(ak47.GetName(), " - power:", ak47.GetPower())
	ak47.MakeSound()

	desert, _ := factory.GetGun(factory.DesertEagleType)
	println(desert.GetName(), " - power:", desert.GetPower())
	desert.MakeSound()

	fmt.Println(factory.Ak47Type)
}

func builderExample() {
	myBuilder := builder.GetBuilder("normal")
	myBuilder.SetWindowType()
	myBuilder.SetDoorType()
	myBuilder.SetNumFloor()
	house := myBuilder.GetHouse()
	fmt.Println(house)

	director := builder.NewDirector(
		builder.GetBuilder("igloo"),
	)

	iglooHouse := director.BuildHouse()
	fmt.Println(iglooHouse)
}

func singletonExample() {
	for i := 0; i < 15; i++ {
		go singleton.GetInstance()
	}

	fmt.Scanln()
}

func commandExample() {
	frango := command.Frango{
		Qtde: 2,
	}
	frangoFrito := command.FrangoFrito{
		Frango: frango,
	}
	invoker := command.Invoker{}
	invoker.SetCommand(&frangoFrito)
	invoker.ExecuteCommand()

	pratoFeito := command.PratoFeito{}
	invoker.SetCommand(&pratoFeito)
	invoker.ExecuteCommand()
}

func someTests() {
	var p Person = Person{
		LastName: "Silva",
		Age:      20,
		Height:   1.75,
		Weight:   85.2,
	}
	p.SetName("Fulano")
	UpdateLastName(&p, "Silva Jr.")
	fmt.Println(p)
}
