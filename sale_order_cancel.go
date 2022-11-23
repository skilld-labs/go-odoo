package odoo

import (
	"fmt"
)

// SaleOrderCancel represents sale.order.cancel model.
type SaleOrderCancel struct {
	LastUpdate          *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate          *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid           *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayInvoiceAlert *Bool     `xmlrpc:"display_invoice_alert,omitempty"`
	DisplayName         *String   `xmlrpc:"display_name,omitempty"`
	Id                  *Int      `xmlrpc:"id,omitempty"`
	OrderId             *Many2One `xmlrpc:"order_id,omitempty"`
	WriteDate           *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid            *Many2One `xmlrpc:"write_uid,omitempty"`
}

// SaleOrderCancels represents array of sale.order.cancel model.
type SaleOrderCancels []SaleOrderCancel

// SaleOrderCancelModel is the odoo model name.
const SaleOrderCancelModel = "sale.order.cancel"

// Many2One convert SaleOrderCancel to *Many2One.
func (soc *SaleOrderCancel) Many2One() *Many2One {
	return NewMany2One(soc.Id.Get(), "")
}

// CreateSaleOrderCancel creates a new sale.order.cancel model and returns its id.
func (c *Client) CreateSaleOrderCancel(soc *SaleOrderCancel) (int64, error) {
	return c.Create(SaleOrderCancelModel, soc)
}

// UpdateSaleOrderCancel updates an existing sale.order.cancel record.
func (c *Client) UpdateSaleOrderCancel(soc *SaleOrderCancel) error {
	return c.UpdateSaleOrderCancels([]int64{soc.Id.Get()}, soc)
}

// UpdateSaleOrderCancels updates existing sale.order.cancel records.
// All records (represented by IDs) will be updated by soc values.
func (c *Client) UpdateSaleOrderCancels(ids []int64, soc *SaleOrderCancel) error {
	return c.Update(SaleOrderCancelModel, ids, soc)
}

// DeleteSaleOrderCancel deletes an existing sale.order.cancel record.
func (c *Client) DeleteSaleOrderCancel(id int64) error {
	return c.DeleteSaleOrderCancels([]int64{id})
}

// DeleteSaleOrderCancels deletes existing sale.order.cancel records.
func (c *Client) DeleteSaleOrderCancels(ids []int64) error {
	return c.Delete(SaleOrderCancelModel, ids)
}

// GetSaleOrderCancel gets sale.order.cancel existing record.
func (c *Client) GetSaleOrderCancel(id int64) (*SaleOrderCancel, error) {
	socs, err := c.GetSaleOrderCancels([]int64{id})
	if err != nil {
		return nil, err
	}
	if socs != nil && len(*socs) > 0 {
		return &((*socs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of sale.order.cancel not found", id)
}

// GetSaleOrderCancels gets sale.order.cancel existing records.
func (c *Client) GetSaleOrderCancels(ids []int64) (*SaleOrderCancels, error) {
	socs := &SaleOrderCancels{}
	if err := c.Read(SaleOrderCancelModel, ids, nil, socs); err != nil {
		return nil, err
	}
	return socs, nil
}

// FindSaleOrderCancel finds sale.order.cancel record by querying it with criteria.
func (c *Client) FindSaleOrderCancel(criteria *Criteria) (*SaleOrderCancel, error) {
	socs := &SaleOrderCancels{}
	if err := c.SearchRead(SaleOrderCancelModel, criteria, NewOptions().Limit(1), socs); err != nil {
		return nil, err
	}
	if socs != nil && len(*socs) > 0 {
		return &((*socs)[0]), nil
	}
	return nil, fmt.Errorf("sale.order.cancel was not found")
}

// FindSaleOrderCancels finds sale.order.cancel records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleOrderCancels(criteria *Criteria, options *Options) (*SaleOrderCancels, error) {
	socs := &SaleOrderCancels{}
	if err := c.SearchRead(SaleOrderCancelModel, criteria, options, socs); err != nil {
		return nil, err
	}
	return socs, nil
}

// FindSaleOrderCancelIds finds records IDs by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleOrderCancelIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SaleOrderCancelModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSaleOrderCancelId finds record id by querying it with criteria.
func (c *Client) FindSaleOrderCancelId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SaleOrderCancelModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("sale.order.cancel was not found")
}
