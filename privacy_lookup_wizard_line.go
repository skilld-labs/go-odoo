package odoo

// PrivacyLookupWizardLine represents privacy.lookup.wizard.line model.
type PrivacyLookupWizardLine struct {
	CreateDate       *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid        *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName      *String   `xmlrpc:"display_name,omitempty"`
	ExecutionDetails *String   `xmlrpc:"execution_details,omitempty"`
	HasActive        *Bool     `xmlrpc:"has_active,omitempty"`
	Id               *Int      `xmlrpc:"id,omitempty"`
	IsActive         *Bool     `xmlrpc:"is_active,omitempty"`
	IsUnlinked       *Bool     `xmlrpc:"is_unlinked,omitempty"`
	ResId            *Int      `xmlrpc:"res_id,omitempty"`
	ResModel         *String   `xmlrpc:"res_model,omitempty"`
	ResModelId       *Many2One `xmlrpc:"res_model_id,omitempty"`
	ResName          *String   `xmlrpc:"res_name,omitempty"`
	ResourceRef      *String   `xmlrpc:"resource_ref,omitempty"`
	WizardId         *Many2One `xmlrpc:"wizard_id,omitempty"`
	WriteDate        *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid         *Many2One `xmlrpc:"write_uid,omitempty"`
}

// PrivacyLookupWizardLines represents array of privacy.lookup.wizard.line model.
type PrivacyLookupWizardLines []PrivacyLookupWizardLine

// PrivacyLookupWizardLineModel is the odoo model name.
const PrivacyLookupWizardLineModel = "privacy.lookup.wizard.line"

// Many2One convert PrivacyLookupWizardLine to *Many2One.
func (plwl *PrivacyLookupWizardLine) Many2One() *Many2One {
	return NewMany2One(plwl.Id.Get(), "")
}

// CreatePrivacyLookupWizardLine creates a new privacy.lookup.wizard.line model and returns its id.
func (c *Client) CreatePrivacyLookupWizardLine(plwl *PrivacyLookupWizardLine) (int64, error) {
	ids, err := c.CreatePrivacyLookupWizardLines([]*PrivacyLookupWizardLine{plwl})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreatePrivacyLookupWizardLine creates a new privacy.lookup.wizard.line model and returns its id.
func (c *Client) CreatePrivacyLookupWizardLines(plwls []*PrivacyLookupWizardLine) ([]int64, error) {
	var vv []interface{}
	for _, v := range plwls {
		vv = append(vv, v)
	}
	return c.Create(PrivacyLookupWizardLineModel, vv, nil)
}

// UpdatePrivacyLookupWizardLine updates an existing privacy.lookup.wizard.line record.
func (c *Client) UpdatePrivacyLookupWizardLine(plwl *PrivacyLookupWizardLine) error {
	return c.UpdatePrivacyLookupWizardLines([]int64{plwl.Id.Get()}, plwl)
}

// UpdatePrivacyLookupWizardLines updates existing privacy.lookup.wizard.line records.
// All records (represented by ids) will be updated by plwl values.
func (c *Client) UpdatePrivacyLookupWizardLines(ids []int64, plwl *PrivacyLookupWizardLine) error {
	return c.Update(PrivacyLookupWizardLineModel, ids, plwl, nil)
}

// DeletePrivacyLookupWizardLine deletes an existing privacy.lookup.wizard.line record.
func (c *Client) DeletePrivacyLookupWizardLine(id int64) error {
	return c.DeletePrivacyLookupWizardLines([]int64{id})
}

// DeletePrivacyLookupWizardLines deletes existing privacy.lookup.wizard.line records.
func (c *Client) DeletePrivacyLookupWizardLines(ids []int64) error {
	return c.Delete(PrivacyLookupWizardLineModel, ids)
}

// GetPrivacyLookupWizardLine gets privacy.lookup.wizard.line existing record.
func (c *Client) GetPrivacyLookupWizardLine(id int64) (*PrivacyLookupWizardLine, error) {
	plwls, err := c.GetPrivacyLookupWizardLines([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*plwls)[0]), nil
}

// GetPrivacyLookupWizardLines gets privacy.lookup.wizard.line existing records.
func (c *Client) GetPrivacyLookupWizardLines(ids []int64) (*PrivacyLookupWizardLines, error) {
	plwls := &PrivacyLookupWizardLines{}
	if err := c.Read(PrivacyLookupWizardLineModel, ids, nil, plwls); err != nil {
		return nil, err
	}
	return plwls, nil
}

// FindPrivacyLookupWizardLine finds privacy.lookup.wizard.line record by querying it with criteria.
func (c *Client) FindPrivacyLookupWizardLine(criteria *Criteria) (*PrivacyLookupWizardLine, error) {
	plwls := &PrivacyLookupWizardLines{}
	if err := c.SearchRead(PrivacyLookupWizardLineModel, criteria, NewOptions().Limit(1), plwls); err != nil {
		return nil, err
	}
	return &((*plwls)[0]), nil
}

// FindPrivacyLookupWizardLines finds privacy.lookup.wizard.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPrivacyLookupWizardLines(criteria *Criteria, options *Options) (*PrivacyLookupWizardLines, error) {
	plwls := &PrivacyLookupWizardLines{}
	if err := c.SearchRead(PrivacyLookupWizardLineModel, criteria, options, plwls); err != nil {
		return nil, err
	}
	return plwls, nil
}

// FindPrivacyLookupWizardLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPrivacyLookupWizardLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(PrivacyLookupWizardLineModel, criteria, options)
}

// FindPrivacyLookupWizardLineId finds record id by querying it with criteria.
func (c *Client) FindPrivacyLookupWizardLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PrivacyLookupWizardLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
