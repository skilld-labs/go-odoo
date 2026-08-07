package odoo

// PropertiesBaseDefinitionMixin represents properties.base.definition.mixin model.
type PropertiesBaseDefinitionMixin struct {
	DisplayName                *String     `xmlrpc:"display_name,omitempty"`
	Id                         *Int        `xmlrpc:"id,omitempty"`
	Properties                 interface{} `xmlrpc:"properties,omitempty"`
	PropertiesBaseDefinitionId *Many2One   `xmlrpc:"properties_base_definition_id,omitempty"`
}

// PropertiesBaseDefinitionMixins represents array of properties.base.definition.mixin model.
type PropertiesBaseDefinitionMixins []PropertiesBaseDefinitionMixin

// PropertiesBaseDefinitionMixinModel is the odoo model name.
const PropertiesBaseDefinitionMixinModel = "properties.base.definition.mixin"

// Many2One convert PropertiesBaseDefinitionMixin to *Many2One.
func (pbdm *PropertiesBaseDefinitionMixin) Many2One() *Many2One {
	return NewMany2One(pbdm.Id.Get(), "")
}

// CreatePropertiesBaseDefinitionMixin creates a new properties.base.definition.mixin model and returns its id.
func (c *Client) CreatePropertiesBaseDefinitionMixin(pbdm *PropertiesBaseDefinitionMixin) (int64, error) {
	ids, err := c.CreatePropertiesBaseDefinitionMixins([]*PropertiesBaseDefinitionMixin{pbdm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreatePropertiesBaseDefinitionMixin creates a new properties.base.definition.mixin model and returns its id.
func (c *Client) CreatePropertiesBaseDefinitionMixins(pbdms []*PropertiesBaseDefinitionMixin) ([]int64, error) {
	var vv []interface{}
	for _, v := range pbdms {
		vv = append(vv, v)
	}
	return c.Create(PropertiesBaseDefinitionMixinModel, vv, nil)
}

// UpdatePropertiesBaseDefinitionMixin updates an existing properties.base.definition.mixin record.
func (c *Client) UpdatePropertiesBaseDefinitionMixin(pbdm *PropertiesBaseDefinitionMixin) error {
	return c.UpdatePropertiesBaseDefinitionMixins([]int64{pbdm.Id.Get()}, pbdm)
}

// UpdatePropertiesBaseDefinitionMixins updates existing properties.base.definition.mixin records.
// All records (represented by ids) will be updated by pbdm values.
func (c *Client) UpdatePropertiesBaseDefinitionMixins(ids []int64, pbdm *PropertiesBaseDefinitionMixin) error {
	return c.Update(PropertiesBaseDefinitionMixinModel, ids, pbdm, nil)
}

// DeletePropertiesBaseDefinitionMixin deletes an existing properties.base.definition.mixin record.
func (c *Client) DeletePropertiesBaseDefinitionMixin(id int64) error {
	return c.DeletePropertiesBaseDefinitionMixins([]int64{id})
}

// DeletePropertiesBaseDefinitionMixins deletes existing properties.base.definition.mixin records.
func (c *Client) DeletePropertiesBaseDefinitionMixins(ids []int64) error {
	return c.Delete(PropertiesBaseDefinitionMixinModel, ids)
}

// GetPropertiesBaseDefinitionMixin gets properties.base.definition.mixin existing record.
func (c *Client) GetPropertiesBaseDefinitionMixin(id int64) (*PropertiesBaseDefinitionMixin, error) {
	pbdms, err := c.GetPropertiesBaseDefinitionMixins([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*pbdms)[0]), nil
}

// GetPropertiesBaseDefinitionMixins gets properties.base.definition.mixin existing records.
func (c *Client) GetPropertiesBaseDefinitionMixins(ids []int64) (*PropertiesBaseDefinitionMixins, error) {
	pbdms := &PropertiesBaseDefinitionMixins{}
	if err := c.Read(PropertiesBaseDefinitionMixinModel, ids, nil, pbdms); err != nil {
		return nil, err
	}
	return pbdms, nil
}

// FindPropertiesBaseDefinitionMixin finds properties.base.definition.mixin record by querying it with criteria.
func (c *Client) FindPropertiesBaseDefinitionMixin(criteria *Criteria) (*PropertiesBaseDefinitionMixin, error) {
	pbdms := &PropertiesBaseDefinitionMixins{}
	if err := c.SearchRead(PropertiesBaseDefinitionMixinModel, criteria, NewOptions().Limit(1), pbdms); err != nil {
		return nil, err
	}
	return &((*pbdms)[0]), nil
}

// FindPropertiesBaseDefinitionMixins finds properties.base.definition.mixin records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPropertiesBaseDefinitionMixins(criteria *Criteria, options *Options) (*PropertiesBaseDefinitionMixins, error) {
	pbdms := &PropertiesBaseDefinitionMixins{}
	if err := c.SearchRead(PropertiesBaseDefinitionMixinModel, criteria, options, pbdms); err != nil {
		return nil, err
	}
	return pbdms, nil
}

// FindPropertiesBaseDefinitionMixinIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPropertiesBaseDefinitionMixinIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(PropertiesBaseDefinitionMixinModel, criteria, options)
}

// FindPropertiesBaseDefinitionMixinId finds record id by querying it with criteria.
func (c *Client) FindPropertiesBaseDefinitionMixinId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PropertiesBaseDefinitionMixinModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
