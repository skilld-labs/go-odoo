package odoo

import (
	"fmt"
)

// IrRule represents ir.rule model.
type IrRule struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	Active      *Bool     `xmlrpc:"active,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	DomainForce *String   `xmlrpc:"domain_force,omitempty"`
	Global      *Bool     `xmlrpc:"global,omitempty"`
	Groups      *Relation `xmlrpc:"groups,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	ModelId     *Many2One `xmlrpc:"model_id,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	PermCreate  *Bool     `xmlrpc:"perm_create,omitempty"`
	PermRead    *Bool     `xmlrpc:"perm_read,omitempty"`
	PermUnlink  *Bool     `xmlrpc:"perm_unlink,omitempty"`
	PermWrite   *Bool     `xmlrpc:"perm_write,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// IrRules represents array of ir.rule model.
type IrRules []IrRule

// IrRuleModel is the odoo model name.
const IrRuleModel = "ir.rule"

// Many2One convert IrRule to *Many2One.
func (ir *IrRule) Many2One() *Many2One {
	return NewMany2One(ir.Id.Get(), "")
}

// CreateIrRule creates a new ir.rule model and returns its id.
func (c *Client) CreateIrRule(ir *IrRule) (int64, error) {
	return c.Create(IrRuleModel, ir)
}

// UpdateIrRule updates an existing ir.rule record.
func (c *Client) UpdateIrRule(ir *IrRule) error {
	return c.UpdateIrRules([]int64{ir.Id.Get()}, ir)
}

// UpdateIrRules updates existing ir.rule records.
// All records (represented by IDs) will be updated by ir values.
func (c *Client) UpdateIrRules(ids []int64, ir *IrRule) error {
	return c.Update(IrRuleModel, ids, ir)
}

// DeleteIrRule deletes an existing ir.rule record.
func (c *Client) DeleteIrRule(id int64) error {
	return c.DeleteIrRules([]int64{id})
}

// DeleteIrRules deletes existing ir.rule records.
func (c *Client) DeleteIrRules(ids []int64) error {
	return c.Delete(IrRuleModel, ids)
}

// GetIrRule gets ir.rule existing record.
func (c *Client) GetIrRule(id int64) (*IrRule, error) {
	irs, err := c.GetIrRules([]int64{id})
	if err != nil {
		return nil, err
	}
	if irs != nil && len(*irs) > 0 {
		return &((*irs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of ir.rule not found", id)
}

// GetIrRules gets ir.rule existing records.
func (c *Client) GetIrRules(ids []int64) (*IrRules, error) {
	irs := &IrRules{}
	if err := c.Read(IrRuleModel, ids, nil, irs); err != nil {
		return nil, err
	}
	return irs, nil
}

// FindIrRule finds ir.rule record by querying it with criteria.
func (c *Client) FindIrRule(criteria *Criteria) (*IrRule, error) {
	irs := &IrRules{}
	if err := c.SearchRead(IrRuleModel, criteria, NewOptions().Limit(1), irs); err != nil {
		return nil, err
	}
	if irs != nil && len(*irs) > 0 {
		return &((*irs)[0]), nil
	}
	return nil, fmt.Errorf("ir.rule was not found")
}

// FindIrRules finds ir.rule records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrRules(criteria *Criteria, options *Options) (*IrRules, error) {
	irs := &IrRules{}
	if err := c.SearchRead(IrRuleModel, criteria, options, irs); err != nil {
		return nil, err
	}
	return irs, nil
}

// FindIrRuleIds finds records IDs by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrRuleIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(IrRuleModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindIrRuleId finds record id by querying it with criteria.
func (c *Client) FindIrRuleId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrRuleModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("ir.rule was not found")
}
