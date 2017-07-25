package api

import (
	"../types"
)

var AccountInvoiceModel string = "account.invoice"

type AccountInvoiceService struct {
	client *Client
}

func (c *Client) NewAccountInvoiceService() {
	c.AccountInvoice = &AccountInvoiceService{client: c}
	return
}

func (s *AccountInvoiceService) GetIdByName(name string) ([]int, error) {
	var args []interface{}
	var ids []int
	args = append(args, []interface{}{"name", "=", name})
	err := s.client.Search(AccountInvoiceModel, args, nil, &ids)
	return ids, err
}

func (s *AccountInvoiceService) GetByIds(ids []int) (*types.AccountInvoices, error) {
	var args []interface{}
	args = append(args, ids)
	so := &types.AccountInvoices{}
	err := s.client.Read(AccountInvoiceModel, args, nil, so)
	return so, err
}

func (s *AccountInvoiceService) GetByName(name string) (*types.AccountInvoices, error) {
	var args []interface{}
	var args2 []interface{}
	args2 = append(args2, []string{"name", "=", name})
	args = append(args, args2)
	so := &types.AccountInvoices{}
	err := s.client.SearchRead(AccountInvoiceModel, args, nil, so)
	return so, err
}

func (s *AccountInvoiceService) GetAll() (*types.AccountInvoices, error) {
	var args []interface{}
	var args2 []interface{}
	args = append(args, args2)
	so := &types.AccountInvoices{}
	err := s.client.SearchRead(AccountInvoiceModel, args, nil, so)
	return so, err
}

func (s *AccountInvoiceService) Create(required *types.RequiredAccountInvoiceFields, fields map[string]interface{}, fields2Many *types.Handle2ManysStruct) (int, error) {
	var args []interface{}
	var id int
	fields["name"] = required.Name
	fields["partner_id"] = required.PartnerID
	if fields2Many != nil {
		fields = types.Handle2Manys(fields, fields2Many)
	}
	args = append(args, fields)
	err := s.client.Create(AccountInvoiceModel, args, &id)
	return id, err
}

func (s *AccountInvoiceService) Update(ids []int, fields map[string]interface{}, fields2Many *types.Handle2ManysStruct) error {
	var args []interface{}
	args = append(args, ids)
	if fields2Many != nil {
		fields = types.Handle2Manys(fields, fields2Many)
	}
	args = append(args, fields)
	err := s.client.Update(AccountInvoiceModel, args)
	return err
}

func (s *AccountInvoiceService) Delete(ids []int) error {
	var args []interface{}
	args = append(args, ids)
	return s.client.Delete("sale.order", args)
}
