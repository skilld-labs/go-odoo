package odoo

import (
	"fmt"
)

// AccountAccountTemplate represents account.account.template model.
type AccountAccountTemplate struct {
	LastUpdate      *Time     `xmlrpc:"__last_update,omitempty"`
	ChartTemplateId *Many2One `xmlrpc:"chart_template_id,omitempty"`
	Code            *String   `xmlrpc:"code,omitempty"`
	CreateDate      *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid       *Many2One `xmlrpc:"create_uid,omitempty"`
	CurrencyId      *Many2One `xmlrpc:"currency_id,omitempty"`
	DisplayName     *String   `xmlrpc:"display_name,omitempty"`
	GroupId         *Many2One `xmlrpc:"group_id,omitempty"`
	Id              *Int      `xmlrpc:"id,omitempty"`
	Name            *String   `xmlrpc:"name,omitempty"`
	Nocreate        *Bool     `xmlrpc:"nocreate,omitempty"`
	Note            *String   `xmlrpc:"note,omitempty"`
	Reconcile       *Bool     `xmlrpc:"reconcile,omitempty"`
	TagIds          *Relation `xmlrpc:"tag_ids,omitempty"`
	TaxIds          *Relation `xmlrpc:"tax_ids,omitempty"`
	UserTypeId      *Many2One `xmlrpc:"user_type_id,omitempty"`
	WriteDate       *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid        *Many2One `xmlrpc:"write_uid,omitempty"`
}

// AccountAccountTemplates represents array of account.account.template model.
type AccountAccountTemplates []AccountAccountTemplate

// AccountAccountTemplateModel is the odoo model name.
const AccountAccountTemplateModel = "account.account.template"

// Many2One convert AccountAccountTemplate to *Many2One.
func (aat *AccountAccountTemplate) Many2One() *Many2One {
	return NewMany2One(aat.Id.Get(), "")
}

// CreateAccountAccountTemplate creates a new account.account.template model and returns its id.
func (c *Client) CreateAccountAccountTemplate(aat *AccountAccountTemplate) (int64, error) {
	return c.Create(AccountAccountTemplateModel, aat)
}

// UpdateAccountAccountTemplate updates an existing account.account.template record.
func (c *Client) UpdateAccountAccountTemplate(aat *AccountAccountTemplate) error {
	return c.UpdateAccountAccountTemplates([]int64{aat.Id.Get()}, aat)
}

// UpdateAccountAccountTemplates updates existing account.account.template records.
// All records (represented by ids) will be updated by aat values.
func (c *Client) UpdateAccountAccountTemplates(ids []int64, aat *AccountAccountTemplate) error {
	return c.Update(AccountAccountTemplateModel, ids, aat)
}

// DeleteAccountAccountTemplate deletes an existing account.account.template record.
func (c *Client) DeleteAccountAccountTemplate(id int64) error {
	return c.DeleteAccountAccountTemplates([]int64{id})
}

// DeleteAccountAccountTemplates deletes existing account.account.template records.
func (c *Client) DeleteAccountAccountTemplates(ids []int64) error {
	return c.Delete(AccountAccountTemplateModel, ids)
}

// GetAccountAccountTemplate gets account.account.template existing record.
func (c *Client) GetAccountAccountTemplate(id int64) (*AccountAccountTemplate, error) {
	aats, err := c.GetAccountAccountTemplates([]int64{id})
	if err != nil {
		return nil, err
	}
	if aats != nil && len(*aats) > 0 {
		return &((*aats)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.account.template not found", id)
}

// GetAccountAccountTemplates gets account.account.template existing records.
func (c *Client) GetAccountAccountTemplates(ids []int64) (*AccountAccountTemplates, error) {
	aats := &AccountAccountTemplates{}
	if err := c.Read(AccountAccountTemplateModel, ids, nil, aats); err != nil {
		return nil, err
	}
	return aats, nil
}

// FindAccountAccountTemplate finds account.account.template record by querying it with criteria.
func (c *Client) FindAccountAccountTemplate(criteria *Criteria) (*AccountAccountTemplate, error) {
	aats := &AccountAccountTemplates{}
	if err := c.SearchRead(AccountAccountTemplateModel, criteria, NewOptions().Limit(1), aats); err != nil {
		return nil, err
	}
	if aats != nil && len(*aats) > 0 {
		return &((*aats)[0]), nil
	}
	return nil, fmt.Errorf("no account.account.template was found with criteria %v", criteria)
}

// FindAccountAccountTemplates finds account.account.template records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAccountTemplates(criteria *Criteria, options *Options) (*AccountAccountTemplates, error) {
	aats := &AccountAccountTemplates{}
	if err := c.SearchRead(AccountAccountTemplateModel, criteria, options, aats); err != nil {
		return nil, err
	}
	return aats, nil
}

// FindAccountAccountTemplateIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAccountTemplateIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountAccountTemplateModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountAccountTemplateId finds record id by querying it with criteria.
func (c *Client) FindAccountAccountTemplateId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountAccountTemplateModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no account.account.template was found with criteria %v and options %v", criteria, options)
}
