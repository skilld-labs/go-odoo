package odoo

import (
	"fmt"
)

// AccountResequenceWizard represents account.resequence.wizard model.
type AccountResequenceWizard struct {
	LastUpdate          *Time      `xmlrpc:"__last_update,omitempty"`
	CreateDate          *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid           *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName         *String    `xmlrpc:"display_name,omitempty"`
	EndDate             *Time      `xmlrpc:"end_date,omitempty"`
	FirstDate           *Time      `xmlrpc:"first_date,omitempty"`
	FirstName           *String    `xmlrpc:"first_name,omitempty"`
	Id                  *Int       `xmlrpc:"id,omitempty"`
	MoveIds             *Relation  `xmlrpc:"move_ids,omitempty"`
	NewValues           *String    `xmlrpc:"new_values,omitempty"`
	Ordering            *Selection `xmlrpc:"ordering,omitempty"`
	PreviewMoves        *String    `xmlrpc:"preview_moves,omitempty"`
	SequenceNumberReset *String    `xmlrpc:"sequence_number_reset,omitempty"`
	WriteDate           *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid            *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// AccountResequenceWizards represents array of account.resequence.wizard model.
type AccountResequenceWizards []AccountResequenceWizard

// AccountResequenceWizardModel is the odoo model name.
const AccountResequenceWizardModel = "account.resequence.wizard"

// Many2One convert AccountResequenceWizard to *Many2One.
func (arw *AccountResequenceWizard) Many2One() *Many2One {
	return NewMany2One(arw.Id.Get(), "")
}

// CreateAccountResequenceWizard creates a new account.resequence.wizard model and returns its id.
func (c *Client) CreateAccountResequenceWizard(arw *AccountResequenceWizard) (int64, error) {
	return c.Create(AccountResequenceWizardModel, arw)
}

// UpdateAccountResequenceWizard updates an existing account.resequence.wizard record.
func (c *Client) UpdateAccountResequenceWizard(arw *AccountResequenceWizard) error {
	return c.UpdateAccountResequenceWizards([]int64{arw.Id.Get()}, arw)
}

// UpdateAccountResequenceWizards updates existing account.resequence.wizard records.
// All records (represented by ids) will be updated by arw values.
func (c *Client) UpdateAccountResequenceWizards(ids []int64, arw *AccountResequenceWizard) error {
	return c.Update(AccountResequenceWizardModel, ids, arw)
}

// DeleteAccountResequenceWizard deletes an existing account.resequence.wizard record.
func (c *Client) DeleteAccountResequenceWizard(id int64) error {
	return c.DeleteAccountResequenceWizards([]int64{id})
}

// DeleteAccountResequenceWizards deletes existing account.resequence.wizard records.
func (c *Client) DeleteAccountResequenceWizards(ids []int64) error {
	return c.Delete(AccountResequenceWizardModel, ids)
}

// GetAccountResequenceWizard gets account.resequence.wizard existing record.
func (c *Client) GetAccountResequenceWizard(id int64) (*AccountResequenceWizard, error) {
	arws, err := c.GetAccountResequenceWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	if arws != nil && len(*arws) > 0 {
		return &((*arws)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.resequence.wizard not found", id)
}

// GetAccountResequenceWizards gets account.resequence.wizard existing records.
func (c *Client) GetAccountResequenceWizards(ids []int64) (*AccountResequenceWizards, error) {
	arws := &AccountResequenceWizards{}
	if err := c.Read(AccountResequenceWizardModel, ids, nil, arws); err != nil {
		return nil, err
	}
	return arws, nil
}

// FindAccountResequenceWizard finds account.resequence.wizard record by querying it with criteria.
func (c *Client) FindAccountResequenceWizard(criteria *Criteria) (*AccountResequenceWizard, error) {
	arws := &AccountResequenceWizards{}
	if err := c.SearchRead(AccountResequenceWizardModel, criteria, NewOptions().Limit(1), arws); err != nil {
		return nil, err
	}
	if arws != nil && len(*arws) > 0 {
		return &((*arws)[0]), nil
	}
	return nil, fmt.Errorf("account.resequence.wizard was not found")
}

// FindAccountResequenceWizards finds account.resequence.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountResequenceWizards(criteria *Criteria, options *Options) (*AccountResequenceWizards, error) {
	arws := &AccountResequenceWizards{}
	if err := c.SearchRead(AccountResequenceWizardModel, criteria, options, arws); err != nil {
		return nil, err
	}
	return arws, nil
}

// FindAccountResequenceWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountResequenceWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountResequenceWizardModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountResequenceWizardId finds record id by querying it with criteria.
func (c *Client) FindAccountResequenceWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountResequenceWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.resequence.wizard was not found")
}
