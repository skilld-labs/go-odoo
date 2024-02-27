package odoo

// ChangePasswordWizard represents change.password.wizard model.
type ChangePasswordWizard struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	UserIds     *Relation `xmlrpc:"user_ids,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// ChangePasswordWizards represents array of change.password.wizard model.
type ChangePasswordWizards []ChangePasswordWizard

// ChangePasswordWizardModel is the odoo model name.
const ChangePasswordWizardModel = "change.password.wizard"

// Many2One convert ChangePasswordWizard to *Many2One.
func (cpw *ChangePasswordWizard) Many2One() *Many2One {
	return NewMany2One(cpw.Id.Get(), "")
}

// CreateChangePasswordWizard creates a new change.password.wizard model and returns its id.
func (c *Client) CreateChangePasswordWizard(cpw *ChangePasswordWizard) (int64, error) {
	ids, err := c.CreateChangePasswordWizards([]*ChangePasswordWizard{cpw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateChangePasswordWizard creates a new change.password.wizard model and returns its id.
func (c *Client) CreateChangePasswordWizards(cpws []*ChangePasswordWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range cpws {
		vv = append(vv, v)
	}
	return c.Create(ChangePasswordWizardModel, vv, nil)
}

// UpdateChangePasswordWizard updates an existing change.password.wizard record.
func (c *Client) UpdateChangePasswordWizard(cpw *ChangePasswordWizard) error {
	return c.UpdateChangePasswordWizards([]int64{cpw.Id.Get()}, cpw)
}

// UpdateChangePasswordWizards updates existing change.password.wizard records.
// All records (represented by ids) will be updated by cpw values.
func (c *Client) UpdateChangePasswordWizards(ids []int64, cpw *ChangePasswordWizard) error {
	return c.Update(ChangePasswordWizardModel, ids, cpw, nil)
}

// DeleteChangePasswordWizard deletes an existing change.password.wizard record.
func (c *Client) DeleteChangePasswordWizard(id int64) error {
	return c.DeleteChangePasswordWizards([]int64{id})
}

// DeleteChangePasswordWizards deletes existing change.password.wizard records.
func (c *Client) DeleteChangePasswordWizards(ids []int64) error {
	return c.Delete(ChangePasswordWizardModel, ids)
}

// GetChangePasswordWizard gets change.password.wizard existing record.
func (c *Client) GetChangePasswordWizard(id int64) (*ChangePasswordWizard, error) {
	cpws, err := c.GetChangePasswordWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*cpws)[0]), nil
}

// GetChangePasswordWizards gets change.password.wizard existing records.
func (c *Client) GetChangePasswordWizards(ids []int64) (*ChangePasswordWizards, error) {
	cpws := &ChangePasswordWizards{}
	if err := c.Read(ChangePasswordWizardModel, ids, nil, cpws); err != nil {
		return nil, err
	}
	return cpws, nil
}

// FindChangePasswordWizard finds change.password.wizard record by querying it with criteria.
func (c *Client) FindChangePasswordWizard(criteria *Criteria) (*ChangePasswordWizard, error) {
	cpws := &ChangePasswordWizards{}
	if err := c.SearchRead(ChangePasswordWizardModel, criteria, NewOptions().Limit(1), cpws); err != nil {
		return nil, err
	}
	return &((*cpws)[0]), nil
}

// FindChangePasswordWizards finds change.password.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindChangePasswordWizards(criteria *Criteria, options *Options) (*ChangePasswordWizards, error) {
	cpws := &ChangePasswordWizards{}
	if err := c.SearchRead(ChangePasswordWizardModel, criteria, options, cpws); err != nil {
		return nil, err
	}
	return cpws, nil
}

// FindChangePasswordWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindChangePasswordWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ChangePasswordWizardModel, criteria, options)
}

// FindChangePasswordWizardId finds record id by querying it with criteria.
func (c *Client) FindChangePasswordWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ChangePasswordWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
