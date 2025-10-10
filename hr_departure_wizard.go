package odoo

// HrDepartureWizard represents hr.departure.wizard model.
type HrDepartureWizard struct {
	CreateDate           *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid            *Many2One `xmlrpc:"create_uid,omitempty"`
	DepartureDate        *Time     `xmlrpc:"departure_date,omitempty"`
	DepartureDescription *String   `xmlrpc:"departure_description,omitempty"`
	DepartureReasonId    *Many2One `xmlrpc:"departure_reason_id,omitempty"`
	DisplayName          *String   `xmlrpc:"display_name,omitempty"`
	EmployeeId           *Many2One `xmlrpc:"employee_id,omitempty"`
	Id                   *Int      `xmlrpc:"id,omitempty"`
	WriteDate            *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid             *Many2One `xmlrpc:"write_uid,omitempty"`
}

// HrDepartureWizards represents array of hr.departure.wizard model.
type HrDepartureWizards []HrDepartureWizard

// HrDepartureWizardModel is the odoo model name.
const HrDepartureWizardModel = "hr.departure.wizard"

// Many2One convert HrDepartureWizard to *Many2One.
func (hdw *HrDepartureWizard) Many2One() *Many2One {
	return NewMany2One(hdw.Id.Get(), "")
}

// CreateHrDepartureWizard creates a new hr.departure.wizard model and returns its id.
func (c *Client) CreateHrDepartureWizard(hdw *HrDepartureWizard) (int64, error) {
	ids, err := c.CreateHrDepartureWizards([]*HrDepartureWizard{hdw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrDepartureWizard creates a new hr.departure.wizard model and returns its id.
func (c *Client) CreateHrDepartureWizards(hdws []*HrDepartureWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range hdws {
		vv = append(vv, v)
	}
	return c.Create(HrDepartureWizardModel, vv, nil)
}

// UpdateHrDepartureWizard updates an existing hr.departure.wizard record.
func (c *Client) UpdateHrDepartureWizard(hdw *HrDepartureWizard) error {
	return c.UpdateHrDepartureWizards([]int64{hdw.Id.Get()}, hdw)
}

// UpdateHrDepartureWizards updates existing hr.departure.wizard records.
// All records (represented by ids) will be updated by hdw values.
func (c *Client) UpdateHrDepartureWizards(ids []int64, hdw *HrDepartureWizard) error {
	return c.Update(HrDepartureWizardModel, ids, hdw, nil)
}

// DeleteHrDepartureWizard deletes an existing hr.departure.wizard record.
func (c *Client) DeleteHrDepartureWizard(id int64) error {
	return c.DeleteHrDepartureWizards([]int64{id})
}

// DeleteHrDepartureWizards deletes existing hr.departure.wizard records.
func (c *Client) DeleteHrDepartureWizards(ids []int64) error {
	return c.Delete(HrDepartureWizardModel, ids)
}

// GetHrDepartureWizard gets hr.departure.wizard existing record.
func (c *Client) GetHrDepartureWizard(id int64) (*HrDepartureWizard, error) {
	hdws, err := c.GetHrDepartureWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*hdws)[0]), nil
}

// GetHrDepartureWizards gets hr.departure.wizard existing records.
func (c *Client) GetHrDepartureWizards(ids []int64) (*HrDepartureWizards, error) {
	hdws := &HrDepartureWizards{}
	if err := c.Read(HrDepartureWizardModel, ids, nil, hdws); err != nil {
		return nil, err
	}
	return hdws, nil
}

// FindHrDepartureWizard finds hr.departure.wizard record by querying it with criteria.
func (c *Client) FindHrDepartureWizard(criteria *Criteria) (*HrDepartureWizard, error) {
	hdws := &HrDepartureWizards{}
	if err := c.SearchRead(HrDepartureWizardModel, criteria, NewOptions().Limit(1), hdws); err != nil {
		return nil, err
	}
	return &((*hdws)[0]), nil
}

// FindHrDepartureWizards finds hr.departure.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrDepartureWizards(criteria *Criteria, options *Options) (*HrDepartureWizards, error) {
	hdws := &HrDepartureWizards{}
	if err := c.SearchRead(HrDepartureWizardModel, criteria, options, hdws); err != nil {
		return nil, err
	}
	return hdws, nil
}

// FindHrDepartureWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrDepartureWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(HrDepartureWizardModel, criteria, options)
}

// FindHrDepartureWizardId finds record id by querying it with criteria.
func (c *Client) FindHrDepartureWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrDepartureWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
