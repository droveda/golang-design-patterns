package bridge

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
}
