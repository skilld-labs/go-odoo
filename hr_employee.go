package odoo

import (
	"fmt"
)

// HrEmployee represents hr.employee model.
type HrEmployee struct {
	LastUpdate               *Time      `xmlrpc:"__last_update,omptempty"`
	Active                   *Bool      `xmlrpc:"active,omptempty"`
	AddressHomeId            *Many2One  `xmlrpc:"address_home_id,omptempty"`
	AddressId                *Many2One  `xmlrpc:"address_id,omptempty"`
	BankAccountId            *Many2One  `xmlrpc:"bank_account_id,omptempty"`
	Birthday                 *Time      `xmlrpc:"birthday,omptempty"`
	CategoryIds              *Relation  `xmlrpc:"category_ids,omptempty"`
	ChildAllCount            *Int       `xmlrpc:"child_all_count,omptempty"`
	ChildIds                 *Relation  `xmlrpc:"child_ids,omptempty"`
	CoachId                  *Many2One  `xmlrpc:"coach_id,omptempty"`
	Color                    *Int       `xmlrpc:"color,omptempty"`
	CompanyId                *Many2One  `xmlrpc:"company_id,omptempty"`
	CountryId                *Many2One  `xmlrpc:"country_id,omptempty"`
	CreateDate               *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrencyId               *Many2One  `xmlrpc:"currency_id,omptempty"`
	CurrentLeaveId           *Many2One  `xmlrpc:"current_leave_id,omptempty"`
	CurrentLeaveState        *Selection `xmlrpc:"current_leave_state,omptempty"`
	DepartmentId             *Many2One  `xmlrpc:"department_id,omptempty"`
	DisplayName              *String    `xmlrpc:"display_name,omptempty"`
	Gender                   *Selection `xmlrpc:"gender,omptempty"`
	Id                       *Int       `xmlrpc:"id,omptempty"`
	IdentificationId         *String    `xmlrpc:"identification_id,omptempty"`
	Image                    *String    `xmlrpc:"image,omptempty"`
	ImageMedium              *String    `xmlrpc:"image_medium,omptempty"`
	ImageSmall               *String    `xmlrpc:"image_small,omptempty"`
	IsAbsentTotay            *Bool      `xmlrpc:"is_absent_totay,omptempty"`
	IsAddressHomeACompany    *Bool      `xmlrpc:"is_address_home_a_company,omptempty"`
	JobId                    *Many2One  `xmlrpc:"job_id,omptempty"`
	LeaveDateFrom            *Time      `xmlrpc:"leave_date_from,omptempty"`
	LeaveDateTo              *Time      `xmlrpc:"leave_date_to,omptempty"`
	LeavesCount              *Float     `xmlrpc:"leaves_count,omptempty"`
	Marital                  *Selection `xmlrpc:"marital,omptempty"`
	MessageChannelIds        *Relation  `xmlrpc:"message_channel_ids,omptempty"`
	MessageFollowerIds       *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageIds               *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower        *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageLastPost          *Time      `xmlrpc:"message_last_post,omptempty"`
	MessageNeedaction        *Bool      `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter *Int       `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds        *Relation  `xmlrpc:"message_partner_ids,omptempty"`
	MessageUnread            *Bool      `xmlrpc:"message_unread,omptempty"`
	MessageUnreadCounter     *Int       `xmlrpc:"message_unread_counter,omptempty"`
	MobilePhone              *String    `xmlrpc:"mobile_phone,omptempty"`
	Name                     *String    `xmlrpc:"name,omptempty"`
	Notes                    *String    `xmlrpc:"notes,omptempty"`
	ParentId                 *Many2One  `xmlrpc:"parent_id,omptempty"`
	PassportId               *String    `xmlrpc:"passport_id,omptempty"`
	PermitNo                 *String    `xmlrpc:"permit_no,omptempty"`
	RemainingLeaves          *Float     `xmlrpc:"remaining_leaves,omptempty"`
	ResourceCalendarId       *Many2One  `xmlrpc:"resource_calendar_id,omptempty"`
	ResourceId               *Many2One  `xmlrpc:"resource_id,omptempty"`
	ShowLeaves               *Bool      `xmlrpc:"show_leaves,omptempty"`
	Sinid                    *String    `xmlrpc:"sinid,omptempty"`
	Ssnid                    *String    `xmlrpc:"ssnid,omptempty"`
	TimesheetCost            *Float     `xmlrpc:"timesheet_cost,omptempty"`
	UserId                   *Many2One  `xmlrpc:"user_id,omptempty"`
	VisaExpire               *Time      `xmlrpc:"visa_expire,omptempty"`
	VisaNo                   *String    `xmlrpc:"visa_no,omptempty"`
	WebsiteMessageIds        *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WorkEmail                *String    `xmlrpc:"work_email,omptempty"`
	WorkLocation             *String    `xmlrpc:"work_location,omptempty"`
	WorkPhone                *String    `xmlrpc:"work_phone,omptempty"`
	WriteDate                *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                 *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// HrEmployees represents array of hr.employee model.
type HrEmployees []HrEmployee

// HrEmployeeModel is the odoo model name.
const HrEmployeeModel = "hr.employee"

// Many2One convert HrEmployee to *Many2One.
func (he *HrEmployee) Many2One() *Many2One {
	return NewMany2One(he.Id.Get(), "")
}

// CreateHrEmployee creates a new hr.employee model and returns its id.
func (c *Client) CreateHrEmployee(he *HrEmployee) (int64, error) {
	return c.Create(HrEmployeeModel, he)
}

// UpdateHrEmployee updates an existing hr.employee record.
func (c *Client) UpdateHrEmployee(he *HrEmployee) error {
	return c.UpdateHrEmployees([]int64{he.Id.Get()}, he)
}

// UpdateHrEmployees updates existing hr.employee records.
// All records (represented by ids) will be updated by he values.
func (c *Client) UpdateHrEmployees(ids []int64, he *HrEmployee) error {
	return c.Update(HrEmployeeModel, ids, he)
}

// DeleteHrEmployee deletes an existing hr.employee record.
func (c *Client) DeleteHrEmployee(id int64) error {
	return c.DeleteHrEmployees([]int64{id})
}

// DeleteHrEmployees deletes existing hr.employee records.
func (c *Client) DeleteHrEmployees(ids []int64) error {
	return c.Delete(HrEmployeeModel, ids)
}

// GetHrEmployee gets hr.employee existing record.
func (c *Client) GetHrEmployee(id int64) (*HrEmployee, error) {
	hes, err := c.GetHrEmployees([]int64{id})
	if err != nil {
		return nil, err
	}
	if hes != nil && len(*hes) > 0 {
		return &((*hes)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.employee not found", id)
}

// GetHrEmployees gets hr.employee existing records.
func (c *Client) GetHrEmployees(ids []int64) (*HrEmployees, error) {
	hes := &HrEmployees{}
	if err := c.Read(HrEmployeeModel, ids, nil, hes); err != nil {
		return nil, err
	}
	return hes, nil
}

// FindHrEmployee finds hr.employee record by querying it with criteria.
func (c *Client) FindHrEmployee(criteria *Criteria) (*HrEmployee, error) {
	hes := &HrEmployees{}
	if err := c.SearchRead(HrEmployeeModel, criteria, NewOptions().Limit(1), hes); err != nil {
		return nil, err
	}
	if hes != nil && len(*hes) > 0 {
		return &((*hes)[0]), nil
	}
	return nil, fmt.Errorf("hr.employee was not found with criteria %v", criteria)
}

// FindHrEmployees finds hr.employee records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrEmployees(criteria *Criteria, options *Options) (*HrEmployees, error) {
	hes := &HrEmployees{}
	if err := c.SearchRead(HrEmployeeModel, criteria, options, hes); err != nil {
		return nil, err
	}
	return hes, nil
}

// FindHrEmployeeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrEmployeeIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrEmployeeModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrEmployeeId finds record id by querying it with criteria.
func (c *Client) FindHrEmployeeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrEmployeeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.employee was not found with criteria %v and options %v", criteria, options)
}
