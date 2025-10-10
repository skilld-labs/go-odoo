package odoo

// ProductTag represents product.tag model.
type ProductTag struct {
	Color              *String   `xmlrpc:"color,omitempty"`
	CreateDate         *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid          *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName        *String   `xmlrpc:"display_name,omitempty"`
	Id                 *Int      `xmlrpc:"id,omitempty"`
	Name               *String   `xmlrpc:"name,omitempty"`
	ProductIds         *Relation `xmlrpc:"product_ids,omitempty"`
	ProductProductIds  *Relation `xmlrpc:"product_product_ids,omitempty"`
	ProductTemplateIds *Relation `xmlrpc:"product_template_ids,omitempty"`
	Sequence           *Int      `xmlrpc:"sequence,omitempty"`
	WriteDate          *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid           *Many2One `xmlrpc:"write_uid,omitempty"`
}

// ProductTags represents array of product.tag model.
type ProductTags []ProductTag

// ProductTagModel is the odoo model name.
const ProductTagModel = "product.tag"

// Many2One convert ProductTag to *Many2One.
func (pt *ProductTag) Many2One() *Many2One {
	return NewMany2One(pt.Id.Get(), "")
}

// CreateProductTag creates a new product.tag model and returns its id.
func (c *Client) CreateProductTag(pt *ProductTag) (int64, error) {
	ids, err := c.CreateProductTags([]*ProductTag{pt})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProductTag creates a new product.tag model and returns its id.
func (c *Client) CreateProductTags(pts []*ProductTag) ([]int64, error) {
	var vv []interface{}
	for _, v := range pts {
		vv = append(vv, v)
	}
	return c.Create(ProductTagModel, vv, nil)
}

// UpdateProductTag updates an existing product.tag record.
func (c *Client) UpdateProductTag(pt *ProductTag) error {
	return c.UpdateProductTags([]int64{pt.Id.Get()}, pt)
}

// UpdateProductTags updates existing product.tag records.
// All records (represented by ids) will be updated by pt values.
func (c *Client) UpdateProductTags(ids []int64, pt *ProductTag) error {
	return c.Update(ProductTagModel, ids, pt, nil)
}

// DeleteProductTag deletes an existing product.tag record.
func (c *Client) DeleteProductTag(id int64) error {
	return c.DeleteProductTags([]int64{id})
}

// DeleteProductTags deletes existing product.tag records.
func (c *Client) DeleteProductTags(ids []int64) error {
	return c.Delete(ProductTagModel, ids)
}

// GetProductTag gets product.tag existing record.
func (c *Client) GetProductTag(id int64) (*ProductTag, error) {
	pts, err := c.GetProductTags([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*pts)[0]), nil
}

// GetProductTags gets product.tag existing records.
func (c *Client) GetProductTags(ids []int64) (*ProductTags, error) {
	pts := &ProductTags{}
	if err := c.Read(ProductTagModel, ids, nil, pts); err != nil {
		return nil, err
	}
	return pts, nil
}

// FindProductTag finds product.tag record by querying it with criteria.
func (c *Client) FindProductTag(criteria *Criteria) (*ProductTag, error) {
	pts := &ProductTags{}
	if err := c.SearchRead(ProductTagModel, criteria, NewOptions().Limit(1), pts); err != nil {
		return nil, err
	}
	return &((*pts)[0]), nil
}

// FindProductTags finds product.tag records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductTags(criteria *Criteria, options *Options) (*ProductTags, error) {
	pts := &ProductTags{}
	if err := c.SearchRead(ProductTagModel, criteria, options, pts); err != nil {
		return nil, err
	}
	return pts, nil
}

// FindProductTagIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductTagIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ProductTagModel, criteria, options)
}

// FindProductTagId finds record id by querying it with criteria.
func (c *Client) FindProductTagId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProductTagModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
