package odoo

import (
	"fmt"
)

// IrDemoFailureWizard represents ir.demo_failure.wizard model.
type IrDemoFailureWizard struct {
	LastUpdate    *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate    *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid     *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName   *String   `xmlrpc:"display_name,omitempty"`
	FailureIds    *Relation `xmlrpc:"failure_ids,omitempty"`
	FailuresCount *Int      `xmlrpc:"failures_count,omitempty"`
	Id            *Int      `xmlrpc:"id,omitempty"`
	WriteDate     *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid      *Many2One `xmlrpc:"write_uid,omitempty"`
}

// IrDemoFailureWizards represents array of ir.demo_failure.wizard model.
type IrDemoFailureWizards []IrDemoFailureWizard

// IrDemoFailureWizardModel is the odoo model name.
const IrDemoFailureWizardModel = "ir.demo_failure.wizard"

// Many2One convert IrDemoFailureWizard to *Many2One.
func (idw *IrDemoFailureWizard) Many2One() *Many2One {
	return NewMany2One(idw.Id.Get(), "")
}

// CreateIrDemoFailureWizard creates a new ir.demo_failure.wizard model and returns its id.
func (c *Client) CreateIrDemoFailureWizard(idw *IrDemoFailureWizard) (int64, error) {
	return c.Create(IrDemoFailureWizardModel, idw)
}

// UpdateIrDemoFailureWizard updates an existing ir.demo_failure.wizard record.
func (c *Client) UpdateIrDemoFailureWizard(idw *IrDemoFailureWizard) error {
	return c.UpdateIrDemoFailureWizards([]int64{idw.Id.Get()}, idw)
}

// UpdateIrDemoFailureWizards updates existing ir.demo_failure.wizard records.
// All records (represented by IDs) will be updated by idw values.
func (c *Client) UpdateIrDemoFailureWizards(ids []int64, idw *IrDemoFailureWizard) error {
	return c.Update(IrDemoFailureWizardModel, ids, idw)
}

// DeleteIrDemoFailureWizard deletes an existing ir.demo_failure.wizard record.
func (c *Client) DeleteIrDemoFailureWizard(id int64) error {
	return c.DeleteIrDemoFailureWizards([]int64{id})
}

// DeleteIrDemoFailureWizards deletes existing ir.demo_failure.wizard records.
func (c *Client) DeleteIrDemoFailureWizards(ids []int64) error {
	return c.Delete(IrDemoFailureWizardModel, ids)
}

// GetIrDemoFailureWizard gets ir.demo_failure.wizard existing record.
func (c *Client) GetIrDemoFailureWizard(id int64) (*IrDemoFailureWizard, error) {
	idws, err := c.GetIrDemoFailureWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	if idws != nil && len(*idws) > 0 {
		return &((*idws)[0]), nil
	}
	return nil, fmt.Errorf("id %v of ir.demo_failure.wizard not found", id)
}

// GetIrDemoFailureWizards gets ir.demo_failure.wizard existing records.
func (c *Client) GetIrDemoFailureWizards(ids []int64) (*IrDemoFailureWizards, error) {
	idws := &IrDemoFailureWizards{}
	if err := c.Read(IrDemoFailureWizardModel, ids, nil, idws); err != nil {
		return nil, err
	}
	return idws, nil
}

// FindIrDemoFailureWizard finds ir.demo_failure.wizard record by querying it with criteria.
func (c *Client) FindIrDemoFailureWizard(criteria *Criteria) (*IrDemoFailureWizard, error) {
	idws := &IrDemoFailureWizards{}
	if err := c.SearchRead(IrDemoFailureWizardModel, criteria, NewOptions().Limit(1), idws); err != nil {
		return nil, err
	}
	if idws != nil && len(*idws) > 0 {
		return &((*idws)[0]), nil
	}
	return nil, fmt.Errorf("ir.demo_failure.wizard was not found")
}

// FindIrDemoFailureWizards finds ir.demo_failure.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrDemoFailureWizards(criteria *Criteria, options *Options) (*IrDemoFailureWizards, error) {
	idws := &IrDemoFailureWizards{}
	if err := c.SearchRead(IrDemoFailureWizardModel, criteria, options, idws); err != nil {
		return nil, err
	}
	return idws, nil
}

// FindIrDemoFailureWizardIds finds records IDs by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrDemoFailureWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(IrDemoFailureWizardModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindIrDemoFailureWizardId finds record id by querying it with criteria.
func (c *Client) FindIrDemoFailureWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrDemoFailureWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("ir.demo_failure.wizard was not found")
}
