package odoo

// ResRequestLink represents res.request.link model.
type ResRequestLink struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	Object      *String   `xmlrpc:"object,omitempty"`
	Priority    *Int      `xmlrpc:"priority,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// ResRequestLinks represents array of res.request.link model.
type ResRequestLinks []ResRequestLink

// ResRequestLinkModel is the odoo model name.
const ResRequestLinkModel = "res.request.link"

// Many2One convert ResRequestLink to *Many2One.
func (rrl *ResRequestLink) Many2One() *Many2One {
	return NewMany2One(rrl.Id.Get(), "")
}

// CreateResRequestLink creates a new res.request.link model and returns its id.
func (c *Client) CreateResRequestLink(rrl *ResRequestLink) (int64, error) {
	ids, err := c.CreateResRequestLinks([]*ResRequestLink{rrl})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateResRequestLinks creates a new res.request.link model and returns its id.
func (c *Client) CreateResRequestLinks(rrls []*ResRequestLink) ([]int64, error) {
	var vv []interface{}
	for _, v := range rrls {
		vv = append(vv, v)
	}
	return c.Create(ResRequestLinkModel, vv, nil)
}

// UpdateResRequestLink updates an existing res.request.link record.
func (c *Client) UpdateResRequestLink(rrl *ResRequestLink) error {
	return c.UpdateResRequestLinks([]int64{rrl.Id.Get()}, rrl)
}

// UpdateResRequestLinks updates existing res.request.link records.
// All records (represented by ids) will be updated by rrl values.
func (c *Client) UpdateResRequestLinks(ids []int64, rrl *ResRequestLink) error {
	return c.Update(ResRequestLinkModel, ids, rrl, nil)
}

// DeleteResRequestLink deletes an existing res.request.link record.
func (c *Client) DeleteResRequestLink(id int64) error {
	return c.DeleteResRequestLinks([]int64{id})
}

// DeleteResRequestLinks deletes existing res.request.link records.
func (c *Client) DeleteResRequestLinks(ids []int64) error {
	return c.Delete(ResRequestLinkModel, ids)
}

// GetResRequestLink gets res.request.link existing record.
func (c *Client) GetResRequestLink(id int64) (*ResRequestLink, error) {
	rrls, err := c.GetResRequestLinks([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*rrls)[0]), nil
}

// GetResRequestLinks gets res.request.link existing records.
func (c *Client) GetResRequestLinks(ids []int64) (*ResRequestLinks, error) {
	rrls := &ResRequestLinks{}
	if err := c.Read(ResRequestLinkModel, ids, nil, rrls); err != nil {
		return nil, err
	}
	return rrls, nil
}

// FindResRequestLink finds res.request.link record by querying it with criteria.
func (c *Client) FindResRequestLink(criteria *Criteria) (*ResRequestLink, error) {
	rrls := &ResRequestLinks{}
	if err := c.SearchRead(ResRequestLinkModel, criteria, NewOptions().Limit(1), rrls); err != nil {
		return nil, err
	}
	return &((*rrls)[0]), nil
}

// FindResRequestLinks finds res.request.link records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResRequestLinks(criteria *Criteria, options *Options) (*ResRequestLinks, error) {
	rrls := &ResRequestLinks{}
	if err := c.SearchRead(ResRequestLinkModel, criteria, options, rrls); err != nil {
		return nil, err
	}
	return rrls, nil
}

// FindResRequestLinkIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResRequestLinkIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ResRequestLinkModel, criteria, options)
}

// FindResRequestLinkId finds record id by querying it with criteria.
func (c *Client) FindResRequestLinkId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResRequestLinkModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
