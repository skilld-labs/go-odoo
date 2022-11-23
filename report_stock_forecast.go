package odoo

import (
	"fmt"
)

// ReportStockForecast represents report.stock.forecast model.
type ReportStockForecast struct {
	LastUpdate         *Time     `xmlrpc:"__last_update,omptempty"`
	CumulativeQuantity *Float    `xmlrpc:"cumulative_quantity,omptempty"`
	Date               *Time     `xmlrpc:"date,omptempty"`
	DisplayName        *String   `xmlrpc:"display_name,omptempty"`
	Id                 *Int      `xmlrpc:"id,omptempty"`
	ProductId          *Many2One `xmlrpc:"product_id,omptempty"`
	ProductTmplId      *Many2One `xmlrpc:"product_tmpl_id,omptempty"`
	Quantity           *Float    `xmlrpc:"quantity,omptempty"`
}

// ReportStockForecasts represents array of report.stock.forecast model.
type ReportStockForecasts []ReportStockForecast

// ReportStockForecastModel is the odoo model name.
const ReportStockForecastModel = "report.stock.forecast"

// Many2One convert ReportStockForecast to *Many2One.
func (rsf *ReportStockForecast) Many2One() *Many2One {
	return NewMany2One(rsf.Id.Get(), "")
}

// CreateReportStockForecast creates a new report.stock.forecast model and returns its id.
func (c *Client) CreateReportStockForecast(rsf *ReportStockForecast) (int64, error) {
	return c.Create(ReportStockForecastModel, rsf)
}

// UpdateReportStockForecast updates an existing report.stock.forecast record.
func (c *Client) UpdateReportStockForecast(rsf *ReportStockForecast) error {
	return c.UpdateReportStockForecasts([]int64{rsf.Id.Get()}, rsf)
}

// UpdateReportStockForecasts updates existing report.stock.forecast records.
// All records (represented by ids) will be updated by rsf values.
func (c *Client) UpdateReportStockForecasts(ids []int64, rsf *ReportStockForecast) error {
	return c.Update(ReportStockForecastModel, ids, rsf)
}

// DeleteReportStockForecast deletes an existing report.stock.forecast record.
func (c *Client) DeleteReportStockForecast(id int64) error {
	return c.DeleteReportStockForecasts([]int64{id})
}

// DeleteReportStockForecasts deletes existing report.stock.forecast records.
func (c *Client) DeleteReportStockForecasts(ids []int64) error {
	return c.Delete(ReportStockForecastModel, ids)
}

// GetReportStockForecast gets report.stock.forecast existing record.
func (c *Client) GetReportStockForecast(id int64) (*ReportStockForecast, error) {
	rsfs, err := c.GetReportStockForecasts([]int64{id})
	if err != nil {
		return nil, err
	}
	if rsfs != nil && len(*rsfs) > 0 {
		return &((*rsfs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of report.stock.forecast not found", id)
}

// GetReportStockForecasts gets report.stock.forecast existing records.
func (c *Client) GetReportStockForecasts(ids []int64) (*ReportStockForecasts, error) {
	rsfs := &ReportStockForecasts{}
	if err := c.Read(ReportStockForecastModel, ids, nil, rsfs); err != nil {
		return nil, err
	}
	return rsfs, nil
}

// FindReportStockForecast finds report.stock.forecast record by querying it with criteria.
func (c *Client) FindReportStockForecast(criteria *Criteria) (*ReportStockForecast, error) {
	rsfs := &ReportStockForecasts{}
	if err := c.SearchRead(ReportStockForecastModel, criteria, NewOptions().Limit(1), rsfs); err != nil {
		return nil, err
	}
	if rsfs != nil && len(*rsfs) > 0 {
		return &((*rsfs)[0]), nil
	}
	return nil, fmt.Errorf("report.stock.forecast was not found with criteria %v", criteria)
}

// FindReportStockForecasts finds report.stock.forecast records by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportStockForecasts(criteria *Criteria, options *Options) (*ReportStockForecasts, error) {
	rsfs := &ReportStockForecasts{}
	if err := c.SearchRead(ReportStockForecastModel, criteria, options, rsfs); err != nil {
		return nil, err
	}
	return rsfs, nil
}

// FindReportStockForecastIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportStockForecastIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ReportStockForecastModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindReportStockForecastId finds record id by querying it with criteria.
func (c *Client) FindReportStockForecastId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ReportStockForecastModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("report.stock.forecast was not found with criteria %v and options %v", criteria, options)
}
