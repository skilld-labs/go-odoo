package odoo

import (
	"fmt"
)

// ReportProductReportPricelist represents report.product.report_pricelist model.
type ReportProductReportPricelist struct {
	LastUpdate  *Time   `xmlrpc:"__last_update,omptempty"`
	DisplayName *String `xmlrpc:"display_name,omptempty"`
	Id          *Int    `xmlrpc:"id,omptempty"`
}

// ReportProductReportPricelists represents array of report.product.report_pricelist model.
type ReportProductReportPricelists []ReportProductReportPricelist

// ReportProductReportPricelistModel is the odoo model name.
const ReportProductReportPricelistModel = "report.product.report_pricelist"

// Many2One convert ReportProductReportPricelist to *Many2One.
func (rpr *ReportProductReportPricelist) Many2One() *Many2One {
	return NewMany2One(rpr.Id.Get(), "")
}

// CreateReportProductReportPricelist creates a new report.product.report_pricelist model and returns its id.
func (c *Client) CreateReportProductReportPricelist(rpr *ReportProductReportPricelist) (int64, error) {
	ids, err := c.CreateReportProductReportPricelists([]*ReportProductReportPricelist{rpr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateReportProductReportPricelist creates a new report.product.report_pricelist model and returns its id.
func (c *Client) CreateReportProductReportPricelists(rprs []*ReportProductReportPricelist) ([]int64, error) {
	var vv []interface{}
	for _, v := range rprs {
		vv = append(vv, v)
	}
	return c.Create(ReportProductReportPricelistModel, vv)
}

// UpdateReportProductReportPricelist updates an existing report.product.report_pricelist record.
func (c *Client) UpdateReportProductReportPricelist(rpr *ReportProductReportPricelist) error {
	return c.UpdateReportProductReportPricelists([]int64{rpr.Id.Get()}, rpr)
}

// UpdateReportProductReportPricelists updates existing report.product.report_pricelist records.
// All records (represented by ids) will be updated by rpr values.
func (c *Client) UpdateReportProductReportPricelists(ids []int64, rpr *ReportProductReportPricelist) error {
	return c.Update(ReportProductReportPricelistModel, ids, rpr)
}

// DeleteReportProductReportPricelist deletes an existing report.product.report_pricelist record.
func (c *Client) DeleteReportProductReportPricelist(id int64) error {
	return c.DeleteReportProductReportPricelists([]int64{id})
}

// DeleteReportProductReportPricelists deletes existing report.product.report_pricelist records.
func (c *Client) DeleteReportProductReportPricelists(ids []int64) error {
	return c.Delete(ReportProductReportPricelistModel, ids)
}

// GetReportProductReportPricelist gets report.product.report_pricelist existing record.
func (c *Client) GetReportProductReportPricelist(id int64) (*ReportProductReportPricelist, error) {
	rprs, err := c.GetReportProductReportPricelists([]int64{id})
	if err != nil {
		return nil, err
	}
	if rprs != nil && len(*rprs) > 0 {
		return &((*rprs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of report.product.report_pricelist not found", id)
}

// GetReportProductReportPricelists gets report.product.report_pricelist existing records.
func (c *Client) GetReportProductReportPricelists(ids []int64) (*ReportProductReportPricelists, error) {
	rprs := &ReportProductReportPricelists{}
	if err := c.Read(ReportProductReportPricelistModel, ids, nil, rprs); err != nil {
		return nil, err
	}
	return rprs, nil
}

// FindReportProductReportPricelist finds report.product.report_pricelist record by querying it with criteria.
func (c *Client) FindReportProductReportPricelist(criteria *Criteria) (*ReportProductReportPricelist, error) {
	rprs := &ReportProductReportPricelists{}
	if err := c.SearchRead(ReportProductReportPricelistModel, criteria, NewOptions().Limit(1), rprs); err != nil {
		return nil, err
	}
	if rprs != nil && len(*rprs) > 0 {
		return &((*rprs)[0]), nil
	}
	return nil, fmt.Errorf("report.product.report_pricelist was not found with criteria %v", criteria)
}

// FindReportProductReportPricelists finds report.product.report_pricelist records by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportProductReportPricelists(criteria *Criteria, options *Options) (*ReportProductReportPricelists, error) {
	rprs := &ReportProductReportPricelists{}
	if err := c.SearchRead(ReportProductReportPricelistModel, criteria, options, rprs); err != nil {
		return nil, err
	}
	return rprs, nil
}

// FindReportProductReportPricelistIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportProductReportPricelistIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ReportProductReportPricelistModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindReportProductReportPricelistId finds record id by querying it with criteria.
func (c *Client) FindReportProductReportPricelistId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ReportProductReportPricelistModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("report.product.report_pricelist was not found with criteria %v and options %v", criteria, options)
}
