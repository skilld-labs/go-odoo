package odoo

import (
	"fmt"
)

// AccountJournalGroup represents account.journal.group model.
type AccountJournalGroup struct {
	LastUpdate         *Time     `xmlrpc:"__last_update,omitempty"`
	CompanyId          *Many2One `xmlrpc:"company_id,omitempty"`
	CreateDate         *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid          *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName        *String   `xmlrpc:"display_name,omitempty"`
	ExcludedJournalIds *Relation `xmlrpc:"excluded_journal_ids,omitempty"`
	Id                 *Int      `xmlrpc:"id,omitempty"`
	Name               *String   `xmlrpc:"name,omitempty"`
	Sequence           *Int      `xmlrpc:"sequence,omitempty"`
	WriteDate          *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid           *Many2One `xmlrpc:"write_uid,omitempty"`
}

// AccountJournalGroups represents array of account.journal.group model.
type AccountJournalGroups []AccountJournalGroup

// AccountJournalGroupModel is the odoo model name.
const AccountJournalGroupModel = "account.journal.group"

// Many2One convert AccountJournalGroup to *Many2One.
func (ajg *AccountJournalGroup) Many2One() *Many2One {
	return NewMany2One(ajg.Id.Get(), "")
}

// CreateAccountJournalGroup creates a new account.journal.group model and returns its id.
func (c *Client) CreateAccountJournalGroup(ajg *AccountJournalGroup) (int64, error) {
	return c.Create(AccountJournalGroupModel, ajg)
}

// UpdateAccountJournalGroup updates an existing account.journal.group record.
func (c *Client) UpdateAccountJournalGroup(ajg *AccountJournalGroup) error {
	return c.UpdateAccountJournalGroups([]int64{ajg.Id.Get()}, ajg)
}

// UpdateAccountJournalGroups updates existing account.journal.group records.
// All records (represented by IDs) will be updated by ajg values.
func (c *Client) UpdateAccountJournalGroups(ids []int64, ajg *AccountJournalGroup) error {
	return c.Update(AccountJournalGroupModel, ids, ajg)
}

// DeleteAccountJournalGroup deletes an existing account.journal.group record.
func (c *Client) DeleteAccountJournalGroup(id int64) error {
	return c.DeleteAccountJournalGroups([]int64{id})
}

// DeleteAccountJournalGroups deletes existing account.journal.group records.
func (c *Client) DeleteAccountJournalGroups(ids []int64) error {
	return c.Delete(AccountJournalGroupModel, ids)
}

// GetAccountJournalGroup gets account.journal.group existing record.
func (c *Client) GetAccountJournalGroup(id int64) (*AccountJournalGroup, error) {
	ajgs, err := c.GetAccountJournalGroups([]int64{id})
	if err != nil {
		return nil, err
	}
	if ajgs != nil && len(*ajgs) > 0 {
		return &((*ajgs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.journal.group not found", id)
}

// GetAccountJournalGroups gets account.journal.group existing records.
func (c *Client) GetAccountJournalGroups(ids []int64) (*AccountJournalGroups, error) {
	ajgs := &AccountJournalGroups{}
	if err := c.Read(AccountJournalGroupModel, ids, nil, ajgs); err != nil {
		return nil, err
	}
	return ajgs, nil
}

// FindAccountJournalGroup finds account.journal.group record by querying it with criteria.
func (c *Client) FindAccountJournalGroup(criteria *Criteria) (*AccountJournalGroup, error) {
	ajgs := &AccountJournalGroups{}
	if err := c.SearchRead(AccountJournalGroupModel, criteria, NewOptions().Limit(1), ajgs); err != nil {
		return nil, err
	}
	if ajgs != nil && len(*ajgs) > 0 {
		return &((*ajgs)[0]), nil
	}
	return nil, fmt.Errorf("account.journal.group was not found")
}

// FindAccountJournalGroups finds account.journal.group records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountJournalGroups(criteria *Criteria, options *Options) (*AccountJournalGroups, error) {
	ajgs := &AccountJournalGroups{}
	if err := c.SearchRead(AccountJournalGroupModel, criteria, options, ajgs); err != nil {
		return nil, err
	}
	return ajgs, nil
}

// FindAccountJournalGroupIds finds records IDs by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountJournalGroupIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountJournalGroupModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountJournalGroupId finds record id by querying it with criteria.
func (c *Client) FindAccountJournalGroupId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountJournalGroupModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.journal.group was not found")
}
