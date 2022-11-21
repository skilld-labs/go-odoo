package odoo

import (
	"fmt"
)

// IrActionsServer represents ir.actions.server model.
type IrActionsServer struct {
	LastUpdate     *Time      `xmlrpc:"__last_update,omptempty"`
	BindingModelId *Many2One  `xmlrpc:"binding_model_id,omptempty"`
	BindingType    *Selection `xmlrpc:"binding_type,omptempty"`
	ChannelIds     *Relation  `xmlrpc:"channel_ids,omptempty"`
	ChildIds       *Relation  `xmlrpc:"child_ids,omptempty"`
	Code           *String    `xmlrpc:"code,omptempty"`
	CreateDate     *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid      *Many2One  `xmlrpc:"create_uid,omptempty"`
	CrudModelId    *Many2One  `xmlrpc:"crud_model_id,omptempty"`
	CrudModelName  *String    `xmlrpc:"crud_model_name,omptempty"`
	DisplayName    *String    `xmlrpc:"display_name,omptempty"`
	FieldsLines    *Relation  `xmlrpc:"fields_lines,omptempty"`
	Help           *String    `xmlrpc:"help,omptempty"`
	Id             *Int       `xmlrpc:"id,omptempty"`
	LinkFieldId    *Many2One  `xmlrpc:"link_field_id,omptempty"`
	ModelId        *Many2One  `xmlrpc:"model_id,omptempty"`
	ModelName      *String    `xmlrpc:"model_name,omptempty"`
	Name           *String    `xmlrpc:"name,omptempty"`
	PartnerIds     *Relation  `xmlrpc:"partner_ids,omptempty"`
	Sequence       *Int       `xmlrpc:"sequence,omptempty"`
	State          *Selection `xmlrpc:"state,omptempty"`
	TemplateId     *Many2One  `xmlrpc:"template_id,omptempty"`
	Type           *String    `xmlrpc:"type,omptempty"`
	Usage          *Selection `xmlrpc:"usage,omptempty"`
	WriteDate      *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid       *Many2One  `xmlrpc:"write_uid,omptempty"`
	XmlId          *String    `xmlrpc:"xml_id,omptempty"`
}

// IrActionsServers represents array of ir.actions.server model.
type IrActionsServers []IrActionsServer

// IrActionsServerModel is the odoo model name.
const IrActionsServerModel = "ir.actions.server"

// Many2One convert IrActionsServer to *Many2One.
func (ias *IrActionsServer) Many2One() *Many2One {
	return NewMany2One(ias.Id.Get(), "")
}

// CreateIrActionsServer creates a new ir.actions.server model and returns its id.
func (c *Client) CreateIrActionsServer(ias *IrActionsServer) (int64, error) {
	return c.Create(IrActionsServerModel, ias)
}

// UpdateIrActionsServer updates an existing ir.actions.server record.
func (c *Client) UpdateIrActionsServer(ias *IrActionsServer) error {
	return c.UpdateIrActionsServers([]int64{ias.Id.Get()}, ias)
}

// UpdateIrActionsServers updates existing ir.actions.server records.
// All records (represented by ids) will be updated by ias values.
func (c *Client) UpdateIrActionsServers(ids []int64, ias *IrActionsServer) error {
	return c.Update(IrActionsServerModel, ids, ias)
}

// DeleteIrActionsServer deletes an existing ir.actions.server record.
func (c *Client) DeleteIrActionsServer(id int64) error {
	return c.DeleteIrActionsServers([]int64{id})
}

// DeleteIrActionsServers deletes existing ir.actions.server records.
func (c *Client) DeleteIrActionsServers(ids []int64) error {
	return c.Delete(IrActionsServerModel, ids)
}

// GetIrActionsServer gets ir.actions.server existing record.
func (c *Client) GetIrActionsServer(id int64) (*IrActionsServer, error) {
	iass, err := c.GetIrActionsServers([]int64{id})
	if err != nil {
		return nil, err
	}
	if iass != nil && len(*iass) > 0 {
		return &((*iass)[0]), nil
	}
	return nil, fmt.Errorf("id %v of ir.actions.server not found", id)
}

// GetIrActionsServers gets ir.actions.server existing records.
func (c *Client) GetIrActionsServers(ids []int64) (*IrActionsServers, error) {
	iass := &IrActionsServers{}
	if err := c.Read(IrActionsServerModel, ids, nil, iass); err != nil {
		return nil, err
	}
	return iass, nil
}

// FindIrActionsServer finds ir.actions.server record by querying it with criteria.
func (c *Client) FindIrActionsServer(criteria *Criteria) (*IrActionsServer, error) {
	iass := &IrActionsServers{}
	if err := c.SearchRead(IrActionsServerModel, criteria, NewOptions().Limit(1), iass); err != nil {
		return nil, err
	}
	if iass != nil && len(*iass) > 0 {
		return &((*iass)[0]), nil
	}
	return nil, fmt.Errorf("no ir.actions.server was found with criteria %v", criteria)
}

// FindIrActionsServers finds ir.actions.server records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrActionsServers(criteria *Criteria, options *Options) (*IrActionsServers, error) {
	iass := &IrActionsServers{}
	if err := c.SearchRead(IrActionsServerModel, criteria, options, iass); err != nil {
		return nil, err
	}
	return iass, nil
}

// FindIrActionsServerIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrActionsServerIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(IrActionsServerModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindIrActionsServerId finds record id by querying it with criteria.
func (c *Client) FindIrActionsServerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrActionsServerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no ir.actions.server was found with criteria %v and options %v", criteria, options)
}
