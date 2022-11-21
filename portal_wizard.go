package odoo

import (
	"fmt"
)

// PortalWizard represents portal.wizard model.
type PortalWizard struct {
	LastUpdate     *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate     *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid      *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName    *String   `xmlrpc:"display_name,omptempty"`
	Id             *Int      `xmlrpc:"id,omptempty"`
	PortalId       *Many2One `xmlrpc:"portal_id,omptempty"`
	UserIds        *Relation `xmlrpc:"user_ids,omptempty"`
	WelcomeMessage *String   `xmlrpc:"welcome_message,omptempty"`
	WriteDate      *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid       *Many2One `xmlrpc:"write_uid,omptempty"`
}

// PortalWizards represents array of portal.wizard model.
type PortalWizards []PortalWizard

// PortalWizardModel is the odoo model name.
const PortalWizardModel = "portal.wizard"

// Many2One convert PortalWizard to *Many2One.
func (pw *PortalWizard) Many2One() *Many2One {
	return NewMany2One(pw.Id.Get(), "")
}

// CreatePortalWizard creates a new portal.wizard model and returns its id.
func (c *Client) CreatePortalWizard(pw *PortalWizard) (int64, error) {
	return c.Create(PortalWizardModel, pw)
}

// UpdatePortalWizard updates an existing portal.wizard record.
func (c *Client) UpdatePortalWizard(pw *PortalWizard) error {
	return c.UpdatePortalWizards([]int64{pw.Id.Get()}, pw)
}

// UpdatePortalWizards updates existing portal.wizard records.
// All records (represented by ids) will be updated by pw values.
func (c *Client) UpdatePortalWizards(ids []int64, pw *PortalWizard) error {
	return c.Update(PortalWizardModel, ids, pw)
}

// DeletePortalWizard deletes an existing portal.wizard record.
func (c *Client) DeletePortalWizard(id int64) error {
	return c.DeletePortalWizards([]int64{id})
}

// DeletePortalWizards deletes existing portal.wizard records.
func (c *Client) DeletePortalWizards(ids []int64) error {
	return c.Delete(PortalWizardModel, ids)
}

// GetPortalWizard gets portal.wizard existing record.
func (c *Client) GetPortalWizard(id int64) (*PortalWizard, error) {
	pws, err := c.GetPortalWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	if pws != nil && len(*pws) > 0 {
		return &((*pws)[0]), nil
	}
	return nil, fmt.Errorf("id %v of portal.wizard not found", id)
}

// GetPortalWizards gets portal.wizard existing records.
func (c *Client) GetPortalWizards(ids []int64) (*PortalWizards, error) {
	pws := &PortalWizards{}
	if err := c.Read(PortalWizardModel, ids, nil, pws); err != nil {
		return nil, err
	}
	return pws, nil
}

// FindPortalWizard finds portal.wizard record by querying it with criteria.
func (c *Client) FindPortalWizard(criteria *Criteria) (*PortalWizard, error) {
	pws := &PortalWizards{}
	if err := c.SearchRead(PortalWizardModel, criteria, NewOptions().Limit(1), pws); err != nil {
		return nil, err
	}
	if pws != nil && len(*pws) > 0 {
		return &((*pws)[0]), nil
	}
	return nil, fmt.Errorf("no portal.wizard was found with criteria %v", criteria)
}

// FindPortalWizards finds portal.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPortalWizards(criteria *Criteria, options *Options) (*PortalWizards, error) {
	pws := &PortalWizards{}
	if err := c.SearchRead(PortalWizardModel, criteria, options, pws); err != nil {
		return nil, err
	}
	return pws, nil
}

// FindPortalWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPortalWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(PortalWizardModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindPortalWizardId finds record id by querying it with criteria.
func (c *Client) FindPortalWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PortalWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no portal.wizard was found with criteria %v and options %v", criteria, options)
}
