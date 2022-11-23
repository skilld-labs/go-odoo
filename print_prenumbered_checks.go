package odoo

import (
	"fmt"
)

// PrintPrenumberedChecks represents print.prenumbered.checks model.
type PrintPrenumberedChecks struct {
	LastUpdate      *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate      *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid       *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName     *String   `xmlrpc:"display_name,omitempty"`
	Id              *Int      `xmlrpc:"id,omitempty"`
	NextCheckNumber *String   `xmlrpc:"next_check_number,omitempty"`
	WriteDate       *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid        *Many2One `xmlrpc:"write_uid,omitempty"`
}

// PrintPrenumberedCheckss represents array of print.prenumbered.checks model.
type PrintPrenumberedCheckss []PrintPrenumberedChecks

// PrintPrenumberedChecksModel is the odoo model name.
const PrintPrenumberedChecksModel = "print.prenumbered.checks"

// Many2One convert PrintPrenumberedChecks to *Many2One.
func (ppc *PrintPrenumberedChecks) Many2One() *Many2One {
	return NewMany2One(ppc.Id.Get(), "")
}

// CreatePrintPrenumberedChecks creates a new print.prenumbered.checks model and returns its id.
func (c *Client) CreatePrintPrenumberedChecks(ppc *PrintPrenumberedChecks) (int64, error) {
	return c.Create(PrintPrenumberedChecksModel, ppc)
}

// UpdatePrintPrenumberedChecks updates an existing print.prenumbered.checks record.
func (c *Client) UpdatePrintPrenumberedChecks(ppc *PrintPrenumberedChecks) error {
	return c.UpdatePrintPrenumberedCheckss([]int64{ppc.Id.Get()}, ppc)
}

// UpdatePrintPrenumberedCheckss updates existing print.prenumbered.checks records.
// All records (represented by IDs) will be updated by ppc values.
func (c *Client) UpdatePrintPrenumberedCheckss(ids []int64, ppc *PrintPrenumberedChecks) error {
	return c.Update(PrintPrenumberedChecksModel, ids, ppc)
}

// DeletePrintPrenumberedChecks deletes an existing print.prenumbered.checks record.
func (c *Client) DeletePrintPrenumberedChecks(id int64) error {
	return c.DeletePrintPrenumberedCheckss([]int64{id})
}

// DeletePrintPrenumberedCheckss deletes existing print.prenumbered.checks records.
func (c *Client) DeletePrintPrenumberedCheckss(ids []int64) error {
	return c.Delete(PrintPrenumberedChecksModel, ids)
}

// GetPrintPrenumberedChecks gets print.prenumbered.checks existing record.
func (c *Client) GetPrintPrenumberedChecks(id int64) (*PrintPrenumberedChecks, error) {
	ppcs, err := c.GetPrintPrenumberedCheckss([]int64{id})
	if err != nil {
		return nil, err
	}
	if ppcs != nil && len(*ppcs) > 0 {
		return &((*ppcs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of print.prenumbered.checks not found", id)
}

// GetPrintPrenumberedCheckss gets print.prenumbered.checks existing records.
func (c *Client) GetPrintPrenumberedCheckss(ids []int64) (*PrintPrenumberedCheckss, error) {
	ppcs := &PrintPrenumberedCheckss{}
	if err := c.Read(PrintPrenumberedChecksModel, ids, nil, ppcs); err != nil {
		return nil, err
	}
	return ppcs, nil
}

// FindPrintPrenumberedChecks finds print.prenumbered.checks record by querying it with criteria.
func (c *Client) FindPrintPrenumberedChecks(criteria *Criteria) (*PrintPrenumberedChecks, error) {
	ppcs := &PrintPrenumberedCheckss{}
	if err := c.SearchRead(PrintPrenumberedChecksModel, criteria, NewOptions().Limit(1), ppcs); err != nil {
		return nil, err
	}
	if ppcs != nil && len(*ppcs) > 0 {
		return &((*ppcs)[0]), nil
	}
	return nil, fmt.Errorf("print.prenumbered.checks was not found")
}

// FindPrintPrenumberedCheckss finds print.prenumbered.checks records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPrintPrenumberedCheckss(criteria *Criteria, options *Options) (*PrintPrenumberedCheckss, error) {
	ppcs := &PrintPrenumberedCheckss{}
	if err := c.SearchRead(PrintPrenumberedChecksModel, criteria, options, ppcs); err != nil {
		return nil, err
	}
	return ppcs, nil
}

// FindPrintPrenumberedChecksIds finds records IDs by querying it
// and filtering it with criteria and options.
func (c *Client) FindPrintPrenumberedChecksIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(PrintPrenumberedChecksModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindPrintPrenumberedChecksId finds record id by querying it with criteria.
func (c *Client) FindPrintPrenumberedChecksId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PrintPrenumberedChecksModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("print.prenumbered.checks was not found")
}
