package odoo

// ProductAttributeValue represents product.attribute.value model.
type ProductAttributeValue struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	AttributeId *Many2One `xmlrpc:"attribute_id,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	Name        *String   `xmlrpc:"name,omptempty"`
	PriceExtra  *Float    `xmlrpc:"price_extra,omptempty"`
	PriceIds    *Relation `xmlrpc:"price_ids,omptempty"`
	ProductIds  *Relation `xmlrpc:"product_ids,omptempty"`
	Sequence    *Int      `xmlrpc:"sequence,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// ProductAttributeValues represents array of product.attribute.value model.
type ProductAttributeValues []ProductAttributeValue

// ProductAttributeValueModel is the odoo model name.
const ProductAttributeValueModel = "product.attribute.value"

// Many2One convert ProductAttributeValue to *Many2One.
func (pav *ProductAttributeValue) Many2One() *Many2One {
	return NewMany2One(pav.Id.Get(), "")
}

// CreateProductAttributeValue creates a new product.attribute.value model and returns its id.
func (c *Client) CreateProductAttributeValue(pav *ProductAttributeValue) (int64, error) {
	ids, err := c.CreateProductAttributeValues([]*ProductAttributeValue{pav})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProductAttributeValue creates a new product.attribute.value model and returns its id.
func (c *Client) CreateProductAttributeValues(pavs []*ProductAttributeValue) ([]int64, error) {
	var vv []interface{}
	for _, v := range pavs {
		vv = append(vv, v)
	}
	return c.Create(ProductAttributeValueModel, vv, nil)
}

// UpdateProductAttributeValue updates an existing product.attribute.value record.
func (c *Client) UpdateProductAttributeValue(pav *ProductAttributeValue) error {
	return c.UpdateProductAttributeValues([]int64{pav.Id.Get()}, pav)
}

// UpdateProductAttributeValues updates existing product.attribute.value records.
// All records (represented by ids) will be updated by pav values.
func (c *Client) UpdateProductAttributeValues(ids []int64, pav *ProductAttributeValue) error {
	return c.Update(ProductAttributeValueModel, ids, pav, nil)
}

// DeleteProductAttributeValue deletes an existing product.attribute.value record.
func (c *Client) DeleteProductAttributeValue(id int64) error {
	return c.DeleteProductAttributeValues([]int64{id})
}

// DeleteProductAttributeValues deletes existing product.attribute.value records.
func (c *Client) DeleteProductAttributeValues(ids []int64) error {
	return c.Delete(ProductAttributeValueModel, ids)
}

// GetProductAttributeValue gets product.attribute.value existing record.
func (c *Client) GetProductAttributeValue(id int64) (*ProductAttributeValue, error) {
	pavs, err := c.GetProductAttributeValues([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*pavs)[0]), nil
}

// GetProductAttributeValues gets product.attribute.value existing records.
func (c *Client) GetProductAttributeValues(ids []int64) (*ProductAttributeValues, error) {
	pavs := &ProductAttributeValues{}
	if err := c.Read(ProductAttributeValueModel, ids, nil, pavs); err != nil {
		return nil, err
	}
	return pavs, nil
}

// FindProductAttributeValue finds product.attribute.value record by querying it with criteria.
func (c *Client) FindProductAttributeValue(criteria *Criteria) (*ProductAttributeValue, error) {
	pavs := &ProductAttributeValues{}
	if err := c.SearchRead(ProductAttributeValueModel, criteria, NewOptions().Limit(1), pavs); err != nil {
		return nil, err
	}
	return &((*pavs)[0]), nil
}

// FindProductAttributeValues finds product.attribute.value records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductAttributeValues(criteria *Criteria, options *Options) (*ProductAttributeValues, error) {
	pavs := &ProductAttributeValues{}
	if err := c.SearchRead(ProductAttributeValueModel, criteria, options, pavs); err != nil {
		return nil, err
	}
	return pavs, nil
}

// FindProductAttributeValueIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductAttributeValueIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ProductAttributeValueModel, criteria, options)
}

// FindProductAttributeValueId finds record id by querying it with criteria.
func (c *Client) FindProductAttributeValueId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProductAttributeValueModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
