package odoo

import (
	"fmt"
)

// ResPartnerBank represents res.partner.bank model.
type ResPartnerBank struct {
	LastUpdate         *Time     `xmlrpc:"__last_update,omptempty"`
	AccNumber          *String   `xmlrpc:"acc_number,omptempty"`
	AccType            *String   `xmlrpc:"acc_type,omptempty"`
	BankBic            *String   `xmlrpc:"bank_bic,omptempty"`
	BankId             *Many2One `xmlrpc:"bank_id,omptempty"`
	BankName           *String   `xmlrpc:"bank_name,omptempty"`
	CompanyId          *Many2One `xmlrpc:"company_id,omptempty"`
	CreateDate         *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid          *Many2One `xmlrpc:"create_uid,omptempty"`
	CurrencyId         *Many2One `xmlrpc:"currency_id,omptempty"`
	DisplayName        *String   `xmlrpc:"display_name,omptempty"`
	Id                 *Int      `xmlrpc:"id,omptempty"`
	JournalId          *Relation `xmlrpc:"journal_id,omptempty"`
	PartnerId          *Many2One `xmlrpc:"partner_id,omptempty"`
	SanitizedAccNumber *String   `xmlrpc:"sanitized_acc_number,omptempty"`
	Sequence           *Int      `xmlrpc:"sequence,omptempty"`
	WriteDate          *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid           *Many2One `xmlrpc:"write_uid,omptempty"`
}

// ResPartnerBanks represents array of res.partner.bank model.
type ResPartnerBanks []ResPartnerBank

// ResPartnerBankModel is the odoo model name.
const ResPartnerBankModel = "res.partner.bank"

// Many2One convert ResPartnerBank to *Many2One.
func (rpb *ResPartnerBank) Many2One() *Many2One {
	return NewMany2One(rpb.Id.Get(), "")
}

// CreateResPartnerBank creates a new res.partner.bank model and returns its id.
func (c *Client) CreateResPartnerBank(rpb *ResPartnerBank) (int64, error) {
	return c.Create(ResPartnerBankModel, rpb)
}

// UpdateResPartnerBank updates an existing res.partner.bank record.
func (c *Client) UpdateResPartnerBank(rpb *ResPartnerBank) error {
	return c.UpdateResPartnerBanks([]int64{rpb.Id.Get()}, rpb)
}

// UpdateResPartnerBanks updates existing res.partner.bank records.
// All records (represented by ids) will be updated by rpb values.
func (c *Client) UpdateResPartnerBanks(ids []int64, rpb *ResPartnerBank) error {
	return c.Update(ResPartnerBankModel, ids, rpb)
}

// DeleteResPartnerBank deletes an existing res.partner.bank record.
func (c *Client) DeleteResPartnerBank(id int64) error {
	return c.DeleteResPartnerBanks([]int64{id})
}

// DeleteResPartnerBanks deletes existing res.partner.bank records.
func (c *Client) DeleteResPartnerBanks(ids []int64) error {
	return c.Delete(ResPartnerBankModel, ids)
}

// GetResPartnerBank gets res.partner.bank existing record.
func (c *Client) GetResPartnerBank(id int64) (*ResPartnerBank, error) {
	rpbs, err := c.GetResPartnerBanks([]int64{id})
	if err != nil {
		return nil, err
	}
	if rpbs != nil && len(*rpbs) > 0 {
		return &((*rpbs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of res.partner.bank not found", id)
}

// GetResPartnerBanks gets res.partner.bank existing records.
func (c *Client) GetResPartnerBanks(ids []int64) (*ResPartnerBanks, error) {
	rpbs := &ResPartnerBanks{}
	if err := c.Read(ResPartnerBankModel, ids, nil, rpbs); err != nil {
		return nil, err
	}
	return rpbs, nil
}

// FindResPartnerBank finds res.partner.bank record by querying it with criteria.
func (c *Client) FindResPartnerBank(criteria *Criteria) (*ResPartnerBank, error) {
	rpbs := &ResPartnerBanks{}
	if err := c.SearchRead(ResPartnerBankModel, criteria, NewOptions().Limit(1), rpbs); err != nil {
		return nil, err
	}
	if rpbs != nil && len(*rpbs) > 0 {
		return &((*rpbs)[0]), nil
	}
	return nil, fmt.Errorf("no res.partner.bank was found with criteria %v", criteria)
}

// FindResPartnerBanks finds res.partner.bank records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResPartnerBanks(criteria *Criteria, options *Options) (*ResPartnerBanks, error) {
	rpbs := &ResPartnerBanks{}
	if err := c.SearchRead(ResPartnerBankModel, criteria, options, rpbs); err != nil {
		return nil, err
	}
	return rpbs, nil
}

// FindResPartnerBankIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResPartnerBankIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ResPartnerBankModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindResPartnerBankId finds record id by querying it with criteria.
func (c *Client) FindResPartnerBankId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResPartnerBankModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no res.partner.bank was found with criteria %v and options %v", criteria, options)
}
