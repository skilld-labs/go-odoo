package odoo

import (
	"fmt"
)

// IrModelData represents ir.model.data model.
type IrModelData struct {
	LastUpdate   *Time     `xmlrpc:"__last_update,omitempty"`
	CompleteName *String   `xmlrpc:"complete_name,omitempty"`
	CreateDate   *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid    *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName  *String   `xmlrpc:"display_name,omitempty"`
	Id           *Int      `xmlrpc:"id,omitempty"`
	Model        *String   `xmlrpc:"model,omitempty"`
	Module       *String   `xmlrpc:"module,omitempty"`
	Name         *String   `xmlrpc:"name,omitempty"`
	Noupdate     *Bool     `xmlrpc:"noupdate,omitempty"`
	Reference    *String   `xmlrpc:"reference,omitempty"`
	ResId        *Many2One `xmlrpc:"res_id,omitempty"`
	WriteDate    *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid     *Many2One `xmlrpc:"write_uid,omitempty"`
}

// IrModelDatas represents array of ir.model.data model.
type IrModelDatas []IrModelData

// IrModelDataModel is the odoo model name.
const IrModelDataModel = "ir.model.data"

// Many2One convert IrModelData to *Many2One.
func (imd *IrModelData) Many2One() *Many2One {
	return NewMany2One(imd.Id.Get(), "")
}

// CreateIrModelData creates a new ir.model.data model and returns its id.
func (c *Client) CreateIrModelData(imd *IrModelData) (int64, error) {
	return c.Create(IrModelDataModel, imd)
}

// UpdateIrModelData updates an existing ir.model.data record.
func (c *Client) UpdateIrModelData(imd *IrModelData) error {
	return c.UpdateIrModelDatas([]int64{imd.Id.Get()}, imd)
}

// UpdateIrModelDatas updates existing ir.model.data records.
// All records (represented by ids) will be updated by imd values.
func (c *Client) UpdateIrModelDatas(ids []int64, imd *IrModelData) error {
	return c.Update(IrModelDataModel, ids, imd)
}

// DeleteIrModelData deletes an existing ir.model.data record.
func (c *Client) DeleteIrModelData(id int64) error {
	return c.DeleteIrModelDatas([]int64{id})
}

// DeleteIrModelDatas deletes existing ir.model.data records.
func (c *Client) DeleteIrModelDatas(ids []int64) error {
	return c.Delete(IrModelDataModel, ids)
}

// GetIrModelData gets ir.model.data existing record.
func (c *Client) GetIrModelData(id int64) (*IrModelData, error) {
	imds, err := c.GetIrModelDatas([]int64{id})
	if err != nil {
		return nil, err
	}
	if imds != nil && len(*imds) > 0 {
		return &((*imds)[0]), nil
	}
	return nil, fmt.Errorf("id %v of ir.model.data not found", id)
}

// GetIrModelDatas gets ir.model.data existing records.
func (c *Client) GetIrModelDatas(ids []int64) (*IrModelDatas, error) {
	imds := &IrModelDatas{}
	if err := c.Read(IrModelDataModel, ids, nil, imds); err != nil {
		return nil, err
	}
	return imds, nil
}

// FindIrModelData finds ir.model.data record by querying it with criteria.
func (c *Client) FindIrModelData(criteria *Criteria) (*IrModelData, error) {
	imds := &IrModelDatas{}
	if err := c.SearchRead(IrModelDataModel, criteria, NewOptions().Limit(1), imds); err != nil {
		return nil, err
	}
	if imds != nil && len(*imds) > 0 {
		return &((*imds)[0]), nil
	}
	return nil, fmt.Errorf("ir.model.data was not found")
}

// FindIrModelDatas finds ir.model.data records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrModelDatas(criteria *Criteria, options *Options) (*IrModelDatas, error) {
	imds := &IrModelDatas{}
	if err := c.SearchRead(IrModelDataModel, criteria, options, imds); err != nil {
		return nil, err
	}
	return imds, nil
}

// FindIrModelDataIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrModelDataIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(IrModelDataModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindIrModelDataId finds record id by querying it with criteria.
func (c *Client) FindIrModelDataId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrModelDataModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("ir.model.data was not found")
}
