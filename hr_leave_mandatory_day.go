package odoo

// HrLeaveMandatoryDay represents hr.leave.mandatory.day model.
type HrLeaveMandatoryDay struct {
	Color              *Int      `xmlrpc:"color,omitempty"`
	CompanyId          *Many2One `xmlrpc:"company_id,omitempty"`
	CreateDate         *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid          *Many2One `xmlrpc:"create_uid,omitempty"`
	DepartmentIds      *Relation `xmlrpc:"department_ids,omitempty"`
	DisplayName        *String   `xmlrpc:"display_name,omitempty"`
	EndDate            *Time     `xmlrpc:"end_date,omitempty"`
	Id                 *Int      `xmlrpc:"id,omitempty"`
	Name               *String   `xmlrpc:"name,omitempty"`
	ResourceCalendarId *Many2One `xmlrpc:"resource_calendar_id,omitempty"`
	StartDate          *Time     `xmlrpc:"start_date,omitempty"`
	WriteDate          *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid           *Many2One `xmlrpc:"write_uid,omitempty"`
}

// HrLeaveMandatoryDays represents array of hr.leave.mandatory.day model.
type HrLeaveMandatoryDays []HrLeaveMandatoryDay

// HrLeaveMandatoryDayModel is the odoo model name.
const HrLeaveMandatoryDayModel = "hr.leave.mandatory.day"

// Many2One convert HrLeaveMandatoryDay to *Many2One.
func (hlmd *HrLeaveMandatoryDay) Many2One() *Many2One {
	return NewMany2One(hlmd.Id.Get(), "")
}

// CreateHrLeaveMandatoryDay creates a new hr.leave.mandatory.day model and returns its id.
func (c *Client) CreateHrLeaveMandatoryDay(hlmd *HrLeaveMandatoryDay) (int64, error) {
	ids, err := c.CreateHrLeaveMandatoryDays([]*HrLeaveMandatoryDay{hlmd})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrLeaveMandatoryDay creates a new hr.leave.mandatory.day model and returns its id.
func (c *Client) CreateHrLeaveMandatoryDays(hlmds []*HrLeaveMandatoryDay) ([]int64, error) {
	var vv []interface{}
	for _, v := range hlmds {
		vv = append(vv, v)
	}
	return c.Create(HrLeaveMandatoryDayModel, vv, nil)
}

// UpdateHrLeaveMandatoryDay updates an existing hr.leave.mandatory.day record.
func (c *Client) UpdateHrLeaveMandatoryDay(hlmd *HrLeaveMandatoryDay) error {
	return c.UpdateHrLeaveMandatoryDays([]int64{hlmd.Id.Get()}, hlmd)
}

// UpdateHrLeaveMandatoryDays updates existing hr.leave.mandatory.day records.
// All records (represented by ids) will be updated by hlmd values.
func (c *Client) UpdateHrLeaveMandatoryDays(ids []int64, hlmd *HrLeaveMandatoryDay) error {
	return c.Update(HrLeaveMandatoryDayModel, ids, hlmd, nil)
}

// DeleteHrLeaveMandatoryDay deletes an existing hr.leave.mandatory.day record.
func (c *Client) DeleteHrLeaveMandatoryDay(id int64) error {
	return c.DeleteHrLeaveMandatoryDays([]int64{id})
}

// DeleteHrLeaveMandatoryDays deletes existing hr.leave.mandatory.day records.
func (c *Client) DeleteHrLeaveMandatoryDays(ids []int64) error {
	return c.Delete(HrLeaveMandatoryDayModel, ids)
}

// GetHrLeaveMandatoryDay gets hr.leave.mandatory.day existing record.
func (c *Client) GetHrLeaveMandatoryDay(id int64) (*HrLeaveMandatoryDay, error) {
	hlmds, err := c.GetHrLeaveMandatoryDays([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*hlmds)[0]), nil
}

// GetHrLeaveMandatoryDays gets hr.leave.mandatory.day existing records.
func (c *Client) GetHrLeaveMandatoryDays(ids []int64) (*HrLeaveMandatoryDays, error) {
	hlmds := &HrLeaveMandatoryDays{}
	if err := c.Read(HrLeaveMandatoryDayModel, ids, nil, hlmds); err != nil {
		return nil, err
	}
	return hlmds, nil
}

// FindHrLeaveMandatoryDay finds hr.leave.mandatory.day record by querying it with criteria.
func (c *Client) FindHrLeaveMandatoryDay(criteria *Criteria) (*HrLeaveMandatoryDay, error) {
	hlmds := &HrLeaveMandatoryDays{}
	if err := c.SearchRead(HrLeaveMandatoryDayModel, criteria, NewOptions().Limit(1), hlmds); err != nil {
		return nil, err
	}
	return &((*hlmds)[0]), nil
}

// FindHrLeaveMandatoryDays finds hr.leave.mandatory.day records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrLeaveMandatoryDays(criteria *Criteria, options *Options) (*HrLeaveMandatoryDays, error) {
	hlmds := &HrLeaveMandatoryDays{}
	if err := c.SearchRead(HrLeaveMandatoryDayModel, criteria, options, hlmds); err != nil {
		return nil, err
	}
	return hlmds, nil
}

// FindHrLeaveMandatoryDayIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrLeaveMandatoryDayIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(HrLeaveMandatoryDayModel, criteria, options)
}

// FindHrLeaveMandatoryDayId finds record id by querying it with criteria.
func (c *Client) FindHrLeaveMandatoryDayId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrLeaveMandatoryDayModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
