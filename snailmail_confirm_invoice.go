package odoo

import (
	"fmt"
)

// SnailmailConfirmInvoice represents snailmail.confirm.invoice model.
type SnailmailConfirmInvoice struct {
	LastUpdate    *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate    *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid     *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName   *String   `xmlrpc:"display_name,omitempty"`
	Id            *Int      `xmlrpc:"id,omitempty"`
	InvoiceSendId *Many2One `xmlrpc:"invoice_send_id,omitempty"`
	ModelName     *String   `xmlrpc:"model_name,omitempty"`
	WriteDate     *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid      *Many2One `xmlrpc:"write_uid,omitempty"`
}

// SnailmailConfirmInvoices represents array of snailmail.confirm.invoice model.
type SnailmailConfirmInvoices []SnailmailConfirmInvoice

// SnailmailConfirmInvoiceModel is the odoo model name.
const SnailmailConfirmInvoiceModel = "snailmail.confirm.invoice"

// Many2One convert SnailmailConfirmInvoice to *Many2One.
func (sci *SnailmailConfirmInvoice) Many2One() *Many2One {
	return NewMany2One(sci.Id.Get(), "")
}

// CreateSnailmailConfirmInvoice creates a new snailmail.confirm.invoice model and returns its id.
func (c *Client) CreateSnailmailConfirmInvoice(sci *SnailmailConfirmInvoice) (int64, error) {
	return c.Create(SnailmailConfirmInvoiceModel, sci)
}

// UpdateSnailmailConfirmInvoice updates an existing snailmail.confirm.invoice record.
func (c *Client) UpdateSnailmailConfirmInvoice(sci *SnailmailConfirmInvoice) error {
	return c.UpdateSnailmailConfirmInvoices([]int64{sci.Id.Get()}, sci)
}

// UpdateSnailmailConfirmInvoices updates existing snailmail.confirm.invoice records.
// All records (represented by IDs) will be updated by sci values.
func (c *Client) UpdateSnailmailConfirmInvoices(ids []int64, sci *SnailmailConfirmInvoice) error {
	return c.Update(SnailmailConfirmInvoiceModel, ids, sci)
}

// DeleteSnailmailConfirmInvoice deletes an existing snailmail.confirm.invoice record.
func (c *Client) DeleteSnailmailConfirmInvoice(id int64) error {
	return c.DeleteSnailmailConfirmInvoices([]int64{id})
}

// DeleteSnailmailConfirmInvoices deletes existing snailmail.confirm.invoice records.
func (c *Client) DeleteSnailmailConfirmInvoices(ids []int64) error {
	return c.Delete(SnailmailConfirmInvoiceModel, ids)
}

// GetSnailmailConfirmInvoice gets snailmail.confirm.invoice existing record.
func (c *Client) GetSnailmailConfirmInvoice(id int64) (*SnailmailConfirmInvoice, error) {
	scis, err := c.GetSnailmailConfirmInvoices([]int64{id})
	if err != nil {
		return nil, err
	}
	if scis != nil && len(*scis) > 0 {
		return &((*scis)[0]), nil
	}
	return nil, fmt.Errorf("id %v of snailmail.confirm.invoice not found", id)
}

// GetSnailmailConfirmInvoices gets snailmail.confirm.invoice existing records.
func (c *Client) GetSnailmailConfirmInvoices(ids []int64) (*SnailmailConfirmInvoices, error) {
	scis := &SnailmailConfirmInvoices{}
	if err := c.Read(SnailmailConfirmInvoiceModel, ids, nil, scis); err != nil {
		return nil, err
	}
	return scis, nil
}

// FindSnailmailConfirmInvoice finds snailmail.confirm.invoice record by querying it with criteria.
func (c *Client) FindSnailmailConfirmInvoice(criteria *Criteria) (*SnailmailConfirmInvoice, error) {
	scis := &SnailmailConfirmInvoices{}
	if err := c.SearchRead(SnailmailConfirmInvoiceModel, criteria, NewOptions().Limit(1), scis); err != nil {
		return nil, err
	}
	if scis != nil && len(*scis) > 0 {
		return &((*scis)[0]), nil
	}
	return nil, fmt.Errorf("snailmail.confirm.invoice was not found")
}

// FindSnailmailConfirmInvoices finds snailmail.confirm.invoice records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSnailmailConfirmInvoices(criteria *Criteria, options *Options) (*SnailmailConfirmInvoices, error) {
	scis := &SnailmailConfirmInvoices{}
	if err := c.SearchRead(SnailmailConfirmInvoiceModel, criteria, options, scis); err != nil {
		return nil, err
	}
	return scis, nil
}

// FindSnailmailConfirmInvoiceIds finds records IDs by querying it
// and filtering it with criteria and options.
func (c *Client) FindSnailmailConfirmInvoiceIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SnailmailConfirmInvoiceModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSnailmailConfirmInvoiceId finds record id by querying it with criteria.
func (c *Client) FindSnailmailConfirmInvoiceId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SnailmailConfirmInvoiceModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("snailmail.confirm.invoice was not found")
}
