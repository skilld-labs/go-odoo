package odoo

import (
	"fmt"
)

// AccountReconcileModel represents account.reconcile.model model.
type AccountReconcileModel struct {
	LastUpdate                 *Time      `xmlrpc:"__last_update,omitempty"`
	Active                     *Bool      `xmlrpc:"active,omitempty"`
	AutoReconcile              *Bool      `xmlrpc:"auto_reconcile,omitempty"`
	CompanyId                  *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate                 *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                  *Many2One  `xmlrpc:"create_uid,omitempty"`
	DecimalSeparator           *String    `xmlrpc:"decimal_separator,omitempty"`
	DisplayName                *String    `xmlrpc:"display_name,omitempty"`
	Id                         *Int       `xmlrpc:"id,omitempty"`
	LineIds                    *Relation  `xmlrpc:"line_ids,omitempty"`
	MatchAmount                *Selection `xmlrpc:"match_amount,omitempty"`
	MatchAmountMax             *Float     `xmlrpc:"match_amount_max,omitempty"`
	MatchAmountMin             *Float     `xmlrpc:"match_amount_min,omitempty"`
	MatchJournalIds            *Relation  `xmlrpc:"match_journal_ids,omitempty"`
	MatchLabel                 *Selection `xmlrpc:"match_label,omitempty"`
	MatchLabelParam            *String    `xmlrpc:"match_label_param,omitempty"`
	MatchNature                *Selection `xmlrpc:"match_nature,omitempty"`
	MatchNote                  *Selection `xmlrpc:"match_note,omitempty"`
	MatchNoteParam             *String    `xmlrpc:"match_note_param,omitempty"`
	MatchPartner               *Bool      `xmlrpc:"match_partner,omitempty"`
	MatchPartnerCategoryIds    *Relation  `xmlrpc:"match_partner_category_ids,omitempty"`
	MatchPartnerIds            *Relation  `xmlrpc:"match_partner_ids,omitempty"`
	MatchSameCurrency          *Bool      `xmlrpc:"match_same_currency,omitempty"`
	MatchTextLocationLabel     *Bool      `xmlrpc:"match_text_location_label,omitempty"`
	MatchTextLocationNote      *Bool      `xmlrpc:"match_text_location_note,omitempty"`
	MatchTextLocationReference *Bool      `xmlrpc:"match_text_location_reference,omitempty"`
	MatchTotalAmount           *Bool      `xmlrpc:"match_total_amount,omitempty"`
	MatchTotalAmountParam      *Float     `xmlrpc:"match_total_amount_param,omitempty"`
	MatchTransactionType       *Selection `xmlrpc:"match_transaction_type,omitempty"`
	MatchTransactionTypeParam  *String    `xmlrpc:"match_transaction_type_param,omitempty"`
	MatchingOrder              *Selection `xmlrpc:"matching_order,omitempty"`
	Name                       *String    `xmlrpc:"name,omitempty"`
	NumberEntries              *Int       `xmlrpc:"number_entries,omitempty"`
	PartnerMappingLineIds      *Relation  `xmlrpc:"partner_mapping_line_ids,omitempty"`
	PastMonthsLimit            *Int       `xmlrpc:"past_months_limit,omitempty"`
	RuleType                   *Selection `xmlrpc:"rule_type,omitempty"`
	Sequence                   *Int       `xmlrpc:"sequence,omitempty"`
	ShowDecimalSeparator       *Bool      `xmlrpc:"show_decimal_separator,omitempty"`
	ToCheck                    *Bool      `xmlrpc:"to_check,omitempty"`
	WriteDate                  *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                   *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// AccountReconcileModels represents array of account.reconcile.model model.
type AccountReconcileModels []AccountReconcileModel

// AccountReconcileModelModel is the odoo model name.
const AccountReconcileModelModel = "account.reconcile.model"

// Many2One convert AccountReconcileModel to *Many2One.
func (arm *AccountReconcileModel) Many2One() *Many2One {
	return NewMany2One(arm.Id.Get(), "")
}

// CreateAccountReconcileModel creates a new account.reconcile.model model and returns its id.
func (c *Client) CreateAccountReconcileModel(arm *AccountReconcileModel) (int64, error) {
	return c.Create(AccountReconcileModelModel, arm)
}

// UpdateAccountReconcileModel updates an existing account.reconcile.model record.
func (c *Client) UpdateAccountReconcileModel(arm *AccountReconcileModel) error {
	return c.UpdateAccountReconcileModels([]int64{arm.Id.Get()}, arm)
}

// UpdateAccountReconcileModels updates existing account.reconcile.model records.
// All records (represented by ids) will be updated by arm values.
func (c *Client) UpdateAccountReconcileModels(ids []int64, arm *AccountReconcileModel) error {
	return c.Update(AccountReconcileModelModel, ids, arm)
}

// DeleteAccountReconcileModel deletes an existing account.reconcile.model record.
func (c *Client) DeleteAccountReconcileModel(id int64) error {
	return c.DeleteAccountReconcileModels([]int64{id})
}

// DeleteAccountReconcileModels deletes existing account.reconcile.model records.
func (c *Client) DeleteAccountReconcileModels(ids []int64) error {
	return c.Delete(AccountReconcileModelModel, ids)
}

// GetAccountReconcileModel gets account.reconcile.model existing record.
func (c *Client) GetAccountReconcileModel(id int64) (*AccountReconcileModel, error) {
	arms, err := c.GetAccountReconcileModels([]int64{id})
	if err != nil {
		return nil, err
	}
	if arms != nil && len(*arms) > 0 {
		return &((*arms)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.reconcile.model not found", id)
}

// GetAccountReconcileModels gets account.reconcile.model existing records.
func (c *Client) GetAccountReconcileModels(ids []int64) (*AccountReconcileModels, error) {
	arms := &AccountReconcileModels{}
	if err := c.Read(AccountReconcileModelModel, ids, nil, arms); err != nil {
		return nil, err
	}
	return arms, nil
}

// FindAccountReconcileModel finds account.reconcile.model record by querying it with criteria.
func (c *Client) FindAccountReconcileModel(criteria *Criteria) (*AccountReconcileModel, error) {
	arms := &AccountReconcileModels{}
	if err := c.SearchRead(AccountReconcileModelModel, criteria, NewOptions().Limit(1), arms); err != nil {
		return nil, err
	}
	if arms != nil && len(*arms) > 0 {
		return &((*arms)[0]), nil
	}
	return nil, fmt.Errorf("account.reconcile.model was not found")
}

// FindAccountReconcileModels finds account.reconcile.model records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReconcileModels(criteria *Criteria, options *Options) (*AccountReconcileModels, error) {
	arms := &AccountReconcileModels{}
	if err := c.SearchRead(AccountReconcileModelModel, criteria, options, arms); err != nil {
		return nil, err
	}
	return arms, nil
}

// FindAccountReconcileModelIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReconcileModelIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountReconcileModelModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountReconcileModelId finds record id by querying it with criteria.
func (c *Client) FindAccountReconcileModelId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountReconcileModelModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.reconcile.model was not found")
}
