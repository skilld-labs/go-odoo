package odoo

import (
	"fmt"
)

// IrActionsClient represents ir.actions.client model.
type IrActionsClient struct {
	LastUpdate     *Time      `xmlrpc:"__last_update,omitempty"`
	BindingModelId *Many2One  `xmlrpc:"binding_model_id,omitempty"`
	BindingType    *Selection `xmlrpc:"binding_type,omitempty"`
	Context        *String    `xmlrpc:"context,omitempty"`
	CreateDate     *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid      *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName    *String    `xmlrpc:"display_name,omitempty"`
	Help           *String    `xmlrpc:"help,omitempty"`
	Id             *Int       `xmlrpc:"id,omitempty"`
	Name           *String    `xmlrpc:"name,omitempty"`
	Params         *String    `xmlrpc:"params,omitempty"`
	ParamsStore    *String    `xmlrpc:"params_store,omitempty"`
	ResModel       *String    `xmlrpc:"res_model,omitempty"`
	Tag            *String    `xmlrpc:"tag,omitempty"`
	Target         *Selection `xmlrpc:"target,omitempty"`
	Type           *String    `xmlrpc:"type,omitempty"`
	WriteDate      *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid       *Many2One  `xmlrpc:"write_uid,omitempty"`
	XmlId          *String    `xmlrpc:"xml_id,omitempty"`
}

// IrActionsClients represents array of ir.actions.client model.
type IrActionsClients []IrActionsClient

// IrActionsClientModel is the odoo model name.
const IrActionsClientModel = "ir.actions.client"

// Many2One convert IrActionsClient to *Many2One.
func (iac *IrActionsClient) Many2One() *Many2One {
	return NewMany2One(iac.Id.Get(), "")
}

// CreateIrActionsClient creates a new ir.actions.client model and returns its id.
func (c *Client) CreateIrActionsClient(iac *IrActionsClient) (int64, error) {
	return c.Create(IrActionsClientModel, iac)
}

// UpdateIrActionsClient updates an existing ir.actions.client record.
func (c *Client) UpdateIrActionsClient(iac *IrActionsClient) error {
	return c.UpdateIrActionsClients([]int64{iac.Id.Get()}, iac)
}

// UpdateIrActionsClients updates existing ir.actions.client records.
// All records (represented by ids) will be updated by iac values.
func (c *Client) UpdateIrActionsClients(ids []int64, iac *IrActionsClient) error {
	return c.Update(IrActionsClientModel, ids, iac)
}

// DeleteIrActionsClient deletes an existing ir.actions.client record.
func (c *Client) DeleteIrActionsClient(id int64) error {
	return c.DeleteIrActionsClients([]int64{id})
}

// DeleteIrActionsClients deletes existing ir.actions.client records.
func (c *Client) DeleteIrActionsClients(ids []int64) error {
	return c.Delete(IrActionsClientModel, ids)
}

// GetIrActionsClient gets ir.actions.client existing record.
func (c *Client) GetIrActionsClient(id int64) (*IrActionsClient, error) {
	iacs, err := c.GetIrActionsClients([]int64{id})
	if err != nil {
		return nil, err
	}
	if iacs != nil && len(*iacs) > 0 {
		return &((*iacs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of ir.actions.client not found", id)
}

// GetIrActionsClients gets ir.actions.client existing records.
func (c *Client) GetIrActionsClients(ids []int64) (*IrActionsClients, error) {
	iacs := &IrActionsClients{}
	if err := c.Read(IrActionsClientModel, ids, nil, iacs); err != nil {
		return nil, err
	}
	return iacs, nil
}

// FindIrActionsClient finds ir.actions.client record by querying it with criteria.
func (c *Client) FindIrActionsClient(criteria *Criteria) (*IrActionsClient, error) {
	iacs := &IrActionsClients{}
	if err := c.SearchRead(IrActionsClientModel, criteria, NewOptions().Limit(1), iacs); err != nil {
		return nil, err
	}
	if iacs != nil && len(*iacs) > 0 {
		return &((*iacs)[0]), nil
	}
	return nil, fmt.Errorf("no ir.actions.client was found with criteria %v", criteria)
}

// FindIrActionsClients finds ir.actions.client records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrActionsClients(criteria *Criteria, options *Options) (*IrActionsClients, error) {
	iacs := &IrActionsClients{}
	if err := c.SearchRead(IrActionsClientModel, criteria, options, iacs); err != nil {
		return nil, err
	}
	return iacs, nil
}

// FindIrActionsClientIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrActionsClientIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(IrActionsClientModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindIrActionsClientId finds record id by querying it with criteria.
func (c *Client) FindIrActionsClientId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrActionsClientModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no ir.actions.client was found with criteria %v and options %v", criteria, options)
}
