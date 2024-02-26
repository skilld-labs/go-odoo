package odoo

// HrHolidays represents hr.holidays model.
type HrHolidays struct {
	LastUpdate               *Time      `xmlrpc:"__last_update,omptempty"`
	CanReset                 *Bool      `xmlrpc:"can_reset,omptempty"`
	CategoryId               *Many2One  `xmlrpc:"category_id,omptempty"`
	CreateDate               *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                *Many2One  `xmlrpc:"create_uid,omptempty"`
	DateFrom                 *Time      `xmlrpc:"date_from,omptempty"`
	DateTo                   *Time      `xmlrpc:"date_to,omptempty"`
	DepartmentId             *Many2One  `xmlrpc:"department_id,omptempty"`
	DisplayName              *String    `xmlrpc:"display_name,omptempty"`
	DoubleValidation         *Bool      `xmlrpc:"double_validation,omptempty"`
	EmployeeId               *Many2One  `xmlrpc:"employee_id,omptempty"`
	FirstApproverId          *Many2One  `xmlrpc:"first_approver_id,omptempty"`
	HolidayStatusId          *Many2One  `xmlrpc:"holiday_status_id,omptempty"`
	HolidayType              *Selection `xmlrpc:"holiday_type,omptempty"`
	Id                       *Int       `xmlrpc:"id,omptempty"`
	LinkedRequestIds         *Relation  `xmlrpc:"linked_request_ids,omptempty"`
	ManagerId                *Many2One  `xmlrpc:"manager_id,omptempty"`
	MeetingId                *Many2One  `xmlrpc:"meeting_id,omptempty"`
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
	Name                     *String    `xmlrpc:"name,omptempty"`
	Notes                    *String    `xmlrpc:"notes,omptempty"`
	NumberOfDays             *Float     `xmlrpc:"number_of_days,omptempty"`
	NumberOfDaysTemp         *Float     `xmlrpc:"number_of_days_temp,omptempty"`
	ParentId                 *Many2One  `xmlrpc:"parent_id,omptempty"`
	PayslipStatus            *Bool      `xmlrpc:"payslip_status,omptempty"`
	ReportNote               *String    `xmlrpc:"report_note,omptempty"`
	SecondApproverId         *Many2One  `xmlrpc:"second_approver_id,omptempty"`
	State                    *Selection `xmlrpc:"state,omptempty"`
	TimesheetIds             *Relation  `xmlrpc:"timesheet_ids,omptempty"`
	Type                     *Selection `xmlrpc:"type,omptempty"`
	UserId                   *Many2One  `xmlrpc:"user_id,omptempty"`
	WebsiteMessageIds        *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                 *Many2One  `xmlrpc:"write_uid,omptempty"`
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
	ids, err := c.CreateHrHolidayss([]*HrHolidays{hh})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrHolidays creates a new hr.holidays model and returns its id.
func (c *Client) CreateHrHolidayss(hhs []*HrHolidays) ([]int64, error) {
	var vv []interface{}
	for _, v := range hhs {
		vv = append(vv, v)
	}
	return c.Create(HrHolidaysModel, vv, nil)
}

// UpdateHrHolidays updates an existing hr.holidays record.
func (c *Client) UpdateHrHolidays(hh *HrHolidays) error {
	return c.UpdateHrHolidayss([]int64{hh.Id.Get()}, hh)
}

// UpdateHrHolidayss updates existing hr.holidays records.
// All records (represented by ids) will be updated by hh values.
func (c *Client) UpdateHrHolidayss(ids []int64, hh *HrHolidays) error {
	return c.Update(HrHolidaysModel, ids, hh, nil)
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
	return &((*hhs)[0]), nil
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
	return &((*hhs)[0]), nil
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
	return c.Search(HrHolidaysModel, criteria, options)
}

// FindHrHolidaysId finds record id by querying it with criteria.
func (c *Client) FindHrHolidaysId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrHolidaysModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
