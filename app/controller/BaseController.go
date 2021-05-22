package controller

import "fmt"

type BaseController struct {

}

func(base BaseController) test() {
	fmt.Println("我是BaseController..")
}