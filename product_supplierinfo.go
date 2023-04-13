package odoo

import (
	"fmt"
)

// ProductSupplierinfo represents product.supplierinfo model.
type ProductSupplierinfo struct {
	LastUpdate          *Time     `xmlrpc:"__last_update,omptempty"`
	CompanyId           *Many2One `xmlrpc:"company_id,omptempty"`
	CreateDate          *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid           *Many2One `xmlrpc:"create_uid,omptempty"`
	CurrencyId          *Many2One `xmlrpc:"currency_id,omptempty"`
	DateEnd             *Time     `xmlrpc:"date_end,omptempty"`
	DateStart           *Time     `xmlrpc:"date_start,omptempty"`
	Delay               *Int      `xmlrpc:"delay,omptempty"`
	DisplayName         *String   `xmlrpc:"display_name,omptempty"`
	Id                  *Int      `xmlrpc:"id,omptempty"`
	MinQty              *Float    `xmlrpc:"min_qty,omptempty"`
	Name                *Many2One `xmlrpc:"name,omptempty"`
	Price               *Float    `xmlrpc:"price,omptempty"`
	ProductCode         *String   `xmlrpc:"product_code,omptempty"`
	ProductId           *Many2One `xmlrpc:"product_id,omptempty"`
	ProductName         *String   `xmlrpc:"product_name,omptempty"`
	ProductTmplId       *Many2One `xmlrpc:"product_tmpl_id,omptempty"`
	ProductUom          *Many2One `xmlrpc:"product_uom,omptempty"`
	ProductVariantCount *Int      `xmlrpc:"product_variant_count,omptempty"`
	Sequence            *Int      `xmlrpc:"sequence,omptempty"`
	WriteDate           *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid            *Many2One `xmlrpc:"write_uid,omptempty"`
}

// ProductSupplierinfos represents array of product.supplierinfo model.
type ProductSupplierinfos []ProductSupplierinfo

// ProductSupplierinfoModel is the odoo model name.
const ProductSupplierinfoModel = "product.supplierinfo"

// Many2One convert ProductSupplierinfo to *Many2One.
func (ps *ProductSupplierinfo) Many2One() *Many2One {
	return NewMany2One(ps.Id.Get(), "")
}

// CreateProductSupplierinfo creates a new product.supplierinfo model and returns its id.
func (c *Client) CreateProductSupplierinfo(ps *ProductSupplierinfo) (int64, error) {
	ids, err := c.Create(ProductSupplierinfoModel, []interface{}{ps})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProductSupplierinfo creates a new product.supplierinfo model and returns its id.
func (c *Client) CreateProductSupplierinfos(pss []*ProductSupplierinfo) ([]int64, error) {
	var vv []interface{}
	for _, v := range pss {
		vv = append(vv, v)
	}
	return c.Create(ProductSupplierinfoModel, vv)
}

// UpdateProductSupplierinfo updates an existing product.supplierinfo record.
func (c *Client) UpdateProductSupplierinfo(ps *ProductSupplierinfo) error {
	return c.UpdateProductSupplierinfos([]int64{ps.Id.Get()}, ps)
}

// UpdateProductSupplierinfos updates existing product.supplierinfo records.
// All records (represented by ids) will be updated by ps values.
func (c *Client) UpdateProductSupplierinfos(ids []int64, ps *ProductSupplierinfo) error {
	return c.Update(ProductSupplierinfoModel, ids, ps)
}

// DeleteProductSupplierinfo deletes an existing product.supplierinfo record.
func (c *Client) DeleteProductSupplierinfo(id int64) error {
	return c.DeleteProductSupplierinfos([]int64{id})
}

// DeleteProductSupplierinfos deletes existing product.supplierinfo records.
func (c *Client) DeleteProductSupplierinfos(ids []int64) error {
	return c.Delete(ProductSupplierinfoModel, ids)
}

// GetProductSupplierinfo gets product.supplierinfo existing record.
func (c *Client) GetProductSupplierinfo(id int64) (*ProductSupplierinfo, error) {
	pss, err := c.GetProductSupplierinfos([]int64{id})
	if err != nil {
		return nil, err
	}
	if pss != nil && len(*pss) > 0 {
		return &((*pss)[0]), nil
	}
	return nil, fmt.Errorf("id %v of product.supplierinfo not found", id)
}

// GetProductSupplierinfos gets product.supplierinfo existing records.
func (c *Client) GetProductSupplierinfos(ids []int64) (*ProductSupplierinfos, error) {
	pss := &ProductSupplierinfos{}
	if err := c.Read(ProductSupplierinfoModel, ids, nil, pss); err != nil {
		return nil, err
	}
	return pss, nil
}

// FindProductSupplierinfo finds product.supplierinfo record by querying it with criteria.
func (c *Client) FindProductSupplierinfo(criteria *Criteria) (*ProductSupplierinfo, error) {
	pss := &ProductSupplierinfos{}
	if err := c.SearchRead(ProductSupplierinfoModel, criteria, NewOptions().Limit(1), pss); err != nil {
		return nil, err
	}
	if pss != nil && len(*pss) > 0 {
		return &((*pss)[0]), nil
	}
	return nil, fmt.Errorf("product.supplierinfo was not found with criteria %v", criteria)
}

// FindProductSupplierinfos finds product.supplierinfo records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductSupplierinfos(criteria *Criteria, options *Options) (*ProductSupplierinfos, error) {
	pss := &ProductSupplierinfos{}
	if err := c.SearchRead(ProductSupplierinfoModel, criteria, options, pss); err != nil {
		return nil, err
	}
	return pss, nil
}

// FindProductSupplierinfoIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductSupplierinfoIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ProductSupplierinfoModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindProductSupplierinfoId finds record id by querying it with criteria.
func (c *Client) FindProductSupplierinfoId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProductSupplierinfoModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("product.supplierinfo was not found with criteria %v and options %v", criteria, options)
}
