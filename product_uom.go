package odoo

// ProductUom represents product.uom model.
type ProductUom struct {
	LastUpdate  *Time      `xmlrpc:"__last_update,omitempty"`
	Active      *Bool      `xmlrpc:"active,omitempty"`
	CategoryId  *Many2One  `xmlrpc:"category_id,omitempty"`
	CreateDate  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName *String    `xmlrpc:"display_name,omitempty"`
	Factor      *Float     `xmlrpc:"factor,omitempty"`
	FactorInv   *Float     `xmlrpc:"factor_inv,omitempty"`
	Id          *Int       `xmlrpc:"id,omitempty"`
	Name        *String    `xmlrpc:"name,omitempty"`
	Rounding    *Float     `xmlrpc:"rounding,omitempty"`
	UomType     *Selection `xmlrpc:"uom_type,omitempty"`
	WriteDate   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// ProductUoms represents array of product.uom model.
type ProductUoms []ProductUom

// ProductUomModel is the odoo model name.
const ProductUomModel = "product.uom"

// Many2One convert ProductUom to *Many2One.
func (pu *ProductUom) Many2One() *Many2One {
	return NewMany2One(pu.Id.Get(), "")
}

// CreateProductUom creates a new product.uom model and returns its id.
func (c *Client) CreateProductUom(pu *ProductUom) (int64, error) {
	ids, err := c.CreateProductUoms([]*ProductUom{pu})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProductUom creates a new product.uom model and returns its id.
func (c *Client) CreateProductUoms(pus []*ProductUom) ([]int64, error) {
	var vv []interface{}
	for _, v := range pus {
		vv = append(vv, v)
	}
	return c.Create(ProductUomModel, vv, nil)
}

// UpdateProductUom updates an existing product.uom record.
func (c *Client) UpdateProductUom(pu *ProductUom) error {
	return c.UpdateProductUoms([]int64{pu.Id.Get()}, pu)
}

// UpdateProductUoms updates existing product.uom records.
// All records (represented by ids) will be updated by pu values.
func (c *Client) UpdateProductUoms(ids []int64, pu *ProductUom) error {
	return c.Update(ProductUomModel, ids, pu, nil)
}

// DeleteProductUom deletes an existing product.uom record.
func (c *Client) DeleteProductUom(id int64) error {
	return c.DeleteProductUoms([]int64{id})
}

// DeleteProductUoms deletes existing product.uom records.
func (c *Client) DeleteProductUoms(ids []int64) error {
	return c.Delete(ProductUomModel, ids)
}

// GetProductUom gets product.uom existing record.
func (c *Client) GetProductUom(id int64) (*ProductUom, error) {
	pus, err := c.GetProductUoms([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*pus)[0]), nil
}

// GetProductUoms gets product.uom existing records.
func (c *Client) GetProductUoms(ids []int64) (*ProductUoms, error) {
	pus := &ProductUoms{}
	if err := c.Read(ProductUomModel, ids, nil, pus); err != nil {
		return nil, err
	}
	return pus, nil
}

// FindProductUom finds product.uom record by querying it with criteria.
func (c *Client) FindProductUom(criteria *Criteria) (*ProductUom, error) {
	pus := &ProductUoms{}
	if err := c.SearchRead(ProductUomModel, criteria, NewOptions().Limit(1), pus); err != nil {
		return nil, err
	}
	return &((*pus)[0]), nil
}

// FindProductUoms finds product.uom records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductUoms(criteria *Criteria, options *Options) (*ProductUoms, error) {
	pus := &ProductUoms{}
	if err := c.SearchRead(ProductUomModel, criteria, options, pus); err != nil {
		return nil, err
	}
	return pus, nil
}

// FindProductUomIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductUomIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ProductUomModel, criteria, options)
}

// FindProductUomId finds record id by querying it with criteria.
func (c *Client) FindProductUomId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProductUomModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
