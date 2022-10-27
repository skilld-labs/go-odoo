package odoo

import (
	"fmt"
)

// PortalShare represents portal.share model.
type PortalShare struct {
	LastUpdate    *Time     `xmlrpc:"__last_update,omitempty"`
	AccessWarning *String   `xmlrpc:"access_warning,omitempty"`
	CreateDate    *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid     *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName   *String   `xmlrpc:"display_name,omitempty"`
	Id            *Int      `xmlrpc:"id,omitempty"`
	Note          *String   `xmlrpc:"note,omitempty"`
	PartnerIds    *Relation `xmlrpc:"partner_ids,omitempty"`
	ResId         *Int      `xmlrpc:"res_id,omitempty"`
	ResModel      *String   `xmlrpc:"res_model,omitempty"`
	ShareLink     *String   `xmlrpc:"share_link,omitempty"`
	WriteDate     *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid      *Many2One `xmlrpc:"write_uid,omitempty"`
}

// PortalShares represents array of portal.share model.
type PortalShares []PortalShare

// PortalShareModel is the odoo model name.
const PortalShareModel = "portal.share"

// Many2One convert PortalShare to *Many2One.
func (ps *PortalShare) Many2One() *Many2One {
	return NewMany2One(ps.Id.Get(), "")
}

// CreatePortalShare creates a new portal.share model and returns its id.
func (c *Client) CreatePortalShare(ps *PortalShare) (int64, error) {
	return c.Create(PortalShareModel, ps)
}

// UpdatePortalShare updates an existing portal.share record.
func (c *Client) UpdatePortalShare(ps *PortalShare) error {
	return c.UpdatePortalShares([]int64{ps.Id.Get()}, ps)
}

// UpdatePortalShares updates existing portal.share records.
// All records (represented by ids) will be updated by ps values.
func (c *Client) UpdatePortalShares(ids []int64, ps *PortalShare) error {
	return c.Update(PortalShareModel, ids, ps)
}

// DeletePortalShare deletes an existing portal.share record.
func (c *Client) DeletePortalShare(id int64) error {
	return c.DeletePortalShares([]int64{id})
}

// DeletePortalShares deletes existing portal.share records.
func (c *Client) DeletePortalShares(ids []int64) error {
	return c.Delete(PortalShareModel, ids)
}

// GetPortalShare gets portal.share existing record.
func (c *Client) GetPortalShare(id int64) (*PortalShare, error) {
	pss, err := c.GetPortalShares([]int64{id})
	if err != nil {
		return nil, err
	}
	if pss != nil && len(*pss) > 0 {
		return &((*pss)[0]), nil
	}
	return nil, fmt.Errorf("id %v of portal.share not found", id)
}

// GetPortalShares gets portal.share existing records.
func (c *Client) GetPortalShares(ids []int64) (*PortalShares, error) {
	pss := &PortalShares{}
	if err := c.Read(PortalShareModel, ids, nil, pss); err != nil {
		return nil, err
	}
	return pss, nil
}

// FindPortalShare finds portal.share record by querying it with criteria.
func (c *Client) FindPortalShare(criteria *Criteria) (*PortalShare, error) {
	pss := &PortalShares{}
	if err := c.SearchRead(PortalShareModel, criteria, NewOptions().Limit(1), pss); err != nil {
		return nil, err
	}
	if pss != nil && len(*pss) > 0 {
		return &((*pss)[0]), nil
	}
	return nil, fmt.Errorf("portal.share was not found")
}

// FindPortalShares finds portal.share records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPortalShares(criteria *Criteria, options *Options) (*PortalShares, error) {
	pss := &PortalShares{}
	if err := c.SearchRead(PortalShareModel, criteria, options, pss); err != nil {
		return nil, err
	}
	return pss, nil
}

// FindPortalShareIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPortalShareIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(PortalShareModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindPortalShareId finds record id by querying it with criteria.
func (c *Client) FindPortalShareId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PortalShareModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("portal.share was not found")
}
