package odoo

// HrMixin represents hr.mixin model.
type HrMixin struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// HrMixins represents array of hr.mixin model.
type HrMixins []HrMixin

// HrMixinModel is the odoo model name.
const HrMixinModel = "hr.mixin"

// Many2One convert HrMixin to *Many2One.
func (hm *HrMixin) Many2One() *Many2One {
	return NewMany2One(hm.Id.Get(), "")
}

// CreateHrMixin creates a new hr.mixin model and returns its id.
func (c *Client) CreateHrMixin(hm *HrMixin) (int64, error) {
	ids, err := c.CreateHrMixins([]*HrMixin{hm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrMixin creates a new hr.mixin model and returns its id.
func (c *Client) CreateHrMixins(hms []*HrMixin) ([]int64, error) {
	var vv []interface{}
	for _, v := range hms {
		vv = append(vv, v)
	}
	return c.Create(HrMixinModel, vv, nil)
}

// UpdateHrMixin updates an existing hr.mixin record.
func (c *Client) UpdateHrMixin(hm *HrMixin) error {
	return c.UpdateHrMixins([]int64{hm.Id.Get()}, hm)
}

// UpdateHrMixins updates existing hr.mixin records.
// All records (represented by ids) will be updated by hm values.
func (c *Client) UpdateHrMixins(ids []int64, hm *HrMixin) error {
	return c.Update(HrMixinModel, ids, hm, nil)
}

// DeleteHrMixin deletes an existing hr.mixin record.
func (c *Client) DeleteHrMixin(id int64) error {
	return c.DeleteHrMixins([]int64{id})
}

// DeleteHrMixins deletes existing hr.mixin records.
func (c *Client) DeleteHrMixins(ids []int64) error {
	return c.Delete(HrMixinModel, ids)
}

// GetHrMixin gets hr.mixin existing record.
func (c *Client) GetHrMixin(id int64) (*HrMixin, error) {
	hms, err := c.GetHrMixins([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*hms)[0]), nil
}

// GetHrMixins gets hr.mixin existing records.
func (c *Client) GetHrMixins(ids []int64) (*HrMixins, error) {
	hms := &HrMixins{}
	if err := c.Read(HrMixinModel, ids, nil, hms); err != nil {
		return nil, err
	}
	return hms, nil
}

// FindHrMixin finds hr.mixin record by querying it with criteria.
func (c *Client) FindHrMixin(criteria *Criteria) (*HrMixin, error) {
	hms := &HrMixins{}
	if err := c.SearchRead(HrMixinModel, criteria, NewOptions().Limit(1), hms); err != nil {
		return nil, err
	}
	return &((*hms)[0]), nil
}

// FindHrMixins finds hr.mixin records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrMixins(criteria *Criteria, options *Options) (*HrMixins, error) {
	hms := &HrMixins{}
	if err := c.SearchRead(HrMixinModel, criteria, options, hms); err != nil {
		return nil, err
	}
	return hms, nil
}

// FindHrMixinIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrMixinIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(HrMixinModel, criteria, options)
}

// FindHrMixinId finds record id by querying it with criteria.
func (c *Client) FindHrMixinId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrMixinModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
