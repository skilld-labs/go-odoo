package odoo

// HrBankAccountAllocationWizard represents hr.bank.account.allocation.wizard model.
type HrBankAccountAllocationWizard struct {
	AllocationIds *Relation `xmlrpc:"allocation_ids,omitempty"`
	CreateDate    *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid     *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName   *String   `xmlrpc:"display_name,omitempty"`
	EmployeeId    *Many2One `xmlrpc:"employee_id,omitempty"`
	Id            *Int      `xmlrpc:"id,omitempty"`
	WriteDate     *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid      *Many2One `xmlrpc:"write_uid,omitempty"`
}

// HrBankAccountAllocationWizards represents array of hr.bank.account.allocation.wizard model.
type HrBankAccountAllocationWizards []HrBankAccountAllocationWizard

// HrBankAccountAllocationWizardModel is the odoo model name.
const HrBankAccountAllocationWizardModel = "hr.bank.account.allocation.wizard"

// Many2One convert HrBankAccountAllocationWizard to *Many2One.
func (hbaaw *HrBankAccountAllocationWizard) Many2One() *Many2One {
	return NewMany2One(hbaaw.Id.Get(), "")
}

// CreateHrBankAccountAllocationWizard creates a new hr.bank.account.allocation.wizard model and returns its id.
func (c *Client) CreateHrBankAccountAllocationWizard(hbaaw *HrBankAccountAllocationWizard) (int64, error) {
	ids, err := c.CreateHrBankAccountAllocationWizards([]*HrBankAccountAllocationWizard{hbaaw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrBankAccountAllocationWizard creates a new hr.bank.account.allocation.wizard model and returns its id.
func (c *Client) CreateHrBankAccountAllocationWizards(hbaaws []*HrBankAccountAllocationWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range hbaaws {
		vv = append(vv, v)
	}
	return c.Create(HrBankAccountAllocationWizardModel, vv, nil)
}

// UpdateHrBankAccountAllocationWizard updates an existing hr.bank.account.allocation.wizard record.
func (c *Client) UpdateHrBankAccountAllocationWizard(hbaaw *HrBankAccountAllocationWizard) error {
	return c.UpdateHrBankAccountAllocationWizards([]int64{hbaaw.Id.Get()}, hbaaw)
}

// UpdateHrBankAccountAllocationWizards updates existing hr.bank.account.allocation.wizard records.
// All records (represented by ids) will be updated by hbaaw values.
func (c *Client) UpdateHrBankAccountAllocationWizards(ids []int64, hbaaw *HrBankAccountAllocationWizard) error {
	return c.Update(HrBankAccountAllocationWizardModel, ids, hbaaw, nil)
}

// DeleteHrBankAccountAllocationWizard deletes an existing hr.bank.account.allocation.wizard record.
func (c *Client) DeleteHrBankAccountAllocationWizard(id int64) error {
	return c.DeleteHrBankAccountAllocationWizards([]int64{id})
}

// DeleteHrBankAccountAllocationWizards deletes existing hr.bank.account.allocation.wizard records.
func (c *Client) DeleteHrBankAccountAllocationWizards(ids []int64) error {
	return c.Delete(HrBankAccountAllocationWizardModel, ids)
}

// GetHrBankAccountAllocationWizard gets hr.bank.account.allocation.wizard existing record.
func (c *Client) GetHrBankAccountAllocationWizard(id int64) (*HrBankAccountAllocationWizard, error) {
	hbaaws, err := c.GetHrBankAccountAllocationWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*hbaaws)[0]), nil
}

// GetHrBankAccountAllocationWizards gets hr.bank.account.allocation.wizard existing records.
func (c *Client) GetHrBankAccountAllocationWizards(ids []int64) (*HrBankAccountAllocationWizards, error) {
	hbaaws := &HrBankAccountAllocationWizards{}
	if err := c.Read(HrBankAccountAllocationWizardModel, ids, nil, hbaaws); err != nil {
		return nil, err
	}
	return hbaaws, nil
}

// FindHrBankAccountAllocationWizard finds hr.bank.account.allocation.wizard record by querying it with criteria.
func (c *Client) FindHrBankAccountAllocationWizard(criteria *Criteria) (*HrBankAccountAllocationWizard, error) {
	hbaaws := &HrBankAccountAllocationWizards{}
	if err := c.SearchRead(HrBankAccountAllocationWizardModel, criteria, NewOptions().Limit(1), hbaaws); err != nil {
		return nil, err
	}
	return &((*hbaaws)[0]), nil
}

// FindHrBankAccountAllocationWizards finds hr.bank.account.allocation.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrBankAccountAllocationWizards(criteria *Criteria, options *Options) (*HrBankAccountAllocationWizards, error) {
	hbaaws := &HrBankAccountAllocationWizards{}
	if err := c.SearchRead(HrBankAccountAllocationWizardModel, criteria, options, hbaaws); err != nil {
		return nil, err
	}
	return hbaaws, nil
}

// FindHrBankAccountAllocationWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrBankAccountAllocationWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(HrBankAccountAllocationWizardModel, criteria, options)
}

// FindHrBankAccountAllocationWizardId finds record id by querying it with criteria.
func (c *Client) FindHrBankAccountAllocationWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrBankAccountAllocationWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
