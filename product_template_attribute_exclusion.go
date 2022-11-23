package odoo

import (
	"fmt"
)

// ProductTemplateAttributeExclusion represents product.template.attribute.exclusion model.
type ProductTemplateAttributeExclusion struct {
	LastUpdate                      *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate                      *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid                       *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName                     *String   `xmlrpc:"display_name,omitempty"`
	Id                              *Int      `xmlrpc:"id,omitempty"`
	ProductTemplateAttributeValueId *Many2One `xmlrpc:"product_template_attribute_value_id,omitempty"`
	ProductTmplId                   *Many2One `xmlrpc:"product_tmpl_id,omitempty"`
	ValueIds                        *Relation `xmlrpc:"value_ids,omitempty"`
	WriteDate                       *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid                        *Many2One `xmlrpc:"write_uid,omitempty"`
}

// ProductTemplateAttributeExclusions represents array of product.template.attribute.exclusion model.
type ProductTemplateAttributeExclusions []ProductTemplateAttributeExclusion

// ProductTemplateAttributeExclusionModel is the odoo model name.
const ProductTemplateAttributeExclusionModel = "product.template.attribute.exclusion"

// Many2One convert ProductTemplateAttributeExclusion to *Many2One.
func (ptae *ProductTemplateAttributeExclusion) Many2One() *Many2One {
	return NewMany2One(ptae.Id.Get(), "")
}

// CreateProductTemplateAttributeExclusion creates a new product.template.attribute.exclusion model and returns its id.
func (c *Client) CreateProductTemplateAttributeExclusion(ptae *ProductTemplateAttributeExclusion) (int64, error) {
	return c.Create(ProductTemplateAttributeExclusionModel, ptae)
}

// UpdateProductTemplateAttributeExclusion updates an existing product.template.attribute.exclusion record.
func (c *Client) UpdateProductTemplateAttributeExclusion(ptae *ProductTemplateAttributeExclusion) error {
	return c.UpdateProductTemplateAttributeExclusions([]int64{ptae.Id.Get()}, ptae)
}

// UpdateProductTemplateAttributeExclusions updates existing product.template.attribute.exclusion records.
// All records (represented by IDs) will be updated by ptae values.
func (c *Client) UpdateProductTemplateAttributeExclusions(ids []int64, ptae *ProductTemplateAttributeExclusion) error {
	return c.Update(ProductTemplateAttributeExclusionModel, ids, ptae)
}

// DeleteProductTemplateAttributeExclusion deletes an existing product.template.attribute.exclusion record.
func (c *Client) DeleteProductTemplateAttributeExclusion(id int64) error {
	return c.DeleteProductTemplateAttributeExclusions([]int64{id})
}

// DeleteProductTemplateAttributeExclusions deletes existing product.template.attribute.exclusion records.
func (c *Client) DeleteProductTemplateAttributeExclusions(ids []int64) error {
	return c.Delete(ProductTemplateAttributeExclusionModel, ids)
}

// GetProductTemplateAttributeExclusion gets product.template.attribute.exclusion existing record.
func (c *Client) GetProductTemplateAttributeExclusion(id int64) (*ProductTemplateAttributeExclusion, error) {
	ptaes, err := c.GetProductTemplateAttributeExclusions([]int64{id})
	if err != nil {
		return nil, err
	}
	if ptaes != nil && len(*ptaes) > 0 {
		return &((*ptaes)[0]), nil
	}
	return nil, fmt.Errorf("id %v of product.template.attribute.exclusion not found", id)
}

// GetProductTemplateAttributeExclusions gets product.template.attribute.exclusion existing records.
func (c *Client) GetProductTemplateAttributeExclusions(ids []int64) (*ProductTemplateAttributeExclusions, error) {
	ptaes := &ProductTemplateAttributeExclusions{}
	if err := c.Read(ProductTemplateAttributeExclusionModel, ids, nil, ptaes); err != nil {
		return nil, err
	}
	return ptaes, nil
}

// FindProductTemplateAttributeExclusion finds product.template.attribute.exclusion record by querying it with criteria.
func (c *Client) FindProductTemplateAttributeExclusion(criteria *Criteria) (*ProductTemplateAttributeExclusion, error) {
	ptaes := &ProductTemplateAttributeExclusions{}
	if err := c.SearchRead(ProductTemplateAttributeExclusionModel, criteria, NewOptions().Limit(1), ptaes); err != nil {
		return nil, err
	}
	if ptaes != nil && len(*ptaes) > 0 {
		return &((*ptaes)[0]), nil
	}
	return nil, fmt.Errorf("product.template.attribute.exclusion was not found")
}

// FindProductTemplateAttributeExclusions finds product.template.attribute.exclusion records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductTemplateAttributeExclusions(criteria *Criteria, options *Options) (*ProductTemplateAttributeExclusions, error) {
	ptaes := &ProductTemplateAttributeExclusions{}
	if err := c.SearchRead(ProductTemplateAttributeExclusionModel, criteria, options, ptaes); err != nil {
		return nil, err
	}
	return ptaes, nil
}

// FindProductTemplateAttributeExclusionIds finds records IDs by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductTemplateAttributeExclusionIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ProductTemplateAttributeExclusionModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindProductTemplateAttributeExclusionId finds record id by querying it with criteria.
func (c *Client) FindProductTemplateAttributeExclusionId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProductTemplateAttributeExclusionModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("product.template.attribute.exclusion was not found")
}
