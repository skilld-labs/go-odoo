package odoo

import (
	"fmt"
)

// IrRule represents ir.rule model.
type IrRule struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	Active      *Bool     `xmlrpc:"active,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	DomainForce *String   `xmlrpc:"domain_force,omptempty"`
	Global      *Bool     `xmlrpc:"global,omptempty"`
	Groups      *Relation `xmlrpc:"groups,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	ModelId     *Many2One `xmlrpc:"model_id,omptempty"`
	Name        *String   `xmlrpc:"name,omptempty"`
	PermCreate  *Bool     `xmlrpc:"perm_create,omptempty"`
	PermRead    *Bool     `xmlrpc:"perm_read,omptempty"`
	PermUnlink  *Bool     `xmlrpc:"perm_unlink,omptempty"`
	PermWrite   *Bool     `xmlrpc:"perm_write,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
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
// All records (represented by ids) will be updated by ir values.
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
	return nil, fmt.Errorf("ir.rule was not found with criteria %v", criteria)
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

// FindIrRuleIds finds records ids by querying it
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
	return -1, fmt.Errorf("ir.rule was not found with criteria %v and options %v", criteria, options)
}
