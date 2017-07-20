package odoo 

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	c, err := NewClient("http://62.210.250.198:8069", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = c.Login("the-s-team", "admin", "admin")
	if err != nil {
		fmt.Println(err.Error())
	}
	c.NewSaleOrderService()
	so, err := c.SaleOrder.GetById([]int{8, 7})
	spew.Dump(so)
	if err != nil {
		fmt.Println(err.Error())
	}
}
