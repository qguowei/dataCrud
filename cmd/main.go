package main

import (
	"dataCrud"
	"log"
)

func main() {

	dataCrud.Cfg.DBName = "newjxc_orders"
	dataCrud.Cfg.BasePath = "../newjxc"
	dataCrud.Cfg.SeverPath = "orderSev"
	dataCrud.Cfg.ModelPath = "orderModels"

	var err error
	err = dataCrud.CreateModel("jxc_order", "jxc_order_pro")
	log.Println()
	if err != nil {
		panic(err)
	}
	err = dataCrud.CreateSever("jxc_order", "jxc_order_pro")
	if err != nil {
		panic(err)
	}
	err = dataCrud.CreateApi("jxc_order")
	if err != nil {
		panic(err)
	}
}
