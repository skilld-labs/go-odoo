package odoo

import (
	"fmt"
)

// HrHolidaysSummaryDept represents hr.holidays.summary.dept model.
type HrHolidaysSummaryDept struct {
	LastUpdate  *Time      `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One  `xmlrpc:"create_uid,omptempty"`
	DateFrom    *Time      `xmlrpc:"date_from,omptempty"`
	Depts       *Relation  `xmlrpc:"depts,omptempty"`
	DisplayName *String    `xmlrpc:"display_name,omptempty"`
	HolidayType *Selection `xmlrpc:"holiday_type,omptempty"`
	Id          *Int       `xmlrpc:"id,omptempty"`
	WriteDate   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// HrHolidaysSummaryDepts represents array of hr.holidays.summary.dept model.
type HrHolidaysSummaryDepts []HrHolidaysSummaryDept

// HrHolidaysSummaryDeptModel is the odoo model name.
const HrHolidaysSummaryDeptModel = "hr.holidays.summary.dept"

// Many2One convert HrHolidaysSummaryDept to *Many2One.
func (hhsd *HrHolidaysSummaryDept) Many2One() *Many2One {
	return NewMany2One(hhsd.Id.Get(), "")
}

// CreateHrHolidaysSummaryDept creates a new hr.holidays.summary.dept model and returns its id.
func (c *Client) CreateHrHolidaysSummaryDept(hhsd *HrHolidaysSummaryDept) (int64, error) {
	ids, err := c.CreateHrHolidaysSummaryDepts([]*HrHolidaysSummaryDept{hhsd})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrHolidaysSummaryDept creates a new hr.holidays.summary.dept model and returns its id.
func (c *Client) CreateHrHolidaysSummaryDepts(hhsds []*HrHolidaysSummaryDept) ([]int64, error) {
	var vv []interface{}
	for _, v := range hhsds {
		vv = append(vv, v)
	}
	return c.Create(HrHolidaysSummaryDeptModel, vv)
}

// UpdateHrHolidaysSummaryDept updates an existing hr.holidays.summary.dept record.
func (c *Client) UpdateHrHolidaysSummaryDept(hhsd *HrHolidaysSummaryDept) error {
	return c.UpdateHrHolidaysSummaryDepts([]int64{hhsd.Id.Get()}, hhsd)
}

// UpdateHrHolidaysSummaryDepts updates existing hr.holidays.summary.dept records.
// All records (represented by ids) will be updated by hhsd values.
func (c *Client) UpdateHrHolidaysSummaryDepts(ids []int64, hhsd *HrHolidaysSummaryDept) error {
	return c.Update(HrHolidaysSummaryDeptModel, ids, hhsd)
}

// DeleteHrHolidaysSummaryDept deletes an existing hr.holidays.summary.dept record.
func (c *Client) DeleteHrHolidaysSummaryDept(id int64) error {
	return c.DeleteHrHolidaysSummaryDepts([]int64{id})
}

// DeleteHrHolidaysSummaryDepts deletes existing hr.holidays.summary.dept records.
func (c *Client) DeleteHrHolidaysSummaryDepts(ids []int64) error {
	return c.Delete(HrHolidaysSummaryDeptModel, ids)
}

// GetHrHolidaysSummaryDept gets hr.holidays.summary.dept existing record.
func (c *Client) GetHrHolidaysSummaryDept(id int64) (*HrHolidaysSummaryDept, error) {
	hhsds, err := c.GetHrHolidaysSummaryDepts([]int64{id})
	if err != nil {
		return nil, err
	}
	if hhsds != nil && len(*hhsds) > 0 {
		return &((*hhsds)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.holidays.summary.dept not found", id)
}

// GetHrHolidaysSummaryDepts gets hr.holidays.summary.dept existing records.
func (c *Client) GetHrHolidaysSummaryDepts(ids []int64) (*HrHolidaysSummaryDepts, error) {
	hhsds := &HrHolidaysSummaryDepts{}
	if err := c.Read(HrHolidaysSummaryDeptModel, ids, nil, hhsds); err != nil {
		return nil, err
	}
	return hhsds, nil
}

// FindHrHolidaysSummaryDept finds hr.holidays.summary.dept record by querying it with criteria.
func (c *Client) FindHrHolidaysSummaryDept(criteria *Criteria) (*HrHolidaysSummaryDept, error) {
	hhsds := &HrHolidaysSummaryDepts{}
	if err := c.SearchRead(HrHolidaysSummaryDeptModel, criteria, NewOptions().Limit(1), hhsds); err != nil {
		return nil, err
	}
	if hhsds != nil && len(*hhsds) > 0 {
		return &((*hhsds)[0]), nil
	}
	return nil, fmt.Errorf("hr.holidays.summary.dept was not found with criteria %v", criteria)
}

// FindHrHolidaysSummaryDepts finds hr.holidays.summary.dept records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrHolidaysSummaryDepts(criteria *Criteria, options *Options) (*HrHolidaysSummaryDepts, error) {
	hhsds := &HrHolidaysSummaryDepts{}
	if err := c.SearchRead(HrHolidaysSummaryDeptModel, criteria, options, hhsds); err != nil {
		return nil, err
	}
	return hhsds, nil
}

// FindHrHolidaysSummaryDeptIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrHolidaysSummaryDeptIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrHolidaysSummaryDeptModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrHolidaysSummaryDeptId finds record id by querying it with criteria.
func (c *Client) FindHrHolidaysSummaryDeptId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrHolidaysSummaryDeptModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.holidays.summary.dept was not found with criteria %v and options %v", criteria, options)
}
