package odoo

// HrEmployeeDeleteWizard represents hr.employee.delete.wizard model.
type HrEmployeeDeleteWizard struct {
	CreateDate        *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid         *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName       *String   `xmlrpc:"display_name,omitempty"`
	EmployeeIds       *Relation `xmlrpc:"employee_ids,omitempty"`
	HasActiveEmployee *Bool     `xmlrpc:"has_active_employee,omitempty"`
	HasTimesheet      *Bool     `xmlrpc:"has_timesheet,omitempty"`
	Id                *Int      `xmlrpc:"id,omitempty"`
	WriteDate         *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid          *Many2One `xmlrpc:"write_uid,omitempty"`
}

// HrEmployeeDeleteWizards represents array of hr.employee.delete.wizard model.
type HrEmployeeDeleteWizards []HrEmployeeDeleteWizard

// HrEmployeeDeleteWizardModel is the odoo model name.
const HrEmployeeDeleteWizardModel = "hr.employee.delete.wizard"

// Many2One convert HrEmployeeDeleteWizard to *Many2One.
func (hedw *HrEmployeeDeleteWizard) Many2One() *Many2One {
	return NewMany2One(hedw.Id.Get(), "")
}

// CreateHrEmployeeDeleteWizard creates a new hr.employee.delete.wizard model and returns its id.
func (c *Client) CreateHrEmployeeDeleteWizard(hedw *HrEmployeeDeleteWizard) (int64, error) {
	ids, err := c.CreateHrEmployeeDeleteWizards([]*HrEmployeeDeleteWizard{hedw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrEmployeeDeleteWizard creates a new hr.employee.delete.wizard model and returns its id.
func (c *Client) CreateHrEmployeeDeleteWizards(hedws []*HrEmployeeDeleteWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range hedws {
		vv = append(vv, v)
	}
	return c.Create(HrEmployeeDeleteWizardModel, vv, nil)
}

// UpdateHrEmployeeDeleteWizard updates an existing hr.employee.delete.wizard record.
func (c *Client) UpdateHrEmployeeDeleteWizard(hedw *HrEmployeeDeleteWizard) error {
	return c.UpdateHrEmployeeDeleteWizards([]int64{hedw.Id.Get()}, hedw)
}

// UpdateHrEmployeeDeleteWizards updates existing hr.employee.delete.wizard records.
// All records (represented by ids) will be updated by hedw values.
func (c *Client) UpdateHrEmployeeDeleteWizards(ids []int64, hedw *HrEmployeeDeleteWizard) error {
	return c.Update(HrEmployeeDeleteWizardModel, ids, hedw, nil)
}

// DeleteHrEmployeeDeleteWizard deletes an existing hr.employee.delete.wizard record.
func (c *Client) DeleteHrEmployeeDeleteWizard(id int64) error {
	return c.DeleteHrEmployeeDeleteWizards([]int64{id})
}

// DeleteHrEmployeeDeleteWizards deletes existing hr.employee.delete.wizard records.
func (c *Client) DeleteHrEmployeeDeleteWizards(ids []int64) error {
	return c.Delete(HrEmployeeDeleteWizardModel, ids)
}

// GetHrEmployeeDeleteWizard gets hr.employee.delete.wizard existing record.
func (c *Client) GetHrEmployeeDeleteWizard(id int64) (*HrEmployeeDeleteWizard, error) {
	hedws, err := c.GetHrEmployeeDeleteWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*hedws)[0]), nil
}

// GetHrEmployeeDeleteWizards gets hr.employee.delete.wizard existing records.
func (c *Client) GetHrEmployeeDeleteWizards(ids []int64) (*HrEmployeeDeleteWizards, error) {
	hedws := &HrEmployeeDeleteWizards{}
	if err := c.Read(HrEmployeeDeleteWizardModel, ids, nil, hedws); err != nil {
		return nil, err
	}
	return hedws, nil
}

// FindHrEmployeeDeleteWizard finds hr.employee.delete.wizard record by querying it with criteria.
func (c *Client) FindHrEmployeeDeleteWizard(criteria *Criteria) (*HrEmployeeDeleteWizard, error) {
	hedws := &HrEmployeeDeleteWizards{}
	if err := c.SearchRead(HrEmployeeDeleteWizardModel, criteria, NewOptions().Limit(1), hedws); err != nil {
		return nil, err
	}
	return &((*hedws)[0]), nil
}

// FindHrEmployeeDeleteWizards finds hr.employee.delete.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrEmployeeDeleteWizards(criteria *Criteria, options *Options) (*HrEmployeeDeleteWizards, error) {
	hedws := &HrEmployeeDeleteWizards{}
	if err := c.SearchRead(HrEmployeeDeleteWizardModel, criteria, options, hedws); err != nil {
		return nil, err
	}
	return hedws, nil
}

// FindHrEmployeeDeleteWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrEmployeeDeleteWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(HrEmployeeDeleteWizardModel, criteria, options)
}

// FindHrEmployeeDeleteWizardId finds record id by querying it with criteria.
func (c *Client) FindHrEmployeeDeleteWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrEmployeeDeleteWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
