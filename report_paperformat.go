package odoo

import (
	"fmt"
)

// ReportPaperformat represents report.paperformat model.
type ReportPaperformat struct {
	LastUpdate    *Time      `xmlrpc:"__last_update,omitempty"`
	CreateDate    *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid     *Many2One  `xmlrpc:"create_uid,omitempty"`
	Default       *Bool      `xmlrpc:"default,omitempty"`
	DisplayName   *String    `xmlrpc:"display_name,omitempty"`
	Dpi           *Int       `xmlrpc:"dpi,omitempty"`
	Format        *Selection `xmlrpc:"format,omitempty"`
	HeaderLine    *Bool      `xmlrpc:"header_line,omitempty"`
	HeaderSpacing *Int       `xmlrpc:"header_spacing,omitempty"`
	Id            *Int       `xmlrpc:"id,omitempty"`
	MarginBottom  *Float     `xmlrpc:"margin_bottom,omitempty"`
	MarginLeft    *Float     `xmlrpc:"margin_left,omitempty"`
	MarginRight   *Float     `xmlrpc:"margin_right,omitempty"`
	MarginTop     *Float     `xmlrpc:"margin_top,omitempty"`
	Name          *String    `xmlrpc:"name,omitempty"`
	Orientation   *Selection `xmlrpc:"orientation,omitempty"`
	PageHeight    *Int       `xmlrpc:"page_height,omitempty"`
	PageWidth     *Int       `xmlrpc:"page_width,omitempty"`
	ReportIds     *Relation  `xmlrpc:"report_ids,omitempty"`
	WriteDate     *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid      *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// ReportPaperformats represents array of report.paperformat model.
type ReportPaperformats []ReportPaperformat

// ReportPaperformatModel is the odoo model name.
const ReportPaperformatModel = "report.paperformat"

// Many2One convert ReportPaperformat to *Many2One.
func (rp *ReportPaperformat) Many2One() *Many2One {
	return NewMany2One(rp.Id.Get(), "")
}

// CreateReportPaperformat creates a new report.paperformat model and returns its id.
func (c *Client) CreateReportPaperformat(rp *ReportPaperformat) (int64, error) {
	return c.Create(ReportPaperformatModel, rp)
}

// UpdateReportPaperformat updates an existing report.paperformat record.
func (c *Client) UpdateReportPaperformat(rp *ReportPaperformat) error {
	return c.UpdateReportPaperformats([]int64{rp.Id.Get()}, rp)
}

// UpdateReportPaperformats updates existing report.paperformat records.
// All records (represented by ids) will be updated by rp values.
func (c *Client) UpdateReportPaperformats(ids []int64, rp *ReportPaperformat) error {
	return c.Update(ReportPaperformatModel, ids, rp)
}

// DeleteReportPaperformat deletes an existing report.paperformat record.
func (c *Client) DeleteReportPaperformat(id int64) error {
	return c.DeleteReportPaperformats([]int64{id})
}

// DeleteReportPaperformats deletes existing report.paperformat records.
func (c *Client) DeleteReportPaperformats(ids []int64) error {
	return c.Delete(ReportPaperformatModel, ids)
}

// GetReportPaperformat gets report.paperformat existing record.
func (c *Client) GetReportPaperformat(id int64) (*ReportPaperformat, error) {
	rps, err := c.GetReportPaperformats([]int64{id})
	if err != nil {
		return nil, err
	}
	if rps != nil && len(*rps) > 0 {
		return &((*rps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of report.paperformat not found", id)
}

// GetReportPaperformats gets report.paperformat existing records.
func (c *Client) GetReportPaperformats(ids []int64) (*ReportPaperformats, error) {
	rps := &ReportPaperformats{}
	if err := c.Read(ReportPaperformatModel, ids, nil, rps); err != nil {
		return nil, err
	}
	return rps, nil
}

// FindReportPaperformat finds report.paperformat record by querying it with criteria.
func (c *Client) FindReportPaperformat(criteria *Criteria) (*ReportPaperformat, error) {
	rps := &ReportPaperformats{}
	if err := c.SearchRead(ReportPaperformatModel, criteria, NewOptions().Limit(1), rps); err != nil {
		return nil, err
	}
	if rps != nil && len(*rps) > 0 {
		return &((*rps)[0]), nil
	}
	return nil, fmt.Errorf("no report.paperformat was found with criteria %v", criteria)
}

// FindReportPaperformats finds report.paperformat records by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportPaperformats(criteria *Criteria, options *Options) (*ReportPaperformats, error) {
	rps := &ReportPaperformats{}
	if err := c.SearchRead(ReportPaperformatModel, criteria, options, rps); err != nil {
		return nil, err
	}
	return rps, nil
}

// FindReportPaperformatIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportPaperformatIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ReportPaperformatModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindReportPaperformatId finds record id by querying it with criteria.
func (c *Client) FindReportPaperformatId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ReportPaperformatModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no report.paperformat was found with criteria %v and options %v", criteria, options)
}
