package odoo

import (
	"fmt"
)

// AccountFiscalPositionAccountTemplate represents account.fiscal.position.account.template model.
type AccountFiscalPositionAccountTemplate struct {
	LastUpdate    *Time     `xmlrpc:"__last_update,omitempty"`
	AccountDestId *Many2One `xmlrpc:"account_dest_id,omitempty"`
	AccountSrcId  *Many2One `xmlrpc:"account_src_id,omitempty"`
	CreateDate    *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid     *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName   *String   `xmlrpc:"display_name,omitempty"`
	Id            *Int      `xmlrpc:"id,omitempty"`
	PositionId    *Many2One `xmlrpc:"position_id,omitempty"`
	WriteDate     *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid      *Many2One `xmlrpc:"write_uid,omitempty"`
}

// AccountFiscalPositionAccountTemplates represents array of account.fiscal.position.account.template model.
type AccountFiscalPositionAccountTemplates []AccountFiscalPositionAccountTemplate

// AccountFiscalPositionAccountTemplateModel is the odoo model name.
const AccountFiscalPositionAccountTemplateModel = "account.fiscal.position.account.template"

// Many2One convert AccountFiscalPositionAccountTemplate to *Many2One.
func (afpat *AccountFiscalPositionAccountTemplate) Many2One() *Many2One {
	return NewMany2One(afpat.Id.Get(), "")
}

// CreateAccountFiscalPositionAccountTemplate creates a new account.fiscal.position.account.template model and returns its id.
func (c *Client) CreateAccountFiscalPositionAccountTemplate(afpat *AccountFiscalPositionAccountTemplate) (int64, error) {
	return c.Create(AccountFiscalPositionAccountTemplateModel, afpat)
}

// UpdateAccountFiscalPositionAccountTemplate updates an existing account.fiscal.position.account.template record.
func (c *Client) UpdateAccountFiscalPositionAccountTemplate(afpat *AccountFiscalPositionAccountTemplate) error {
	return c.UpdateAccountFiscalPositionAccountTemplates([]int64{afpat.Id.Get()}, afpat)
}

// UpdateAccountFiscalPositionAccountTemplates updates existing account.fiscal.position.account.template records.
// All records (represented by IDs) will be updated by afpat values.
func (c *Client) UpdateAccountFiscalPositionAccountTemplates(ids []int64, afpat *AccountFiscalPositionAccountTemplate) error {
	return c.Update(AccountFiscalPositionAccountTemplateModel, ids, afpat)
}

// DeleteAccountFiscalPositionAccountTemplate deletes an existing account.fiscal.position.account.template record.
func (c *Client) DeleteAccountFiscalPositionAccountTemplate(id int64) error {
	return c.DeleteAccountFiscalPositionAccountTemplates([]int64{id})
}

// DeleteAccountFiscalPositionAccountTemplates deletes existing account.fiscal.position.account.template records.
func (c *Client) DeleteAccountFiscalPositionAccountTemplates(ids []int64) error {
	return c.Delete(AccountFiscalPositionAccountTemplateModel, ids)
}

// GetAccountFiscalPositionAccountTemplate gets account.fiscal.position.account.template existing record.
func (c *Client) GetAccountFiscalPositionAccountTemplate(id int64) (*AccountFiscalPositionAccountTemplate, error) {
	afpats, err := c.GetAccountFiscalPositionAccountTemplates([]int64{id})
	if err != nil {
		return nil, err
	}
	if afpats != nil && len(*afpats) > 0 {
		return &((*afpats)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.fiscal.position.account.template not found", id)
}

// GetAccountFiscalPositionAccountTemplates gets account.fiscal.position.account.template existing records.
func (c *Client) GetAccountFiscalPositionAccountTemplates(ids []int64) (*AccountFiscalPositionAccountTemplates, error) {
	afpats := &AccountFiscalPositionAccountTemplates{}
	if err := c.Read(AccountFiscalPositionAccountTemplateModel, ids, nil, afpats); err != nil {
		return nil, err
	}
	return afpats, nil
}

// FindAccountFiscalPositionAccountTemplate finds account.fiscal.position.account.template record by querying it with criteria.
func (c *Client) FindAccountFiscalPositionAccountTemplate(criteria *Criteria) (*AccountFiscalPositionAccountTemplate, error) {
	afpats := &AccountFiscalPositionAccountTemplates{}
	if err := c.SearchRead(AccountFiscalPositionAccountTemplateModel, criteria, NewOptions().Limit(1), afpats); err != nil {
		return nil, err
	}
	if afpats != nil && len(*afpats) > 0 {
		return &((*afpats)[0]), nil
	}
	return nil, fmt.Errorf("account.fiscal.position.account.template was not found")
}

// FindAccountFiscalPositionAccountTemplates finds account.fiscal.position.account.template records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountFiscalPositionAccountTemplates(criteria *Criteria, options *Options) (*AccountFiscalPositionAccountTemplates, error) {
	afpats := &AccountFiscalPositionAccountTemplates{}
	if err := c.SearchRead(AccountFiscalPositionAccountTemplateModel, criteria, options, afpats); err != nil {
		return nil, err
	}
	return afpats, nil
}

// FindAccountFiscalPositionAccountTemplateIds finds records IDs by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountFiscalPositionAccountTemplateIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountFiscalPositionAccountTemplateModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountFiscalPositionAccountTemplateId finds record id by querying it with criteria.
func (c *Client) FindAccountFiscalPositionAccountTemplateId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountFiscalPositionAccountTemplateModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.fiscal.position.account.template was not found")
}
