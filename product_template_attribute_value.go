package odoo

// ProductTemplateAttributeValue represents product.template.attribute.value model.
type ProductTemplateAttributeValue struct {
	AttributeId             *Many2One  `xmlrpc:"attribute_id,omitempty"`
	AttributeLineId         *Many2One  `xmlrpc:"attribute_line_id,omitempty"`
	Color                   *Int       `xmlrpc:"color,omitempty"`
	CreateDate              *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid               *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrencyId              *Many2One  `xmlrpc:"currency_id,omitempty"`
	DisplayName             *String    `xmlrpc:"display_name,omitempty"`
	DisplayType             *Selection `xmlrpc:"display_type,omitempty"`
	ExcludeFor              *Relation  `xmlrpc:"exclude_for,omitempty"`
	HtmlColor               *String    `xmlrpc:"html_color,omitempty"`
	Id                      *Int       `xmlrpc:"id,omitempty"`
	Image                   *String    `xmlrpc:"image,omitempty"`
	IsCustom                *Bool      `xmlrpc:"is_custom,omitempty"`
	Name                    *String    `xmlrpc:"name,omitempty"`
	PriceExtra              *Float     `xmlrpc:"price_extra,omitempty"`
	ProductAttributeValueId *Many2One  `xmlrpc:"product_attribute_value_id,omitempty"`
	ProductTmplId           *Many2One  `xmlrpc:"product_tmpl_id,omitempty"`
	PtavActive              *Bool      `xmlrpc:"ptav_active,omitempty"`
	PtavProductVariantIds   *Relation  `xmlrpc:"ptav_product_variant_ids,omitempty"`
	WriteDate               *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// ProductTemplateAttributeValues represents array of product.template.attribute.value model.
type ProductTemplateAttributeValues []ProductTemplateAttributeValue

// ProductTemplateAttributeValueModel is the odoo model name.
const ProductTemplateAttributeValueModel = "product.template.attribute.value"

// Many2One convert ProductTemplateAttributeValue to *Many2One.
func (ptav *ProductTemplateAttributeValue) Many2One() *Many2One {
	return NewMany2One(ptav.Id.Get(), "")
}

// CreateProductTemplateAttributeValue creates a new product.template.attribute.value model and returns its id.
func (c *Client) CreateProductTemplateAttributeValue(ptav *ProductTemplateAttributeValue) (int64, error) {
	ids, err := c.CreateProductTemplateAttributeValues([]*ProductTemplateAttributeValue{ptav})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProductTemplateAttributeValue creates a new product.template.attribute.value model and returns its id.
func (c *Client) CreateProductTemplateAttributeValues(ptavs []*ProductTemplateAttributeValue) ([]int64, error) {
	var vv []interface{}
	for _, v := range ptavs {
		vv = append(vv, v)
	}
	return c.Create(ProductTemplateAttributeValueModel, vv, nil)
}

// UpdateProductTemplateAttributeValue updates an existing product.template.attribute.value record.
func (c *Client) UpdateProductTemplateAttributeValue(ptav *ProductTemplateAttributeValue) error {
	return c.UpdateProductTemplateAttributeValues([]int64{ptav.Id.Get()}, ptav)
}

// UpdateProductTemplateAttributeValues updates existing product.template.attribute.value records.
// All records (represented by ids) will be updated by ptav values.
func (c *Client) UpdateProductTemplateAttributeValues(ids []int64, ptav *ProductTemplateAttributeValue) error {
	return c.Update(ProductTemplateAttributeValueModel, ids, ptav, nil)
}

// DeleteProductTemplateAttributeValue deletes an existing product.template.attribute.value record.
func (c *Client) DeleteProductTemplateAttributeValue(id int64) error {
	return c.DeleteProductTemplateAttributeValues([]int64{id})
}

// DeleteProductTemplateAttributeValues deletes existing product.template.attribute.value records.
func (c *Client) DeleteProductTemplateAttributeValues(ids []int64) error {
	return c.Delete(ProductTemplateAttributeValueModel, ids)
}

// GetProductTemplateAttributeValue gets product.template.attribute.value existing record.
func (c *Client) GetProductTemplateAttributeValue(id int64) (*ProductTemplateAttributeValue, error) {
	ptavs, err := c.GetProductTemplateAttributeValues([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ptavs)[0]), nil
}

// GetProductTemplateAttributeValues gets product.template.attribute.value existing records.
func (c *Client) GetProductTemplateAttributeValues(ids []int64) (*ProductTemplateAttributeValues, error) {
	ptavs := &ProductTemplateAttributeValues{}
	if err := c.Read(ProductTemplateAttributeValueModel, ids, nil, ptavs); err != nil {
		return nil, err
	}
	return ptavs, nil
}

// FindProductTemplateAttributeValue finds product.template.attribute.value record by querying it with criteria.
func (c *Client) FindProductTemplateAttributeValue(criteria *Criteria) (*ProductTemplateAttributeValue, error) {
	ptavs := &ProductTemplateAttributeValues{}
	if err := c.SearchRead(ProductTemplateAttributeValueModel, criteria, NewOptions().Limit(1), ptavs); err != nil {
		return nil, err
	}
	return &((*ptavs)[0]), nil
}

// FindProductTemplateAttributeValues finds product.template.attribute.value records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductTemplateAttributeValues(criteria *Criteria, options *Options) (*ProductTemplateAttributeValues, error) {
	ptavs := &ProductTemplateAttributeValues{}
	if err := c.SearchRead(ProductTemplateAttributeValueModel, criteria, options, ptavs); err != nil {
		return nil, err
	}
	return ptavs, nil
}

// FindProductTemplateAttributeValueIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductTemplateAttributeValueIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ProductTemplateAttributeValueModel, criteria, options)
}

// FindProductTemplateAttributeValueId finds record id by querying it with criteria.
func (c *Client) FindProductTemplateAttributeValueId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProductTemplateAttributeValueModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
