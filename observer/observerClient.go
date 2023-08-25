package observer

import "fmt"

func ObserverExample() {
	tvLg := NewItem(
		"LG Oled TV Evo 2023",
	)

	observer1 := Customer{
		Id: "01",
	}

	observer2 := Customer{
		Id: "02",
	}

	tvLg.Register(&observer1)
	tvLg.Register(&observer2)

	tvLg.UpdateAvailability()

	tvLg.DeRegister(&observer1)

	fmt.Println("---####---")
	tvLg.UpdateAvailability()
}
