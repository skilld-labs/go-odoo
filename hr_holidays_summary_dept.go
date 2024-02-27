package odoo

// HrHolidaysSummaryDept represents hr.holidays.summary.dept model.
type HrHolidaysSummaryDept struct {
	LastUpdate  *Time      `xmlrpc:"__last_update,omitempty"`
	CreateDate  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One  `xmlrpc:"create_uid,omitempty"`
	DateFrom    *Time      `xmlrpc:"date_from,omitempty"`
	Depts       *Relation  `xmlrpc:"depts,omitempty"`
	DisplayName *String    `xmlrpc:"display_name,omitempty"`
	HolidayType *Selection `xmlrpc:"holiday_type,omitempty"`
	Id          *Int       `xmlrpc:"id,omitempty"`
	WriteDate   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One  `xmlrpc:"write_uid,omitempty"`
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

// CreateHrHolidaysSummaryDepts creates a new hr.holidays.summary.dept model and returns its id.
func (c *Client) CreateHrHolidaysSummaryDepts(hhsds []*HrHolidaysSummaryDept) ([]int64, error) {
	var vv []interface{}
	for _, v := range hhsds {
		vv = append(vv, v)
	}
	return c.Create(HrHolidaysSummaryDeptModel, vv, nil)
}

// UpdateHrHolidaysSummaryDept updates an existing hr.holidays.summary.dept record.
func (c *Client) UpdateHrHolidaysSummaryDept(hhsd *HrHolidaysSummaryDept) error {
	return c.UpdateHrHolidaysSummaryDepts([]int64{hhsd.Id.Get()}, hhsd)
}

// UpdateHrHolidaysSummaryDepts updates existing hr.holidays.summary.dept records.
// All records (represented by ids) will be updated by hhsd values.
func (c *Client) UpdateHrHolidaysSummaryDepts(ids []int64, hhsd *HrHolidaysSummaryDept) error {
	return c.Update(HrHolidaysSummaryDeptModel, ids, hhsd, nil)
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
	return &((*hhsds)[0]), nil
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
	return &((*hhsds)[0]), nil
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
	return c.Search(HrHolidaysSummaryDeptModel, criteria, options)
}

// FindHrHolidaysSummaryDeptId finds record id by querying it with criteria.
func (c *Client) FindHrHolidaysSummaryDeptId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrHolidaysSummaryDeptModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
