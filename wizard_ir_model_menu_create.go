package odoo

// WizardIrModelMenuCreate represents wizard.ir.model.menu.create model.
type WizardIrModelMenuCreate struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	MenuId      *Many2One `xmlrpc:"menu_id,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// WizardIrModelMenuCreates represents array of wizard.ir.model.menu.create model.
type WizardIrModelMenuCreates []WizardIrModelMenuCreate

// WizardIrModelMenuCreateModel is the odoo model name.
const WizardIrModelMenuCreateModel = "wizard.ir.model.menu.create"

// Many2One convert WizardIrModelMenuCreate to *Many2One.
func (wimmc *WizardIrModelMenuCreate) Many2One() *Many2One {
	return NewMany2One(wimmc.Id.Get(), "")
}

// CreateWizardIrModelMenuCreate creates a new wizard.ir.model.menu.create model and returns its id.
func (c *Client) CreateWizardIrModelMenuCreate(wimmc *WizardIrModelMenuCreate) (int64, error) {
	ids, err := c.CreateWizardIrModelMenuCreates([]*WizardIrModelMenuCreate{wimmc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateWizardIrModelMenuCreates creates a new wizard.ir.model.menu.create model and returns its id.
func (c *Client) CreateWizardIrModelMenuCreates(wimmcs []*WizardIrModelMenuCreate) ([]int64, error) {
	var vv []interface{}
	for _, v := range wimmcs {
		vv = append(vv, v)
	}
	return c.Create(WizardIrModelMenuCreateModel, vv, nil)
}

// UpdateWizardIrModelMenuCreate updates an existing wizard.ir.model.menu.create record.
func (c *Client) UpdateWizardIrModelMenuCreate(wimmc *WizardIrModelMenuCreate) error {
	return c.UpdateWizardIrModelMenuCreates([]int64{wimmc.Id.Get()}, wimmc)
}

// UpdateWizardIrModelMenuCreates updates existing wizard.ir.model.menu.create records.
// All records (represented by ids) will be updated by wimmc values.
func (c *Client) UpdateWizardIrModelMenuCreates(ids []int64, wimmc *WizardIrModelMenuCreate) error {
	return c.Update(WizardIrModelMenuCreateModel, ids, wimmc, nil)
}

// DeleteWizardIrModelMenuCreate deletes an existing wizard.ir.model.menu.create record.
func (c *Client) DeleteWizardIrModelMenuCreate(id int64) error {
	return c.DeleteWizardIrModelMenuCreates([]int64{id})
}

// DeleteWizardIrModelMenuCreates deletes existing wizard.ir.model.menu.create records.
func (c *Client) DeleteWizardIrModelMenuCreates(ids []int64) error {
	return c.Delete(WizardIrModelMenuCreateModel, ids)
}

// GetWizardIrModelMenuCreate gets wizard.ir.model.menu.create existing record.
func (c *Client) GetWizardIrModelMenuCreate(id int64) (*WizardIrModelMenuCreate, error) {
	wimmcs, err := c.GetWizardIrModelMenuCreates([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*wimmcs)[0]), nil
}

// GetWizardIrModelMenuCreates gets wizard.ir.model.menu.create existing records.
func (c *Client) GetWizardIrModelMenuCreates(ids []int64) (*WizardIrModelMenuCreates, error) {
	wimmcs := &WizardIrModelMenuCreates{}
	if err := c.Read(WizardIrModelMenuCreateModel, ids, nil, wimmcs); err != nil {
		return nil, err
	}
	return wimmcs, nil
}

// FindWizardIrModelMenuCreate finds wizard.ir.model.menu.create record by querying it with criteria.
func (c *Client) FindWizardIrModelMenuCreate(criteria *Criteria) (*WizardIrModelMenuCreate, error) {
	wimmcs := &WizardIrModelMenuCreates{}
	if err := c.SearchRead(WizardIrModelMenuCreateModel, criteria, NewOptions().Limit(1), wimmcs); err != nil {
		return nil, err
	}
	return &((*wimmcs)[0]), nil
}

// FindWizardIrModelMenuCreates finds wizard.ir.model.menu.create records by querying it
// and filtering it with criteria and options.
func (c *Client) FindWizardIrModelMenuCreates(criteria *Criteria, options *Options) (*WizardIrModelMenuCreates, error) {
	wimmcs := &WizardIrModelMenuCreates{}
	if err := c.SearchRead(WizardIrModelMenuCreateModel, criteria, options, wimmcs); err != nil {
		return nil, err
	}
	return wimmcs, nil
}

// FindWizardIrModelMenuCreateIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindWizardIrModelMenuCreateIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(WizardIrModelMenuCreateModel, criteria, options)
}

// FindWizardIrModelMenuCreateId finds record id by querying it with criteria.
func (c *Client) FindWizardIrModelMenuCreateId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(WizardIrModelMenuCreateModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
