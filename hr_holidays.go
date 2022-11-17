package odoo

import (
	"fmt"
)

// HrHolidays represents hr.holidays model.
type HrHolidays struct {
	LastUpdate               *Time      `xmlrpc:"__last_update,omitempty"`
	CanReset                 *Bool      `xmlrpc:"can_reset,omitempty"`
	CategoryId               *Many2One  `xmlrpc:"category_id,omitempty"`
	CreateDate               *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                *Many2One  `xmlrpc:"create_uid,omitempty"`
	DateFrom                 *Time      `xmlrpc:"date_from,omitempty"`
	DateTo                   *Time      `xmlrpc:"date_to,omitempty"`
	DepartmentId             *Many2One  `xmlrpc:"department_id,omitempty"`
	DisplayName              *String    `xmlrpc:"display_name,omitempty"`
	DoubleValidation         *Bool      `xmlrpc:"double_validation,omitempty"`
	EmployeeId               *Many2One  `xmlrpc:"employee_id,omitempty"`
	FirstApproverId          *Many2One  `xmlrpc:"first_approver_id,omitempty"`
	HolidayStatusId          *Many2One  `xmlrpc:"holiday_status_id,omitempty"`
	HolidayType              *Selection `xmlrpc:"holiday_type,omitempty"`
	Id                       *Int       `xmlrpc:"id,omitempty"`
	LinkedRequestIds         *Relation  `xmlrpc:"linked_request_ids,omitempty"`
	ManagerId                *Many2One  `xmlrpc:"manager_id,omitempty"`
	MeetingId                *Many2One  `xmlrpc:"meeting_id,omitempty"`
	MessageChannelIds        *Relation  `xmlrpc:"message_channel_ids,omitempty"`
	MessageFollowerIds       *Relation  `xmlrpc:"message_follower_ids,omitempty"`
	MessageIds               *Relation  `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower        *Bool      `xmlrpc:"message_is_follower,omitempty"`
	MessageLastPost          *Time      `xmlrpc:"message_last_post,omitempty"`
	MessageNeedaction        *Bool      `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter *Int       `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds        *Relation  `xmlrpc:"message_partner_ids,omitempty"`
	MessageUnread            *Bool      `xmlrpc:"message_unread,omitempty"`
	MessageUnreadCounter     *Int       `xmlrpc:"message_unread_counter,omitempty"`
	Name                     *String    `xmlrpc:"name,omitempty"`
	Notes                    *String    `xmlrpc:"notes,omitempty"`
	NumberOfDays             *Float     `xmlrpc:"number_of_days,omitempty"`
	NumberOfDaysTemp         *Float     `xmlrpc:"number_of_days_temp,omitempty"`
	ParentId                 *Many2One  `xmlrpc:"parent_id,omitempty"`
	PayslipStatus            *Bool      `xmlrpc:"payslip_status,omitempty"`
	ReportNote               *String    `xmlrpc:"report_note,omitempty"`
	SecondApproverId         *Many2One  `xmlrpc:"second_approver_id,omitempty"`
	State                    *Selection `xmlrpc:"state,omitempty"`
	TimesheetIds             *Relation  `xmlrpc:"timesheet_ids,omitempty"`
	Type                     *Selection `xmlrpc:"type,omitempty"`
	UserId                   *Many2One  `xmlrpc:"user_id,omitempty"`
	WebsiteMessageIds        *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                 *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// HrHolidayss represents array of hr.holidays model.
type HrHolidayss []HrHolidays

// HrHolidaysModel is the odoo model name.
const HrHolidaysModel = "hr.holidays"

// Many2One convert HrHolidays to *Many2One.
func (hh *HrHolidays) Many2One() *Many2One {
	return NewMany2One(hh.Id.Get(), "")
}

// CreateHrHolidays creates a new hr.holidays model and returns its id.
func (c *Client) CreateHrHolidays(hh *HrHolidays) (int64, error) {
	return c.Create(HrHolidaysModel, hh)
}

// UpdateHrHolidays updates an existing hr.holidays record.
func (c *Client) UpdateHrHolidays(hh *HrHolidays) error {
	return c.UpdateHrHolidayss([]int64{hh.Id.Get()}, hh)
}

// UpdateHrHolidayss updates existing hr.holidays records.
// All records (represented by ids) will be updated by hh values.
func (c *Client) UpdateHrHolidayss(ids []int64, hh *HrHolidays) error {
	return c.Update(HrHolidaysModel, ids, hh)
}

// DeleteHrHolidays deletes an existing hr.holidays record.
func (c *Client) DeleteHrHolidays(id int64) error {
	return c.DeleteHrHolidayss([]int64{id})
}

// DeleteHrHolidayss deletes existing hr.holidays records.
func (c *Client) DeleteHrHolidayss(ids []int64) error {
	return c.Delete(HrHolidaysModel, ids)
}

// GetHrHolidays gets hr.holidays existing record.
func (c *Client) GetHrHolidays(id int64) (*HrHolidays, error) {
	hhs, err := c.GetHrHolidayss([]int64{id})
	if err != nil {
		return nil, err
	}
	if hhs != nil && len(*hhs) > 0 {
		return &((*hhs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.holidays not found", id)
}

// GetHrHolidayss gets hr.holidays existing records.
func (c *Client) GetHrHolidayss(ids []int64) (*HrHolidayss, error) {
	hhs := &HrHolidayss{}
	if err := c.Read(HrHolidaysModel, ids, nil, hhs); err != nil {
		return nil, err
	}
	return hhs, nil
}

// FindHrHolidays finds hr.holidays record by querying it with criteria.
func (c *Client) FindHrHolidays(criteria *Criteria) (*HrHolidays, error) {
	hhs := &HrHolidayss{}
	if err := c.SearchRead(HrHolidaysModel, criteria, NewOptions().Limit(1), hhs); err != nil {
		return nil, err
	}
	if hhs != nil && len(*hhs) > 0 {
		return &((*hhs)[0]), nil
	}
	return nil, fmt.Errorf("no hr.holidays was found with criteria %v", criteria)
}

// FindHrHolidayss finds hr.holidays records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrHolidayss(criteria *Criteria, options *Options) (*HrHolidayss, error) {
	hhs := &HrHolidayss{}
	if err := c.SearchRead(HrHolidaysModel, criteria, options, hhs); err != nil {
		return nil, err
	}
	return hhs, nil
}

// FindHrHolidaysIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrHolidaysIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrHolidaysModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrHolidaysId finds record id by querying it with criteria.
func (c *Client) FindHrHolidaysId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrHolidaysModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no hr.holidays was found with criteria %v and options %v", criteria, options)
}
