package odoo

// HrLeaveGenerateMultiWizard represents hr.leave.generate.multi.wizard model.
type HrLeaveGenerateMultiWizard struct {
	AllocationMode  *Selection `xmlrpc:"allocation_mode,omitempty"`
	CategoryId      *Many2One  `xmlrpc:"category_id,omitempty"`
	CompanyId       *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate      *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid       *Many2One  `xmlrpc:"create_uid,omitempty"`
	DateFrom        *Time      `xmlrpc:"date_from,omitempty"`
	DateTo          *Time      `xmlrpc:"date_to,omitempty"`
	DepartmentId    *Many2One  `xmlrpc:"department_id,omitempty"`
	DisplayName     *String    `xmlrpc:"display_name,omitempty"`
	EmployeeIds     *Relation  `xmlrpc:"employee_ids,omitempty"`
	HolidayStatusId *Many2One  `xmlrpc:"holiday_status_id,omitempty"`
	Id              *Int       `xmlrpc:"id,omitempty"`
	Name            *String    `xmlrpc:"name,omitempty"`
	WriteDate       *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid        *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// HrLeaveGenerateMultiWizards represents array of hr.leave.generate.multi.wizard model.
type HrLeaveGenerateMultiWizards []HrLeaveGenerateMultiWizard

// HrLeaveGenerateMultiWizardModel is the odoo model name.
const HrLeaveGenerateMultiWizardModel = "hr.leave.generate.multi.wizard"

// Many2One convert HrLeaveGenerateMultiWizard to *Many2One.
func (hlgmw *HrLeaveGenerateMultiWizard) Many2One() *Many2One {
	return NewMany2One(hlgmw.Id.Get(), "")
}

// CreateHrLeaveGenerateMultiWizard creates a new hr.leave.generate.multi.wizard model and returns its id.
func (c *Client) CreateHrLeaveGenerateMultiWizard(hlgmw *HrLeaveGenerateMultiWizard) (int64, error) {
	ids, err := c.CreateHrLeaveGenerateMultiWizards([]*HrLeaveGenerateMultiWizard{hlgmw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrLeaveGenerateMultiWizard creates a new hr.leave.generate.multi.wizard model and returns its id.
func (c *Client) CreateHrLeaveGenerateMultiWizards(hlgmws []*HrLeaveGenerateMultiWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range hlgmws {
		vv = append(vv, v)
	}
	return c.Create(HrLeaveGenerateMultiWizardModel, vv, nil)
}

// UpdateHrLeaveGenerateMultiWizard updates an existing hr.leave.generate.multi.wizard record.
func (c *Client) UpdateHrLeaveGenerateMultiWizard(hlgmw *HrLeaveGenerateMultiWizard) error {
	return c.UpdateHrLeaveGenerateMultiWizards([]int64{hlgmw.Id.Get()}, hlgmw)
}

// UpdateHrLeaveGenerateMultiWizards updates existing hr.leave.generate.multi.wizard records.
// All records (represented by ids) will be updated by hlgmw values.
func (c *Client) UpdateHrLeaveGenerateMultiWizards(ids []int64, hlgmw *HrLeaveGenerateMultiWizard) error {
	return c.Update(HrLeaveGenerateMultiWizardModel, ids, hlgmw, nil)
}

// DeleteHrLeaveGenerateMultiWizard deletes an existing hr.leave.generate.multi.wizard record.
func (c *Client) DeleteHrLeaveGenerateMultiWizard(id int64) error {
	return c.DeleteHrLeaveGenerateMultiWizards([]int64{id})
}

// DeleteHrLeaveGenerateMultiWizards deletes existing hr.leave.generate.multi.wizard records.
func (c *Client) DeleteHrLeaveGenerateMultiWizards(ids []int64) error {
	return c.Delete(HrLeaveGenerateMultiWizardModel, ids)
}

// GetHrLeaveGenerateMultiWizard gets hr.leave.generate.multi.wizard existing record.
func (c *Client) GetHrLeaveGenerateMultiWizard(id int64) (*HrLeaveGenerateMultiWizard, error) {
	hlgmws, err := c.GetHrLeaveGenerateMultiWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*hlgmws)[0]), nil
}

// GetHrLeaveGenerateMultiWizards gets hr.leave.generate.multi.wizard existing records.
func (c *Client) GetHrLeaveGenerateMultiWizards(ids []int64) (*HrLeaveGenerateMultiWizards, error) {
	hlgmws := &HrLeaveGenerateMultiWizards{}
	if err := c.Read(HrLeaveGenerateMultiWizardModel, ids, nil, hlgmws); err != nil {
		return nil, err
	}
	return hlgmws, nil
}

// FindHrLeaveGenerateMultiWizard finds hr.leave.generate.multi.wizard record by querying it with criteria.
func (c *Client) FindHrLeaveGenerateMultiWizard(criteria *Criteria) (*HrLeaveGenerateMultiWizard, error) {
	hlgmws := &HrLeaveGenerateMultiWizards{}
	if err := c.SearchRead(HrLeaveGenerateMultiWizardModel, criteria, NewOptions().Limit(1), hlgmws); err != nil {
		return nil, err
	}
	return &((*hlgmws)[0]), nil
}

// FindHrLeaveGenerateMultiWizards finds hr.leave.generate.multi.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrLeaveGenerateMultiWizards(criteria *Criteria, options *Options) (*HrLeaveGenerateMultiWizards, error) {
	hlgmws := &HrLeaveGenerateMultiWizards{}
	if err := c.SearchRead(HrLeaveGenerateMultiWizardModel, criteria, options, hlgmws); err != nil {
		return nil, err
	}
	return hlgmws, nil
}

// FindHrLeaveGenerateMultiWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrLeaveGenerateMultiWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(HrLeaveGenerateMultiWizardModel, criteria, options)
}

// FindHrLeaveGenerateMultiWizardId finds record id by querying it with criteria.
func (c *Client) FindHrLeaveGenerateMultiWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrLeaveGenerateMultiWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
