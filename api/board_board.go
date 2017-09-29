package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type BoardBoardService struct {
	client *Client
}

func NewBoardBoardService(c *Client) *BoardBoardService {
	return &BoardBoardService{client: c}
}

func (svc *BoardBoardService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.BoardBoardModel, name)
}

func (svc *BoardBoardService) GetByIds(ids []int64) (*types.BoardBoards, error) {
	b := &types.BoardBoards{}
	return b, svc.client.getByIds(types.BoardBoardModel, ids, b)
}

func (svc *BoardBoardService) GetByName(name string) (*types.BoardBoards, error) {
	b := &types.BoardBoards{}
	return b, svc.client.getByName(types.BoardBoardModel, name, b)
}

func (svc *BoardBoardService) GetByField(field string, value string) (*types.BoardBoards, error) {
	b := &types.BoardBoards{}
	return b, svc.client.getByField(types.BoardBoardModel, field, value, b)
}

func (svc *BoardBoardService) GetAll() (*types.BoardBoards, error) {
	b := &types.BoardBoards{}
	return b, svc.client.getAll(types.BoardBoardModel, b)
}

func (svc *BoardBoardService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.BoardBoardModel, fields, relations)
}

func (svc *BoardBoardService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.BoardBoardModel, ids, fields, relations)
}

func (svc *BoardBoardService) Delete(ids []int64) error {
	return svc.client.delete(types.BoardBoardModel, ids)
}
