package api

import (
	"../types"
)

var AccountInvoiceLineModel string = "account.invoice.line"

type AccountInvoiceLineService struct {
	client *Client
}

func (c *Client) NewAccountInvoiceLineService() {
	c.AccountInvoiceLine = &AccountInvoiceLineService{client: c}
	return
}

func (s *AccountInvoiceLineService) GetIdByName(name string) ([]int, error) {
	var args []interface{}
	var ids []int
	args = append(args, []interface{}{"name", "=", name})
	err := s.client.Search(AccountInvoiceLineModel, args, nil, &ids)
	return ids, err
}

func (s *AccountInvoiceLineService) GetByIds(ids []int) (*types.AccountInvoiceLines, error) {
	var args []interface{}
	args = append(args, ids)
	so := &types.AccountInvoiceLines{}
	err := s.client.Read(AccountInvoiceLineModel, args, nil, so)
	return so, err
}

func (s *AccountInvoiceLineService) GetByName(name string) (*types.AccountInvoiceLines, error) {
	var args []interface{}
	var args2 []interface{}
	args2 = append(args2, []string{"name", "=", name})
	args = append(args, args2)
	so := &types.AccountInvoiceLines{}
	err := s.client.SearchRead(AccountInvoiceLineModel, args, nil, so)
	return so, err
}

func (s *AccountInvoiceLineService) GetAll() (*types.AccountInvoiceLines, error) {
	var args []interface{}
	var args2 []interface{}
	args = append(args, args2)
	so := &types.AccountInvoiceLines{}
	err := s.client.SearchRead(AccountInvoiceLineModel, args, nil, so)
	return so, err
}

func (s *AccountInvoiceLineService) Create(required *types.RequiredAccountInvoiceLineFields, fields map[string]interface{}, fields2Many *types.Handle2ManysStruct) (int, error) {
	var args []interface{}
	var id int
	fields["name"] = required.Name
	fields["price_unit"] = required.PriceUnit
	fields["account_id"] = required.AccountID
	if fields2Many != nil {
		fields = types.Handle2Manys(fields, fields2Many)
	}
	args = append(args, fields)
	err := s.client.Create(AccountInvoiceLineModel, args, &id)
	return id, err
}

func (s *AccountInvoiceLineService) Update(ids []int, fields map[string]interface{}, fields2Many *types.Handle2ManysStruct) error {
	var args []interface{}
	args = append(args, ids)
	if fields2Many != nil {
		fields = types.Handle2Manys(fields, fields2Many)
	}
	args = append(args, fields)
	err := s.client.Update(AccountInvoiceLineModel, args)
	return err
}

func (s *AccountInvoiceLineService) Delete(ids []int) error {
	var args []interface{}
	args = append(args, ids)
	return s.client.Delete("sale.order", args)
}
