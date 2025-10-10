package odoo

// HrLeaveReportCalendar represents hr.leave.report.calendar model.
type HrLeaveReportCalendar struct {
	CompanyId       *Many2One  `xmlrpc:"company_id,omitempty"`
	DepartmentId    *Many2One  `xmlrpc:"department_id,omitempty"`
	Description     *String    `xmlrpc:"description,omitempty"`
	DisplayName     *String    `xmlrpc:"display_name,omitempty"`
	Duration        *Float     `xmlrpc:"duration,omitempty"`
	EmployeeId      *Many2One  `xmlrpc:"employee_id,omitempty"`
	HolidayStatusId *Many2One  `xmlrpc:"holiday_status_id,omitempty"`
	Id              *Int       `xmlrpc:"id,omitempty"`
	IsAbsent        *Bool      `xmlrpc:"is_absent,omitempty"`
	IsHatched       *Bool      `xmlrpc:"is_hatched,omitempty"`
	IsManager       *Bool      `xmlrpc:"is_manager,omitempty"`
	IsStriked       *Bool      `xmlrpc:"is_striked,omitempty"`
	JobId           *Many2One  `xmlrpc:"job_id,omitempty"`
	LeaveId         *Many2One  `xmlrpc:"leave_id,omitempty"`
	LeaveManagerId  *Many2One  `xmlrpc:"leave_manager_id,omitempty"`
	Name            *String    `xmlrpc:"name,omitempty"`
	StartDatetime   *Time      `xmlrpc:"start_datetime,omitempty"`
	State           *Selection `xmlrpc:"state,omitempty"`
	StopDatetime    *Time      `xmlrpc:"stop_datetime,omitempty"`
	Tz              *Selection `xmlrpc:"tz,omitempty"`
}

// HrLeaveReportCalendars represents array of hr.leave.report.calendar model.
type HrLeaveReportCalendars []HrLeaveReportCalendar

// HrLeaveReportCalendarModel is the odoo model name.
const HrLeaveReportCalendarModel = "hr.leave.report.calendar"

// Many2One convert HrLeaveReportCalendar to *Many2One.
func (hlrc *HrLeaveReportCalendar) Many2One() *Many2One {
	return NewMany2One(hlrc.Id.Get(), "")
}

// CreateHrLeaveReportCalendar creates a new hr.leave.report.calendar model and returns its id.
func (c *Client) CreateHrLeaveReportCalendar(hlrc *HrLeaveReportCalendar) (int64, error) {
	ids, err := c.CreateHrLeaveReportCalendars([]*HrLeaveReportCalendar{hlrc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrLeaveReportCalendar creates a new hr.leave.report.calendar model and returns its id.
func (c *Client) CreateHrLeaveReportCalendars(hlrcs []*HrLeaveReportCalendar) ([]int64, error) {
	var vv []interface{}
	for _, v := range hlrcs {
		vv = append(vv, v)
	}
	return c.Create(HrLeaveReportCalendarModel, vv, nil)
}

// UpdateHrLeaveReportCalendar updates an existing hr.leave.report.calendar record.
func (c *Client) UpdateHrLeaveReportCalendar(hlrc *HrLeaveReportCalendar) error {
	return c.UpdateHrLeaveReportCalendars([]int64{hlrc.Id.Get()}, hlrc)
}

// UpdateHrLeaveReportCalendars updates existing hr.leave.report.calendar records.
// All records (represented by ids) will be updated by hlrc values.
func (c *Client) UpdateHrLeaveReportCalendars(ids []int64, hlrc *HrLeaveReportCalendar) error {
	return c.Update(HrLeaveReportCalendarModel, ids, hlrc, nil)
}

// DeleteHrLeaveReportCalendar deletes an existing hr.leave.report.calendar record.
func (c *Client) DeleteHrLeaveReportCalendar(id int64) error {
	return c.DeleteHrLeaveReportCalendars([]int64{id})
}

// DeleteHrLeaveReportCalendars deletes existing hr.leave.report.calendar records.
func (c *Client) DeleteHrLeaveReportCalendars(ids []int64) error {
	return c.Delete(HrLeaveReportCalendarModel, ids)
}

// GetHrLeaveReportCalendar gets hr.leave.report.calendar existing record.
func (c *Client) GetHrLeaveReportCalendar(id int64) (*HrLeaveReportCalendar, error) {
	hlrcs, err := c.GetHrLeaveReportCalendars([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*hlrcs)[0]), nil
}

// GetHrLeaveReportCalendars gets hr.leave.report.calendar existing records.
func (c *Client) GetHrLeaveReportCalendars(ids []int64) (*HrLeaveReportCalendars, error) {
	hlrcs := &HrLeaveReportCalendars{}
	if err := c.Read(HrLeaveReportCalendarModel, ids, nil, hlrcs); err != nil {
		return nil, err
	}
	return hlrcs, nil
}

// FindHrLeaveReportCalendar finds hr.leave.report.calendar record by querying it with criteria.
func (c *Client) FindHrLeaveReportCalendar(criteria *Criteria) (*HrLeaveReportCalendar, error) {
	hlrcs := &HrLeaveReportCalendars{}
	if err := c.SearchRead(HrLeaveReportCalendarModel, criteria, NewOptions().Limit(1), hlrcs); err != nil {
		return nil, err
	}
	return &((*hlrcs)[0]), nil
}

// FindHrLeaveReportCalendars finds hr.leave.report.calendar records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrLeaveReportCalendars(criteria *Criteria, options *Options) (*HrLeaveReportCalendars, error) {
	hlrcs := &HrLeaveReportCalendars{}
	if err := c.SearchRead(HrLeaveReportCalendarModel, criteria, options, hlrcs); err != nil {
		return nil, err
	}
	return hlrcs, nil
}

// FindHrLeaveReportCalendarIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrLeaveReportCalendarIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(HrLeaveReportCalendarModel, criteria, options)
}

// FindHrLeaveReportCalendarId finds record id by querying it with criteria.
func (c *Client) FindHrLeaveReportCalendarId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrLeaveReportCalendarModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
