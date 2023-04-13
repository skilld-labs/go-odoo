package odoo

import (
	"fmt"
)

// AccountFiscalPosition represents account.fiscal.position model.
type AccountFiscalPosition struct {
	LastUpdate     *Time     `xmlrpc:"__last_update,omptempty"`
	AccountIds     *Relation `xmlrpc:"account_ids,omptempty"`
	Active         *Bool     `xmlrpc:"active,omptempty"`
	AutoApply      *Bool     `xmlrpc:"auto_apply,omptempty"`
	CompanyId      *Many2One `xmlrpc:"company_id,omptempty"`
	CountryGroupId *Many2One `xmlrpc:"country_group_id,omptempty"`
	CountryId      *Many2One `xmlrpc:"country_id,omptempty"`
	CreateDate     *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid      *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName    *String   `xmlrpc:"display_name,omptempty"`
	Id             *Int      `xmlrpc:"id,omptempty"`
	Name           *String   `xmlrpc:"name,omptempty"`
	Note           *String   `xmlrpc:"note,omptempty"`
	Sequence       *Int      `xmlrpc:"sequence,omptempty"`
	StateIds       *Relation `xmlrpc:"state_ids,omptempty"`
	StatesCount    *Int      `xmlrpc:"states_count,omptempty"`
	TaxIds         *Relation `xmlrpc:"tax_ids,omptempty"`
	VatRequired    *Bool     `xmlrpc:"vat_required,omptempty"`
	WriteDate      *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid       *Many2One `xmlrpc:"write_uid,omptempty"`
	ZipFrom        *Int      `xmlrpc:"zip_from,omptempty"`
	ZipTo          *Int      `xmlrpc:"zip_to,omptempty"`
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
	ids, err := c.Create(AccountFiscalPositionModel, []interface{}{afp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountFiscalPosition creates a new account.fiscal.position model and returns its id.
func (c *Client) CreateAccountFiscalPositions(afps []*AccountFiscalPosition) ([]int64, error) {
	var vv []interface{}
	for _, v := range afps {
		vv = append(vv, v)
	}
	return c.Create(AccountFiscalPositionModel, vv)
}

// UpdateAccountFiscalPosition updates an existing account.fiscal.position record.
func (c *Client) UpdateAccountFiscalPosition(afp *AccountFiscalPosition) error {
	return c.UpdateAccountFiscalPositions([]int64{afp.Id.Get()}, afp)
}

// UpdateAccountFiscalPositions updates existing account.fiscal.position records.
// All records (represented by ids) will be updated by afp values.
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
	return nil, fmt.Errorf("account.fiscal.position was not found with criteria %v", criteria)
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

// FindAccountFiscalPositionIds finds records ids by querying it
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
	return -1, fmt.Errorf("account.fiscal.position was not found with criteria %v and options %v", criteria, options)
}
