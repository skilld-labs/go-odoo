package odoo

// ProductTemplateAttributeLine represents product.template.attribute.line model.
type ProductTemplateAttributeLine struct {
	Active                  *Bool     `xmlrpc:"active,omitempty"`
	AttributeId             *Many2One `xmlrpc:"attribute_id,omitempty"`
	CreateDate              *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid               *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName             *String   `xmlrpc:"display_name,omitempty"`
	Id                      *Int      `xmlrpc:"id,omitempty"`
	ProductTemplateValueIds *Relation `xmlrpc:"product_template_value_ids,omitempty"`
	ProductTmplId           *Many2One `xmlrpc:"product_tmpl_id,omitempty"`
	Sequence                *Int      `xmlrpc:"sequence,omitempty"`
	ValueCount              *Int      `xmlrpc:"value_count,omitempty"`
	ValueIds                *Relation `xmlrpc:"value_ids,omitempty"`
	WriteDate               *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid                *Many2One `xmlrpc:"write_uid,omitempty"`
}

// ProductTemplateAttributeLines represents array of product.template.attribute.line model.
type ProductTemplateAttributeLines []ProductTemplateAttributeLine

// ProductTemplateAttributeLineModel is the odoo model name.
const ProductTemplateAttributeLineModel = "product.template.attribute.line"

// Many2One convert ProductTemplateAttributeLine to *Many2One.
func (ptal *ProductTemplateAttributeLine) Many2One() *Many2One {
	return NewMany2One(ptal.Id.Get(), "")
}

// CreateProductTemplateAttributeLine creates a new product.template.attribute.line model and returns its id.
func (c *Client) CreateProductTemplateAttributeLine(ptal *ProductTemplateAttributeLine) (int64, error) {
	ids, err := c.CreateProductTemplateAttributeLines([]*ProductTemplateAttributeLine{ptal})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProductTemplateAttributeLine creates a new product.template.attribute.line model and returns its id.
func (c *Client) CreateProductTemplateAttributeLines(ptals []*ProductTemplateAttributeLine) ([]int64, error) {
	var vv []interface{}
	for _, v := range ptals {
		vv = append(vv, v)
	}
	return c.Create(ProductTemplateAttributeLineModel, vv, nil)
}

// UpdateProductTemplateAttributeLine updates an existing product.template.attribute.line record.
func (c *Client) UpdateProductTemplateAttributeLine(ptal *ProductTemplateAttributeLine) error {
	return c.UpdateProductTemplateAttributeLines([]int64{ptal.Id.Get()}, ptal)
}

// UpdateProductTemplateAttributeLines updates existing product.template.attribute.line records.
// All records (represented by ids) will be updated by ptal values.
func (c *Client) UpdateProductTemplateAttributeLines(ids []int64, ptal *ProductTemplateAttributeLine) error {
	return c.Update(ProductTemplateAttributeLineModel, ids, ptal, nil)
}

// DeleteProductTemplateAttributeLine deletes an existing product.template.attribute.line record.
func (c *Client) DeleteProductTemplateAttributeLine(id int64) error {
	return c.DeleteProductTemplateAttributeLines([]int64{id})
}

// DeleteProductTemplateAttributeLines deletes existing product.template.attribute.line records.
func (c *Client) DeleteProductTemplateAttributeLines(ids []int64) error {
	return c.Delete(ProductTemplateAttributeLineModel, ids)
}

// GetProductTemplateAttributeLine gets product.template.attribute.line existing record.
func (c *Client) GetProductTemplateAttributeLine(id int64) (*ProductTemplateAttributeLine, error) {
	ptals, err := c.GetProductTemplateAttributeLines([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ptals)[0]), nil
}

// GetProductTemplateAttributeLines gets product.template.attribute.line existing records.
func (c *Client) GetProductTemplateAttributeLines(ids []int64) (*ProductTemplateAttributeLines, error) {
	ptals := &ProductTemplateAttributeLines{}
	if err := c.Read(ProductTemplateAttributeLineModel, ids, nil, ptals); err != nil {
		return nil, err
	}
	return ptals, nil
}

// FindProductTemplateAttributeLine finds product.template.attribute.line record by querying it with criteria.
func (c *Client) FindProductTemplateAttributeLine(criteria *Criteria) (*ProductTemplateAttributeLine, error) {
	ptals := &ProductTemplateAttributeLines{}
	if err := c.SearchRead(ProductTemplateAttributeLineModel, criteria, NewOptions().Limit(1), ptals); err != nil {
		return nil, err
	}
	return &((*ptals)[0]), nil
}

// FindProductTemplateAttributeLines finds product.template.attribute.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductTemplateAttributeLines(criteria *Criteria, options *Options) (*ProductTemplateAttributeLines, error) {
	ptals := &ProductTemplateAttributeLines{}
	if err := c.SearchRead(ProductTemplateAttributeLineModel, criteria, options, ptals); err != nil {
		return nil, err
	}
	return ptals, nil
}

// FindProductTemplateAttributeLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductTemplateAttributeLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ProductTemplateAttributeLineModel, criteria, options)
}

// FindProductTemplateAttributeLineId finds record id by querying it with criteria.
func (c *Client) FindProductTemplateAttributeLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProductTemplateAttributeLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
