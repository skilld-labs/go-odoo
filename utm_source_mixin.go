package odoo

// UtmSourceMixin represents utm.source.mixin model.
type UtmSourceMixin struct {
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	SourceId    *Many2One `xmlrpc:"source_id,omitempty"`
}

// UtmSourceMixins represents array of utm.source.mixin model.
type UtmSourceMixins []UtmSourceMixin

// UtmSourceMixinModel is the odoo model name.
const UtmSourceMixinModel = "utm.source.mixin"

// Many2One convert UtmSourceMixin to *Many2One.
func (usm *UtmSourceMixin) Many2One() *Many2One {
	return NewMany2One(usm.Id.Get(), "")
}

// CreateUtmSourceMixin creates a new utm.source.mixin model and returns its id.
func (c *Client) CreateUtmSourceMixin(usm *UtmSourceMixin) (int64, error) {
	ids, err := c.CreateUtmSourceMixins([]*UtmSourceMixin{usm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateUtmSourceMixin creates a new utm.source.mixin model and returns its id.
func (c *Client) CreateUtmSourceMixins(usms []*UtmSourceMixin) ([]int64, error) {
	var vv []interface{}
	for _, v := range usms {
		vv = append(vv, v)
	}
	return c.Create(UtmSourceMixinModel, vv, nil)
}

// UpdateUtmSourceMixin updates an existing utm.source.mixin record.
func (c *Client) UpdateUtmSourceMixin(usm *UtmSourceMixin) error {
	return c.UpdateUtmSourceMixins([]int64{usm.Id.Get()}, usm)
}

// UpdateUtmSourceMixins updates existing utm.source.mixin records.
// All records (represented by ids) will be updated by usm values.
func (c *Client) UpdateUtmSourceMixins(ids []int64, usm *UtmSourceMixin) error {
	return c.Update(UtmSourceMixinModel, ids, usm, nil)
}

// DeleteUtmSourceMixin deletes an existing utm.source.mixin record.
func (c *Client) DeleteUtmSourceMixin(id int64) error {
	return c.DeleteUtmSourceMixins([]int64{id})
}

// DeleteUtmSourceMixins deletes existing utm.source.mixin records.
func (c *Client) DeleteUtmSourceMixins(ids []int64) error {
	return c.Delete(UtmSourceMixinModel, ids)
}

// GetUtmSourceMixin gets utm.source.mixin existing record.
func (c *Client) GetUtmSourceMixin(id int64) (*UtmSourceMixin, error) {
	usms, err := c.GetUtmSourceMixins([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*usms)[0]), nil
}

// GetUtmSourceMixins gets utm.source.mixin existing records.
func (c *Client) GetUtmSourceMixins(ids []int64) (*UtmSourceMixins, error) {
	usms := &UtmSourceMixins{}
	if err := c.Read(UtmSourceMixinModel, ids, nil, usms); err != nil {
		return nil, err
	}
	return usms, nil
}

// FindUtmSourceMixin finds utm.source.mixin record by querying it with criteria.
func (c *Client) FindUtmSourceMixin(criteria *Criteria) (*UtmSourceMixin, error) {
	usms := &UtmSourceMixins{}
	if err := c.SearchRead(UtmSourceMixinModel, criteria, NewOptions().Limit(1), usms); err != nil {
		return nil, err
	}
	return &((*usms)[0]), nil
}

// FindUtmSourceMixins finds utm.source.mixin records by querying it
// and filtering it with criteria and options.
func (c *Client) FindUtmSourceMixins(criteria *Criteria, options *Options) (*UtmSourceMixins, error) {
	usms := &UtmSourceMixins{}
	if err := c.SearchRead(UtmSourceMixinModel, criteria, options, usms); err != nil {
		return nil, err
	}
	return usms, nil
}

// FindUtmSourceMixinIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindUtmSourceMixinIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(UtmSourceMixinModel, criteria, options)
}

// FindUtmSourceMixinId finds record id by querying it with criteria.
func (c *Client) FindUtmSourceMixinId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(UtmSourceMixinModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
