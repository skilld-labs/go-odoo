package odoo

// BusListenerMixin represents bus.listener.mixin model.
type BusListenerMixin struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// BusListenerMixins represents array of bus.listener.mixin model.
type BusListenerMixins []BusListenerMixin

// BusListenerMixinModel is the odoo model name.
const BusListenerMixinModel = "bus.listener.mixin"

// Many2One convert BusListenerMixin to *Many2One.
func (blm *BusListenerMixin) Many2One() *Many2One {
	return NewMany2One(blm.Id.Get(), "")
}

// CreateBusListenerMixin creates a new bus.listener.mixin model and returns its id.
func (c *Client) CreateBusListenerMixin(blm *BusListenerMixin) (int64, error) {
	ids, err := c.CreateBusListenerMixins([]*BusListenerMixin{blm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateBusListenerMixin creates a new bus.listener.mixin model and returns its id.
func (c *Client) CreateBusListenerMixins(blms []*BusListenerMixin) ([]int64, error) {
	var vv []interface{}
	for _, v := range blms {
		vv = append(vv, v)
	}
	return c.Create(BusListenerMixinModel, vv, nil)
}

// UpdateBusListenerMixin updates an existing bus.listener.mixin record.
func (c *Client) UpdateBusListenerMixin(blm *BusListenerMixin) error {
	return c.UpdateBusListenerMixins([]int64{blm.Id.Get()}, blm)
}

// UpdateBusListenerMixins updates existing bus.listener.mixin records.
// All records (represented by ids) will be updated by blm values.
func (c *Client) UpdateBusListenerMixins(ids []int64, blm *BusListenerMixin) error {
	return c.Update(BusListenerMixinModel, ids, blm, nil)
}

// DeleteBusListenerMixin deletes an existing bus.listener.mixin record.
func (c *Client) DeleteBusListenerMixin(id int64) error {
	return c.DeleteBusListenerMixins([]int64{id})
}

// DeleteBusListenerMixins deletes existing bus.listener.mixin records.
func (c *Client) DeleteBusListenerMixins(ids []int64) error {
	return c.Delete(BusListenerMixinModel, ids)
}

// GetBusListenerMixin gets bus.listener.mixin existing record.
func (c *Client) GetBusListenerMixin(id int64) (*BusListenerMixin, error) {
	blms, err := c.GetBusListenerMixins([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*blms)[0]), nil
}

// GetBusListenerMixins gets bus.listener.mixin existing records.
func (c *Client) GetBusListenerMixins(ids []int64) (*BusListenerMixins, error) {
	blms := &BusListenerMixins{}
	if err := c.Read(BusListenerMixinModel, ids, nil, blms); err != nil {
		return nil, err
	}
	return blms, nil
}

// FindBusListenerMixin finds bus.listener.mixin record by querying it with criteria.
func (c *Client) FindBusListenerMixin(criteria *Criteria) (*BusListenerMixin, error) {
	blms := &BusListenerMixins{}
	if err := c.SearchRead(BusListenerMixinModel, criteria, NewOptions().Limit(1), blms); err != nil {
		return nil, err
	}
	return &((*blms)[0]), nil
}

// FindBusListenerMixins finds bus.listener.mixin records by querying it
// and filtering it with criteria and options.
func (c *Client) FindBusListenerMixins(criteria *Criteria, options *Options) (*BusListenerMixins, error) {
	blms := &BusListenerMixins{}
	if err := c.SearchRead(BusListenerMixinModel, criteria, options, blms); err != nil {
		return nil, err
	}
	return blms, nil
}

// FindBusListenerMixinIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindBusListenerMixinIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(BusListenerMixinModel, criteria, options)
}

// FindBusListenerMixinId finds record id by querying it with criteria.
func (c *Client) FindBusListenerMixinId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(BusListenerMixinModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
