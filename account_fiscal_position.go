package odoo

import (
	"fmt"
)

// AccountFiscalPosition represents account.fiscal.position model.
type AccountFiscalPosition struct {
	LastUpdate     *Time     `xmlrpc:"__last_update,omitempty"`
	AccountIds     *Relation `xmlrpc:"account_ids,omitempty"`
	Active         *Bool     `xmlrpc:"active,omitempty"`
	AutoApply      *Bool     `xmlrpc:"auto_apply,omitempty"`
	CompanyId      *Many2One `xmlrpc:"company_id,omitempty"`
	CountryGroupId *Many2One `xmlrpc:"country_group_id,omitempty"`
	CountryId      *Many2One `xmlrpc:"country_id,omitempty"`
	CreateDate     *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid      *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName    *String   `xmlrpc:"display_name,omitempty"`
	Id             *Int      `xmlrpc:"id,omitempty"`
	Name           *String   `xmlrpc:"name,omitempty"`
	Note           *String   `xmlrpc:"note,omitempty"`
	Sequence       *Int      `xmlrpc:"sequence,omitempty"`
	StateIds       *Relation `xmlrpc:"state_ids,omitempty"`
	StatesCount    *Int      `xmlrpc:"states_count,omitempty"`
	TaxIds         *Relation `xmlrpc:"tax_ids,omitempty"`
	VatRequired    *Bool     `xmlrpc:"vat_required,omitempty"`
	WriteDate      *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid       *Many2One `xmlrpc:"write_uid,omitempty"`
	ZipFrom        *String   `xmlrpc:"zip_from,omitempty"`
	ZipTo          *String   `xmlrpc:"zip_to,omitempty"`
}

// AccountFiscalPositions represents array of account.fiscal.position model.
type AccountFiscalPositions []AccountFiscalPosition

// AccountFiscalPositionModel is the odoo model name.
const AccountFiscalPositionModel = "account.fiscal.position"

// Many2One convert AccountFiscalPosition to *Many2One.
func (afp *AccountFiscalPosition) Many2One() *Many2One {
	return NewMany2One(afp.Id.Get(), "")
}

// CreateAccountFiscalPosition creates a new account.fiscal.position model and returns its id.
func (c *Client) CreateAccountFiscalPosition(afp *AccountFiscalPosition) (int64, error) {
	return c.Create(AccountFiscalPositionModel, afp)
}

// UpdateAccountFiscalPosition updates an existing account.fiscal.position record.
func (c *Client) UpdateAccountFiscalPosition(afp *AccountFiscalPosition) error {
	return c.UpdateAccountFiscalPositions([]int64{afp.Id.Get()}, afp)
}

// UpdateAccountFiscalPositions updates existing account.fiscal.position records.
// All records (represented by IDs) will be updated by afp values.
func (c *Client) UpdateAccountFiscalPositions(ids []int64, afp *AccountFiscalPosition) error {
	return c.Update(AccountFiscalPositionModel, ids, afp)
}

// DeleteAccountFiscalPosition deletes an existing account.fiscal.position record.
func (c *Client) DeleteAccountFiscalPosition(id int64) error {
	return c.DeleteAccountFiscalPositions([]int64{id})
}

// DeleteAccountFiscalPositions deletes existing account.fiscal.position records.
func (c *Client) DeleteAccountFiscalPositions(ids []int64) error {
	return c.Delete(AccountFiscalPositionModel, ids)
}

// GetAccountFiscalPosition gets account.fiscal.position existing record.
func (c *Client) GetAccountFiscalPosition(id int64) (*AccountFiscalPosition, error) {
	afps, err := c.GetAccountFiscalPositions([]int64{id})
	if err != nil {
		return nil, err
	}
	if afps != nil && len(*afps) > 0 {
		return &((*afps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.fiscal.position not found", id)
}

// GetAccountFiscalPositions gets account.fiscal.position existing records.
func (c *Client) GetAccountFiscalPositions(ids []int64) (*AccountFiscalPositions, error) {
	afps := &AccountFiscalPositions{}
	if err := c.Read(AccountFiscalPositionModel, ids, nil, afps); err != nil {
		return nil, err
	}
	return afps, nil
}

// FindAccountFiscalPosition finds account.fiscal.position record by querying it with criteria.
func (c *Client) FindAccountFiscalPosition(criteria *Criteria) (*AccountFiscalPosition, error) {
	afps := &AccountFiscalPositions{}
	if err := c.SearchRead(AccountFiscalPositionModel, criteria, NewOptions().Limit(1), afps); err != nil {
		return nil, err
	}
	if afps != nil && len(*afps) > 0 {
		return &((*afps)[0]), nil
	}
	return nil, fmt.Errorf("account.fiscal.position was not found")
}

// FindAccountFiscalPositions finds account.fiscal.position records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountFiscalPositions(criteria *Criteria, options *Options) (*AccountFiscalPositions, error) {
	afps := &AccountFiscalPositions{}
	if err := c.SearchRead(AccountFiscalPositionModel, criteria, options, afps); err != nil {
		return nil, err
	}
	return afps, nil
}

// FindAccountFiscalPositionIds finds records IDs by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountFiscalPositionIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountFiscalPositionModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountFiscalPositionId finds record id by querying it with criteria.
func (c *Client) FindAccountFiscalPositionId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountFiscalPositionModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.fiscal.position was not found")
}
