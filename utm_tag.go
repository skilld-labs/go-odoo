package odoo

// UtmTag represents utm.tag model.
type UtmTag struct {
	Color       *Int      `xmlrpc:"color,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// UtmTags represents array of utm.tag model.
type UtmTags []UtmTag

// UtmTagModel is the odoo model name.
const UtmTagModel = "utm.tag"

// Many2One convert UtmTag to *Many2One.
func (ut *UtmTag) Many2One() *Many2One {
	return NewMany2One(ut.Id.Get(), "")
}

// CreateUtmTag creates a new utm.tag model and returns its id.
func (c *Client) CreateUtmTag(ut *UtmTag) (int64, error) {
	ids, err := c.CreateUtmTags([]*UtmTag{ut})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateUtmTag creates a new utm.tag model and returns its id.
func (c *Client) CreateUtmTags(uts []*UtmTag) ([]int64, error) {
	var vv []interface{}
	for _, v := range uts {
		vv = append(vv, v)
	}
	return c.Create(UtmTagModel, vv, nil)
}

// UpdateUtmTag updates an existing utm.tag record.
func (c *Client) UpdateUtmTag(ut *UtmTag) error {
	return c.UpdateUtmTags([]int64{ut.Id.Get()}, ut)
}

// UpdateUtmTags updates existing utm.tag records.
// All records (represented by ids) will be updated by ut values.
func (c *Client) UpdateUtmTags(ids []int64, ut *UtmTag) error {
	return c.Update(UtmTagModel, ids, ut, nil)
}

// DeleteUtmTag deletes an existing utm.tag record.
func (c *Client) DeleteUtmTag(id int64) error {
	return c.DeleteUtmTags([]int64{id})
}

// DeleteUtmTags deletes existing utm.tag records.
func (c *Client) DeleteUtmTags(ids []int64) error {
	return c.Delete(UtmTagModel, ids)
}

// GetUtmTag gets utm.tag existing record.
func (c *Client) GetUtmTag(id int64) (*UtmTag, error) {
	uts, err := c.GetUtmTags([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*uts)[0]), nil
}

// GetUtmTags gets utm.tag existing records.
func (c *Client) GetUtmTags(ids []int64) (*UtmTags, error) {
	uts := &UtmTags{}
	if err := c.Read(UtmTagModel, ids, nil, uts); err != nil {
		return nil, err
	}
	return uts, nil
}

// FindUtmTag finds utm.tag record by querying it with criteria.
func (c *Client) FindUtmTag(criteria *Criteria) (*UtmTag, error) {
	uts := &UtmTags{}
	if err := c.SearchRead(UtmTagModel, criteria, NewOptions().Limit(1), uts); err != nil {
		return nil, err
	}
	return &((*uts)[0]), nil
}

// FindUtmTags finds utm.tag records by querying it
// and filtering it with criteria and options.
func (c *Client) FindUtmTags(criteria *Criteria, options *Options) (*UtmTags, error) {
	uts := &UtmTags{}
	if err := c.SearchRead(UtmTagModel, criteria, options, uts); err != nil {
		return nil, err
	}
	return uts, nil
}

// FindUtmTagIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindUtmTagIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(UtmTagModel, criteria, options)
}

// FindUtmTagId finds record id by querying it with criteria.
func (c *Client) FindUtmTagId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(UtmTagModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
