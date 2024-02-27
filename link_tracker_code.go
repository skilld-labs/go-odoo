package odoo

// LinkTrackerCode represents link.tracker.code model.
type LinkTrackerCode struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	Code        *String   `xmlrpc:"code,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	LinkId      *Many2One `xmlrpc:"link_id,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// LinkTrackerCodes represents array of link.tracker.code model.
type LinkTrackerCodes []LinkTrackerCode

// LinkTrackerCodeModel is the odoo model name.
const LinkTrackerCodeModel = "link.tracker.code"

// Many2One convert LinkTrackerCode to *Many2One.
func (ltc *LinkTrackerCode) Many2One() *Many2One {
	return NewMany2One(ltc.Id.Get(), "")
}

// CreateLinkTrackerCode creates a new link.tracker.code model and returns its id.
func (c *Client) CreateLinkTrackerCode(ltc *LinkTrackerCode) (int64, error) {
	ids, err := c.CreateLinkTrackerCodes([]*LinkTrackerCode{ltc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateLinkTrackerCodes creates a new link.tracker.code model and returns its id.
func (c *Client) CreateLinkTrackerCodes(ltcs []*LinkTrackerCode) ([]int64, error) {
	var vv []interface{}
	for _, v := range ltcs {
		vv = append(vv, v)
	}
	return c.Create(LinkTrackerCodeModel, vv, nil)
}

// UpdateLinkTrackerCode updates an existing link.tracker.code record.
func (c *Client) UpdateLinkTrackerCode(ltc *LinkTrackerCode) error {
	return c.UpdateLinkTrackerCodes([]int64{ltc.Id.Get()}, ltc)
}

// UpdateLinkTrackerCodes updates existing link.tracker.code records.
// All records (represented by ids) will be updated by ltc values.
func (c *Client) UpdateLinkTrackerCodes(ids []int64, ltc *LinkTrackerCode) error {
	return c.Update(LinkTrackerCodeModel, ids, ltc, nil)
}

// DeleteLinkTrackerCode deletes an existing link.tracker.code record.
func (c *Client) DeleteLinkTrackerCode(id int64) error {
	return c.DeleteLinkTrackerCodes([]int64{id})
}

// DeleteLinkTrackerCodes deletes existing link.tracker.code records.
func (c *Client) DeleteLinkTrackerCodes(ids []int64) error {
	return c.Delete(LinkTrackerCodeModel, ids)
}

// GetLinkTrackerCode gets link.tracker.code existing record.
func (c *Client) GetLinkTrackerCode(id int64) (*LinkTrackerCode, error) {
	ltcs, err := c.GetLinkTrackerCodes([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ltcs)[0]), nil
}

// GetLinkTrackerCodes gets link.tracker.code existing records.
func (c *Client) GetLinkTrackerCodes(ids []int64) (*LinkTrackerCodes, error) {
	ltcs := &LinkTrackerCodes{}
	if err := c.Read(LinkTrackerCodeModel, ids, nil, ltcs); err != nil {
		return nil, err
	}
	return ltcs, nil
}

// FindLinkTrackerCode finds link.tracker.code record by querying it with criteria.
func (c *Client) FindLinkTrackerCode(criteria *Criteria) (*LinkTrackerCode, error) {
	ltcs := &LinkTrackerCodes{}
	if err := c.SearchRead(LinkTrackerCodeModel, criteria, NewOptions().Limit(1), ltcs); err != nil {
		return nil, err
	}
	return &((*ltcs)[0]), nil
}

// FindLinkTrackerCodes finds link.tracker.code records by querying it
// and filtering it with criteria and options.
func (c *Client) FindLinkTrackerCodes(criteria *Criteria, options *Options) (*LinkTrackerCodes, error) {
	ltcs := &LinkTrackerCodes{}
	if err := c.SearchRead(LinkTrackerCodeModel, criteria, options, ltcs); err != nil {
		return nil, err
	}
	return ltcs, nil
}

// FindLinkTrackerCodeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindLinkTrackerCodeIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(LinkTrackerCodeModel, criteria, options)
}

// FindLinkTrackerCodeId finds record id by querying it with criteria.
func (c *Client) FindLinkTrackerCodeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(LinkTrackerCodeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
