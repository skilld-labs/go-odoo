package odoo

// HrHolidaysStatus represents hr.holidays.status model.
type HrHolidaysStatus struct {
	LastUpdate             *Time      `xmlrpc:"__last_update,omitempty"`
	Active                 *Bool      `xmlrpc:"active,omitempty"`
	CategId                *Many2One  `xmlrpc:"categ_id,omitempty"`
	ColorName              *Selection `xmlrpc:"color_name,omitempty"`
	CompanyId              *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate             *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid              *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName            *String    `xmlrpc:"display_name,omitempty"`
	DoubleValidation       *Bool      `xmlrpc:"double_validation,omitempty"`
	Id                     *Int       `xmlrpc:"id,omitempty"`
	LeavesTaken            *Float     `xmlrpc:"leaves_taken,omitempty"`
	Limit                  *Bool      `xmlrpc:"limit,omitempty"`
	MaxLeaves              *Float     `xmlrpc:"max_leaves,omitempty"`
	Name                   *String    `xmlrpc:"name,omitempty"`
	RemainingLeaves        *Float     `xmlrpc:"remaining_leaves,omitempty"`
	TimesheetGenerate      *Bool      `xmlrpc:"timesheet_generate,omitempty"`
	TimesheetProjectId     *Many2One  `xmlrpc:"timesheet_project_id,omitempty"`
	TimesheetTaskId        *Many2One  `xmlrpc:"timesheet_task_id,omitempty"`
	VirtualRemainingLeaves *Float     `xmlrpc:"virtual_remaining_leaves,omitempty"`
	WriteDate              *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid               *Many2One  `xmlrpc:"write_uid,omitempty"`
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
	ids, err := c.CreateHrHolidaysStatuss([]*HrHolidaysStatus{hhs})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrHolidaysStatuss creates a new hr.holidays.status model and returns its id.
func (c *Client) CreateHrHolidaysStatuss(hhss []*HrHolidaysStatus) ([]int64, error) {
	var vv []interface{}
	for _, v := range hhss {
		vv = append(vv, v)
	}
	return c.Create(HrHolidaysStatusModel, vv, nil)
}

// UpdateHrHolidaysStatus updates an existing hr.holidays.status record.
func (c *Client) UpdateHrHolidaysStatus(hhs *HrHolidaysStatus) error {
	return c.UpdateHrHolidaysStatuss([]int64{hhs.Id.Get()}, hhs)
}

// UpdateHrHolidaysStatuss updates existing hr.holidays.status records.
// All records (represented by ids) will be updated by hhs values.
func (c *Client) UpdateHrHolidaysStatuss(ids []int64, hhs *HrHolidaysStatus) error {
	return c.Update(HrHolidaysStatusModel, ids, hhs, nil)
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
	return &((*hhss)[0]), nil
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
	return &((*hhss)[0]), nil
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
	return c.Search(HrHolidaysStatusModel, criteria, options)
}

// FindHrHolidaysStatusId finds record id by querying it with criteria.
func (c *Client) FindHrHolidaysStatusId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrHolidaysStatusModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
