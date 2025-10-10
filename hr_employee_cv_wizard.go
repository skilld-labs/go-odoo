package odoo

// HrEmployeeCvWizard represents hr.employee.cv.wizard model.
type HrEmployeeCvWizard struct {
	CanShowOthers  *Bool     `xmlrpc:"can_show_others,omitempty"`
	CanShowSkills  *Bool     `xmlrpc:"can_show_skills,omitempty"`
	ColorPrimary   *String   `xmlrpc:"color_primary,omitempty"`
	ColorSecondary *String   `xmlrpc:"color_secondary,omitempty"`
	CreateDate     *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid      *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName    *String   `xmlrpc:"display_name,omitempty"`
	EmployeeIds    *Relation `xmlrpc:"employee_ids,omitempty"`
	Id             *Int      `xmlrpc:"id,omitempty"`
	ShowContact    *Bool     `xmlrpc:"show_contact,omitempty"`
	ShowOthers     *Bool     `xmlrpc:"show_others,omitempty"`
	ShowSkills     *Bool     `xmlrpc:"show_skills,omitempty"`
	WriteDate      *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid       *Many2One `xmlrpc:"write_uid,omitempty"`
}

// HrEmployeeCvWizards represents array of hr.employee.cv.wizard model.
type HrEmployeeCvWizards []HrEmployeeCvWizard

// HrEmployeeCvWizardModel is the odoo model name.
const HrEmployeeCvWizardModel = "hr.employee.cv.wizard"

// Many2One convert HrEmployeeCvWizard to *Many2One.
func (hecw *HrEmployeeCvWizard) Many2One() *Many2One {
	return NewMany2One(hecw.Id.Get(), "")
}

// CreateHrEmployeeCvWizard creates a new hr.employee.cv.wizard model and returns its id.
func (c *Client) CreateHrEmployeeCvWizard(hecw *HrEmployeeCvWizard) (int64, error) {
	ids, err := c.CreateHrEmployeeCvWizards([]*HrEmployeeCvWizard{hecw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrEmployeeCvWizard creates a new hr.employee.cv.wizard model and returns its id.
func (c *Client) CreateHrEmployeeCvWizards(hecws []*HrEmployeeCvWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range hecws {
		vv = append(vv, v)
	}
	return c.Create(HrEmployeeCvWizardModel, vv, nil)
}

// UpdateHrEmployeeCvWizard updates an existing hr.employee.cv.wizard record.
func (c *Client) UpdateHrEmployeeCvWizard(hecw *HrEmployeeCvWizard) error {
	return c.UpdateHrEmployeeCvWizards([]int64{hecw.Id.Get()}, hecw)
}

// UpdateHrEmployeeCvWizards updates existing hr.employee.cv.wizard records.
// All records (represented by ids) will be updated by hecw values.
func (c *Client) UpdateHrEmployeeCvWizards(ids []int64, hecw *HrEmployeeCvWizard) error {
	return c.Update(HrEmployeeCvWizardModel, ids, hecw, nil)
}

// DeleteHrEmployeeCvWizard deletes an existing hr.employee.cv.wizard record.
func (c *Client) DeleteHrEmployeeCvWizard(id int64) error {
	return c.DeleteHrEmployeeCvWizards([]int64{id})
}

// DeleteHrEmployeeCvWizards deletes existing hr.employee.cv.wizard records.
func (c *Client) DeleteHrEmployeeCvWizards(ids []int64) error {
	return c.Delete(HrEmployeeCvWizardModel, ids)
}

// GetHrEmployeeCvWizard gets hr.employee.cv.wizard existing record.
func (c *Client) GetHrEmployeeCvWizard(id int64) (*HrEmployeeCvWizard, error) {
	hecws, err := c.GetHrEmployeeCvWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*hecws)[0]), nil
}

// GetHrEmployeeCvWizards gets hr.employee.cv.wizard existing records.
func (c *Client) GetHrEmployeeCvWizards(ids []int64) (*HrEmployeeCvWizards, error) {
	hecws := &HrEmployeeCvWizards{}
	if err := c.Read(HrEmployeeCvWizardModel, ids, nil, hecws); err != nil {
		return nil, err
	}
	return hecws, nil
}

// FindHrEmployeeCvWizard finds hr.employee.cv.wizard record by querying it with criteria.
func (c *Client) FindHrEmployeeCvWizard(criteria *Criteria) (*HrEmployeeCvWizard, error) {
	hecws := &HrEmployeeCvWizards{}
	if err := c.SearchRead(HrEmployeeCvWizardModel, criteria, NewOptions().Limit(1), hecws); err != nil {
		return nil, err
	}
	return &((*hecws)[0]), nil
}

// FindHrEmployeeCvWizards finds hr.employee.cv.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrEmployeeCvWizards(criteria *Criteria, options *Options) (*HrEmployeeCvWizards, error) {
	hecws := &HrEmployeeCvWizards{}
	if err := c.SearchRead(HrEmployeeCvWizardModel, criteria, options, hecws); err != nil {
		return nil, err
	}
	return hecws, nil
}

// FindHrEmployeeCvWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrEmployeeCvWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(HrEmployeeCvWizardModel, criteria, options)
}

// FindHrEmployeeCvWizardId finds record id by querying it with criteria.
func (c *Client) FindHrEmployeeCvWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrEmployeeCvWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
