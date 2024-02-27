package odoo

// ProductSupplierinfo represents product.supplierinfo model.
type ProductSupplierinfo struct {
	LastUpdate          *Time     `xmlrpc:"__last_update,omitempty"`
	CompanyId           *Many2One `xmlrpc:"company_id,omitempty"`
	CreateDate          *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid           *Many2One `xmlrpc:"create_uid,omitempty"`
	CurrencyId          *Many2One `xmlrpc:"currency_id,omitempty"`
	DateEnd             *Time     `xmlrpc:"date_end,omitempty"`
	DateStart           *Time     `xmlrpc:"date_start,omitempty"`
	Delay               *Int      `xmlrpc:"delay,omitempty"`
	DisplayName         *String   `xmlrpc:"display_name,omitempty"`
	Id                  *Int      `xmlrpc:"id,omitempty"`
	MinQty              *Float    `xmlrpc:"min_qty,omitempty"`
	Name                *Many2One `xmlrpc:"name,omitempty"`
	Price               *Float    `xmlrpc:"price,omitempty"`
	ProductCode         *String   `xmlrpc:"product_code,omitempty"`
	ProductId           *Many2One `xmlrpc:"product_id,omitempty"`
	ProductName         *String   `xmlrpc:"product_name,omitempty"`
	ProductTmplId       *Many2One `xmlrpc:"product_tmpl_id,omitempty"`
	ProductUom          *Many2One `xmlrpc:"product_uom,omitempty"`
	ProductVariantCount *Int      `xmlrpc:"product_variant_count,omitempty"`
	Sequence            *Int      `xmlrpc:"sequence,omitempty"`
	WriteDate           *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid            *Many2One `xmlrpc:"write_uid,omitempty"`
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
	ids, err := c.CreateProductSupplierinfos([]*ProductSupplierinfo{ps})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProductSupplierinfos creates a new product.supplierinfo model and returns its id.
func (c *Client) CreateProductSupplierinfos(pss []*ProductSupplierinfo) ([]int64, error) {
	var vv []interface{}
	for _, v := range pss {
		vv = append(vv, v)
	}
	return c.Create(ProductSupplierinfoModel, vv, nil)
}

// UpdateProductSupplierinfo updates an existing product.supplierinfo record.
func (c *Client) UpdateProductSupplierinfo(ps *ProductSupplierinfo) error {
	return c.UpdateProductSupplierinfos([]int64{ps.Id.Get()}, ps)
}

// UpdateProductSupplierinfos updates existing product.supplierinfo records.
// All records (represented by ids) will be updated by ps values.
func (c *Client) UpdateProductSupplierinfos(ids []int64, ps *ProductSupplierinfo) error {
	return c.Update(ProductSupplierinfoModel, ids, ps, nil)
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
	return &((*pss)[0]), nil
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
	return &((*pss)[0]), nil
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
	return c.Search(ProductSupplierinfoModel, criteria, options)
}

// FindProductSupplierinfoId finds record id by querying it with criteria.
func (c *Client) FindProductSupplierinfoId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProductSupplierinfoModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
