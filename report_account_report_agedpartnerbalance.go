package odoo

import (
	"fmt"
)

// ReportAccountReportAgedpartnerbalance represents report.account.report_agedpartnerbalance model.
type ReportAccountReportAgedpartnerbalance struct {
	LastUpdate  *Time   `xmlrpc:"__last_update,omitempty"`
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// ReportAccountReportAgedpartnerbalances represents array of report.account.report_agedpartnerbalance model.
type ReportAccountReportAgedpartnerbalances []ReportAccountReportAgedpartnerbalance

// ReportAccountReportAgedpartnerbalanceModel is the odoo model name.
const ReportAccountReportAgedpartnerbalanceModel = "report.account.report_agedpartnerbalance"

// Many2One convert ReportAccountReportAgedpartnerbalance to *Many2One.
func (rar *ReportAccountReportAgedpartnerbalance) Many2One() *Many2One {
	return NewMany2One(rar.Id.Get(), "")
}

// CreateReportAccountReportAgedpartnerbalance creates a new report.account.report_agedpartnerbalance model and returns its id.
func (c *Client) CreateReportAccountReportAgedpartnerbalance(rar *ReportAccountReportAgedpartnerbalance) (int64, error) {
	return c.Create(ReportAccountReportAgedpartnerbalanceModel, rar)
}

// UpdateReportAccountReportAgedpartnerbalance updates an existing report.account.report_agedpartnerbalance record.
func (c *Client) UpdateReportAccountReportAgedpartnerbalance(rar *ReportAccountReportAgedpartnerbalance) error {
	return c.UpdateReportAccountReportAgedpartnerbalances([]int64{rar.Id.Get()}, rar)
}

// UpdateReportAccountReportAgedpartnerbalances updates existing report.account.report_agedpartnerbalance records.
// All records (represented by ids) will be updated by rar values.
func (c *Client) UpdateReportAccountReportAgedpartnerbalances(ids []int64, rar *ReportAccountReportAgedpartnerbalance) error {
	return c.Update(ReportAccountReportAgedpartnerbalanceModel, ids, rar)
}

// DeleteReportAccountReportAgedpartnerbalance deletes an existing report.account.report_agedpartnerbalance record.
func (c *Client) DeleteReportAccountReportAgedpartnerbalance(id int64) error {
	return c.DeleteReportAccountReportAgedpartnerbalances([]int64{id})
}

// DeleteReportAccountReportAgedpartnerbalances deletes existing report.account.report_agedpartnerbalance records.
func (c *Client) DeleteReportAccountReportAgedpartnerbalances(ids []int64) error {
	return c.Delete(ReportAccountReportAgedpartnerbalanceModel, ids)
}

// GetReportAccountReportAgedpartnerbalance gets report.account.report_agedpartnerbalance existing record.
func (c *Client) GetReportAccountReportAgedpartnerbalance(id int64) (*ReportAccountReportAgedpartnerbalance, error) {
	rars, err := c.GetReportAccountReportAgedpartnerbalances([]int64{id})
	if err != nil {
		return nil, err
	}
	if rars != nil && len(*rars) > 0 {
		return &((*rars)[0]), nil
	}
	return nil, fmt.Errorf("id %v of report.account.report_agedpartnerbalance not found", id)
}

// GetReportAccountReportAgedpartnerbalances gets report.account.report_agedpartnerbalance existing records.
func (c *Client) GetReportAccountReportAgedpartnerbalances(ids []int64) (*ReportAccountReportAgedpartnerbalances, error) {
	rars := &ReportAccountReportAgedpartnerbalances{}
	if err := c.Read(ReportAccountReportAgedpartnerbalanceModel, ids, nil, rars); err != nil {
		return nil, err
	}
	return rars, nil
}

// FindReportAccountReportAgedpartnerbalance finds report.account.report_agedpartnerbalance record by querying it with criteria.
func (c *Client) FindReportAccountReportAgedpartnerbalance(criteria *Criteria) (*ReportAccountReportAgedpartnerbalance, error) {
	rars := &ReportAccountReportAgedpartnerbalances{}
	if err := c.SearchRead(ReportAccountReportAgedpartnerbalanceModel, criteria, NewOptions().Limit(1), rars); err != nil {
		return nil, err
	}
	if rars != nil && len(*rars) > 0 {
		return &((*rars)[0]), nil
	}
	return nil, fmt.Errorf("no report.account.report_agedpartnerbalance was found with criteria %v", criteria)
}

// FindReportAccountReportAgedpartnerbalances finds report.account.report_agedpartnerbalance records by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportAccountReportAgedpartnerbalances(criteria *Criteria, options *Options) (*ReportAccountReportAgedpartnerbalances, error) {
	rars := &ReportAccountReportAgedpartnerbalances{}
	if err := c.SearchRead(ReportAccountReportAgedpartnerbalanceModel, criteria, options, rars); err != nil {
		return nil, err
	}
	return rars, nil
}

// FindReportAccountReportAgedpartnerbalanceIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportAccountReportAgedpartnerbalanceIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ReportAccountReportAgedpartnerbalanceModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindReportAccountReportAgedpartnerbalanceId finds record id by querying it with criteria.
func (c *Client) FindReportAccountReportAgedpartnerbalanceId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ReportAccountReportAgedpartnerbalanceModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no report.account.report_agedpartnerbalance was found with criteria %v and options %v", criteria, options)
}
