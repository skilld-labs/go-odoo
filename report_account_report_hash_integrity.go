package odoo

import (
	"fmt"
)

// ReportAccountReportHashIntegrity represents report.account.report_hash_integrity model.
type ReportAccountReportHashIntegrity struct {
	LastUpdate  *Time   `xmlrpc:"__last_update,omitempty"`
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// ReportAccountReportHashIntegritys represents array of report.account.report_hash_integrity model.
type ReportAccountReportHashIntegritys []ReportAccountReportHashIntegrity

// ReportAccountReportHashIntegrityModel is the odoo model name.
const ReportAccountReportHashIntegrityModel = "report.account.report_hash_integrity"

// Many2One convert ReportAccountReportHashIntegrity to *Many2One.
func (rar *ReportAccountReportHashIntegrity) Many2One() *Many2One {
	return NewMany2One(rar.Id.Get(), "")
}

// CreateReportAccountReportHashIntegrity creates a new report.account.report_hash_integrity model and returns its id.
func (c *Client) CreateReportAccountReportHashIntegrity(rar *ReportAccountReportHashIntegrity) (int64, error) {
	return c.Create(ReportAccountReportHashIntegrityModel, rar)
}

// UpdateReportAccountReportHashIntegrity updates an existing report.account.report_hash_integrity record.
func (c *Client) UpdateReportAccountReportHashIntegrity(rar *ReportAccountReportHashIntegrity) error {
	return c.UpdateReportAccountReportHashIntegritys([]int64{rar.Id.Get()}, rar)
}

// UpdateReportAccountReportHashIntegritys updates existing report.account.report_hash_integrity records.
// All records (represented by IDs) will be updated by rar values.
func (c *Client) UpdateReportAccountReportHashIntegritys(ids []int64, rar *ReportAccountReportHashIntegrity) error {
	return c.Update(ReportAccountReportHashIntegrityModel, ids, rar)
}

// DeleteReportAccountReportHashIntegrity deletes an existing report.account.report_hash_integrity record.
func (c *Client) DeleteReportAccountReportHashIntegrity(id int64) error {
	return c.DeleteReportAccountReportHashIntegritys([]int64{id})
}

// DeleteReportAccountReportHashIntegritys deletes existing report.account.report_hash_integrity records.
func (c *Client) DeleteReportAccountReportHashIntegritys(ids []int64) error {
	return c.Delete(ReportAccountReportHashIntegrityModel, ids)
}

// GetReportAccountReportHashIntegrity gets report.account.report_hash_integrity existing record.
func (c *Client) GetReportAccountReportHashIntegrity(id int64) (*ReportAccountReportHashIntegrity, error) {
	rars, err := c.GetReportAccountReportHashIntegritys([]int64{id})
	if err != nil {
		return nil, err
	}
	if rars != nil && len(*rars) > 0 {
		return &((*rars)[0]), nil
	}
	return nil, fmt.Errorf("id %v of report.account.report_hash_integrity not found", id)
}

// GetReportAccountReportHashIntegritys gets report.account.report_hash_integrity existing records.
func (c *Client) GetReportAccountReportHashIntegritys(ids []int64) (*ReportAccountReportHashIntegritys, error) {
	rars := &ReportAccountReportHashIntegritys{}
	if err := c.Read(ReportAccountReportHashIntegrityModel, ids, nil, rars); err != nil {
		return nil, err
	}
	return rars, nil
}

// FindReportAccountReportHashIntegrity finds report.account.report_hash_integrity record by querying it with criteria.
func (c *Client) FindReportAccountReportHashIntegrity(criteria *Criteria) (*ReportAccountReportHashIntegrity, error) {
	rars := &ReportAccountReportHashIntegritys{}
	if err := c.SearchRead(ReportAccountReportHashIntegrityModel, criteria, NewOptions().Limit(1), rars); err != nil {
		return nil, err
	}
	if rars != nil && len(*rars) > 0 {
		return &((*rars)[0]), nil
	}
	return nil, fmt.Errorf("report.account.report_hash_integrity was not found")
}

// FindReportAccountReportHashIntegritys finds report.account.report_hash_integrity records by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportAccountReportHashIntegritys(criteria *Criteria, options *Options) (*ReportAccountReportHashIntegritys, error) {
	rars := &ReportAccountReportHashIntegritys{}
	if err := c.SearchRead(ReportAccountReportHashIntegrityModel, criteria, options, rars); err != nil {
		return nil, err
	}
	return rars, nil
}

// FindReportAccountReportHashIntegrityIds finds records IDs by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportAccountReportHashIntegrityIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ReportAccountReportHashIntegrityModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindReportAccountReportHashIntegrityId finds record id by querying it with criteria.
func (c *Client) FindReportAccountReportHashIntegrityId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ReportAccountReportHashIntegrityModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("report.account.report_hash_integrity was not found")
}
