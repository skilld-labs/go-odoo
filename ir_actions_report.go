package odoo

import (
	"fmt"
)

// IrActionsReport represents ir.actions.report model.
type IrActionsReport struct {
	LastUpdate       *Time      `xmlrpc:"__last_update,omitempty"`
	Attachment       *String    `xmlrpc:"attachment,omitempty"`
	AttachmentUse    *Bool      `xmlrpc:"attachment_use,omitempty"`
	BindingModelId   *Many2One  `xmlrpc:"binding_model_id,omitempty"`
	BindingType      *Selection `xmlrpc:"binding_type,omitempty"`
	BindingViewTypes *String    `xmlrpc:"binding_view_types,omitempty"`
	CreateDate       *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid        *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName      *String    `xmlrpc:"display_name,omitempty"`
	GroupsId         *Relation  `xmlrpc:"groups_id,omitempty"`
	Help             *String    `xmlrpc:"help,omitempty"`
	Id               *Int       `xmlrpc:"id,omitempty"`
	Model            *String    `xmlrpc:"model,omitempty"`
	ModelId          *Many2One  `xmlrpc:"model_id,omitempty"`
	Multi            *Bool      `xmlrpc:"multi,omitempty"`
	Name             *String    `xmlrpc:"name,omitempty"`
	PaperformatId    *Many2One  `xmlrpc:"paperformat_id,omitempty"`
	PrintReportName  *String    `xmlrpc:"print_report_name,omitempty"`
	ReportFile       *String    `xmlrpc:"report_file,omitempty"`
	ReportName       *String    `xmlrpc:"report_name,omitempty"`
	ReportType       *Selection `xmlrpc:"report_type,omitempty"`
	Type             *String    `xmlrpc:"type,omitempty"`
	WriteDate        *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid         *Many2One  `xmlrpc:"write_uid,omitempty"`
	XmlId            *String    `xmlrpc:"xml_id,omitempty"`
}

// IrActionsReports represents array of ir.actions.report model.
type IrActionsReports []IrActionsReport

// IrActionsReportModel is the odoo model name.
const IrActionsReportModel = "ir.actions.report"

// Many2One convert IrActionsReport to *Many2One.
func (iar *IrActionsReport) Many2One() *Many2One {
	return NewMany2One(iar.Id.Get(), "")
}

// CreateIrActionsReport creates a new ir.actions.report model and returns its id.
func (c *Client) CreateIrActionsReport(iar *IrActionsReport) (int64, error) {
	return c.Create(IrActionsReportModel, iar)
}

// UpdateIrActionsReport updates an existing ir.actions.report record.
func (c *Client) UpdateIrActionsReport(iar *IrActionsReport) error {
	return c.UpdateIrActionsReports([]int64{iar.Id.Get()}, iar)
}

// UpdateIrActionsReports updates existing ir.actions.report records.
// All records (represented by IDs) will be updated by iar values.
func (c *Client) UpdateIrActionsReports(ids []int64, iar *IrActionsReport) error {
	return c.Update(IrActionsReportModel, ids, iar)
}

// DeleteIrActionsReport deletes an existing ir.actions.report record.
func (c *Client) DeleteIrActionsReport(id int64) error {
	return c.DeleteIrActionsReports([]int64{id})
}

// DeleteIrActionsReports deletes existing ir.actions.report records.
func (c *Client) DeleteIrActionsReports(ids []int64) error {
	return c.Delete(IrActionsReportModel, ids)
}

// GetIrActionsReport gets ir.actions.report existing record.
func (c *Client) GetIrActionsReport(id int64) (*IrActionsReport, error) {
	iars, err := c.GetIrActionsReports([]int64{id})
	if err != nil {
		return nil, err
	}
	if iars != nil && len(*iars) > 0 {
		return &((*iars)[0]), nil
	}
	return nil, fmt.Errorf("id %v of ir.actions.report not found", id)
}

// GetIrActionsReports gets ir.actions.report existing records.
func (c *Client) GetIrActionsReports(ids []int64) (*IrActionsReports, error) {
	iars := &IrActionsReports{}
	if err := c.Read(IrActionsReportModel, ids, nil, iars); err != nil {
		return nil, err
	}
	return iars, nil
}

// FindIrActionsReport finds ir.actions.report record by querying it with criteria.
func (c *Client) FindIrActionsReport(criteria *Criteria) (*IrActionsReport, error) {
	iars := &IrActionsReports{}
	if err := c.SearchRead(IrActionsReportModel, criteria, NewOptions().Limit(1), iars); err != nil {
		return nil, err
	}
	if iars != nil && len(*iars) > 0 {
		return &((*iars)[0]), nil
	}
	return nil, fmt.Errorf("ir.actions.report was not found")
}

// FindIrActionsReports finds ir.actions.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrActionsReports(criteria *Criteria, options *Options) (*IrActionsReports, error) {
	iars := &IrActionsReports{}
	if err := c.SearchRead(IrActionsReportModel, criteria, options, iars); err != nil {
		return nil, err
	}
	return iars, nil
}

// FindIrActionsReportIds finds records IDs by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrActionsReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(IrActionsReportModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindIrActionsReportId finds record id by querying it with criteria.
func (c *Client) FindIrActionsReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrActionsReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("ir.actions.report was not found")
}
