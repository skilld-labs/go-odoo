package odoo

import (
	"fmt"
)

// HrHolidaysStatus represents hr.holidays.status model.
type HrHolidaysStatus struct {
	LastUpdate             *Time      `xmlrpc:"__last_update,omptempty"`
	Active                 *Bool      `xmlrpc:"active,omptempty"`
	CategId                *Many2One  `xmlrpc:"categ_id,omptempty"`
	ColorName              *Selection `xmlrpc:"color_name,omptempty"`
	CompanyId              *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate             *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid              *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName            *String    `xmlrpc:"display_name,omptempty"`
	DoubleValidation       *Bool      `xmlrpc:"double_validation,omptempty"`
	Id                     *Int       `xmlrpc:"id,omptempty"`
	LeavesTaken            *Float     `xmlrpc:"leaves_taken,omptempty"`
	Limit                  *Bool      `xmlrpc:"limit,omptempty"`
	MaxLeaves              *Float     `xmlrpc:"max_leaves,omptempty"`
	Name                   *String    `xmlrpc:"name,omptempty"`
	RemainingLeaves        *Float     `xmlrpc:"remaining_leaves,omptempty"`
	TimesheetGenerate      *Bool      `xmlrpc:"timesheet_generate,omptempty"`
	TimesheetProjectId     *Many2One  `xmlrpc:"timesheet_project_id,omptempty"`
	TimesheetTaskId        *Many2One  `xmlrpc:"timesheet_task_id,omptempty"`
	VirtualRemainingLeaves *Float     `xmlrpc:"virtual_remaining_leaves,omptempty"`
	WriteDate              *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid               *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// HrHolidaysStatuss represents array of hr.holidays.status model.
type HrHolidaysStatuss []HrHolidaysStatus

// HrHolidaysStatusModel is the odoo model name.
const HrHolidaysStatusModel = "hr.holidays.status"

// Many2One convert HrHolidaysStatus to *Many2One.
func (hhs *HrHolidaysStatus) Many2One() *Many2One {
	return NewMany2One(hhs.Id.Get(), "")
}

// CreateHrHolidaysStatus creates a new hr.holidays.status model and returns its id.
func (c *Client) CreateHrHolidaysStatus(hhs *HrHolidaysStatus) (int64, error) {
	ids, err := c.Create(HrHolidaysStatusModel, []interface{}{hhs})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrHolidaysStatus creates a new hr.holidays.status model and returns its id.
func (c *Client) CreateHrHolidaysStatuss(hhss []*HrHolidaysStatus) ([]int64, error) {
	var vv []interface{}
	for _, v := range hhss {
		vv = append(vv, v)
	}
	return c.Create(HrHolidaysStatusModel, vv)
}

// UpdateHrHolidaysStatus updates an existing hr.holidays.status record.
func (c *Client) UpdateHrHolidaysStatus(hhs *HrHolidaysStatus) error {
	return c.UpdateHrHolidaysStatuss([]int64{hhs.Id.Get()}, hhs)
}

// UpdateHrHolidaysStatuss updates existing hr.holidays.status records.
// All records (represented by ids) will be updated by hhs values.
func (c *Client) UpdateHrHolidaysStatuss(ids []int64, hhs *HrHolidaysStatus) error {
	return c.Update(HrHolidaysStatusModel, ids, hhs)
}

// DeleteHrHolidaysStatus deletes an existing hr.holidays.status record.
func (c *Client) DeleteHrHolidaysStatus(id int64) error {
	return c.DeleteHrHolidaysStatuss([]int64{id})
}

// DeleteHrHolidaysStatuss deletes existing hr.holidays.status records.
func (c *Client) DeleteHrHolidaysStatuss(ids []int64) error {
	return c.Delete(HrHolidaysStatusModel, ids)
}

// GetHrHolidaysStatus gets hr.holidays.status existing record.
func (c *Client) GetHrHolidaysStatus(id int64) (*HrHolidaysStatus, error) {
	hhss, err := c.GetHrHolidaysStatuss([]int64{id})
	if err != nil {
		return nil, err
	}
	if hhss != nil && len(*hhss) > 0 {
		return &((*hhss)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.holidays.status not found", id)
}

// GetHrHolidaysStatuss gets hr.holidays.status existing records.
func (c *Client) GetHrHolidaysStatuss(ids []int64) (*HrHolidaysStatuss, error) {
	hhss := &HrHolidaysStatuss{}
	if err := c.Read(HrHolidaysStatusModel, ids, nil, hhss); err != nil {
		return nil, err
	}
	return hhss, nil
}

// FindHrHolidaysStatus finds hr.holidays.status record by querying it with criteria.
func (c *Client) FindHrHolidaysStatus(criteria *Criteria) (*HrHolidaysStatus, error) {
	hhss := &HrHolidaysStatuss{}
	if err := c.SearchRead(HrHolidaysStatusModel, criteria, NewOptions().Limit(1), hhss); err != nil {
		return nil, err
	}
	if hhss != nil && len(*hhss) > 0 {
		return &((*hhss)[0]), nil
	}
	return nil, fmt.Errorf("hr.holidays.status was not found with criteria %v", criteria)
}

// FindHrHolidaysStatuss finds hr.holidays.status records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrHolidaysStatuss(criteria *Criteria, options *Options) (*HrHolidaysStatuss, error) {
	hhss := &HrHolidaysStatuss{}
	if err := c.SearchRead(HrHolidaysStatusModel, criteria, options, hhss); err != nil {
		return nil, err
	}
	return hhss, nil
}

// FindHrHolidaysStatusIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrHolidaysStatusIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrHolidaysStatusModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrHolidaysStatusId finds record id by querying it with criteria.
func (c *Client) FindHrHolidaysStatusId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrHolidaysStatusModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.holidays.status was not found with criteria %v and options %v", criteria, options)
}
