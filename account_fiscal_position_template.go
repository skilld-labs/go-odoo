package odoo

import (
	"fmt"
)

// AccountFiscalPositionTemplate represents account.fiscal.position.template model.
type AccountFiscalPositionTemplate struct {
	LastUpdate      *Time     `xmlrpc:"__last_update,omptempty"`
	AccountIds      *Relation `xmlrpc:"account_ids,omptempty"`
	AutoApply       *Bool     `xmlrpc:"auto_apply,omptempty"`
	ChartTemplateId *Many2One `xmlrpc:"chart_template_id,omptempty"`
	CountryGroupId  *Many2One `xmlrpc:"country_group_id,omptempty"`
	CountryId       *Many2One `xmlrpc:"country_id,omptempty"`
	CreateDate      *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid       *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName     *String   `xmlrpc:"display_name,omptempty"`
	Id              *Int      `xmlrpc:"id,omptempty"`
	Name            *String   `xmlrpc:"name,omptempty"`
	Note            *String   `xmlrpc:"note,omptempty"`
	Sequence        *Int      `xmlrpc:"sequence,omptempty"`
	StateIds        *Relation `xmlrpc:"state_ids,omptempty"`
	TaxIds          *Relation `xmlrpc:"tax_ids,omptempty"`
	VatRequired     *Bool     `xmlrpc:"vat_required,omptempty"`
	WriteDate       *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid        *Many2One `xmlrpc:"write_uid,omptempty"`
	ZipFrom         *Int      `xmlrpc:"zip_from,omptempty"`
	ZipTo           *Int      `xmlrpc:"zip_to,omptempty"`
}

// AccountFiscalPositionTemplates represents array of account.fiscal.position.template model.
type AccountFiscalPositionTemplates []AccountFiscalPositionTemplate

// AccountFiscalPositionTemplateModel is the odoo model name.
const AccountFiscalPositionTemplateModel = "account.fiscal.position.template"

// Many2One convert AccountFiscalPositionTemplate to *Many2One.
func (afpt *AccountFiscalPositionTemplate) Many2One() *Many2One {
	return NewMany2One(afpt.Id.Get(), "")
}

// CreateAccountFiscalPositionTemplate creates a new account.fiscal.position.template model and returns its id.
func (c *Client) CreateAccountFiscalPositionTemplate(afpt *AccountFiscalPositionTemplate) (int64, error) {
	return c.Create(AccountFiscalPositionTemplateModel, afpt)
}

// UpdateAccountFiscalPositionTemplate updates an existing account.fiscal.position.template record.
func (c *Client) UpdateAccountFiscalPositionTemplate(afpt *AccountFiscalPositionTemplate) error {
	return c.UpdateAccountFiscalPositionTemplates([]int64{afpt.Id.Get()}, afpt)
}

// UpdateAccountFiscalPositionTemplates updates existing account.fiscal.position.template records.
// All records (represented by ids) will be updated by afpt values.
func (c *Client) UpdateAccountFiscalPositionTemplates(ids []int64, afpt *AccountFiscalPositionTemplate) error {
	return c.Update(AccountFiscalPositionTemplateModel, ids, afpt)
}

// DeleteAccountFiscalPositionTemplate deletes an existing account.fiscal.position.template record.
func (c *Client) DeleteAccountFiscalPositionTemplate(id int64) error {
	return c.DeleteAccountFiscalPositionTemplates([]int64{id})
}

// DeleteAccountFiscalPositionTemplates deletes existing account.fiscal.position.template records.
func (c *Client) DeleteAccountFiscalPositionTemplates(ids []int64) error {
	return c.Delete(AccountFiscalPositionTemplateModel, ids)
}

// GetAccountFiscalPositionTemplate gets account.fiscal.position.template existing record.
func (c *Client) GetAccountFiscalPositionTemplate(id int64) (*AccountFiscalPositionTemplate, error) {
	afpts, err := c.GetAccountFiscalPositionTemplates([]int64{id})
	if err != nil {
		return nil, err
	}
	if afpts != nil && len(*afpts) > 0 {
		return &((*afpts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.fiscal.position.template not found", id)
}

// GetAccountFiscalPositionTemplates gets account.fiscal.position.template existing records.
func (c *Client) GetAccountFiscalPositionTemplates(ids []int64) (*AccountFiscalPositionTemplates, error) {
	afpts := &AccountFiscalPositionTemplates{}
	if err := c.Read(AccountFiscalPositionTemplateModel, ids, nil, afpts); err != nil {
		return nil, err
	}
	return afpts, nil
}

// FindAccountFiscalPositionTemplate finds account.fiscal.position.template record by querying it with criteria.
func (c *Client) FindAccountFiscalPositionTemplate(criteria *Criteria) (*AccountFiscalPositionTemplate, error) {
	afpts := &AccountFiscalPositionTemplates{}
	if err := c.SearchRead(AccountFiscalPositionTemplateModel, criteria, NewOptions().Limit(1), afpts); err != nil {
		return nil, err
	}
	if afpts != nil && len(*afpts) > 0 {
		return &((*afpts)[0]), nil
	}
	return nil, fmt.Errorf("account.fiscal.position.template was not found with criteria %v", criteria)
}

// FindAccountFiscalPositionTemplates finds account.fiscal.position.template records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountFiscalPositionTemplates(criteria *Criteria, options *Options) (*AccountFiscalPositionTemplates, error) {
	afpts := &AccountFiscalPositionTemplates{}
	if err := c.SearchRead(AccountFiscalPositionTemplateModel, criteria, options, afpts); err != nil {
		return nil, err
	}
	return afpts, nil
}

// FindAccountFiscalPositionTemplateIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountFiscalPositionTemplateIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountFiscalPositionTemplateModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountFiscalPositionTemplateId finds record id by querying it with criteria.
func (c *Client) FindAccountFiscalPositionTemplateId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountFiscalPositionTemplateModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.fiscal.position.template was not found with criteria %v and options %v", criteria, options)
}
