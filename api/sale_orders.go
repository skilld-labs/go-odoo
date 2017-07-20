package api 

import (
	"github.com/antony360/go-odoo/odoo"
	"github.com/antony360/go-odoo/types"
)

type SaleOrderService struct {
	client *odoo.Client
}

func (c *odoo.Client) NewSaleOrderService() {
	c.SaleOrder = &SaleOrderService{client: c}
	return
}

func (s *SaleOrderService) GetByIds (ids []int) (*types.SaleOrders, error) {
	var args []interface{}
	args = append(args, ids)
	so := &types.SaleOrders{}
	err := s.client.Read("sale.order", args, nil, so)
	return so, err
}
