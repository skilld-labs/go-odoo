package odoo

import (
	"fmt"
)

// AccountReconcileModelTemplate represents account.reconcile.model.template model.
type AccountReconcileModelTemplate struct {
	LastUpdate       *Time      `xmlrpc:"__last_update,omptempty"`
	AccountId        *Many2One  `xmlrpc:"account_id,omptempty"`
	Amount           *Float     `xmlrpc:"amount,omptempty"`
	AmountType       *Selection `xmlrpc:"amount_type,omptempty"`
	ChartTemplateId  *Many2One  `xmlrpc:"chart_template_id,omptempty"`
	CreateDate       *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid        *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName      *String    `xmlrpc:"display_name,omptempty"`
	HasSecondLine    *Bool      `xmlrpc:"has_second_line,omptempty"`
	Id               *Int       `xmlrpc:"id,omptempty"`
	Label            *String    `xmlrpc:"label,omptempty"`
	Name             *String    `xmlrpc:"name,omptempty"`
	SecondAccountId  *Many2One  `xmlrpc:"second_account_id,omptempty"`
	SecondAmount     *Float     `xmlrpc:"second_amount,omptempty"`
	SecondAmountType *Selection `xmlrpc:"second_amount_type,omptempty"`
	SecondLabel      *String    `xmlrpc:"second_label,omptempty"`
	SecondTaxId      *Many2One  `xmlrpc:"second_tax_id,omptempty"`
	Sequence         *Int       `xmlrpc:"sequence,omptempty"`
	TaxId            *Many2One  `xmlrpc:"tax_id,omptempty"`
	WriteDate        *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid         *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// AccountReconcileModelTemplates represents array of account.reconcile.model.template model.
type AccountReconcileModelTemplates []AccountReconcileModelTemplate

// AccountReconcileModelTemplateModel is the odoo model name.
const AccountReconcileModelTemplateModel = "account.reconcile.model.template"

// Many2One convert AccountReconcileModelTemplate to *Many2One.
func (armt *AccountReconcileModelTemplate) Many2One() *Many2One {
	return NewMany2One(armt.Id.Get(), "")
}

// CreateAccountReconcileModelTemplate creates a new account.reconcile.model.template model and returns its id.
func (c *Client) CreateAccountReconcileModelTemplate(armt *AccountReconcileModelTemplate) (int64, error) {
	return c.Create(AccountReconcileModelTemplateModel, armt)
}

// UpdateAccountReconcileModelTemplate updates an existing account.reconcile.model.template record.
func (c *Client) UpdateAccountReconcileModelTemplate(armt *AccountReconcileModelTemplate) error {
	return c.UpdateAccountReconcileModelTemplates([]int64{armt.Id.Get()}, armt)
}

// UpdateAccountReconcileModelTemplates updates existing account.reconcile.model.template records.
// All records (represented by ids) will be updated by armt values.
func (c *Client) UpdateAccountReconcileModelTemplates(ids []int64, armt *AccountReconcileModelTemplate) error {
	return c.Update(AccountReconcileModelTemplateModel, ids, armt)
}

// DeleteAccountReconcileModelTemplate deletes an existing account.reconcile.model.template record.
func (c *Client) DeleteAccountReconcileModelTemplate(id int64) error {
	return c.DeleteAccountReconcileModelTemplates([]int64{id})
}

// DeleteAccountReconcileModelTemplates deletes existing account.reconcile.model.template records.
func (c *Client) DeleteAccountReconcileModelTemplates(ids []int64) error {
	return c.Delete(AccountReconcileModelTemplateModel, ids)
}

// GetAccountReconcileModelTemplate gets account.reconcile.model.template existing record.
func (c *Client) GetAccountReconcileModelTemplate(id int64) (*AccountReconcileModelTemplate, error) {
	armts, err := c.GetAccountReconcileModelTemplates([]int64{id})
	if err != nil {
		return nil, err
	}
	if armts != nil && len(*armts) > 0 {
		return &((*armts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.reconcile.model.template not found", id)
}

// GetAccountReconcileModelTemplates gets account.reconcile.model.template existing records.
func (c *Client) GetAccountReconcileModelTemplates(ids []int64) (*AccountReconcileModelTemplates, error) {
	armts := &AccountReconcileModelTemplates{}
	if err := c.Read(AccountReconcileModelTemplateModel, ids, nil, armts); err != nil {
		return nil, err
	}
	return armts, nil
}

// FindAccountReconcileModelTemplate finds account.reconcile.model.template record by querying it with criteria.
func (c *Client) FindAccountReconcileModelTemplate(criteria *Criteria) (*AccountReconcileModelTemplate, error) {
	armts := &AccountReconcileModelTemplates{}
	if err := c.SearchRead(AccountReconcileModelTemplateModel, criteria, NewOptions().Limit(1), armts); err != nil {
		return nil, err
	}
	if armts != nil && len(*armts) > 0 {
		return &((*armts)[0]), nil
	}
	return nil, fmt.Errorf("no account.reconcile.model.template was found with criteria %v", criteria)
}

// FindAccountReconcileModelTemplates finds account.reconcile.model.template records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReconcileModelTemplates(criteria *Criteria, options *Options) (*AccountReconcileModelTemplates, error) {
	armts := &AccountReconcileModelTemplates{}
	if err := c.SearchRead(AccountReconcileModelTemplateModel, criteria, options, armts); err != nil {
		return nil, err
	}
	return armts, nil
}

// FindAccountReconcileModelTemplateIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReconcileModelTemplateIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountReconcileModelTemplateModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountReconcileModelTemplateId finds record id by querying it with criteria.
func (c *Client) FindAccountReconcileModelTemplateId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountReconcileModelTemplateModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no account.reconcile.model.template was found with criteria %v and options %v", criteria, options)
}
