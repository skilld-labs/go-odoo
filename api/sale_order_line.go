package api

import (
	"../types"
)

var SaleOrderLineModel string = "sale.order.line"

type SaleOrderLineService struct {
	client *Client
}

func (c *Client) NewSaleOrderLineService() {
	c.SaleOrderLine = &SaleOrderLineService{client: c}
	return
}

func (s *SaleOrderLineService) GetIdByName(name string) ([]int, error) {
	var args []interface{}
	var ids []int
	args = append(args, []interface{}{"name", "=", name})
	err := s.client.Search(SaleOrderLineModel, args, nil, &ids)
	return ids, err
}

func (s *SaleOrderLineService) GetByIds(ids []int) (*types.SaleOrderLines, error) {
	var args []interface{}
	args = append(args, ids)
	so := &types.SaleOrderLines{}
	err := s.client.Read(SaleOrderLineModel, args, nil, so)
	return so, err
}

func (s *SaleOrderLineService) GetByName(name string) (*types.SaleOrderLines, error) {
	var args []interface{}
	var args2 []interface{}
	args2 = append(args2, []string{"name", "=", name})
	args = append(args, args2)
	so := &types.SaleOrderLines{}
	err := s.client.SearchRead(SaleOrderLineModel, args, nil, so)
	return so, err
}

func (s *SaleOrderLineService) GetAll() (*types.SaleOrderLines, error) {
	var args []interface{}
	var args2 []interface{}
	args = append(args, args2)
	so := &types.SaleOrderLines{}
	err := s.client.SearchRead(SaleOrderLineModel, args, nil, so)
	return so, err
}

func (s *SaleOrderLineService) Create(required *types.RequiredSaleOrderLineFields, fields map[string]interface{}, fields2Many *types.Handle2ManysStruct) (int, error) {
	var args []interface{}
	var id int
	fields["order_id"] = required.OrderID
	fields["product_id"] = required.ProductID
	if fields2Many != nil {
		fields = types.Handle2Manys(fields, fields2Many)
	}
	args = append(args, fields)
	err := s.client.Create(SaleOrderLineModel, args, &id)
	return id, err
}

func (s *SaleOrderLineService) Update(ids []int, fields map[string]interface{}, fields2Many *types.Handle2ManysStruct) error {
	var args []interface{}
	args = append(args, ids)
	if fields2Many != nil {
		fields = types.Handle2Manys(fields, fields2Many)
	}
	args = append(args, fields)
	err := s.client.Update(SaleOrderLineModel, args)
	return err
}

func (s *SaleOrderLineService) Delete(ids []int) error {
	var args []interface{}
	args = append(args, ids)
	return s.client.Delete("sale.order", args)
}
