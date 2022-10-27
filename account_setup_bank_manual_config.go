package odoo

import (
	"fmt"
)

// AccountSetupBankManualConfig represents account.setup.bank.manual.config model.
type AccountSetupBankManualConfig struct {
	LastUpdate                *Time      `xmlrpc:"__last_update,omitempty"`
	AbaRouting                *String    `xmlrpc:"aba_routing,omitempty"`
	AccHolderName             *String    `xmlrpc:"acc_holder_name,omitempty"`
	AccNumber                 *String    `xmlrpc:"acc_number,omitempty"`
	AccType                   *Selection `xmlrpc:"acc_type,omitempty"`
	Active                    *Bool      `xmlrpc:"active,omitempty"`
	BankBic                   *String    `xmlrpc:"bank_bic,omitempty"`
	BankId                    *Many2One  `xmlrpc:"bank_id,omitempty"`
	BankName                  *String    `xmlrpc:"bank_name,omitempty"`
	CompanyId                 *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate                *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                 *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrencyId                *Many2One  `xmlrpc:"currency_id,omitempty"`
	DisplayName               *String    `xmlrpc:"display_name,omitempty"`
	Id                        *Int       `xmlrpc:"id,omitempty"`
	JournalId                 *Relation  `xmlrpc:"journal_id,omitempty"`
	LinkedJournalId           *Many2One  `xmlrpc:"linked_journal_id,omitempty"`
	NewJournalName            *String    `xmlrpc:"new_journal_name,omitempty"`
	NumJournalsWithoutAccount *Int       `xmlrpc:"num_journals_without_account,omitempty"`
	PartnerId                 *Many2One  `xmlrpc:"partner_id,omitempty"`
	ResPartnerBankId          *Many2One  `xmlrpc:"res_partner_bank_id,omitempty"`
	SanitizedAccNumber        *String    `xmlrpc:"sanitized_acc_number,omitempty"`
	Sequence                  *Int       `xmlrpc:"sequence,omitempty"`
	WriteDate                 *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                  *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// AccountSetupBankManualConfigs represents array of account.setup.bank.manual.config model.
type AccountSetupBankManualConfigs []AccountSetupBankManualConfig

// AccountSetupBankManualConfigModel is the odoo model name.
const AccountSetupBankManualConfigModel = "account.setup.bank.manual.config"

// Many2One convert AccountSetupBankManualConfig to *Many2One.
func (asbmc *AccountSetupBankManualConfig) Many2One() *Many2One {
	return NewMany2One(asbmc.Id.Get(), "")
}

// CreateAccountSetupBankManualConfig creates a new account.setup.bank.manual.config model and returns its id.
func (c *Client) CreateAccountSetupBankManualConfig(asbmc *AccountSetupBankManualConfig) (int64, error) {
	return c.Create(AccountSetupBankManualConfigModel, asbmc)
}

// UpdateAccountSetupBankManualConfig updates an existing account.setup.bank.manual.config record.
func (c *Client) UpdateAccountSetupBankManualConfig(asbmc *AccountSetupBankManualConfig) error {
	return c.UpdateAccountSetupBankManualConfigs([]int64{asbmc.Id.Get()}, asbmc)
}

// UpdateAccountSetupBankManualConfigs updates existing account.setup.bank.manual.config records.
// All records (represented by ids) will be updated by asbmc values.
func (c *Client) UpdateAccountSetupBankManualConfigs(ids []int64, asbmc *AccountSetupBankManualConfig) error {
	return c.Update(AccountSetupBankManualConfigModel, ids, asbmc)
}

// DeleteAccountSetupBankManualConfig deletes an existing account.setup.bank.manual.config record.
func (c *Client) DeleteAccountSetupBankManualConfig(id int64) error {
	return c.DeleteAccountSetupBankManualConfigs([]int64{id})
}

// DeleteAccountSetupBankManualConfigs deletes existing account.setup.bank.manual.config records.
func (c *Client) DeleteAccountSetupBankManualConfigs(ids []int64) error {
	return c.Delete(AccountSetupBankManualConfigModel, ids)
}

// GetAccountSetupBankManualConfig gets account.setup.bank.manual.config existing record.
func (c *Client) GetAccountSetupBankManualConfig(id int64) (*AccountSetupBankManualConfig, error) {
	asbmcs, err := c.GetAccountSetupBankManualConfigs([]int64{id})
	if err != nil {
		return nil, err
	}
	if asbmcs != nil && len(*asbmcs) > 0 {
		return &((*asbmcs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.setup.bank.manual.config not found", id)
}

// GetAccountSetupBankManualConfigs gets account.setup.bank.manual.config existing records.
func (c *Client) GetAccountSetupBankManualConfigs(ids []int64) (*AccountSetupBankManualConfigs, error) {
	asbmcs := &AccountSetupBankManualConfigs{}
	if err := c.Read(AccountSetupBankManualConfigModel, ids, nil, asbmcs); err != nil {
		return nil, err
	}
	return asbmcs, nil
}

// FindAccountSetupBankManualConfig finds account.setup.bank.manual.config record by querying it with criteria.
func (c *Client) FindAccountSetupBankManualConfig(criteria *Criteria) (*AccountSetupBankManualConfig, error) {
	asbmcs := &AccountSetupBankManualConfigs{}
	if err := c.SearchRead(AccountSetupBankManualConfigModel, criteria, NewOptions().Limit(1), asbmcs); err != nil {
		return nil, err
	}
	if asbmcs != nil && len(*asbmcs) > 0 {
		return &((*asbmcs)[0]), nil
	}
	return nil, fmt.Errorf("account.setup.bank.manual.config was not found")
}

// FindAccountSetupBankManualConfigs finds account.setup.bank.manual.config records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountSetupBankManualConfigs(criteria *Criteria, options *Options) (*AccountSetupBankManualConfigs, error) {
	asbmcs := &AccountSetupBankManualConfigs{}
	if err := c.SearchRead(AccountSetupBankManualConfigModel, criteria, options, asbmcs); err != nil {
		return nil, err
	}
	return asbmcs, nil
}

// FindAccountSetupBankManualConfigIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountSetupBankManualConfigIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountSetupBankManualConfigModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountSetupBankManualConfigId finds record id by querying it with criteria.
func (c *Client) FindAccountSetupBankManualConfigId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountSetupBankManualConfigModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.setup.bank.manual.config was not found")
}
