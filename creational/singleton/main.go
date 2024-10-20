package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type Singleton struct {
	id int
}

var (
	instance *Singleton
	once     sync.Once
)

func GetInstance() *Singleton {
	randomID := rand.Intn(100)
	once.Do(func() {
		instance = &Singleton{
			id: randomID,
		}
	})
	return instance
}

func main() {
	s1 := GetInstance()
	s2 := GetInstance()

	if s1 != s2 {
		panic("s1 != s2")
	}

	fmt.Printf("s1 %v\t s2 %v\n", s1, s2)

	fmt.Println("Both instances are the same")
}
