package api

import (
	"../types"
)

var SaleOrderModel string = "sale.order"

type SaleOrderService struct {
	client *Client
}

func (c *Client) NewSaleOrderService() {
	c.SaleOrder = &SaleOrderService{client: c}
	return
}

func (s *SaleOrderService) GetIdByName(name string) ([]int, error) {
	var args []interface{}
	var ids []int
	args = append(args, []interface{}{"name", "=", name})
	err := s.client.Search(SaleOrderModel, args, nil, &ids)
	return ids, err
}

func (s *SaleOrderService) GetByIds(ids []int) (*types.SaleOrders, error) {
	var args []interface{}
	args = append(args, ids)
	so := &types.SaleOrders{}
	err := s.client.Read(SaleOrderModel, args, nil, so)
	return so, err
}

func (s *SaleOrderService) GetByName(name string) (*types.SaleOrders, error) {
	var args []interface{}
	var args2 []interface{}
	args2 = append(args2, []string{"name", "=", name})
	args = append(args, args2)
	so := &types.SaleOrders{}
	err := s.client.SearchRead(SaleOrderModel, args, nil, so)
	return so, err
}

func (s *SaleOrderService) GetAll() (*types.SaleOrders, error) {
	var args []interface{}
	var args2 []interface{}
	args = append(args, args2)
	so := &types.SaleOrders{}
	err := s.client.SearchRead(SaleOrderModel, args, nil, so)
	return so, err
}

func (s *SaleOrderService) Create(required *types.RequiredSaleOrderFields, fields map[string]interface{}, fields2Many *types.Handle2ManysStruct) (int, error) {
	var args []interface{}
	var id int
	fields["partner_id"] = required.PartnerId
	if fields2Many != nil {
		fields = types.Handle2Manys(fields, fields2Many)
	}
	args = append(args, fields)
	err := s.client.Create(SaleOrderModel, args, &id)
	return id, err
}

func (s *SaleOrderService) Update(ids []int, fields map[string]interface{}, fields2Many *types.Handle2ManysStruct) error {
	var args []interface{}
	args = append(args, ids)
	if fields2Many != nil {
		fields = types.Handle2Manys(fields, fields2Many)
	}
	args = append(args, fields)
	err := s.client.Update(SaleOrderModel, args)
	return err
}

func (s *SaleOrderService) Delete(ids []int) error {
	var args []interface{}
	args = append(args, ids)
	return s.client.Delete("sale.order", args)
}
