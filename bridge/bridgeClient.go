package bridge

import (
	"github.com/droveda/design-patterns/bridge/bidge2"
)

func BridgeExample() {
	epson := &Epson{}
	hp := &Hp{}

	mac := Mac{}
	mac.SetPrinter(epson)
	mac.Print()
	mac.SetPrinter(hp)
	mac.Print()

	windows := Windows{}
	windows.SetPrinter(hp)
	windows.Print()
	windows.SetPrinter(epson)
	windows.Print()

	mapa := bidge2.GoogleMap{}
	mapClient := bidge2.MapClient{
		Mapa: &mapa,
	}
	mapClient.ImprimirMapa()

	mapa2 := bidge2.MapLink{}
	mapClient.SetMapType(&mapa2)
	mapClient.ImprimirMapa()

}
