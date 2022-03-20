package main

import "C"
import (
	"fmt"
)

func hello() {

	fmt.Println("Hello world goroutine")
}
func main() {
	go hello()
	fmt.Println("main function")
	var fragment1 Fragment =new(GetPodAction)


	var fragment3 Fragment = &GetPodAction{}

	var fragment4 Fragment = GetPodAction{}
	fragment1.Exec(1)
	fragment3.Exec(2)
	fragment4.Exec(3)
}

type TransInfo int
type Fragment interface {
	Exec(transInfo TransInfo) error
}
type GetPodAction struct {}
func (g GetPodAction) Exec(transInfo TransInfo) error {

	return nil
}