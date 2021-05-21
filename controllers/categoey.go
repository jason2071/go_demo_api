package controllers

import "fmt"

type CategoryStruct struct {
	width  int
	height int
}

func (d CategoryStruct) DemoFunc() {

	r := CategoryStruct{4, 5}

	fmt.Println(r)
}
