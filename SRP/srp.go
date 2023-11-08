package main

import "fmt"

type Singleton struct{}

var singleton *Singleton

func getInstance() *Singleton {
	if singleton == nil {
		return &Singleton{}
	}
	return singleton
}

func main() {
	instance1 := getInstance()
	instance2 := getInstance()

	if instance1 == instance2 {
		fmt.Println("Both Instance is same")
	} else {
		fmt.Println("Both Instance different ")
	}

}
