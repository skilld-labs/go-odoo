package odoo

// UpdateProductAttributeValue represents update.product.attribute.value model.
type UpdateProductAttributeValue struct {
	AttributeValueId *Many2One  `xmlrpc:"attribute_value_id,omitempty"`
	CreateDate       *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid        *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName      *String    `xmlrpc:"display_name,omitempty"`
	Id               *Int       `xmlrpc:"id,omitempty"`
	Message          *String    `xmlrpc:"message,omitempty"`
	Mode             *Selection `xmlrpc:"mode,omitempty"`
	ProductCount     *Int       `xmlrpc:"product_count,omitempty"`
	WriteDate        *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid         *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// UpdateProductAttributeValues represents array of update.product.attribute.value model.
type UpdateProductAttributeValues []UpdateProductAttributeValue

// UpdateProductAttributeValueModel is the odoo model name.
const UpdateProductAttributeValueModel = "update.product.attribute.value"

// Many2One convert UpdateProductAttributeValue to *Many2One.
func (upav *UpdateProductAttributeValue) Many2One() *Many2One {
	return NewMany2One(upav.Id.Get(), "")
}

// CreateUpdateProductAttributeValue creates a new update.product.attribute.value model and returns its id.
func (c *Client) CreateUpdateProductAttributeValue(upav *UpdateProductAttributeValue) (int64, error) {
	ids, err := c.CreateUpdateProductAttributeValues([]*UpdateProductAttributeValue{upav})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateUpdateProductAttributeValue creates a new update.product.attribute.value model and returns its id.
func (c *Client) CreateUpdateProductAttributeValues(upavs []*UpdateProductAttributeValue) ([]int64, error) {
	var vv []interface{}
	for _, v := range upavs {
		vv = append(vv, v)
	}
	return c.Create(UpdateProductAttributeValueModel, vv, nil)
}

// UpdateUpdateProductAttributeValue updates an existing update.product.attribute.value record.
func (c *Client) UpdateUpdateProductAttributeValue(upav *UpdateProductAttributeValue) error {
	return c.UpdateUpdateProductAttributeValues([]int64{upav.Id.Get()}, upav)
}

// UpdateUpdateProductAttributeValues updates existing update.product.attribute.value records.
// All records (represented by ids) will be updated by upav values.
func (c *Client) UpdateUpdateProductAttributeValues(ids []int64, upav *UpdateProductAttributeValue) error {
	return c.Update(UpdateProductAttributeValueModel, ids, upav, nil)
}

// DeleteUpdateProductAttributeValue deletes an existing update.product.attribute.value record.
func (c *Client) DeleteUpdateProductAttributeValue(id int64) error {
	return c.DeleteUpdateProductAttributeValues([]int64{id})
}

// DeleteUpdateProductAttributeValues deletes existing update.product.attribute.value records.
func (c *Client) DeleteUpdateProductAttributeValues(ids []int64) error {
	return c.Delete(UpdateProductAttributeValueModel, ids)
}

// GetUpdateProductAttributeValue gets update.product.attribute.value existing record.
func (c *Client) GetUpdateProductAttributeValue(id int64) (*UpdateProductAttributeValue, error) {
	upavs, err := c.GetUpdateProductAttributeValues([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*upavs)[0]), nil
}

// GetUpdateProductAttributeValues gets update.product.attribute.value existing records.
func (c *Client) GetUpdateProductAttributeValues(ids []int64) (*UpdateProductAttributeValues, error) {
	upavs := &UpdateProductAttributeValues{}
	if err := c.Read(UpdateProductAttributeValueModel, ids, nil, upavs); err != nil {
		return nil, err
	}
	return upavs, nil
}

// FindUpdateProductAttributeValue finds update.product.attribute.value record by querying it with criteria.
func (c *Client) FindUpdateProductAttributeValue(criteria *Criteria) (*UpdateProductAttributeValue, error) {
	upavs := &UpdateProductAttributeValues{}
	if err := c.SearchRead(UpdateProductAttributeValueModel, criteria, NewOptions().Limit(1), upavs); err != nil {
		return nil, err
	}
	return &((*upavs)[0]), nil
}

// FindUpdateProductAttributeValues finds update.product.attribute.value records by querying it
// and filtering it with criteria and options.
func (c *Client) FindUpdateProductAttributeValues(criteria *Criteria, options *Options) (*UpdateProductAttributeValues, error) {
	upavs := &UpdateProductAttributeValues{}
	if err := c.SearchRead(UpdateProductAttributeValueModel, criteria, options, upavs); err != nil {
		return nil, err
	}
	return upavs, nil
}

// FindUpdateProductAttributeValueIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindUpdateProductAttributeValueIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(UpdateProductAttributeValueModel, criteria, options)
}

// FindUpdateProductAttributeValueId finds record id by querying it with criteria.
func (c *Client) FindUpdateProductAttributeValueId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(UpdateProductAttributeValueModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
