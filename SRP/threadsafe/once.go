package main

import (
	"fmt"
	"sync"
)

type singleton struct{}

var singletonObj *singleton
var once sync.Once

func getInstance() *singleton {
	once.Do(func() {
		singletonObj = &singleton{}
	})
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
