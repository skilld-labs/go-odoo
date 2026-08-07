package odoo

// TemplateResetMixin represents template.reset.mixin model.
type TemplateResetMixin struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
	TemplateFs  *String `xmlrpc:"template_fs,omitempty"`
}

// TemplateResetMixins represents array of template.reset.mixin model.
type TemplateResetMixins []TemplateResetMixin

// TemplateResetMixinModel is the odoo model name.
const TemplateResetMixinModel = "template.reset.mixin"

// Many2One convert TemplateResetMixin to *Many2One.
func (trm *TemplateResetMixin) Many2One() *Many2One {
	return NewMany2One(trm.Id.Get(), "")
}

// CreateTemplateResetMixin creates a new template.reset.mixin model and returns its id.
func (c *Client) CreateTemplateResetMixin(trm *TemplateResetMixin) (int64, error) {
	ids, err := c.CreateTemplateResetMixins([]*TemplateResetMixin{trm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateTemplateResetMixin creates a new template.reset.mixin model and returns its id.
func (c *Client) CreateTemplateResetMixins(trms []*TemplateResetMixin) ([]int64, error) {
	var vv []interface{}
	for _, v := range trms {
		vv = append(vv, v)
	}
	return c.Create(TemplateResetMixinModel, vv, nil)
}

// UpdateTemplateResetMixin updates an existing template.reset.mixin record.
func (c *Client) UpdateTemplateResetMixin(trm *TemplateResetMixin) error {
	return c.UpdateTemplateResetMixins([]int64{trm.Id.Get()}, trm)
}

// UpdateTemplateResetMixins updates existing template.reset.mixin records.
// All records (represented by ids) will be updated by trm values.
func (c *Client) UpdateTemplateResetMixins(ids []int64, trm *TemplateResetMixin) error {
	return c.Update(TemplateResetMixinModel, ids, trm, nil)
}

// DeleteTemplateResetMixin deletes an existing template.reset.mixin record.
func (c *Client) DeleteTemplateResetMixin(id int64) error {
	return c.DeleteTemplateResetMixins([]int64{id})
}

// DeleteTemplateResetMixins deletes existing template.reset.mixin records.
func (c *Client) DeleteTemplateResetMixins(ids []int64) error {
	return c.Delete(TemplateResetMixinModel, ids)
}

// GetTemplateResetMixin gets template.reset.mixin existing record.
func (c *Client) GetTemplateResetMixin(id int64) (*TemplateResetMixin, error) {
	trms, err := c.GetTemplateResetMixins([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*trms)[0]), nil
}

// GetTemplateResetMixins gets template.reset.mixin existing records.
func (c *Client) GetTemplateResetMixins(ids []int64) (*TemplateResetMixins, error) {
	trms := &TemplateResetMixins{}
	if err := c.Read(TemplateResetMixinModel, ids, nil, trms); err != nil {
		return nil, err
	}
	return trms, nil
}

// FindTemplateResetMixin finds template.reset.mixin record by querying it with criteria.
func (c *Client) FindTemplateResetMixin(criteria *Criteria) (*TemplateResetMixin, error) {
	trms := &TemplateResetMixins{}
	if err := c.SearchRead(TemplateResetMixinModel, criteria, NewOptions().Limit(1), trms); err != nil {
		return nil, err
	}
	return &((*trms)[0]), nil
}

// FindTemplateResetMixins finds template.reset.mixin records by querying it
// and filtering it with criteria and options.
func (c *Client) FindTemplateResetMixins(criteria *Criteria, options *Options) (*TemplateResetMixins, error) {
	trms := &TemplateResetMixins{}
	if err := c.SearchRead(TemplateResetMixinModel, criteria, options, trms); err != nil {
		return nil, err
	}
	return trms, nil
}

// FindTemplateResetMixinIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindTemplateResetMixinIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(TemplateResetMixinModel, criteria, options)
}

// FindTemplateResetMixinId finds record id by querying it with criteria.
func (c *Client) FindTemplateResetMixinId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(TemplateResetMixinModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
