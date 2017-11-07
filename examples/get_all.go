package main

import (
	"fmt"

	"github.com/skilld-labs/go-odoo/api"
)

func getAll() {
	c, err := api.NewClient("http://localhost:8069", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = c.Login("dbName", "admin", "password")
	if err != nil {
		fmt.Prinln(err.Error())
	}
	//get the sale order service
	s := api.NewSaleOrderService(c)
	//call the function GetAll() linked to the sale order service
	so, err := s.GetAll()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(so)
}
