package odoo

// HrVersionWizard represents hr.version.wizard model.
type HrVersionWizard struct {
	ContractTemplateId *Many2One `xmlrpc:"contract_template_id,omitempty"`
	CreateDate         *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid          *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName        *String   `xmlrpc:"display_name,omitempty"`
	Id                 *Int      `xmlrpc:"id,omitempty"`
	WriteDate          *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid           *Many2One `xmlrpc:"write_uid,omitempty"`
}

// HrVersionWizards represents array of hr.version.wizard model.
type HrVersionWizards []HrVersionWizard

// HrVersionWizardModel is the odoo model name.
const HrVersionWizardModel = "hr.version.wizard"

// Many2One convert HrVersionWizard to *Many2One.
func (hvw *HrVersionWizard) Many2One() *Many2One {
	return NewMany2One(hvw.Id.Get(), "")
}

// CreateHrVersionWizard creates a new hr.version.wizard model and returns its id.
func (c *Client) CreateHrVersionWizard(hvw *HrVersionWizard) (int64, error) {
	ids, err := c.CreateHrVersionWizards([]*HrVersionWizard{hvw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrVersionWizard creates a new hr.version.wizard model and returns its id.
func (c *Client) CreateHrVersionWizards(hvws []*HrVersionWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range hvws {
		vv = append(vv, v)
	}
	return c.Create(HrVersionWizardModel, vv, nil)
}

// UpdateHrVersionWizard updates an existing hr.version.wizard record.
func (c *Client) UpdateHrVersionWizard(hvw *HrVersionWizard) error {
	return c.UpdateHrVersionWizards([]int64{hvw.Id.Get()}, hvw)
}

// UpdateHrVersionWizards updates existing hr.version.wizard records.
// All records (represented by ids) will be updated by hvw values.
func (c *Client) UpdateHrVersionWizards(ids []int64, hvw *HrVersionWizard) error {
	return c.Update(HrVersionWizardModel, ids, hvw, nil)
}

// DeleteHrVersionWizard deletes an existing hr.version.wizard record.
func (c *Client) DeleteHrVersionWizard(id int64) error {
	return c.DeleteHrVersionWizards([]int64{id})
}

// DeleteHrVersionWizards deletes existing hr.version.wizard records.
func (c *Client) DeleteHrVersionWizards(ids []int64) error {
	return c.Delete(HrVersionWizardModel, ids)
}

// GetHrVersionWizard gets hr.version.wizard existing record.
func (c *Client) GetHrVersionWizard(id int64) (*HrVersionWizard, error) {
	hvws, err := c.GetHrVersionWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*hvws)[0]), nil
}

// GetHrVersionWizards gets hr.version.wizard existing records.
func (c *Client) GetHrVersionWizards(ids []int64) (*HrVersionWizards, error) {
	hvws := &HrVersionWizards{}
	if err := c.Read(HrVersionWizardModel, ids, nil, hvws); err != nil {
		return nil, err
	}
	return hvws, nil
}

// FindHrVersionWizard finds hr.version.wizard record by querying it with criteria.
func (c *Client) FindHrVersionWizard(criteria *Criteria) (*HrVersionWizard, error) {
	hvws := &HrVersionWizards{}
	if err := c.SearchRead(HrVersionWizardModel, criteria, NewOptions().Limit(1), hvws); err != nil {
		return nil, err
	}
	return &((*hvws)[0]), nil
}

// FindHrVersionWizards finds hr.version.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrVersionWizards(criteria *Criteria, options *Options) (*HrVersionWizards, error) {
	hvws := &HrVersionWizards{}
	if err := c.SearchRead(HrVersionWizardModel, criteria, options, hvws); err != nil {
		return nil, err
	}
	return hvws, nil
}

// FindHrVersionWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrVersionWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(HrVersionWizardModel, criteria, options)
}

// FindHrVersionWizardId finds record id by querying it with criteria.
func (c *Client) FindHrVersionWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrVersionWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
