package odoo

// BaseEnableProfilingWizard represents base.enable.profiling.wizard model.
type BaseEnableProfilingWizard struct {
	CreateDate  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName *String    `xmlrpc:"display_name,omitempty"`
	Duration    *Selection `xmlrpc:"duration,omitempty"`
	Expiration  *Time      `xmlrpc:"expiration,omitempty"`
	Id          *Int       `xmlrpc:"id,omitempty"`
	WriteDate   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// BaseEnableProfilingWizards represents array of base.enable.profiling.wizard model.
type BaseEnableProfilingWizards []BaseEnableProfilingWizard

// BaseEnableProfilingWizardModel is the odoo model name.
const BaseEnableProfilingWizardModel = "base.enable.profiling.wizard"

// Many2One convert BaseEnableProfilingWizard to *Many2One.
func (bepw *BaseEnableProfilingWizard) Many2One() *Many2One {
	return NewMany2One(bepw.Id.Get(), "")
}

// CreateBaseEnableProfilingWizard creates a new base.enable.profiling.wizard model and returns its id.
func (c *Client) CreateBaseEnableProfilingWizard(bepw *BaseEnableProfilingWizard) (int64, error) {
	ids, err := c.CreateBaseEnableProfilingWizards([]*BaseEnableProfilingWizard{bepw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateBaseEnableProfilingWizard creates a new base.enable.profiling.wizard model and returns its id.
func (c *Client) CreateBaseEnableProfilingWizards(bepws []*BaseEnableProfilingWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range bepws {
		vv = append(vv, v)
	}
	return c.Create(BaseEnableProfilingWizardModel, vv, nil)
}

// UpdateBaseEnableProfilingWizard updates an existing base.enable.profiling.wizard record.
func (c *Client) UpdateBaseEnableProfilingWizard(bepw *BaseEnableProfilingWizard) error {
	return c.UpdateBaseEnableProfilingWizards([]int64{bepw.Id.Get()}, bepw)
}

// UpdateBaseEnableProfilingWizards updates existing base.enable.profiling.wizard records.
// All records (represented by ids) will be updated by bepw values.
func (c *Client) UpdateBaseEnableProfilingWizards(ids []int64, bepw *BaseEnableProfilingWizard) error {
	return c.Update(BaseEnableProfilingWizardModel, ids, bepw, nil)
}

// DeleteBaseEnableProfilingWizard deletes an existing base.enable.profiling.wizard record.
func (c *Client) DeleteBaseEnableProfilingWizard(id int64) error {
	return c.DeleteBaseEnableProfilingWizards([]int64{id})
}

// DeleteBaseEnableProfilingWizards deletes existing base.enable.profiling.wizard records.
func (c *Client) DeleteBaseEnableProfilingWizards(ids []int64) error {
	return c.Delete(BaseEnableProfilingWizardModel, ids)
}

// GetBaseEnableProfilingWizard gets base.enable.profiling.wizard existing record.
func (c *Client) GetBaseEnableProfilingWizard(id int64) (*BaseEnableProfilingWizard, error) {
	bepws, err := c.GetBaseEnableProfilingWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*bepws)[0]), nil
}

// GetBaseEnableProfilingWizards gets base.enable.profiling.wizard existing records.
func (c *Client) GetBaseEnableProfilingWizards(ids []int64) (*BaseEnableProfilingWizards, error) {
	bepws := &BaseEnableProfilingWizards{}
	if err := c.Read(BaseEnableProfilingWizardModel, ids, nil, bepws); err != nil {
		return nil, err
	}
	return bepws, nil
}

// FindBaseEnableProfilingWizard finds base.enable.profiling.wizard record by querying it with criteria.
func (c *Client) FindBaseEnableProfilingWizard(criteria *Criteria) (*BaseEnableProfilingWizard, error) {
	bepws := &BaseEnableProfilingWizards{}
	if err := c.SearchRead(BaseEnableProfilingWizardModel, criteria, NewOptions().Limit(1), bepws); err != nil {
		return nil, err
	}
	return &((*bepws)[0]), nil
}

// FindBaseEnableProfilingWizards finds base.enable.profiling.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseEnableProfilingWizards(criteria *Criteria, options *Options) (*BaseEnableProfilingWizards, error) {
	bepws := &BaseEnableProfilingWizards{}
	if err := c.SearchRead(BaseEnableProfilingWizardModel, criteria, options, bepws); err != nil {
		return nil, err
	}
	return bepws, nil
}

// FindBaseEnableProfilingWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseEnableProfilingWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(BaseEnableProfilingWizardModel, criteria, options)
}

// FindBaseEnableProfilingWizardId finds record id by querying it with criteria.
func (c *Client) FindBaseEnableProfilingWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(BaseEnableProfilingWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
