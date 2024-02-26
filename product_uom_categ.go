package odoo

// ProductUomCateg represents product.uom.categ model.
type ProductUomCateg struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	Name        *String   `xmlrpc:"name,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// ProductUomCategs represents array of product.uom.categ model.
type ProductUomCategs []ProductUomCateg

// ProductUomCategModel is the odoo model name.
const ProductUomCategModel = "product.uom.categ"

// Many2One convert ProductUomCateg to *Many2One.
func (puc *ProductUomCateg) Many2One() *Many2One {
	return NewMany2One(puc.Id.Get(), "")
}

// CreateProductUomCateg creates a new product.uom.categ model and returns its id.
func (c *Client) CreateProductUomCateg(puc *ProductUomCateg) (int64, error) {
	ids, err := c.CreateProductUomCategs([]*ProductUomCateg{puc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProductUomCateg creates a new product.uom.categ model and returns its id.
func (c *Client) CreateProductUomCategs(pucs []*ProductUomCateg) ([]int64, error) {
	var vv []interface{}
	for _, v := range pucs {
		vv = append(vv, v)
	}
	return c.Create(ProductUomCategModel, vv, nil)
}

// UpdateProductUomCateg updates an existing product.uom.categ record.
func (c *Client) UpdateProductUomCateg(puc *ProductUomCateg) error {
	return c.UpdateProductUomCategs([]int64{puc.Id.Get()}, puc)
}

// UpdateProductUomCategs updates existing product.uom.categ records.
// All records (represented by ids) will be updated by puc values.
func (c *Client) UpdateProductUomCategs(ids []int64, puc *ProductUomCateg) error {
	return c.Update(ProductUomCategModel, ids, puc, nil)
}

// DeleteProductUomCateg deletes an existing product.uom.categ record.
func (c *Client) DeleteProductUomCateg(id int64) error {
	return c.DeleteProductUomCategs([]int64{id})
}

// DeleteProductUomCategs deletes existing product.uom.categ records.
func (c *Client) DeleteProductUomCategs(ids []int64) error {
	return c.Delete(ProductUomCategModel, ids)
}

// GetProductUomCateg gets product.uom.categ existing record.
func (c *Client) GetProductUomCateg(id int64) (*ProductUomCateg, error) {
	pucs, err := c.GetProductUomCategs([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*pucs)[0]), nil
}

// GetProductUomCategs gets product.uom.categ existing records.
func (c *Client) GetProductUomCategs(ids []int64) (*ProductUomCategs, error) {
	pucs := &ProductUomCategs{}
	if err := c.Read(ProductUomCategModel, ids, nil, pucs); err != nil {
		return nil, err
	}
	return pucs, nil
}

// FindProductUomCateg finds product.uom.categ record by querying it with criteria.
func (c *Client) FindProductUomCateg(criteria *Criteria) (*ProductUomCateg, error) {
	pucs := &ProductUomCategs{}
	if err := c.SearchRead(ProductUomCategModel, criteria, NewOptions().Limit(1), pucs); err != nil {
		return nil, err
	}
	return &((*pucs)[0]), nil
}

// FindProductUomCategs finds product.uom.categ records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductUomCategs(criteria *Criteria, options *Options) (*ProductUomCategs, error) {
	pucs := &ProductUomCategs{}
	if err := c.SearchRead(ProductUomCategModel, criteria, options, pucs); err != nil {
		return nil, err
	}
	return pucs, nil
}

// FindProductUomCategIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductUomCategIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ProductUomCategModel, criteria, options)
}

// FindProductUomCategId finds record id by querying it with criteria.
func (c *Client) FindProductUomCategId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProductUomCategModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
