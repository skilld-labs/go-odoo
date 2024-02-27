package odoo

// HrEmployee represents hr.employee model.
type HrEmployee struct {
	LastUpdate               *Time      `xmlrpc:"__last_update,omitempty"`
	Active                   *Bool      `xmlrpc:"active,omitempty"`
	AddressHomeId            *Many2One  `xmlrpc:"address_home_id,omitempty"`
	AddressId                *Many2One  `xmlrpc:"address_id,omitempty"`
	BankAccountId            *Many2One  `xmlrpc:"bank_account_id,omitempty"`
	Birthday                 *Time      `xmlrpc:"birthday,omitempty"`
	CategoryIds              *Relation  `xmlrpc:"category_ids,omitempty"`
	ChildAllCount            *Int       `xmlrpc:"child_all_count,omitempty"`
	ChildIds                 *Relation  `xmlrpc:"child_ids,omitempty"`
	CoachId                  *Many2One  `xmlrpc:"coach_id,omitempty"`
	Color                    *Int       `xmlrpc:"color,omitempty"`
	CompanyId                *Many2One  `xmlrpc:"company_id,omitempty"`
	CountryId                *Many2One  `xmlrpc:"country_id,omitempty"`
	CreateDate               *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrencyId               *Many2One  `xmlrpc:"currency_id,omitempty"`
	CurrentLeaveId           *Many2One  `xmlrpc:"current_leave_id,omitempty"`
	CurrentLeaveState        *Selection `xmlrpc:"current_leave_state,omitempty"`
	DepartmentId             *Many2One  `xmlrpc:"department_id,omitempty"`
	DisplayName              *String    `xmlrpc:"display_name,omitempty"`
	Gender                   *Selection `xmlrpc:"gender,omitempty"`
	Id                       *Int       `xmlrpc:"id,omitempty"`
	IdentificationId         *String    `xmlrpc:"identification_id,omitempty"`
	Image                    *String    `xmlrpc:"image,omitempty"`
	ImageMedium              *String    `xmlrpc:"image_medium,omitempty"`
	ImageSmall               *String    `xmlrpc:"image_small,omitempty"`
	IsAbsentTotay            *Bool      `xmlrpc:"is_absent_totay,omitempty"`
	IsAddressHomeACompany    *Bool      `xmlrpc:"is_address_home_a_company,omitempty"`
	JobId                    *Many2One  `xmlrpc:"job_id,omitempty"`
	LeaveDateFrom            *Time      `xmlrpc:"leave_date_from,omitempty"`
	LeaveDateTo              *Time      `xmlrpc:"leave_date_to,omitempty"`
	LeavesCount              *Float     `xmlrpc:"leaves_count,omitempty"`
	Marital                  *Selection `xmlrpc:"marital,omitempty"`
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
	MobilePhone              *String    `xmlrpc:"mobile_phone,omitempty"`
	Name                     *String    `xmlrpc:"name,omitempty"`
	Notes                    *String    `xmlrpc:"notes,omitempty"`
	ParentId                 *Many2One  `xmlrpc:"parent_id,omitempty"`
	PassportId               *String    `xmlrpc:"passport_id,omitempty"`
	PermitNo                 *String    `xmlrpc:"permit_no,omitempty"`
	RemainingLeaves          *Float     `xmlrpc:"remaining_leaves,omitempty"`
	ResourceCalendarId       *Many2One  `xmlrpc:"resource_calendar_id,omitempty"`
	ResourceId               *Many2One  `xmlrpc:"resource_id,omitempty"`
	ShowLeaves               *Bool      `xmlrpc:"show_leaves,omitempty"`
	Sinid                    *String    `xmlrpc:"sinid,omitempty"`
	Ssnid                    *String    `xmlrpc:"ssnid,omitempty"`
	TimesheetCost            *Float     `xmlrpc:"timesheet_cost,omitempty"`
	UserId                   *Many2One  `xmlrpc:"user_id,omitempty"`
	VisaExpire               *Time      `xmlrpc:"visa_expire,omitempty"`
	VisaNo                   *String    `xmlrpc:"visa_no,omitempty"`
	WebsiteMessageIds        *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WorkEmail                *String    `xmlrpc:"work_email,omitempty"`
	WorkLocation             *String    `xmlrpc:"work_location,omitempty"`
	WorkPhone                *String    `xmlrpc:"work_phone,omitempty"`
	WriteDate                *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                 *Many2One  `xmlrpc:"write_uid,omitempty"`
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
	ids, err := c.CreateHrEmployees([]*HrEmployee{he})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrEmployee creates a new hr.employee model and returns its id.
func (c *Client) CreateHrEmployees(hes []*HrEmployee) ([]int64, error) {
	var vv []interface{}
	for _, v := range hes {
		vv = append(vv, v)
	}
	return c.Create(HrEmployeeModel, vv, nil)
}

// UpdateHrEmployee updates an existing hr.employee record.
func (c *Client) UpdateHrEmployee(he *HrEmployee) error {
	return c.UpdateHrEmployees([]int64{he.Id.Get()}, he)
}

// UpdateHrEmployees updates existing hr.employee records.
// All records (represented by ids) will be updated by he values.
func (c *Client) UpdateHrEmployees(ids []int64, he *HrEmployee) error {
	return c.Update(HrEmployeeModel, ids, he, nil)
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
	return &((*hes)[0]), nil
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
	return &((*hes)[0]), nil
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
	return c.Search(HrEmployeeModel, criteria, options)
}

// FindHrEmployeeId finds record id by querying it with criteria.
func (c *Client) FindHrEmployeeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrEmployeeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
