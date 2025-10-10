package odoo

// HrLeaveAllocationGenerateMultiWizard represents hr.leave.allocation.generate.multi.wizard model.
type HrLeaveAllocationGenerateMultiWizard struct {
	AccrualPlanId   *Many2One  `xmlrpc:"accrual_plan_id,omitempty"`
	AllocationMode  *Selection `xmlrpc:"allocation_mode,omitempty"`
	AllocationType  *Selection `xmlrpc:"allocation_type,omitempty"`
	CategoryId      *Many2One  `xmlrpc:"category_id,omitempty"`
	CompanyId       *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate      *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid       *Many2One  `xmlrpc:"create_uid,omitempty"`
	DateFrom        *Time      `xmlrpc:"date_from,omitempty"`
	DateTo          *Time      `xmlrpc:"date_to,omitempty"`
	DepartmentId    *Many2One  `xmlrpc:"department_id,omitempty"`
	DisplayName     *String    `xmlrpc:"display_name,omitempty"`
	Duration        *Float     `xmlrpc:"duration,omitempty"`
	EmployeeIds     *Relation  `xmlrpc:"employee_ids,omitempty"`
	HolidayStatusId *Many2One  `xmlrpc:"holiday_status_id,omitempty"`
	Id              *Int       `xmlrpc:"id,omitempty"`
	Name            *String    `xmlrpc:"name,omitempty"`
	Notes           *String    `xmlrpc:"notes,omitempty"`
	RequestUnit     *Selection `xmlrpc:"request_unit,omitempty"`
	WriteDate       *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid        *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// HrLeaveAllocationGenerateMultiWizards represents array of hr.leave.allocation.generate.multi.wizard model.
type HrLeaveAllocationGenerateMultiWizards []HrLeaveAllocationGenerateMultiWizard

// HrLeaveAllocationGenerateMultiWizardModel is the odoo model name.
const HrLeaveAllocationGenerateMultiWizardModel = "hr.leave.allocation.generate.multi.wizard"

// Many2One convert HrLeaveAllocationGenerateMultiWizard to *Many2One.
func (hlagmw *HrLeaveAllocationGenerateMultiWizard) Many2One() *Many2One {
	return NewMany2One(hlagmw.Id.Get(), "")
}

// CreateHrLeaveAllocationGenerateMultiWizard creates a new hr.leave.allocation.generate.multi.wizard model and returns its id.
func (c *Client) CreateHrLeaveAllocationGenerateMultiWizard(hlagmw *HrLeaveAllocationGenerateMultiWizard) (int64, error) {
	ids, err := c.CreateHrLeaveAllocationGenerateMultiWizards([]*HrLeaveAllocationGenerateMultiWizard{hlagmw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrLeaveAllocationGenerateMultiWizard creates a new hr.leave.allocation.generate.multi.wizard model and returns its id.
func (c *Client) CreateHrLeaveAllocationGenerateMultiWizards(hlagmws []*HrLeaveAllocationGenerateMultiWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range hlagmws {
		vv = append(vv, v)
	}
	return c.Create(HrLeaveAllocationGenerateMultiWizardModel, vv, nil)
}

// UpdateHrLeaveAllocationGenerateMultiWizard updates an existing hr.leave.allocation.generate.multi.wizard record.
func (c *Client) UpdateHrLeaveAllocationGenerateMultiWizard(hlagmw *HrLeaveAllocationGenerateMultiWizard) error {
	return c.UpdateHrLeaveAllocationGenerateMultiWizards([]int64{hlagmw.Id.Get()}, hlagmw)
}

// UpdateHrLeaveAllocationGenerateMultiWizards updates existing hr.leave.allocation.generate.multi.wizard records.
// All records (represented by ids) will be updated by hlagmw values.
func (c *Client) UpdateHrLeaveAllocationGenerateMultiWizards(ids []int64, hlagmw *HrLeaveAllocationGenerateMultiWizard) error {
	return c.Update(HrLeaveAllocationGenerateMultiWizardModel, ids, hlagmw, nil)
}

// DeleteHrLeaveAllocationGenerateMultiWizard deletes an existing hr.leave.allocation.generate.multi.wizard record.
func (c *Client) DeleteHrLeaveAllocationGenerateMultiWizard(id int64) error {
	return c.DeleteHrLeaveAllocationGenerateMultiWizards([]int64{id})
}

// DeleteHrLeaveAllocationGenerateMultiWizards deletes existing hr.leave.allocation.generate.multi.wizard records.
func (c *Client) DeleteHrLeaveAllocationGenerateMultiWizards(ids []int64) error {
	return c.Delete(HrLeaveAllocationGenerateMultiWizardModel, ids)
}

// GetHrLeaveAllocationGenerateMultiWizard gets hr.leave.allocation.generate.multi.wizard existing record.
func (c *Client) GetHrLeaveAllocationGenerateMultiWizard(id int64) (*HrLeaveAllocationGenerateMultiWizard, error) {
	hlagmws, err := c.GetHrLeaveAllocationGenerateMultiWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*hlagmws)[0]), nil
}

// GetHrLeaveAllocationGenerateMultiWizards gets hr.leave.allocation.generate.multi.wizard existing records.
func (c *Client) GetHrLeaveAllocationGenerateMultiWizards(ids []int64) (*HrLeaveAllocationGenerateMultiWizards, error) {
	hlagmws := &HrLeaveAllocationGenerateMultiWizards{}
	if err := c.Read(HrLeaveAllocationGenerateMultiWizardModel, ids, nil, hlagmws); err != nil {
		return nil, err
	}
	return hlagmws, nil
}

// FindHrLeaveAllocationGenerateMultiWizard finds hr.leave.allocation.generate.multi.wizard record by querying it with criteria.
func (c *Client) FindHrLeaveAllocationGenerateMultiWizard(criteria *Criteria) (*HrLeaveAllocationGenerateMultiWizard, error) {
	hlagmws := &HrLeaveAllocationGenerateMultiWizards{}
	if err := c.SearchRead(HrLeaveAllocationGenerateMultiWizardModel, criteria, NewOptions().Limit(1), hlagmws); err != nil {
		return nil, err
	}
	return &((*hlagmws)[0]), nil
}

// FindHrLeaveAllocationGenerateMultiWizards finds hr.leave.allocation.generate.multi.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrLeaveAllocationGenerateMultiWizards(criteria *Criteria, options *Options) (*HrLeaveAllocationGenerateMultiWizards, error) {
	hlagmws := &HrLeaveAllocationGenerateMultiWizards{}
	if err := c.SearchRead(HrLeaveAllocationGenerateMultiWizardModel, criteria, options, hlagmws); err != nil {
		return nil, err
	}
	return hlagmws, nil
}

// FindHrLeaveAllocationGenerateMultiWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrLeaveAllocationGenerateMultiWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(HrLeaveAllocationGenerateMultiWizardModel, criteria, options)
}

// FindHrLeaveAllocationGenerateMultiWizardId finds record id by querying it with criteria.
func (c *Client) FindHrLeaveAllocationGenerateMultiWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrLeaveAllocationGenerateMultiWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
