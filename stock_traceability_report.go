package odoo

import (
	"fmt"
)

// StockTraceabilityReport represents stock.traceability.report model.
type StockTraceabilityReport struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// StockTraceabilityReports represents array of stock.traceability.report model.
type StockTraceabilityReports []StockTraceabilityReport

// StockTraceabilityReportModel is the odoo model name.
const StockTraceabilityReportModel = "stock.traceability.report"

// Many2One convert StockTraceabilityReport to *Many2One.
func (str *StockTraceabilityReport) Many2One() *Many2One {
	return NewMany2One(str.Id.Get(), "")
}

// CreateStockTraceabilityReport creates a new stock.traceability.report model and returns its id.
func (c *Client) CreateStockTraceabilityReport(str *StockTraceabilityReport) (int64, error) {
	return c.Create(StockTraceabilityReportModel, str)
}

// UpdateStockTraceabilityReport updates an existing stock.traceability.report record.
func (c *Client) UpdateStockTraceabilityReport(str *StockTraceabilityReport) error {
	return c.UpdateStockTraceabilityReports([]int64{str.Id.Get()}, str)
}

// UpdateStockTraceabilityReports updates existing stock.traceability.report records.
// All records (represented by ids) will be updated by str values.
func (c *Client) UpdateStockTraceabilityReports(ids []int64, str *StockTraceabilityReport) error {
	return c.Update(StockTraceabilityReportModel, ids, str)
}

// DeleteStockTraceabilityReport deletes an existing stock.traceability.report record.
func (c *Client) DeleteStockTraceabilityReport(id int64) error {
	return c.DeleteStockTraceabilityReports([]int64{id})
}

// DeleteStockTraceabilityReports deletes existing stock.traceability.report records.
func (c *Client) DeleteStockTraceabilityReports(ids []int64) error {
	return c.Delete(StockTraceabilityReportModel, ids)
}

// GetStockTraceabilityReport gets stock.traceability.report existing record.
func (c *Client) GetStockTraceabilityReport(id int64) (*StockTraceabilityReport, error) {
	strs, err := c.GetStockTraceabilityReports([]int64{id})
	if err != nil {
		return nil, err
	}
	if strs != nil && len(*strs) > 0 {
		return &((*strs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of stock.traceability.report not found", id)
}

// GetStockTraceabilityReports gets stock.traceability.report existing records.
func (c *Client) GetStockTraceabilityReports(ids []int64) (*StockTraceabilityReports, error) {
	strs := &StockTraceabilityReports{}
	if err := c.Read(StockTraceabilityReportModel, ids, nil, strs); err != nil {
		return nil, err
	}
	return strs, nil
}

// FindStockTraceabilityReport finds stock.traceability.report record by querying it with criteria.
func (c *Client) FindStockTraceabilityReport(criteria *Criteria) (*StockTraceabilityReport, error) {
	strs := &StockTraceabilityReports{}
	if err := c.SearchRead(StockTraceabilityReportModel, criteria, NewOptions().Limit(1), strs); err != nil {
		return nil, err
	}
	if strs != nil && len(*strs) > 0 {
		return &((*strs)[0]), nil
	}
	return nil, fmt.Errorf("no stock.traceability.report was found with criteria %v", criteria)
}

// FindStockTraceabilityReports finds stock.traceability.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockTraceabilityReports(criteria *Criteria, options *Options) (*StockTraceabilityReports, error) {
	strs := &StockTraceabilityReports{}
	if err := c.SearchRead(StockTraceabilityReportModel, criteria, options, strs); err != nil {
		return nil, err
	}
	return strs, nil
}

// FindStockTraceabilityReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockTraceabilityReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(StockTraceabilityReportModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindStockTraceabilityReportId finds record id by querying it with criteria.
func (c *Client) FindStockTraceabilityReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockTraceabilityReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no stock.traceability.report was found with criteria %v and options %v", criteria, options)
}
