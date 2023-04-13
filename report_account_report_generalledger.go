package odoo

import (
	"fmt"
)

// ReportAccountReportGeneralledger represents report.account.report_generalledger model.
type ReportAccountReportGeneralledger struct {
	LastUpdate  *Time   `xmlrpc:"__last_update,omptempty"`
	DisplayName *String `xmlrpc:"display_name,omptempty"`
	Id          *Int    `xmlrpc:"id,omptempty"`
}

// ReportAccountReportGeneralledgers represents array of report.account.report_generalledger model.
type ReportAccountReportGeneralledgers []ReportAccountReportGeneralledger

// ReportAccountReportGeneralledgerModel is the odoo model name.
const ReportAccountReportGeneralledgerModel = "report.account.report_generalledger"

// Many2One convert ReportAccountReportGeneralledger to *Many2One.
func (rar *ReportAccountReportGeneralledger) Many2One() *Many2One {
	return NewMany2One(rar.Id.Get(), "")
}

// CreateReportAccountReportGeneralledger creates a new report.account.report_generalledger model and returns its id.
func (c *Client) CreateReportAccountReportGeneralledger(rar *ReportAccountReportGeneralledger) (int64, error) {
	ids, err := c.Create(ReportAccountReportGeneralledgerModel, []interface{}{rar})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateReportAccountReportGeneralledger creates a new report.account.report_generalledger model and returns its id.
func (c *Client) CreateReportAccountReportGeneralledgers(rars []*ReportAccountReportGeneralledger) ([]int64, error) {
	var vv []interface{}
	for _, v := range rars {
		vv = append(vv, v)
	}
	return c.Create(ReportAccountReportGeneralledgerModel, vv)
}

// UpdateReportAccountReportGeneralledger updates an existing report.account.report_generalledger record.
func (c *Client) UpdateReportAccountReportGeneralledger(rar *ReportAccountReportGeneralledger) error {
	return c.UpdateReportAccountReportGeneralledgers([]int64{rar.Id.Get()}, rar)
}

// UpdateReportAccountReportGeneralledgers updates existing report.account.report_generalledger records.
// All records (represented by ids) will be updated by rar values.
func (c *Client) UpdateReportAccountReportGeneralledgers(ids []int64, rar *ReportAccountReportGeneralledger) error {
	return c.Update(ReportAccountReportGeneralledgerModel, ids, rar)
}

// DeleteReportAccountReportGeneralledger deletes an existing report.account.report_generalledger record.
func (c *Client) DeleteReportAccountReportGeneralledger(id int64) error {
	return c.DeleteReportAccountReportGeneralledgers([]int64{id})
}

// DeleteReportAccountReportGeneralledgers deletes existing report.account.report_generalledger records.
func (c *Client) DeleteReportAccountReportGeneralledgers(ids []int64) error {
	return c.Delete(ReportAccountReportGeneralledgerModel, ids)
}

// GetReportAccountReportGeneralledger gets report.account.report_generalledger existing record.
func (c *Client) GetReportAccountReportGeneralledger(id int64) (*ReportAccountReportGeneralledger, error) {
	rars, err := c.GetReportAccountReportGeneralledgers([]int64{id})
	if err != nil {
		return nil, err
	}
	if rars != nil && len(*rars) > 0 {
		return &((*rars)[0]), nil
	}
	return nil, fmt.Errorf("id %v of report.account.report_generalledger not found", id)
}

// GetReportAccountReportGeneralledgers gets report.account.report_generalledger existing records.
func (c *Client) GetReportAccountReportGeneralledgers(ids []int64) (*ReportAccountReportGeneralledgers, error) {
	rars := &ReportAccountReportGeneralledgers{}
	if err := c.Read(ReportAccountReportGeneralledgerModel, ids, nil, rars); err != nil {
		return nil, err
	}
	return rars, nil
}

// FindReportAccountReportGeneralledger finds report.account.report_generalledger record by querying it with criteria.
func (c *Client) FindReportAccountReportGeneralledger(criteria *Criteria) (*ReportAccountReportGeneralledger, error) {
	rars := &ReportAccountReportGeneralledgers{}
	if err := c.SearchRead(ReportAccountReportGeneralledgerModel, criteria, NewOptions().Limit(1), rars); err != nil {
		return nil, err
	}
	if rars != nil && len(*rars) > 0 {
		return &((*rars)[0]), nil
	}
	return nil, fmt.Errorf("report.account.report_generalledger was not found with criteria %v", criteria)
}

// FindReportAccountReportGeneralledgers finds report.account.report_generalledger records by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportAccountReportGeneralledgers(criteria *Criteria, options *Options) (*ReportAccountReportGeneralledgers, error) {
	rars := &ReportAccountReportGeneralledgers{}
	if err := c.SearchRead(ReportAccountReportGeneralledgerModel, criteria, options, rars); err != nil {
		return nil, err
	}
	return rars, nil
}

// FindReportAccountReportGeneralledgerIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportAccountReportGeneralledgerIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ReportAccountReportGeneralledgerModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindReportAccountReportGeneralledgerId finds record id by querying it with criteria.
func (c *Client) FindReportAccountReportGeneralledgerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ReportAccountReportGeneralledgerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("report.account.report_generalledger was not found with criteria %v and options %v", criteria, options)
}
