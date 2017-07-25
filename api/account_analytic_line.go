package api

import (
	"../types"
)

var AccountAnalyticLineModel string = "account.analytic.line"

type AccountAnalyticLineService struct {
	client *Client
}

func (c *Client) NewAccountAnalyticLineService() {
	c.AccountAnalyticLine = &AccountAnalyticLineService{client: c}
	return
}

func (s *AccountAnalyticLineService) GetIdByName(name string) ([]int, error) {
	var args []interface{}
	var ids []int
	args = append(args, []interface{}{"name", "=", name})
	err := s.client.Search(AccountAnalyticLineModel, args, nil, &ids)
	return ids, err
}

func (s *AccountAnalyticLineService) GetByIds(ids []int) (*types.AccountAnalyticLines, error) {
	var args []interface{}
	args = append(args, ids)
	so := &types.AccountAnalyticLines{}
	err := s.client.Read(AccountAnalyticLineModel, args, nil, so)
	return so, err
}

func (s *AccountAnalyticLineService) GetByName(name string) (*types.AccountAnalyticLines, error) {
	var args []interface{}
	var args2 []interface{}
	args2 = append(args2, []string{"name", "=", name})
	args = append(args, args2)
	so := &types.AccountAnalyticLines{}
	err := s.client.SearchRead(AccountAnalyticLineModel, args, nil, so)
	return so, err
}

func (s *AccountAnalyticLineService) GetAll() (*types.AccountAnalyticLines, error) {
	var args []interface{}
	var args2 []interface{}
	args = append(args, args2)
	so := &types.AccountAnalyticLines{}
	err := s.client.SearchRead(AccountAnalyticLineModel, args, nil, so)
	return so, err
}

func (s *AccountAnalyticLineService) Create(required *types.RequiredAccountAnalyticLineFields, fields map[string]interface{}, fields2Many *types.Handle2ManysStruct) (int, error) {
	var args []interface{}
	var id int
	fields["name"] = required.Name
	fields["account_id"] = require.AccountID
	if fields2Many != nil {
		fields = types.Handle2Manys(fields, fields2Many)
	}
	args = append(args, fields)
	err := s.client.Create(AccountAnalyticLineModel, args, &id)
	return id, err
}

func (s *AccountAnalyticLineService) Update(ids []int, fields map[string]interface{}, fields2Many *types.Handle2ManysStruct) error {
	var args []interface{}
	args = append(args, ids)
	if fields2Many != nil {
		fields = types.Handle2Manys(fields, fields2Many)
	}
	args = append(args, fields)
	err := s.client.Update(AccountAnalyticLineModel, args)
	return err
}

func (s *AccountAnalyticLineService) Delete(ids []int) error {
	var args []interface{}
	args = append(args, ids)
	return s.client.Delete("sale.order", args)
}
