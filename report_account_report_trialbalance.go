package odoo

import (
	"fmt"
)

// ReportAccountReportTrialbalance represents report.account.report_trialbalance model.
type ReportAccountReportTrialbalance struct {
	LastUpdate  *Time   `xmlrpc:"__last_update,omitempty"`
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// ReportAccountReportTrialbalances represents array of report.account.report_trialbalance model.
type ReportAccountReportTrialbalances []ReportAccountReportTrialbalance

// ReportAccountReportTrialbalanceModel is the odoo model name.
const ReportAccountReportTrialbalanceModel = "report.account.report_trialbalance"

// Many2One convert ReportAccountReportTrialbalance to *Many2One.
func (rar *ReportAccountReportTrialbalance) Many2One() *Many2One {
	return NewMany2One(rar.Id.Get(), "")
}

// CreateReportAccountReportTrialbalance creates a new report.account.report_trialbalance model and returns its id.
func (c *Client) CreateReportAccountReportTrialbalance(rar *ReportAccountReportTrialbalance) (int64, error) {
	return c.Create(ReportAccountReportTrialbalanceModel, rar)
}

// UpdateReportAccountReportTrialbalance updates an existing report.account.report_trialbalance record.
func (c *Client) UpdateReportAccountReportTrialbalance(rar *ReportAccountReportTrialbalance) error {
	return c.UpdateReportAccountReportTrialbalances([]int64{rar.Id.Get()}, rar)
}

// UpdateReportAccountReportTrialbalances updates existing report.account.report_trialbalance records.
// All records (represented by ids) will be updated by rar values.
func (c *Client) UpdateReportAccountReportTrialbalances(ids []int64, rar *ReportAccountReportTrialbalance) error {
	return c.Update(ReportAccountReportTrialbalanceModel, ids, rar)
}

// DeleteReportAccountReportTrialbalance deletes an existing report.account.report_trialbalance record.
func (c *Client) DeleteReportAccountReportTrialbalance(id int64) error {
	return c.DeleteReportAccountReportTrialbalances([]int64{id})
}

// DeleteReportAccountReportTrialbalances deletes existing report.account.report_trialbalance records.
func (c *Client) DeleteReportAccountReportTrialbalances(ids []int64) error {
	return c.Delete(ReportAccountReportTrialbalanceModel, ids)
}

// GetReportAccountReportTrialbalance gets report.account.report_trialbalance existing record.
func (c *Client) GetReportAccountReportTrialbalance(id int64) (*ReportAccountReportTrialbalance, error) {
	rars, err := c.GetReportAccountReportTrialbalances([]int64{id})
	if err != nil {
		return nil, err
	}
	if rars != nil && len(*rars) > 0 {
		return &((*rars)[0]), nil
	}
	return nil, fmt.Errorf("id %v of report.account.report_trialbalance not found", id)
}

// GetReportAccountReportTrialbalances gets report.account.report_trialbalance existing records.
func (c *Client) GetReportAccountReportTrialbalances(ids []int64) (*ReportAccountReportTrialbalances, error) {
	rars := &ReportAccountReportTrialbalances{}
	if err := c.Read(ReportAccountReportTrialbalanceModel, ids, nil, rars); err != nil {
		return nil, err
	}
	return rars, nil
}

// FindReportAccountReportTrialbalance finds report.account.report_trialbalance record by querying it with criteria.
func (c *Client) FindReportAccountReportTrialbalance(criteria *Criteria) (*ReportAccountReportTrialbalance, error) {
	rars := &ReportAccountReportTrialbalances{}
	if err := c.SearchRead(ReportAccountReportTrialbalanceModel, criteria, NewOptions().Limit(1), rars); err != nil {
		return nil, err
	}
	if rars != nil && len(*rars) > 0 {
		return &((*rars)[0]), nil
	}
	return nil, fmt.Errorf("no report.account.report_trialbalance was found with criteria %v", criteria)
}

// FindReportAccountReportTrialbalances finds report.account.report_trialbalance records by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportAccountReportTrialbalances(criteria *Criteria, options *Options) (*ReportAccountReportTrialbalances, error) {
	rars := &ReportAccountReportTrialbalances{}
	if err := c.SearchRead(ReportAccountReportTrialbalanceModel, criteria, options, rars); err != nil {
		return nil, err
	}
	return rars, nil
}

// FindReportAccountReportTrialbalanceIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportAccountReportTrialbalanceIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ReportAccountReportTrialbalanceModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindReportAccountReportTrialbalanceId finds record id by querying it with criteria.
func (c *Client) FindReportAccountReportTrialbalanceId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ReportAccountReportTrialbalanceModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no report.account.report_trialbalance was found with criteria %v and options %v", criteria, options)
}
