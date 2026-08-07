package odoo

// AnalyticPlanFieldsMixin represents analytic.plan.fields.mixin model.
type AnalyticPlanFieldsMixin struct {
	AccountId     *Many2One `xmlrpc:"account_id,omitempty"`
	AutoAccountId *Many2One `xmlrpc:"auto_account_id,omitempty"`
	DisplayName   *String   `xmlrpc:"display_name,omitempty"`
	Id            *Int      `xmlrpc:"id,omitempty"`
}

// AnalyticPlanFieldsMixins represents array of analytic.plan.fields.mixin model.
type AnalyticPlanFieldsMixins []AnalyticPlanFieldsMixin

// AnalyticPlanFieldsMixinModel is the odoo model name.
const AnalyticPlanFieldsMixinModel = "analytic.plan.fields.mixin"

// Many2One convert AnalyticPlanFieldsMixin to *Many2One.
func (apfm *AnalyticPlanFieldsMixin) Many2One() *Many2One {
	return NewMany2One(apfm.Id.Get(), "")
}

// CreateAnalyticPlanFieldsMixin creates a new analytic.plan.fields.mixin model and returns its id.
func (c *Client) CreateAnalyticPlanFieldsMixin(apfm *AnalyticPlanFieldsMixin) (int64, error) {
	ids, err := c.CreateAnalyticPlanFieldsMixins([]*AnalyticPlanFieldsMixin{apfm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAnalyticPlanFieldsMixin creates a new analytic.plan.fields.mixin model and returns its id.
func (c *Client) CreateAnalyticPlanFieldsMixins(apfms []*AnalyticPlanFieldsMixin) ([]int64, error) {
	var vv []interface{}
	for _, v := range apfms {
		vv = append(vv, v)
	}
	return c.Create(AnalyticPlanFieldsMixinModel, vv, nil)
}

// UpdateAnalyticPlanFieldsMixin updates an existing analytic.plan.fields.mixin record.
func (c *Client) UpdateAnalyticPlanFieldsMixin(apfm *AnalyticPlanFieldsMixin) error {
	return c.UpdateAnalyticPlanFieldsMixins([]int64{apfm.Id.Get()}, apfm)
}

// UpdateAnalyticPlanFieldsMixins updates existing analytic.plan.fields.mixin records.
// All records (represented by ids) will be updated by apfm values.
func (c *Client) UpdateAnalyticPlanFieldsMixins(ids []int64, apfm *AnalyticPlanFieldsMixin) error {
	return c.Update(AnalyticPlanFieldsMixinModel, ids, apfm, nil)
}

// DeleteAnalyticPlanFieldsMixin deletes an existing analytic.plan.fields.mixin record.
func (c *Client) DeleteAnalyticPlanFieldsMixin(id int64) error {
	return c.DeleteAnalyticPlanFieldsMixins([]int64{id})
}

// DeleteAnalyticPlanFieldsMixins deletes existing analytic.plan.fields.mixin records.
func (c *Client) DeleteAnalyticPlanFieldsMixins(ids []int64) error {
	return c.Delete(AnalyticPlanFieldsMixinModel, ids)
}

// GetAnalyticPlanFieldsMixin gets analytic.plan.fields.mixin existing record.
func (c *Client) GetAnalyticPlanFieldsMixin(id int64) (*AnalyticPlanFieldsMixin, error) {
	apfms, err := c.GetAnalyticPlanFieldsMixins([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*apfms)[0]), nil
}

// GetAnalyticPlanFieldsMixins gets analytic.plan.fields.mixin existing records.
func (c *Client) GetAnalyticPlanFieldsMixins(ids []int64) (*AnalyticPlanFieldsMixins, error) {
	apfms := &AnalyticPlanFieldsMixins{}
	if err := c.Read(AnalyticPlanFieldsMixinModel, ids, nil, apfms); err != nil {
		return nil, err
	}
	return apfms, nil
}

// FindAnalyticPlanFieldsMixin finds analytic.plan.fields.mixin record by querying it with criteria.
func (c *Client) FindAnalyticPlanFieldsMixin(criteria *Criteria) (*AnalyticPlanFieldsMixin, error) {
	apfms := &AnalyticPlanFieldsMixins{}
	if err := c.SearchRead(AnalyticPlanFieldsMixinModel, criteria, NewOptions().Limit(1), apfms); err != nil {
		return nil, err
	}
	return &((*apfms)[0]), nil
}

// FindAnalyticPlanFieldsMixins finds analytic.plan.fields.mixin records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAnalyticPlanFieldsMixins(criteria *Criteria, options *Options) (*AnalyticPlanFieldsMixins, error) {
	apfms := &AnalyticPlanFieldsMixins{}
	if err := c.SearchRead(AnalyticPlanFieldsMixinModel, criteria, options, apfms); err != nil {
		return nil, err
	}
	return apfms, nil
}

// FindAnalyticPlanFieldsMixinIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAnalyticPlanFieldsMixinIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AnalyticPlanFieldsMixinModel, criteria, options)
}

// FindAnalyticPlanFieldsMixinId finds record id by querying it with criteria.
func (c *Client) FindAnalyticPlanFieldsMixinId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AnalyticPlanFieldsMixinModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
