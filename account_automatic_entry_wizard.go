package odoo

import (
	"fmt"
)

// AccountAutomaticEntryWizard represents account.automatic.entry.wizard model.
type AccountAutomaticEntryWizard struct {
	LastUpdate            *Time      `xmlrpc:"__last_update,omitempty"`
	AccountType           *Selection `xmlrpc:"account_type,omitempty"`
	Action                *Selection `xmlrpc:"action,omitempty"`
	CompanyCurrencyId     *Many2One  `xmlrpc:"company_currency_id,omitempty"`
	CompanyId             *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate            *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid             *Many2One  `xmlrpc:"create_uid,omitempty"`
	Date                  *Time      `xmlrpc:"date,omitempty"`
	DestinationAccountId  *Many2One  `xmlrpc:"destination_account_id,omitempty"`
	DisplayCurrencyHelper *Bool      `xmlrpc:"display_currency_helper,omitempty"`
	DisplayName           *String    `xmlrpc:"display_name,omitempty"`
	ExpenseAccrualAccount *Many2One  `xmlrpc:"expense_accrual_account,omitempty"`
	Id                    *Int       `xmlrpc:"id,omitempty"`
	JournalId             *Many2One  `xmlrpc:"journal_id,omitempty"`
	MoveData              *String    `xmlrpc:"move_data,omitempty"`
	MoveLineIds           *Relation  `xmlrpc:"move_line_ids,omitempty"`
	Percentage            *Float     `xmlrpc:"percentage,omitempty"`
	PreviewMoveData       *String    `xmlrpc:"preview_move_data,omitempty"`
	RevenueAccrualAccount *Many2One  `xmlrpc:"revenue_accrual_account,omitempty"`
	TotalAmount           *Float     `xmlrpc:"total_amount,omitempty"`
	WriteDate             *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid              *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// AccountAutomaticEntryWizards represents array of account.automatic.entry.wizard model.
type AccountAutomaticEntryWizards []AccountAutomaticEntryWizard

// AccountAutomaticEntryWizardModel is the odoo model name.
const AccountAutomaticEntryWizardModel = "account.automatic.entry.wizard"

// Many2One convert AccountAutomaticEntryWizard to *Many2One.
func (aaew *AccountAutomaticEntryWizard) Many2One() *Many2One {
	return NewMany2One(aaew.Id.Get(), "")
}

// CreateAccountAutomaticEntryWizard creates a new account.automatic.entry.wizard model and returns its id.
func (c *Client) CreateAccountAutomaticEntryWizard(aaew *AccountAutomaticEntryWizard) (int64, error) {
	return c.Create(AccountAutomaticEntryWizardModel, aaew)
}

// UpdateAccountAutomaticEntryWizard updates an existing account.automatic.entry.wizard record.
func (c *Client) UpdateAccountAutomaticEntryWizard(aaew *AccountAutomaticEntryWizard) error {
	return c.UpdateAccountAutomaticEntryWizards([]int64{aaew.Id.Get()}, aaew)
}

// UpdateAccountAutomaticEntryWizards updates existing account.automatic.entry.wizard records.
// All records (represented by ids) will be updated by aaew values.
func (c *Client) UpdateAccountAutomaticEntryWizards(ids []int64, aaew *AccountAutomaticEntryWizard) error {
	return c.Update(AccountAutomaticEntryWizardModel, ids, aaew)
}

// DeleteAccountAutomaticEntryWizard deletes an existing account.automatic.entry.wizard record.
func (c *Client) DeleteAccountAutomaticEntryWizard(id int64) error {
	return c.DeleteAccountAutomaticEntryWizards([]int64{id})
}

// DeleteAccountAutomaticEntryWizards deletes existing account.automatic.entry.wizard records.
func (c *Client) DeleteAccountAutomaticEntryWizards(ids []int64) error {
	return c.Delete(AccountAutomaticEntryWizardModel, ids)
}

// GetAccountAutomaticEntryWizard gets account.automatic.entry.wizard existing record.
func (c *Client) GetAccountAutomaticEntryWizard(id int64) (*AccountAutomaticEntryWizard, error) {
	aaews, err := c.GetAccountAutomaticEntryWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	if aaews != nil && len(*aaews) > 0 {
		return &((*aaews)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.automatic.entry.wizard not found", id)
}

// GetAccountAutomaticEntryWizards gets account.automatic.entry.wizard existing records.
func (c *Client) GetAccountAutomaticEntryWizards(ids []int64) (*AccountAutomaticEntryWizards, error) {
	aaews := &AccountAutomaticEntryWizards{}
	if err := c.Read(AccountAutomaticEntryWizardModel, ids, nil, aaews); err != nil {
		return nil, err
	}
	return aaews, nil
}

// FindAccountAutomaticEntryWizard finds account.automatic.entry.wizard record by querying it with criteria.
func (c *Client) FindAccountAutomaticEntryWizard(criteria *Criteria) (*AccountAutomaticEntryWizard, error) {
	aaews := &AccountAutomaticEntryWizards{}
	if err := c.SearchRead(AccountAutomaticEntryWizardModel, criteria, NewOptions().Limit(1), aaews); err != nil {
		return nil, err
	}
	if aaews != nil && len(*aaews) > 0 {
		return &((*aaews)[0]), nil
	}
	return nil, fmt.Errorf("account.automatic.entry.wizard was not found")
}

// FindAccountAutomaticEntryWizards finds account.automatic.entry.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAutomaticEntryWizards(criteria *Criteria, options *Options) (*AccountAutomaticEntryWizards, error) {
	aaews := &AccountAutomaticEntryWizards{}
	if err := c.SearchRead(AccountAutomaticEntryWizardModel, criteria, options, aaews); err != nil {
		return nil, err
	}
	return aaews, nil
}

// FindAccountAutomaticEntryWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAutomaticEntryWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountAutomaticEntryWizardModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountAutomaticEntryWizardId finds record id by querying it with criteria.
func (c *Client) FindAccountAutomaticEntryWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountAutomaticEntryWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.automatic.entry.wizard was not found")
}
