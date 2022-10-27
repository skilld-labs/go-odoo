package odoo

import (
	"fmt"
)

// AccountAnalyticDistribution represents account.analytic.distribution model.
type AccountAnalyticDistribution struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	AccountId   *Many2One `xmlrpc:"account_id,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	Percentage  *Float    `xmlrpc:"percentage,omitempty"`
	TagId       *Many2One `xmlrpc:"tag_id,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// AccountAnalyticDistributions represents array of account.analytic.distribution model.
type AccountAnalyticDistributions []AccountAnalyticDistribution

// AccountAnalyticDistributionModel is the odoo model name.
const AccountAnalyticDistributionModel = "account.analytic.distribution"

// Many2One convert AccountAnalyticDistribution to *Many2One.
func (aad *AccountAnalyticDistribution) Many2One() *Many2One {
	return NewMany2One(aad.Id.Get(), "")
}

// CreateAccountAnalyticDistribution creates a new account.analytic.distribution model and returns its id.
func (c *Client) CreateAccountAnalyticDistribution(aad *AccountAnalyticDistribution) (int64, error) {
	return c.Create(AccountAnalyticDistributionModel, aad)
}

// UpdateAccountAnalyticDistribution updates an existing account.analytic.distribution record.
func (c *Client) UpdateAccountAnalyticDistribution(aad *AccountAnalyticDistribution) error {
	return c.UpdateAccountAnalyticDistributions([]int64{aad.Id.Get()}, aad)
}

// UpdateAccountAnalyticDistributions updates existing account.analytic.distribution records.
// All records (represented by ids) will be updated by aad values.
func (c *Client) UpdateAccountAnalyticDistributions(ids []int64, aad *AccountAnalyticDistribution) error {
	return c.Update(AccountAnalyticDistributionModel, ids, aad)
}

// DeleteAccountAnalyticDistribution deletes an existing account.analytic.distribution record.
func (c *Client) DeleteAccountAnalyticDistribution(id int64) error {
	return c.DeleteAccountAnalyticDistributions([]int64{id})
}

// DeleteAccountAnalyticDistributions deletes existing account.analytic.distribution records.
func (c *Client) DeleteAccountAnalyticDistributions(ids []int64) error {
	return c.Delete(AccountAnalyticDistributionModel, ids)
}

// GetAccountAnalyticDistribution gets account.analytic.distribution existing record.
func (c *Client) GetAccountAnalyticDistribution(id int64) (*AccountAnalyticDistribution, error) {
	aads, err := c.GetAccountAnalyticDistributions([]int64{id})
	if err != nil {
		return nil, err
	}
	if aads != nil && len(*aads) > 0 {
		return &((*aads)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.analytic.distribution not found", id)
}

// GetAccountAnalyticDistributions gets account.analytic.distribution existing records.
func (c *Client) GetAccountAnalyticDistributions(ids []int64) (*AccountAnalyticDistributions, error) {
	aads := &AccountAnalyticDistributions{}
	if err := c.Read(AccountAnalyticDistributionModel, ids, nil, aads); err != nil {
		return nil, err
	}
	return aads, nil
}

// FindAccountAnalyticDistribution finds account.analytic.distribution record by querying it with criteria.
func (c *Client) FindAccountAnalyticDistribution(criteria *Criteria) (*AccountAnalyticDistribution, error) {
	aads := &AccountAnalyticDistributions{}
	if err := c.SearchRead(AccountAnalyticDistributionModel, criteria, NewOptions().Limit(1), aads); err != nil {
		return nil, err
	}
	if aads != nil && len(*aads) > 0 {
		return &((*aads)[0]), nil
	}
	return nil, fmt.Errorf("account.analytic.distribution was not found")
}

// FindAccountAnalyticDistributions finds account.analytic.distribution records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAnalyticDistributions(criteria *Criteria, options *Options) (*AccountAnalyticDistributions, error) {
	aads := &AccountAnalyticDistributions{}
	if err := c.SearchRead(AccountAnalyticDistributionModel, criteria, options, aads); err != nil {
		return nil, err
	}
	return aads, nil
}

// FindAccountAnalyticDistributionIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAnalyticDistributionIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountAnalyticDistributionModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountAnalyticDistributionId finds record id by querying it with criteria.
func (c *Client) FindAccountAnalyticDistributionId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountAnalyticDistributionModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.analytic.distribution was not found")
}
