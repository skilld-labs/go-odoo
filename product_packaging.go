package odoo

import (
	"fmt"
)

// ProductPackaging represents product.packaging model.
type ProductPackaging struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	Barcode     *String   `xmlrpc:"barcode,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	Name        *String   `xmlrpc:"name,omptempty"`
	ProductId   *Many2One `xmlrpc:"product_id,omptempty"`
	Qty         *Float    `xmlrpc:"qty,omptempty"`
	Sequence    *Int      `xmlrpc:"sequence,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// ProductPackagings represents array of product.packaging model.
type ProductPackagings []ProductPackaging

// ProductPackagingModel is the odoo model name.
const ProductPackagingModel = "product.packaging"

// Many2One convert ProductPackaging to *Many2One.
func (pp *ProductPackaging) Many2One() *Many2One {
	return NewMany2One(pp.Id.Get(), "")
}

// CreateProductPackaging creates a new product.packaging model and returns its id.
func (c *Client) CreateProductPackaging(pp *ProductPackaging) (int64, error) {
	ids, err := c.CreateProductPackagings([]*ProductPackaging{pp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProductPackaging creates a new product.packaging model and returns its id.
func (c *Client) CreateProductPackagings(pps []*ProductPackaging) ([]int64, error) {
	var vv []interface{}
	for _, v := range pps {
		vv = append(vv, v)
	}
	return c.Create(ProductPackagingModel, vv)
}

// UpdateProductPackaging updates an existing product.packaging record.
func (c *Client) UpdateProductPackaging(pp *ProductPackaging) error {
	return c.UpdateProductPackagings([]int64{pp.Id.Get()}, pp)
}

// UpdateProductPackagings updates existing product.packaging records.
// All records (represented by ids) will be updated by pp values.
func (c *Client) UpdateProductPackagings(ids []int64, pp *ProductPackaging) error {
	return c.Update(ProductPackagingModel, ids, pp)
}

// DeleteProductPackaging deletes an existing product.packaging record.
func (c *Client) DeleteProductPackaging(id int64) error {
	return c.DeleteProductPackagings([]int64{id})
}

// DeleteProductPackagings deletes existing product.packaging records.
func (c *Client) DeleteProductPackagings(ids []int64) error {
	return c.Delete(ProductPackagingModel, ids)
}

// GetProductPackaging gets product.packaging existing record.
func (c *Client) GetProductPackaging(id int64) (*ProductPackaging, error) {
	pps, err := c.GetProductPackagings([]int64{id})
	if err != nil {
		return nil, err
	}
	if pps != nil && len(*pps) > 0 {
		return &((*pps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of product.packaging not found", id)
}

// GetProductPackagings gets product.packaging existing records.
func (c *Client) GetProductPackagings(ids []int64) (*ProductPackagings, error) {
	pps := &ProductPackagings{}
	if err := c.Read(ProductPackagingModel, ids, nil, pps); err != nil {
		return nil, err
	}
	return pps, nil
}

// FindProductPackaging finds product.packaging record by querying it with criteria.
func (c *Client) FindProductPackaging(criteria *Criteria) (*ProductPackaging, error) {
	pps := &ProductPackagings{}
	if err := c.SearchRead(ProductPackagingModel, criteria, NewOptions().Limit(1), pps); err != nil {
		return nil, err
	}
	if pps != nil && len(*pps) > 0 {
		return &((*pps)[0]), nil
	}
	return nil, fmt.Errorf("product.packaging was not found with criteria %v", criteria)
}

// FindProductPackagings finds product.packaging records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductPackagings(criteria *Criteria, options *Options) (*ProductPackagings, error) {
	pps := &ProductPackagings{}
	if err := c.SearchRead(ProductPackagingModel, criteria, options, pps); err != nil {
		return nil, err
	}
	return pps, nil
}

// FindProductPackagingIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductPackagingIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ProductPackagingModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindProductPackagingId finds record id by querying it with criteria.
func (c *Client) FindProductPackagingId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProductPackagingModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("product.packaging was not found with criteria %v and options %v", criteria, options)
}
