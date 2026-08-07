package odoo

// HrEmployeeLocation represents hr.employee.location model.
type HrEmployeeLocation struct {
	CreateDate       *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid        *Many2One  `xmlrpc:"create_uid,omitempty"`
	Date             *Time      `xmlrpc:"date,omitempty"`
	DayWeekString    *String    `xmlrpc:"day_week_string,omitempty"`
	DisplayName      *String    `xmlrpc:"display_name,omitempty"`
	EmployeeId       *Many2One  `xmlrpc:"employee_id,omitempty"`
	EmployeeName     *String    `xmlrpc:"employee_name,omitempty"`
	Id               *Int       `xmlrpc:"id,omitempty"`
	WorkLocationId   *Many2One  `xmlrpc:"work_location_id,omitempty"`
	WorkLocationName *String    `xmlrpc:"work_location_name,omitempty"`
	WorkLocationType *Selection `xmlrpc:"work_location_type,omitempty"`
	WriteDate        *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid         *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// HrEmployeeLocations represents array of hr.employee.location model.
type HrEmployeeLocations []HrEmployeeLocation

// HrEmployeeLocationModel is the odoo model name.
const HrEmployeeLocationModel = "hr.employee.location"

// Many2One convert HrEmployeeLocation to *Many2One.
func (hel *HrEmployeeLocation) Many2One() *Many2One {
	return NewMany2One(hel.Id.Get(), "")
}

// CreateHrEmployeeLocation creates a new hr.employee.location model and returns its id.
func (c *Client) CreateHrEmployeeLocation(hel *HrEmployeeLocation) (int64, error) {
	ids, err := c.CreateHrEmployeeLocations([]*HrEmployeeLocation{hel})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrEmployeeLocation creates a new hr.employee.location model and returns its id.
func (c *Client) CreateHrEmployeeLocations(hels []*HrEmployeeLocation) ([]int64, error) {
	var vv []interface{}
	for _, v := range hels {
		vv = append(vv, v)
	}
	return c.Create(HrEmployeeLocationModel, vv, nil)
}

// UpdateHrEmployeeLocation updates an existing hr.employee.location record.
func (c *Client) UpdateHrEmployeeLocation(hel *HrEmployeeLocation) error {
	return c.UpdateHrEmployeeLocations([]int64{hel.Id.Get()}, hel)
}

// UpdateHrEmployeeLocations updates existing hr.employee.location records.
// All records (represented by ids) will be updated by hel values.
func (c *Client) UpdateHrEmployeeLocations(ids []int64, hel *HrEmployeeLocation) error {
	return c.Update(HrEmployeeLocationModel, ids, hel, nil)
}

// DeleteHrEmployeeLocation deletes an existing hr.employee.location record.
func (c *Client) DeleteHrEmployeeLocation(id int64) error {
	return c.DeleteHrEmployeeLocations([]int64{id})
}

// DeleteHrEmployeeLocations deletes existing hr.employee.location records.
func (c *Client) DeleteHrEmployeeLocations(ids []int64) error {
	return c.Delete(HrEmployeeLocationModel, ids)
}

// GetHrEmployeeLocation gets hr.employee.location existing record.
func (c *Client) GetHrEmployeeLocation(id int64) (*HrEmployeeLocation, error) {
	hels, err := c.GetHrEmployeeLocations([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*hels)[0]), nil
}

// GetHrEmployeeLocations gets hr.employee.location existing records.
func (c *Client) GetHrEmployeeLocations(ids []int64) (*HrEmployeeLocations, error) {
	hels := &HrEmployeeLocations{}
	if err := c.Read(HrEmployeeLocationModel, ids, nil, hels); err != nil {
		return nil, err
	}
	return hels, nil
}

// FindHrEmployeeLocation finds hr.employee.location record by querying it with criteria.
func (c *Client) FindHrEmployeeLocation(criteria *Criteria) (*HrEmployeeLocation, error) {
	hels := &HrEmployeeLocations{}
	if err := c.SearchRead(HrEmployeeLocationModel, criteria, NewOptions().Limit(1), hels); err != nil {
		return nil, err
	}
	return &((*hels)[0]), nil
}

// FindHrEmployeeLocations finds hr.employee.location records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrEmployeeLocations(criteria *Criteria, options *Options) (*HrEmployeeLocations, error) {
	hels := &HrEmployeeLocations{}
	if err := c.SearchRead(HrEmployeeLocationModel, criteria, options, hels); err != nil {
		return nil, err
	}
	return hels, nil
}

// FindHrEmployeeLocationIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrEmployeeLocationIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(HrEmployeeLocationModel, criteria, options)
}

// FindHrEmployeeLocationId finds record id by querying it with criteria.
func (c *Client) FindHrEmployeeLocationId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrEmployeeLocationModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
