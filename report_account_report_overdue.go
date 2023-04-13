package odoo

import (
	"fmt"
)

// ReportAccountReportOverdue represents report.account.report_overdue model.
type ReportAccountReportOverdue struct {
	LastUpdate  *Time   `xmlrpc:"__last_update,omptempty"`
	DisplayName *String `xmlrpc:"display_name,omptempty"`
	Id          *Int    `xmlrpc:"id,omptempty"`
}

// ReportAccountReportOverdues represents array of report.account.report_overdue model.
type ReportAccountReportOverdues []ReportAccountReportOverdue

// ReportAccountReportOverdueModel is the odoo model name.
const ReportAccountReportOverdueModel = "report.account.report_overdue"

// Many2One convert ReportAccountReportOverdue to *Many2One.
func (rar *ReportAccountReportOverdue) Many2One() *Many2One {
	return NewMany2One(rar.Id.Get(), "")
}

// CreateReportAccountReportOverdue creates a new report.account.report_overdue model and returns its id.
func (c *Client) CreateReportAccountReportOverdue(rar *ReportAccountReportOverdue) (int64, error) {
	ids, err := c.CreateReportAccountReportOverdues([]*ReportAccountReportOverdue{rar})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateReportAccountReportOverdue creates a new report.account.report_overdue model and returns its id.
func (c *Client) CreateReportAccountReportOverdues(rars []*ReportAccountReportOverdue) ([]int64, error) {
	var vv []interface{}
	for _, v := range rars {
		vv = append(vv, v)
	}
	return c.Create(ReportAccountReportOverdueModel, vv)
}

// UpdateReportAccountReportOverdue updates an existing report.account.report_overdue record.
func (c *Client) UpdateReportAccountReportOverdue(rar *ReportAccountReportOverdue) error {
	return c.UpdateReportAccountReportOverdues([]int64{rar.Id.Get()}, rar)
}

// UpdateReportAccountReportOverdues updates existing report.account.report_overdue records.
// All records (represented by ids) will be updated by rar values.
func (c *Client) UpdateReportAccountReportOverdues(ids []int64, rar *ReportAccountReportOverdue) error {
	return c.Update(ReportAccountReportOverdueModel, ids, rar)
}

// DeleteReportAccountReportOverdue deletes an existing report.account.report_overdue record.
func (c *Client) DeleteReportAccountReportOverdue(id int64) error {
	return c.DeleteReportAccountReportOverdues([]int64{id})
}

// DeleteReportAccountReportOverdues deletes existing report.account.report_overdue records.
func (c *Client) DeleteReportAccountReportOverdues(ids []int64) error {
	return c.Delete(ReportAccountReportOverdueModel, ids)
}

// GetReportAccountReportOverdue gets report.account.report_overdue existing record.
func (c *Client) GetReportAccountReportOverdue(id int64) (*ReportAccountReportOverdue, error) {
	rars, err := c.GetReportAccountReportOverdues([]int64{id})
	if err != nil {
		return nil, err
	}
	if rars != nil && len(*rars) > 0 {
		return &((*rars)[0]), nil
	}
	return nil, fmt.Errorf("id %v of report.account.report_overdue not found", id)
}

// GetReportAccountReportOverdues gets report.account.report_overdue existing records.
func (c *Client) GetReportAccountReportOverdues(ids []int64) (*ReportAccountReportOverdues, error) {
	rars := &ReportAccountReportOverdues{}
	if err := c.Read(ReportAccountReportOverdueModel, ids, nil, rars); err != nil {
		return nil, err
	}
	return rars, nil
}

// FindReportAccountReportOverdue finds report.account.report_overdue record by querying it with criteria.
func (c *Client) FindReportAccountReportOverdue(criteria *Criteria) (*ReportAccountReportOverdue, error) {
	rars := &ReportAccountReportOverdues{}
	if err := c.SearchRead(ReportAccountReportOverdueModel, criteria, NewOptions().Limit(1), rars); err != nil {
		return nil, err
	}
	if rars != nil && len(*rars) > 0 {
		return &((*rars)[0]), nil
	}
	return nil, fmt.Errorf("report.account.report_overdue was not found with criteria %v", criteria)
}

// FindReportAccountReportOverdues finds report.account.report_overdue records by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportAccountReportOverdues(criteria *Criteria, options *Options) (*ReportAccountReportOverdues, error) {
	rars := &ReportAccountReportOverdues{}
	if err := c.SearchRead(ReportAccountReportOverdueModel, criteria, options, rars); err != nil {
		return nil, err
	}
	return rars, nil
}

// FindReportAccountReportOverdueIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportAccountReportOverdueIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ReportAccountReportOverdueModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindReportAccountReportOverdueId finds record id by querying it with criteria.
func (c *Client) FindReportAccountReportOverdueId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ReportAccountReportOverdueModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("report.account.report_overdue was not found with criteria %v and options %v", criteria, options)
}
