package odoo

// PrivacyLookupWizard represents privacy.lookup.wizard model.
type PrivacyLookupWizard struct {
	CreateDate         *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid          *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName        *String   `xmlrpc:"display_name,omitempty"`
	Email              *String   `xmlrpc:"email,omitempty"`
	ExecutionDetails   *String   `xmlrpc:"execution_details,omitempty"`
	Id                 *Int      `xmlrpc:"id,omitempty"`
	LineCount          *Int      `xmlrpc:"line_count,omitempty"`
	LineIds            *Relation `xmlrpc:"line_ids,omitempty"`
	LogId              *Many2One `xmlrpc:"log_id,omitempty"`
	Name               *String   `xmlrpc:"name,omitempty"`
	RecordsDescription *String   `xmlrpc:"records_description,omitempty"`
	WriteDate          *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid           *Many2One `xmlrpc:"write_uid,omitempty"`
}

// PrivacyLookupWizards represents array of privacy.lookup.wizard model.
type PrivacyLookupWizards []PrivacyLookupWizard

// PrivacyLookupWizardModel is the odoo model name.
const PrivacyLookupWizardModel = "privacy.lookup.wizard"

// Many2One convert PrivacyLookupWizard to *Many2One.
func (plw *PrivacyLookupWizard) Many2One() *Many2One {
	return NewMany2One(plw.Id.Get(), "")
}

// CreatePrivacyLookupWizard creates a new privacy.lookup.wizard model and returns its id.
func (c *Client) CreatePrivacyLookupWizard(plw *PrivacyLookupWizard) (int64, error) {
	ids, err := c.CreatePrivacyLookupWizards([]*PrivacyLookupWizard{plw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreatePrivacyLookupWizard creates a new privacy.lookup.wizard model and returns its id.
func (c *Client) CreatePrivacyLookupWizards(plws []*PrivacyLookupWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range plws {
		vv = append(vv, v)
	}
	return c.Create(PrivacyLookupWizardModel, vv, nil)
}

// UpdatePrivacyLookupWizard updates an existing privacy.lookup.wizard record.
func (c *Client) UpdatePrivacyLookupWizard(plw *PrivacyLookupWizard) error {
	return c.UpdatePrivacyLookupWizards([]int64{plw.Id.Get()}, plw)
}

// UpdatePrivacyLookupWizards updates existing privacy.lookup.wizard records.
// All records (represented by ids) will be updated by plw values.
func (c *Client) UpdatePrivacyLookupWizards(ids []int64, plw *PrivacyLookupWizard) error {
	return c.Update(PrivacyLookupWizardModel, ids, plw, nil)
}

// DeletePrivacyLookupWizard deletes an existing privacy.lookup.wizard record.
func (c *Client) DeletePrivacyLookupWizard(id int64) error {
	return c.DeletePrivacyLookupWizards([]int64{id})
}

// DeletePrivacyLookupWizards deletes existing privacy.lookup.wizard records.
func (c *Client) DeletePrivacyLookupWizards(ids []int64) error {
	return c.Delete(PrivacyLookupWizardModel, ids)
}

// GetPrivacyLookupWizard gets privacy.lookup.wizard existing record.
func (c *Client) GetPrivacyLookupWizard(id int64) (*PrivacyLookupWizard, error) {
	plws, err := c.GetPrivacyLookupWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*plws)[0]), nil
}

// GetPrivacyLookupWizards gets privacy.lookup.wizard existing records.
func (c *Client) GetPrivacyLookupWizards(ids []int64) (*PrivacyLookupWizards, error) {
	plws := &PrivacyLookupWizards{}
	if err := c.Read(PrivacyLookupWizardModel, ids, nil, plws); err != nil {
		return nil, err
	}
	return plws, nil
}

// FindPrivacyLookupWizard finds privacy.lookup.wizard record by querying it with criteria.
func (c *Client) FindPrivacyLookupWizard(criteria *Criteria) (*PrivacyLookupWizard, error) {
	plws := &PrivacyLookupWizards{}
	if err := c.SearchRead(PrivacyLookupWizardModel, criteria, NewOptions().Limit(1), plws); err != nil {
		return nil, err
	}
	return &((*plws)[0]), nil
}

// FindPrivacyLookupWizards finds privacy.lookup.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPrivacyLookupWizards(criteria *Criteria, options *Options) (*PrivacyLookupWizards, error) {
	plws := &PrivacyLookupWizards{}
	if err := c.SearchRead(PrivacyLookupWizardModel, criteria, options, plws); err != nil {
		return nil, err
	}
	return plws, nil
}

// FindPrivacyLookupWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPrivacyLookupWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(PrivacyLookupWizardModel, criteria, options)
}

// FindPrivacyLookupWizardId finds record id by querying it with criteria.
func (c *Client) FindPrivacyLookupWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PrivacyLookupWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
