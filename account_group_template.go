package odoo

import (
	"fmt"
)

// AccountGroupTemplate represents account.group.template model.
type AccountGroupTemplate struct {
	LastUpdate      *Time     `xmlrpc:"__last_update,omitempty"`
	ChartTemplateId *Many2One `xmlrpc:"chart_template_id,omitempty"`
	CodePrefixEnd   *String   `xmlrpc:"code_prefix_end,omitempty"`
	CodePrefixStart *String   `xmlrpc:"code_prefix_start,omitempty"`
	CreateDate      *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid       *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName     *String   `xmlrpc:"display_name,omitempty"`
	Id              *Int      `xmlrpc:"id,omitempty"`
	Name            *String   `xmlrpc:"name,omitempty"`
	ParentId        *Many2One `xmlrpc:"parent_id,omitempty"`
	WriteDate       *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid        *Many2One `xmlrpc:"write_uid,omitempty"`
}

// AccountGroupTemplates represents array of account.group.template model.
type AccountGroupTemplates []AccountGroupTemplate

// AccountGroupTemplateModel is the odoo model name.
const AccountGroupTemplateModel = "account.group.template"

// Many2One convert AccountGroupTemplate to *Many2One.
func (agt *AccountGroupTemplate) Many2One() *Many2One {
	return NewMany2One(agt.Id.Get(), "")
}

// CreateAccountGroupTemplate creates a new account.group.template model and returns its id.
func (c *Client) CreateAccountGroupTemplate(agt *AccountGroupTemplate) (int64, error) {
	return c.Create(AccountGroupTemplateModel, agt)
}

// UpdateAccountGroupTemplate updates an existing account.group.template record.
func (c *Client) UpdateAccountGroupTemplate(agt *AccountGroupTemplate) error {
	return c.UpdateAccountGroupTemplates([]int64{agt.Id.Get()}, agt)
}

// UpdateAccountGroupTemplates updates existing account.group.template records.
// All records (represented by ids) will be updated by agt values.
func (c *Client) UpdateAccountGroupTemplates(ids []int64, agt *AccountGroupTemplate) error {
	return c.Update(AccountGroupTemplateModel, ids, agt)
}

// DeleteAccountGroupTemplate deletes an existing account.group.template record.
func (c *Client) DeleteAccountGroupTemplate(id int64) error {
	return c.DeleteAccountGroupTemplates([]int64{id})
}

// DeleteAccountGroupTemplates deletes existing account.group.template records.
func (c *Client) DeleteAccountGroupTemplates(ids []int64) error {
	return c.Delete(AccountGroupTemplateModel, ids)
}

// GetAccountGroupTemplate gets account.group.template existing record.
func (c *Client) GetAccountGroupTemplate(id int64) (*AccountGroupTemplate, error) {
	agts, err := c.GetAccountGroupTemplates([]int64{id})
	if err != nil {
		return nil, err
	}
	if agts != nil && len(*agts) > 0 {
		return &((*agts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.group.template not found", id)
}

// GetAccountGroupTemplates gets account.group.template existing records.
func (c *Client) GetAccountGroupTemplates(ids []int64) (*AccountGroupTemplates, error) {
	agts := &AccountGroupTemplates{}
	if err := c.Read(AccountGroupTemplateModel, ids, nil, agts); err != nil {
		return nil, err
	}
	return agts, nil
}

// FindAccountGroupTemplate finds account.group.template record by querying it with criteria.
func (c *Client) FindAccountGroupTemplate(criteria *Criteria) (*AccountGroupTemplate, error) {
	agts := &AccountGroupTemplates{}
	if err := c.SearchRead(AccountGroupTemplateModel, criteria, NewOptions().Limit(1), agts); err != nil {
		return nil, err
	}
	if agts != nil && len(*agts) > 0 {
		return &((*agts)[0]), nil
	}
	return nil, fmt.Errorf("account.group.template was not found")
}

// FindAccountGroupTemplates finds account.group.template records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountGroupTemplates(criteria *Criteria, options *Options) (*AccountGroupTemplates, error) {
	agts := &AccountGroupTemplates{}
	if err := c.SearchRead(AccountGroupTemplateModel, criteria, options, agts); err != nil {
		return nil, err
	}
	return agts, nil
}

// FindAccountGroupTemplateIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountGroupTemplateIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountGroupTemplateModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountGroupTemplateId finds record id by querying it with criteria.
func (c *Client) FindAccountGroupTemplateId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountGroupTemplateModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.group.template was not found")
}
