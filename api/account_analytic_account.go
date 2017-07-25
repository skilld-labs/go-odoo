package api

import (
	"../types"
)

var AccountAnalyticAccountModel string = "account.analytic.account"

type AccountAnalyticAccountService struct {
	client *Client
}

func (c *Client) NewAccountAnalyticAccountService() {
	c.AccountAnalyticAccount = &AccountAnalyticAccountService{client: c}
	return
}

func (s *AccountAnalyticAccountService) GetIdByName(name string) ([]int, error) {
	var args []interface{}
	var ids []int
	args = append(args, []interface{}{"name", "=", name})
	err := s.client.Search(AccountAnalyticAccountModel, args, nil, &ids)
	return ids, err
}

func (s *AccountAnalyticAccountService) GetByIds(ids []int) (*types.AccountAnalyticAccounts, error) {
	var args []interface{}
	args = append(args, ids)
	so := &types.AccountAnalyticAccounts{}
	err := s.client.Read(AccountAnalyticAccountModel, args, nil, so)
	return so, err
}

func (s *AccountAnalyticAccountService) GetByName(name string) (*types.AccountAnalyticAccounts, error) {
	var args []interface{}
	var args2 []interface{}
	args2 = append(args2, []string{"name", "=", name})
	args = append(args, args2)
	so := &types.AccountAnalyticAccounts{}
	err := s.client.SearchRead(AccountAnalyticAccountModel, args, nil, so)
	return so, err
}

func (s *AccountAnalyticAccountService) GetAll() (*types.AccountAnalyticAccounts, error) {
	var args []interface{}
	var args2 []interface{}
	args = append(args, args2)
	so := &types.AccountAnalyticAccounts{}
	err := s.client.SearchRead(AccountAnalyticAccountModel, args, nil, so)
	return so, err
}

func (s *AccountAnalyticAccountService) Create(required *types.RequiredAccountAnalyticAccountFields, fields map[string]interface{}, fields2Many *types.Handle2ManysStruct) (int, error) {
	var args []interface{}
	var id int
	fields["name"] = required.Name
	if fields2Many != nil {
		fields = types.Handle2Manys(fields, fields2Many)
	}
	args = append(args, fields)
	err := s.client.Create(AccountAnalyticAccountModel, args, &id)
	return id, err
}

func (s *AccountAnalyticAccountService) Update(ids []int, fields map[string]interface{}, fields2Many *types.Handle2ManysStruct) error {
	var args []interface{}
	args = append(args, ids)
	if fields2Many != nil {
		fields = types.Handle2Manys(fields, fields2Many)
	}
	args = append(args, fields)
	err := s.client.Update(AccountAnalyticAccountModel, args)
	return err
}

func (s *AccountAnalyticAccountService) Delete(ids []int) error {
	var args []interface{}
	args = append(args, ids)
	return s.client.Delete("sale.order", args)
}
