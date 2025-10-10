package odoo

// DigestTip represents digest.tip model.
type DigestTip struct {
	CreateDate     *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid      *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName    *String   `xmlrpc:"display_name,omitempty"`
	GroupId        *Many2One `xmlrpc:"group_id,omitempty"`
	Id             *Int      `xmlrpc:"id,omitempty"`
	Name           *String   `xmlrpc:"name,omitempty"`
	Sequence       *Int      `xmlrpc:"sequence,omitempty"`
	TipDescription *String   `xmlrpc:"tip_description,omitempty"`
	UserIds        *Relation `xmlrpc:"user_ids,omitempty"`
	WriteDate      *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid       *Many2One `xmlrpc:"write_uid,omitempty"`
}

// DigestTips represents array of digest.tip model.
type DigestTips []DigestTip

// DigestTipModel is the odoo model name.
const DigestTipModel = "digest.tip"

// Many2One convert DigestTip to *Many2One.
func (dt *DigestTip) Many2One() *Many2One {
	return NewMany2One(dt.Id.Get(), "")
}

// CreateDigestTip creates a new digest.tip model and returns its id.
func (c *Client) CreateDigestTip(dt *DigestTip) (int64, error) {
	ids, err := c.CreateDigestTips([]*DigestTip{dt})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateDigestTip creates a new digest.tip model and returns its id.
func (c *Client) CreateDigestTips(dts []*DigestTip) ([]int64, error) {
	var vv []interface{}
	for _, v := range dts {
		vv = append(vv, v)
	}
	return c.Create(DigestTipModel, vv, nil)
}

// UpdateDigestTip updates an existing digest.tip record.
func (c *Client) UpdateDigestTip(dt *DigestTip) error {
	return c.UpdateDigestTips([]int64{dt.Id.Get()}, dt)
}

// UpdateDigestTips updates existing digest.tip records.
// All records (represented by ids) will be updated by dt values.
func (c *Client) UpdateDigestTips(ids []int64, dt *DigestTip) error {
	return c.Update(DigestTipModel, ids, dt, nil)
}

// DeleteDigestTip deletes an existing digest.tip record.
func (c *Client) DeleteDigestTip(id int64) error {
	return c.DeleteDigestTips([]int64{id})
}

// DeleteDigestTips deletes existing digest.tip records.
func (c *Client) DeleteDigestTips(ids []int64) error {
	return c.Delete(DigestTipModel, ids)
}

// GetDigestTip gets digest.tip existing record.
func (c *Client) GetDigestTip(id int64) (*DigestTip, error) {
	dts, err := c.GetDigestTips([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*dts)[0]), nil
}

// GetDigestTips gets digest.tip existing records.
func (c *Client) GetDigestTips(ids []int64) (*DigestTips, error) {
	dts := &DigestTips{}
	if err := c.Read(DigestTipModel, ids, nil, dts); err != nil {
		return nil, err
	}
	return dts, nil
}

// FindDigestTip finds digest.tip record by querying it with criteria.
func (c *Client) FindDigestTip(criteria *Criteria) (*DigestTip, error) {
	dts := &DigestTips{}
	if err := c.SearchRead(DigestTipModel, criteria, NewOptions().Limit(1), dts); err != nil {
		return nil, err
	}
	return &((*dts)[0]), nil
}

// FindDigestTips finds digest.tip records by querying it
// and filtering it with criteria and options.
func (c *Client) FindDigestTips(criteria *Criteria, options *Options) (*DigestTips, error) {
	dts := &DigestTips{}
	if err := c.SearchRead(DigestTipModel, criteria, options, dts); err != nil {
		return nil, err
	}
	return dts, nil
}

// FindDigestTipIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindDigestTipIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(DigestTipModel, criteria, options)
}

// FindDigestTipId finds record id by querying it with criteria.
func (c *Client) FindDigestTipId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(DigestTipModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
