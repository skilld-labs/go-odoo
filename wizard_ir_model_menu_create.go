package odoo

import (
	"fmt"
)

// WizardIrModelMenuCreate represents wizard.ir.model.menu.create model.
type WizardIrModelMenuCreate struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	MenuId      *Many2One `xmlrpc:"menu_id,omptempty"`
	Name        *String   `xmlrpc:"name,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
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
	return c.Create(WizardIrModelMenuCreateModel, wimmc)
}

// UpdateWizardIrModelMenuCreate updates an existing wizard.ir.model.menu.create record.
func (c *Client) UpdateWizardIrModelMenuCreate(wimmc *WizardIrModelMenuCreate) error {
	return c.UpdateWizardIrModelMenuCreates([]int64{wimmc.Id.Get()}, wimmc)
}

// UpdateWizardIrModelMenuCreates updates existing wizard.ir.model.menu.create records.
// All records (represented by ids) will be updated by wimmc values.
func (c *Client) UpdateWizardIrModelMenuCreates(ids []int64, wimmc *WizardIrModelMenuCreate) error {
	return c.Update(WizardIrModelMenuCreateModel, ids, wimmc)
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
	if wimmcs != nil && len(*wimmcs) > 0 {
		return &((*wimmcs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of wizard.ir.model.menu.create not found", id)
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
	if wimmcs != nil && len(*wimmcs) > 0 {
		return &((*wimmcs)[0]), nil
	}
	return nil, fmt.Errorf("no wizard.ir.model.menu.create was found with criteria %v", criteria)
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
	ids, err := c.Search(WizardIrModelMenuCreateModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindWizardIrModelMenuCreateId finds record id by querying it with criteria.
func (c *Client) FindWizardIrModelMenuCreateId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(WizardIrModelMenuCreateModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no wizard.ir.model.menu.create was found with criteria %v and options %v", criteria, options)
}
