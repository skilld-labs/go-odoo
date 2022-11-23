package odoo

import (
	"fmt"
)

// AccountReconcileModelPartnerMapping represents account.reconcile.model.partner.mapping model.
type AccountReconcileModelPartnerMapping struct {
	LastUpdate      *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate      *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid       *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName     *String   `xmlrpc:"display_name,omitempty"`
	Id              *Int      `xmlrpc:"id,omitempty"`
	ModelId         *Many2One `xmlrpc:"model_id,omitempty"`
	NarrationRegex  *String   `xmlrpc:"narration_regex,omitempty"`
	PartnerId       *Many2One `xmlrpc:"partner_id,omitempty"`
	PaymentRefRegex *String   `xmlrpc:"payment_ref_regex,omitempty"`
	WriteDate       *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid        *Many2One `xmlrpc:"write_uid,omitempty"`
}

// AccountReconcileModelPartnerMappings represents array of account.reconcile.model.partner.mapping model.
type AccountReconcileModelPartnerMappings []AccountReconcileModelPartnerMapping

// AccountReconcileModelPartnerMappingModel is the odoo model name.
const AccountReconcileModelPartnerMappingModel = "account.reconcile.model.partner.mapping"

// Many2One convert AccountReconcileModelPartnerMapping to *Many2One.
func (armpm *AccountReconcileModelPartnerMapping) Many2One() *Many2One {
	return NewMany2One(armpm.Id.Get(), "")
}

// CreateAccountReconcileModelPartnerMapping creates a new account.reconcile.model.partner.mapping model and returns its id.
func (c *Client) CreateAccountReconcileModelPartnerMapping(armpm *AccountReconcileModelPartnerMapping) (int64, error) {
	return c.Create(AccountReconcileModelPartnerMappingModel, armpm)
}

// UpdateAccountReconcileModelPartnerMapping updates an existing account.reconcile.model.partner.mapping record.
func (c *Client) UpdateAccountReconcileModelPartnerMapping(armpm *AccountReconcileModelPartnerMapping) error {
	return c.UpdateAccountReconcileModelPartnerMappings([]int64{armpm.Id.Get()}, armpm)
}

// UpdateAccountReconcileModelPartnerMappings updates existing account.reconcile.model.partner.mapping records.
// All records (represented by IDs) will be updated by armpm values.
func (c *Client) UpdateAccountReconcileModelPartnerMappings(ids []int64, armpm *AccountReconcileModelPartnerMapping) error {
	return c.Update(AccountReconcileModelPartnerMappingModel, ids, armpm)
}

// DeleteAccountReconcileModelPartnerMapping deletes an existing account.reconcile.model.partner.mapping record.
func (c *Client) DeleteAccountReconcileModelPartnerMapping(id int64) error {
	return c.DeleteAccountReconcileModelPartnerMappings([]int64{id})
}

// DeleteAccountReconcileModelPartnerMappings deletes existing account.reconcile.model.partner.mapping records.
func (c *Client) DeleteAccountReconcileModelPartnerMappings(ids []int64) error {
	return c.Delete(AccountReconcileModelPartnerMappingModel, ids)
}

// GetAccountReconcileModelPartnerMapping gets account.reconcile.model.partner.mapping existing record.
func (c *Client) GetAccountReconcileModelPartnerMapping(id int64) (*AccountReconcileModelPartnerMapping, error) {
	armpms, err := c.GetAccountReconcileModelPartnerMappings([]int64{id})
	if err != nil {
		return nil, err
	}
	if armpms != nil && len(*armpms) > 0 {
		return &((*armpms)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.reconcile.model.partner.mapping not found", id)
}

// GetAccountReconcileModelPartnerMappings gets account.reconcile.model.partner.mapping existing records.
func (c *Client) GetAccountReconcileModelPartnerMappings(ids []int64) (*AccountReconcileModelPartnerMappings, error) {
	armpms := &AccountReconcileModelPartnerMappings{}
	if err := c.Read(AccountReconcileModelPartnerMappingModel, ids, nil, armpms); err != nil {
		return nil, err
	}
	return armpms, nil
}

// FindAccountReconcileModelPartnerMapping finds account.reconcile.model.partner.mapping record by querying it with criteria.
func (c *Client) FindAccountReconcileModelPartnerMapping(criteria *Criteria) (*AccountReconcileModelPartnerMapping, error) {
	armpms := &AccountReconcileModelPartnerMappings{}
	if err := c.SearchRead(AccountReconcileModelPartnerMappingModel, criteria, NewOptions().Limit(1), armpms); err != nil {
		return nil, err
	}
	if armpms != nil && len(*armpms) > 0 {
		return &((*armpms)[0]), nil
	}
	return nil, fmt.Errorf("account.reconcile.model.partner.mapping was not found")
}

// FindAccountReconcileModelPartnerMappings finds account.reconcile.model.partner.mapping records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReconcileModelPartnerMappings(criteria *Criteria, options *Options) (*AccountReconcileModelPartnerMappings, error) {
	armpms := &AccountReconcileModelPartnerMappings{}
	if err := c.SearchRead(AccountReconcileModelPartnerMappingModel, criteria, options, armpms); err != nil {
		return nil, err
	}
	return armpms, nil
}

// FindAccountReconcileModelPartnerMappingIds finds records IDs by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReconcileModelPartnerMappingIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountReconcileModelPartnerMappingModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountReconcileModelPartnerMappingId finds record id by querying it with criteria.
func (c *Client) FindAccountReconcileModelPartnerMappingId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountReconcileModelPartnerMappingModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.reconcile.model.partner.mapping was not found")
}
