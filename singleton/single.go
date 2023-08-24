package singleton

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}
var singleInstance *Single

type Single struct {
}

func init() {
	fmt.Println("Init called")
}

func GetInstance() *Single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()

		if singleInstance == nil {
			fmt.Println("Creating single instance now")
			singleInstance = &Single{}
		} else {
			fmt.Println("Single instance already created 1")
		}

	} else {
		fmt.Println("Single instance already created 2")
	}

	return singleInstance
}
