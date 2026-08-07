package odoo

// ReportProductReportProducttemplatelabel4X12Noprice represents report.product.report_producttemplatelabel4x12noprice model.
type ReportProductReportProducttemplatelabel4X12Noprice struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// ReportProductReportProducttemplatelabel4X12Noprices represents array of report.product.report_producttemplatelabel4x12noprice model.
type ReportProductReportProducttemplatelabel4X12Noprices []ReportProductReportProducttemplatelabel4X12Noprice

// ReportProductReportProducttemplatelabel4X12NopriceModel is the odoo model name.
const ReportProductReportProducttemplatelabel4X12NopriceModel = "report.product.report_producttemplatelabel4x12noprice"

// Many2One convert ReportProductReportProducttemplatelabel4X12Noprice to *Many2One.
func (rpr *ReportProductReportProducttemplatelabel4X12Noprice) Many2One() *Many2One {
	return NewMany2One(rpr.Id.Get(), "")
}

// CreateReportProductReportProducttemplatelabel4X12Noprice creates a new report.product.report_producttemplatelabel4x12noprice model and returns its id.
func (c *Client) CreateReportProductReportProducttemplatelabel4X12Noprice(rpr *ReportProductReportProducttemplatelabel4X12Noprice) (int64, error) {
	ids, err := c.CreateReportProductReportProducttemplatelabel4X12Noprices([]*ReportProductReportProducttemplatelabel4X12Noprice{rpr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateReportProductReportProducttemplatelabel4X12Noprice creates a new report.product.report_producttemplatelabel4x12noprice model and returns its id.
func (c *Client) CreateReportProductReportProducttemplatelabel4X12Noprices(rprs []*ReportProductReportProducttemplatelabel4X12Noprice) ([]int64, error) {
	var vv []interface{}
	for _, v := range rprs {
		vv = append(vv, v)
	}
	return c.Create(ReportProductReportProducttemplatelabel4X12NopriceModel, vv, nil)
}

// UpdateReportProductReportProducttemplatelabel4X12Noprice updates an existing report.product.report_producttemplatelabel4x12noprice record.
func (c *Client) UpdateReportProductReportProducttemplatelabel4X12Noprice(rpr *ReportProductReportProducttemplatelabel4X12Noprice) error {
	return c.UpdateReportProductReportProducttemplatelabel4X12Noprices([]int64{rpr.Id.Get()}, rpr)
}

// UpdateReportProductReportProducttemplatelabel4X12Noprices updates existing report.product.report_producttemplatelabel4x12noprice records.
// All records (represented by ids) will be updated by rpr values.
func (c *Client) UpdateReportProductReportProducttemplatelabel4X12Noprices(ids []int64, rpr *ReportProductReportProducttemplatelabel4X12Noprice) error {
	return c.Update(ReportProductReportProducttemplatelabel4X12NopriceModel, ids, rpr, nil)
}

// DeleteReportProductReportProducttemplatelabel4X12Noprice deletes an existing report.product.report_producttemplatelabel4x12noprice record.
func (c *Client) DeleteReportProductReportProducttemplatelabel4X12Noprice(id int64) error {
	return c.DeleteReportProductReportProducttemplatelabel4X12Noprices([]int64{id})
}

// DeleteReportProductReportProducttemplatelabel4X12Noprices deletes existing report.product.report_producttemplatelabel4x12noprice records.
func (c *Client) DeleteReportProductReportProducttemplatelabel4X12Noprices(ids []int64) error {
	return c.Delete(ReportProductReportProducttemplatelabel4X12NopriceModel, ids)
}

// GetReportProductReportProducttemplatelabel4X12Noprice gets report.product.report_producttemplatelabel4x12noprice existing record.
func (c *Client) GetReportProductReportProducttemplatelabel4X12Noprice(id int64) (*ReportProductReportProducttemplatelabel4X12Noprice, error) {
	rprs, err := c.GetReportProductReportProducttemplatelabel4X12Noprices([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*rprs)[0]), nil
}

// GetReportProductReportProducttemplatelabel4X12Noprices gets report.product.report_producttemplatelabel4x12noprice existing records.
func (c *Client) GetReportProductReportProducttemplatelabel4X12Noprices(ids []int64) (*ReportProductReportProducttemplatelabel4X12Noprices, error) {
	rprs := &ReportProductReportProducttemplatelabel4X12Noprices{}
	if err := c.Read(ReportProductReportProducttemplatelabel4X12NopriceModel, ids, nil, rprs); err != nil {
		return nil, err
	}
	return rprs, nil
}

// FindReportProductReportProducttemplatelabel4X12Noprice finds report.product.report_producttemplatelabel4x12noprice record by querying it with criteria.
func (c *Client) FindReportProductReportProducttemplatelabel4X12Noprice(criteria *Criteria) (*ReportProductReportProducttemplatelabel4X12Noprice, error) {
	rprs := &ReportProductReportProducttemplatelabel4X12Noprices{}
	if err := c.SearchRead(ReportProductReportProducttemplatelabel4X12NopriceModel, criteria, NewOptions().Limit(1), rprs); err != nil {
		return nil, err
	}
	return &((*rprs)[0]), nil
}

// FindReportProductReportProducttemplatelabel4X12Noprices finds report.product.report_producttemplatelabel4x12noprice records by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportProductReportProducttemplatelabel4X12Noprices(criteria *Criteria, options *Options) (*ReportProductReportProducttemplatelabel4X12Noprices, error) {
	rprs := &ReportProductReportProducttemplatelabel4X12Noprices{}
	if err := c.SearchRead(ReportProductReportProducttemplatelabel4X12NopriceModel, criteria, options, rprs); err != nil {
		return nil, err
	}
	return rprs, nil
}

// FindReportProductReportProducttemplatelabel4X12NopriceIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportProductReportProducttemplatelabel4X12NopriceIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ReportProductReportProducttemplatelabel4X12NopriceModel, criteria, options)
}

// FindReportProductReportProducttemplatelabel4X12NopriceId finds record id by querying it with criteria.
func (c *Client) FindReportProductReportProducttemplatelabel4X12NopriceId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ReportProductReportProducttemplatelabel4X12NopriceModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
