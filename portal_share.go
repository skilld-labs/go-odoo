package odoo

// PortalShare represents portal.share model.
type PortalShare struct {
	AccessWarning *String   `xmlrpc:"access_warning,omitempty"`
	CreateDate    *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid     *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName   *String   `xmlrpc:"display_name,omitempty"`
	Id            *Int      `xmlrpc:"id,omitempty"`
	Note          *String   `xmlrpc:"note,omitempty"`
	PartnerIds    *Relation `xmlrpc:"partner_ids,omitempty"`
	ResId         *Int      `xmlrpc:"res_id,omitempty"`
	ResModel      *String   `xmlrpc:"res_model,omitempty"`
	ResourceRef   *String   `xmlrpc:"resource_ref,omitempty"`
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
	ids, err := c.CreatePortalShares([]*PortalShare{ps})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreatePortalShare creates a new portal.share model and returns its id.
func (c *Client) CreatePortalShares(pss []*PortalShare) ([]int64, error) {
	var vv []interface{}
	for _, v := range pss {
		vv = append(vv, v)
	}
	return c.Create(PortalShareModel, vv, nil)
}

// UpdatePortalShare updates an existing portal.share record.
func (c *Client) UpdatePortalShare(ps *PortalShare) error {
	return c.UpdatePortalShares([]int64{ps.Id.Get()}, ps)
}

// UpdatePortalShares updates existing portal.share records.
// All records (represented by ids) will be updated by ps values.
func (c *Client) UpdatePortalShares(ids []int64, ps *PortalShare) error {
	return c.Update(PortalShareModel, ids, ps, nil)
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
	return &((*pss)[0]), nil
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
	return &((*pss)[0]), nil
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
	return c.Search(PortalShareModel, criteria, options)
}

// FindPortalShareId finds record id by querying it with criteria.
func (c *Client) FindPortalShareId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PortalShareModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
