package strategy

import "fmt"

func DoStrategyExample() {

	lfu := Lfu{}

	cache := InitCache(&lfu)
	cache.Add("a", "1")
	cache.Add("b", "2")
	cache.Add("c", "3")

	fmt.Println(cache.Get("c"))

	cache.SetEvictionAlgo(&Lru{})
	cache.Add("d", "4")

	cache.SetEvictionAlgo(&Fifo{})
	cache.Add("e", "5")

}
