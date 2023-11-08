package main

import (
	"fmt"
	"sync"
)

type singleton struct{}

var singletonObj *singleton
var once sync.Once

func getInstance() *singleton {

	if singletonObj == nil {
		once.Do(
			func() {
				fmt.Println("Creating single instance now.")
				singletonObj = &singleton{}
			})
	} else {
		fmt.Println("Single instance already created.")
	}

	return singletonObj
}

func main() {
	instannce1 := getInstance()
	instannce2 := getInstance()
	if instannce1 == instannce2 {
		fmt.Println("Both Instance is same ")
	} else {
		fmt.Println("Both instance is different")
	}

}
