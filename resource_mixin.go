package odoo

// ResourceMixin represents resource.mixin model.
type ResourceMixin struct {
	LastUpdate         *Time     `xmlrpc:"__last_update,omitempty"`
	CompanyId          *Many2One `xmlrpc:"company_id,omitempty"`
	DisplayName        *String   `xmlrpc:"display_name,omitempty"`
	Id                 *Int      `xmlrpc:"id,omitempty"`
	ResourceCalendarId *Many2One `xmlrpc:"resource_calendar_id,omitempty"`
	ResourceId         *Many2One `xmlrpc:"resource_id,omitempty"`
}

// ResourceMixins represents array of resource.mixin model.
type ResourceMixins []ResourceMixin

// ResourceMixinModel is the odoo model name.
const ResourceMixinModel = "resource.mixin"

// Many2One convert ResourceMixin to *Many2One.
func (rm *ResourceMixin) Many2One() *Many2One {
	return NewMany2One(rm.Id.Get(), "")
}

// CreateResourceMixin creates a new resource.mixin model and returns its id.
func (c *Client) CreateResourceMixin(rm *ResourceMixin) (int64, error) {
	ids, err := c.CreateResourceMixins([]*ResourceMixin{rm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateResourceMixins creates a new resource.mixin model and returns its id.
func (c *Client) CreateResourceMixins(rms []*ResourceMixin) ([]int64, error) {
	var vv []interface{}
	for _, v := range rms {
		vv = append(vv, v)
	}
	return c.Create(ResourceMixinModel, vv, nil)
}

// UpdateResourceMixin updates an existing resource.mixin record.
func (c *Client) UpdateResourceMixin(rm *ResourceMixin) error {
	return c.UpdateResourceMixins([]int64{rm.Id.Get()}, rm)
}

// UpdateResourceMixins updates existing resource.mixin records.
// All records (represented by ids) will be updated by rm values.
func (c *Client) UpdateResourceMixins(ids []int64, rm *ResourceMixin) error {
	return c.Update(ResourceMixinModel, ids, rm, nil)
}

// DeleteResourceMixin deletes an existing resource.mixin record.
func (c *Client) DeleteResourceMixin(id int64) error {
	return c.DeleteResourceMixins([]int64{id})
}

// DeleteResourceMixins deletes existing resource.mixin records.
func (c *Client) DeleteResourceMixins(ids []int64) error {
	return c.Delete(ResourceMixinModel, ids)
}

// GetResourceMixin gets resource.mixin existing record.
func (c *Client) GetResourceMixin(id int64) (*ResourceMixin, error) {
	rms, err := c.GetResourceMixins([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*rms)[0]), nil
}

// GetResourceMixins gets resource.mixin existing records.
func (c *Client) GetResourceMixins(ids []int64) (*ResourceMixins, error) {
	rms := &ResourceMixins{}
	if err := c.Read(ResourceMixinModel, ids, nil, rms); err != nil {
		return nil, err
	}
	return rms, nil
}

// FindResourceMixin finds resource.mixin record by querying it with criteria.
func (c *Client) FindResourceMixin(criteria *Criteria) (*ResourceMixin, error) {
	rms := &ResourceMixins{}
	if err := c.SearchRead(ResourceMixinModel, criteria, NewOptions().Limit(1), rms); err != nil {
		return nil, err
	}
	return &((*rms)[0]), nil
}

// FindResourceMixins finds resource.mixin records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResourceMixins(criteria *Criteria, options *Options) (*ResourceMixins, error) {
	rms := &ResourceMixins{}
	if err := c.SearchRead(ResourceMixinModel, criteria, options, rms); err != nil {
		return nil, err
	}
	return rms, nil
}

// FindResourceMixinIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResourceMixinIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ResourceMixinModel, criteria, options)
}

// FindResourceMixinId finds record id by querying it with criteria.
func (c *Client) FindResourceMixinId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResourceMixinModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
