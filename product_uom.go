package odoo

import (
	"fmt"
)

// ProductUom represents product.uom model.
type ProductUom struct {
	LastUpdate  *Time      `xmlrpc:"__last_update,omptempty"`
	Active      *Bool      `xmlrpc:"active,omptempty"`
	CategoryId  *Many2One  `xmlrpc:"category_id,omptempty"`
	CreateDate  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName *String    `xmlrpc:"display_name,omptempty"`
	Factor      *Float     `xmlrpc:"factor,omptempty"`
	FactorInv   *Float     `xmlrpc:"factor_inv,omptempty"`
	Id          *Int       `xmlrpc:"id,omptempty"`
	Name        *String    `xmlrpc:"name,omptempty"`
	Rounding    *Float     `xmlrpc:"rounding,omptempty"`
	UomType     *Selection `xmlrpc:"uom_type,omptempty"`
	WriteDate   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One  `xmlrpc:"write_uid,omptempty"`
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
	return c.Create(ProductUomModel, vv)
}

// UpdateProductUom updates an existing product.uom record.
func (c *Client) UpdateProductUom(pu *ProductUom) error {
	return c.UpdateProductUoms([]int64{pu.Id.Get()}, pu)
}

// UpdateProductUoms updates existing product.uom records.
// All records (represented by ids) will be updated by pu values.
func (c *Client) UpdateProductUoms(ids []int64, pu *ProductUom) error {
	return c.Update(ProductUomModel, ids, pu)
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
	if pus != nil && len(*pus) > 0 {
		return &((*pus)[0]), nil
	}
	return nil, fmt.Errorf("id %v of product.uom not found", id)
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
	if pus != nil && len(*pus) > 0 {
		return &((*pus)[0]), nil
	}
	return nil, fmt.Errorf("product.uom was not found with criteria %v", criteria)
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
	ids, err := c.Search(ProductUomModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindProductUomId finds record id by querying it with criteria.
func (c *Client) FindProductUomId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProductUomModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("product.uom was not found with criteria %v and options %v", criteria, options)
}
