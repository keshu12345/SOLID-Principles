package main

import (
	"fmt"
	"sync"
)

type Singleton struct{}

var mutex = &sync.Mutex{}
var singleInstance *Singleton

func getInstance() *Singleton {
	if singleInstance == nil {
		mutex.Lock()
		defer mutex.Unlock()
		if singleInstance == nil {
			fmt.Println("Creating single instance now.")
			singleInstance = &Singleton{}
		} else {
			fmt.Println("Single instance already created.")
		}
	}
	return singleInstance
}

func main() {
	instannce1 := getInstance()
	instannce2 := getInstance()
	if instannce1 == instannce2 {
		fmt.Println("Both Instance is same ")
	} else {
		fmt.Println("Both instance is different")
	}
	instannce3 := getInstance()
	if instannce1 == instannce3 {
		fmt.Println("Both Instance is same ")
	} else {
		fmt.Println("Both instance is different")
	}
}
