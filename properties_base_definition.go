package odoo

// PropertiesBaseDefinition represents properties.base.definition model.
type PropertiesBaseDefinition struct {
	CreateDate           *Time       `xmlrpc:"create_date,omitempty"`
	CreateUid            *Many2One   `xmlrpc:"create_uid,omitempty"`
	DisplayName          *String     `xmlrpc:"display_name,omitempty"`
	Id                   *Int        `xmlrpc:"id,omitempty"`
	PropertiesDefinition interface{} `xmlrpc:"properties_definition,omitempty"`
	PropertiesFieldId    *Many2One   `xmlrpc:"properties_field_id,omitempty"`
	WriteDate            *Time       `xmlrpc:"write_date,omitempty"`
	WriteUid             *Many2One   `xmlrpc:"write_uid,omitempty"`
}

// PropertiesBaseDefinitions represents array of properties.base.definition model.
type PropertiesBaseDefinitions []PropertiesBaseDefinition

// PropertiesBaseDefinitionModel is the odoo model name.
const PropertiesBaseDefinitionModel = "properties.base.definition"

// Many2One convert PropertiesBaseDefinition to *Many2One.
func (pbd *PropertiesBaseDefinition) Many2One() *Many2One {
	return NewMany2One(pbd.Id.Get(), "")
}

// CreatePropertiesBaseDefinition creates a new properties.base.definition model and returns its id.
func (c *Client) CreatePropertiesBaseDefinition(pbd *PropertiesBaseDefinition) (int64, error) {
	ids, err := c.CreatePropertiesBaseDefinitions([]*PropertiesBaseDefinition{pbd})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreatePropertiesBaseDefinition creates a new properties.base.definition model and returns its id.
func (c *Client) CreatePropertiesBaseDefinitions(pbds []*PropertiesBaseDefinition) ([]int64, error) {
	var vv []interface{}
	for _, v := range pbds {
		vv = append(vv, v)
	}
	return c.Create(PropertiesBaseDefinitionModel, vv, nil)
}

// UpdatePropertiesBaseDefinition updates an existing properties.base.definition record.
func (c *Client) UpdatePropertiesBaseDefinition(pbd *PropertiesBaseDefinition) error {
	return c.UpdatePropertiesBaseDefinitions([]int64{pbd.Id.Get()}, pbd)
}

// UpdatePropertiesBaseDefinitions updates existing properties.base.definition records.
// All records (represented by ids) will be updated by pbd values.
func (c *Client) UpdatePropertiesBaseDefinitions(ids []int64, pbd *PropertiesBaseDefinition) error {
	return c.Update(PropertiesBaseDefinitionModel, ids, pbd, nil)
}

// DeletePropertiesBaseDefinition deletes an existing properties.base.definition record.
func (c *Client) DeletePropertiesBaseDefinition(id int64) error {
	return c.DeletePropertiesBaseDefinitions([]int64{id})
}

// DeletePropertiesBaseDefinitions deletes existing properties.base.definition records.
func (c *Client) DeletePropertiesBaseDefinitions(ids []int64) error {
	return c.Delete(PropertiesBaseDefinitionModel, ids)
}

// GetPropertiesBaseDefinition gets properties.base.definition existing record.
func (c *Client) GetPropertiesBaseDefinition(id int64) (*PropertiesBaseDefinition, error) {
	pbds, err := c.GetPropertiesBaseDefinitions([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*pbds)[0]), nil
}

// GetPropertiesBaseDefinitions gets properties.base.definition existing records.
func (c *Client) GetPropertiesBaseDefinitions(ids []int64) (*PropertiesBaseDefinitions, error) {
	pbds := &PropertiesBaseDefinitions{}
	if err := c.Read(PropertiesBaseDefinitionModel, ids, nil, pbds); err != nil {
		return nil, err
	}
	return pbds, nil
}

// FindPropertiesBaseDefinition finds properties.base.definition record by querying it with criteria.
func (c *Client) FindPropertiesBaseDefinition(criteria *Criteria) (*PropertiesBaseDefinition, error) {
	pbds := &PropertiesBaseDefinitions{}
	if err := c.SearchRead(PropertiesBaseDefinitionModel, criteria, NewOptions().Limit(1), pbds); err != nil {
		return nil, err
	}
	return &((*pbds)[0]), nil
}

// FindPropertiesBaseDefinitions finds properties.base.definition records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPropertiesBaseDefinitions(criteria *Criteria, options *Options) (*PropertiesBaseDefinitions, error) {
	pbds := &PropertiesBaseDefinitions{}
	if err := c.SearchRead(PropertiesBaseDefinitionModel, criteria, options, pbds); err != nil {
		return nil, err
	}
	return pbds, nil
}

// FindPropertiesBaseDefinitionIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPropertiesBaseDefinitionIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(PropertiesBaseDefinitionModel, criteria, options)
}

// FindPropertiesBaseDefinitionId finds record id by querying it with criteria.
func (c *Client) FindPropertiesBaseDefinitionId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PropertiesBaseDefinitionModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
