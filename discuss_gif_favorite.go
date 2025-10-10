package odoo

// DiscussGifFavorite represents discuss.gif.favorite model.
type DiscussGifFavorite struct {
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	TenorGifId  *String   `xmlrpc:"tenor_gif_id,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// DiscussGifFavorites represents array of discuss.gif.favorite model.
type DiscussGifFavorites []DiscussGifFavorite

// DiscussGifFavoriteModel is the odoo model name.
const DiscussGifFavoriteModel = "discuss.gif.favorite"

// Many2One convert DiscussGifFavorite to *Many2One.
func (dgf *DiscussGifFavorite) Many2One() *Many2One {
	return NewMany2One(dgf.Id.Get(), "")
}

// CreateDiscussGifFavorite creates a new discuss.gif.favorite model and returns its id.
func (c *Client) CreateDiscussGifFavorite(dgf *DiscussGifFavorite) (int64, error) {
	ids, err := c.CreateDiscussGifFavorites([]*DiscussGifFavorite{dgf})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateDiscussGifFavorite creates a new discuss.gif.favorite model and returns its id.
func (c *Client) CreateDiscussGifFavorites(dgfs []*DiscussGifFavorite) ([]int64, error) {
	var vv []interface{}
	for _, v := range dgfs {
		vv = append(vv, v)
	}
	return c.Create(DiscussGifFavoriteModel, vv, nil)
}

// UpdateDiscussGifFavorite updates an existing discuss.gif.favorite record.
func (c *Client) UpdateDiscussGifFavorite(dgf *DiscussGifFavorite) error {
	return c.UpdateDiscussGifFavorites([]int64{dgf.Id.Get()}, dgf)
}

// UpdateDiscussGifFavorites updates existing discuss.gif.favorite records.
// All records (represented by ids) will be updated by dgf values.
func (c *Client) UpdateDiscussGifFavorites(ids []int64, dgf *DiscussGifFavorite) error {
	return c.Update(DiscussGifFavoriteModel, ids, dgf, nil)
}

// DeleteDiscussGifFavorite deletes an existing discuss.gif.favorite record.
func (c *Client) DeleteDiscussGifFavorite(id int64) error {
	return c.DeleteDiscussGifFavorites([]int64{id})
}

// DeleteDiscussGifFavorites deletes existing discuss.gif.favorite records.
func (c *Client) DeleteDiscussGifFavorites(ids []int64) error {
	return c.Delete(DiscussGifFavoriteModel, ids)
}

// GetDiscussGifFavorite gets discuss.gif.favorite existing record.
func (c *Client) GetDiscussGifFavorite(id int64) (*DiscussGifFavorite, error) {
	dgfs, err := c.GetDiscussGifFavorites([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*dgfs)[0]), nil
}

// GetDiscussGifFavorites gets discuss.gif.favorite existing records.
func (c *Client) GetDiscussGifFavorites(ids []int64) (*DiscussGifFavorites, error) {
	dgfs := &DiscussGifFavorites{}
	if err := c.Read(DiscussGifFavoriteModel, ids, nil, dgfs); err != nil {
		return nil, err
	}
	return dgfs, nil
}

// FindDiscussGifFavorite finds discuss.gif.favorite record by querying it with criteria.
func (c *Client) FindDiscussGifFavorite(criteria *Criteria) (*DiscussGifFavorite, error) {
	dgfs := &DiscussGifFavorites{}
	if err := c.SearchRead(DiscussGifFavoriteModel, criteria, NewOptions().Limit(1), dgfs); err != nil {
		return nil, err
	}
	return &((*dgfs)[0]), nil
}

// FindDiscussGifFavorites finds discuss.gif.favorite records by querying it
// and filtering it with criteria and options.
func (c *Client) FindDiscussGifFavorites(criteria *Criteria, options *Options) (*DiscussGifFavorites, error) {
	dgfs := &DiscussGifFavorites{}
	if err := c.SearchRead(DiscussGifFavoriteModel, criteria, options, dgfs); err != nil {
		return nil, err
	}
	return dgfs, nil
}

// FindDiscussGifFavoriteIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindDiscussGifFavoriteIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(DiscussGifFavoriteModel, criteria, options)
}

// FindDiscussGifFavoriteId finds record id by querying it with criteria.
func (c *Client) FindDiscussGifFavoriteId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(DiscussGifFavoriteModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
