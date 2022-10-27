package odoo

import (
	"fmt"
)

// ResetViewArchWizard represents reset.view.arch.wizard model.
type ResetViewArchWizard struct {
	LastUpdate    *Time      `xmlrpc:"__last_update,omitempty"`
	ArchDiff      *String    `xmlrpc:"arch_diff,omitempty"`
	ArchToCompare *String    `xmlrpc:"arch_to_compare,omitempty"`
	CompareViewId *Many2One  `xmlrpc:"compare_view_id,omitempty"`
	CreateDate    *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid     *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName   *String    `xmlrpc:"display_name,omitempty"`
	HasDiff       *Bool      `xmlrpc:"has_diff,omitempty"`
	Id            *Int       `xmlrpc:"id,omitempty"`
	ResetMode     *Selection `xmlrpc:"reset_mode,omitempty"`
	ViewId        *Many2One  `xmlrpc:"view_id,omitempty"`
	ViewName      *String    `xmlrpc:"view_name,omitempty"`
	WriteDate     *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid      *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// ResetViewArchWizards represents array of reset.view.arch.wizard model.
type ResetViewArchWizards []ResetViewArchWizard

// ResetViewArchWizardModel is the odoo model name.
const ResetViewArchWizardModel = "reset.view.arch.wizard"

// Many2One convert ResetViewArchWizard to *Many2One.
func (rvaw *ResetViewArchWizard) Many2One() *Many2One {
	return NewMany2One(rvaw.Id.Get(), "")
}

// CreateResetViewArchWizard creates a new reset.view.arch.wizard model and returns its id.
func (c *Client) CreateResetViewArchWizard(rvaw *ResetViewArchWizard) (int64, error) {
	return c.Create(ResetViewArchWizardModel, rvaw)
}

// UpdateResetViewArchWizard updates an existing reset.view.arch.wizard record.
func (c *Client) UpdateResetViewArchWizard(rvaw *ResetViewArchWizard) error {
	return c.UpdateResetViewArchWizards([]int64{rvaw.Id.Get()}, rvaw)
}

// UpdateResetViewArchWizards updates existing reset.view.arch.wizard records.
// All records (represented by ids) will be updated by rvaw values.
func (c *Client) UpdateResetViewArchWizards(ids []int64, rvaw *ResetViewArchWizard) error {
	return c.Update(ResetViewArchWizardModel, ids, rvaw)
}

// DeleteResetViewArchWizard deletes an existing reset.view.arch.wizard record.
func (c *Client) DeleteResetViewArchWizard(id int64) error {
	return c.DeleteResetViewArchWizards([]int64{id})
}

// DeleteResetViewArchWizards deletes existing reset.view.arch.wizard records.
func (c *Client) DeleteResetViewArchWizards(ids []int64) error {
	return c.Delete(ResetViewArchWizardModel, ids)
}

// GetResetViewArchWizard gets reset.view.arch.wizard existing record.
func (c *Client) GetResetViewArchWizard(id int64) (*ResetViewArchWizard, error) {
	rvaws, err := c.GetResetViewArchWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	if rvaws != nil && len(*rvaws) > 0 {
		return &((*rvaws)[0]), nil
	}
	return nil, fmt.Errorf("id %v of reset.view.arch.wizard not found", id)
}

// GetResetViewArchWizards gets reset.view.arch.wizard existing records.
func (c *Client) GetResetViewArchWizards(ids []int64) (*ResetViewArchWizards, error) {
	rvaws := &ResetViewArchWizards{}
	if err := c.Read(ResetViewArchWizardModel, ids, nil, rvaws); err != nil {
		return nil, err
	}
	return rvaws, nil
}

// FindResetViewArchWizard finds reset.view.arch.wizard record by querying it with criteria.
func (c *Client) FindResetViewArchWizard(criteria *Criteria) (*ResetViewArchWizard, error) {
	rvaws := &ResetViewArchWizards{}
	if err := c.SearchRead(ResetViewArchWizardModel, criteria, NewOptions().Limit(1), rvaws); err != nil {
		return nil, err
	}
	if rvaws != nil && len(*rvaws) > 0 {
		return &((*rvaws)[0]), nil
	}
	return nil, fmt.Errorf("reset.view.arch.wizard was not found")
}

// FindResetViewArchWizards finds reset.view.arch.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResetViewArchWizards(criteria *Criteria, options *Options) (*ResetViewArchWizards, error) {
	rvaws := &ResetViewArchWizards{}
	if err := c.SearchRead(ResetViewArchWizardModel, criteria, options, rvaws); err != nil {
		return nil, err
	}
	return rvaws, nil
}

// FindResetViewArchWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResetViewArchWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ResetViewArchWizardModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindResetViewArchWizardId finds record id by querying it with criteria.
func (c *Client) FindResetViewArchWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResetViewArchWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("reset.view.arch.wizard was not found")
}
