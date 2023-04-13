package odoo

import (
	"fmt"
)

// CrmPartnerBinding represents crm.partner.binding model.
type CrmPartnerBinding struct {
	LastUpdate  *Time      `xmlrpc:"__last_update,omptempty"`
	Action      *Selection `xmlrpc:"action,omptempty"`
	CreateDate  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName *String    `xmlrpc:"display_name,omptempty"`
	Id          *Int       `xmlrpc:"id,omptempty"`
	PartnerId   *Many2One  `xmlrpc:"partner_id,omptempty"`
	WriteDate   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// CrmPartnerBindings represents array of crm.partner.binding model.
type CrmPartnerBindings []CrmPartnerBinding

// CrmPartnerBindingModel is the odoo model name.
const CrmPartnerBindingModel = "crm.partner.binding"

// Many2One convert CrmPartnerBinding to *Many2One.
func (cpb *CrmPartnerBinding) Many2One() *Many2One {
	return NewMany2One(cpb.Id.Get(), "")
}

// CreateCrmPartnerBinding creates a new crm.partner.binding model and returns its id.
func (c *Client) CreateCrmPartnerBinding(cpb *CrmPartnerBinding) (int64, error) {
	ids, err := c.Create(CrmPartnerBindingModel, []interface{}{cpb})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateCrmPartnerBinding creates a new crm.partner.binding model and returns its id.
func (c *Client) CreateCrmPartnerBindings(cpbs []*CrmPartnerBinding) ([]int64, error) {
	var vv []interface{}
	for _, v := range cpbs {
		vv = append(vv, v)
	}
	return c.Create(CrmPartnerBindingModel, vv)
}

// UpdateCrmPartnerBinding updates an existing crm.partner.binding record.
func (c *Client) UpdateCrmPartnerBinding(cpb *CrmPartnerBinding) error {
	return c.UpdateCrmPartnerBindings([]int64{cpb.Id.Get()}, cpb)
}

// UpdateCrmPartnerBindings updates existing crm.partner.binding records.
// All records (represented by ids) will be updated by cpb values.
func (c *Client) UpdateCrmPartnerBindings(ids []int64, cpb *CrmPartnerBinding) error {
	return c.Update(CrmPartnerBindingModel, ids, cpb)
}

// DeleteCrmPartnerBinding deletes an existing crm.partner.binding record.
func (c *Client) DeleteCrmPartnerBinding(id int64) error {
	return c.DeleteCrmPartnerBindings([]int64{id})
}

// DeleteCrmPartnerBindings deletes existing crm.partner.binding records.
func (c *Client) DeleteCrmPartnerBindings(ids []int64) error {
	return c.Delete(CrmPartnerBindingModel, ids)
}

// GetCrmPartnerBinding gets crm.partner.binding existing record.
func (c *Client) GetCrmPartnerBinding(id int64) (*CrmPartnerBinding, error) {
	cpbs, err := c.GetCrmPartnerBindings([]int64{id})
	if err != nil {
		return nil, err
	}
	if cpbs != nil && len(*cpbs) > 0 {
		return &((*cpbs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of crm.partner.binding not found", id)
}

// GetCrmPartnerBindings gets crm.partner.binding existing records.
func (c *Client) GetCrmPartnerBindings(ids []int64) (*CrmPartnerBindings, error) {
	cpbs := &CrmPartnerBindings{}
	if err := c.Read(CrmPartnerBindingModel, ids, nil, cpbs); err != nil {
		return nil, err
	}
	return cpbs, nil
}

// FindCrmPartnerBinding finds crm.partner.binding record by querying it with criteria.
func (c *Client) FindCrmPartnerBinding(criteria *Criteria) (*CrmPartnerBinding, error) {
	cpbs := &CrmPartnerBindings{}
	if err := c.SearchRead(CrmPartnerBindingModel, criteria, NewOptions().Limit(1), cpbs); err != nil {
		return nil, err
	}
	if cpbs != nil && len(*cpbs) > 0 {
		return &((*cpbs)[0]), nil
	}
	return nil, fmt.Errorf("crm.partner.binding was not found with criteria %v", criteria)
}

// FindCrmPartnerBindings finds crm.partner.binding records by querying it
// and filtering it with criteria and options.
func (c *Client) FindCrmPartnerBindings(criteria *Criteria, options *Options) (*CrmPartnerBindings, error) {
	cpbs := &CrmPartnerBindings{}
	if err := c.SearchRead(CrmPartnerBindingModel, criteria, options, cpbs); err != nil {
		return nil, err
	}
	return cpbs, nil
}

// FindCrmPartnerBindingIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindCrmPartnerBindingIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(CrmPartnerBindingModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindCrmPartnerBindingId finds record id by querying it with criteria.
func (c *Client) FindCrmPartnerBindingId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CrmPartnerBindingModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("crm.partner.binding was not found with criteria %v and options %v", criteria, options)
}
