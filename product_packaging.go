package odoo

// ProductPackaging represents product.packaging model.
type ProductPackaging struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	Barcode     *String   `xmlrpc:"barcode,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	ProductId   *Many2One `xmlrpc:"product_id,omitempty"`
	Qty         *Float    `xmlrpc:"qty,omitempty"`
	Sequence    *Int      `xmlrpc:"sequence,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
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

// CreateProductPackagings creates a new product.packaging model and returns its id.
func (c *Client) CreateProductPackagings(pps []*ProductPackaging) ([]int64, error) {
	var vv []interface{}
	for _, v := range pps {
		vv = append(vv, v)
	}
	return c.Create(ProductPackagingModel, vv, nil)
}

// UpdateProductPackaging updates an existing product.packaging record.
func (c *Client) UpdateProductPackaging(pp *ProductPackaging) error {
	return c.UpdateProductPackagings([]int64{pp.Id.Get()}, pp)
}

// UpdateProductPackagings updates existing product.packaging records.
// All records (represented by ids) will be updated by pp values.
func (c *Client) UpdateProductPackagings(ids []int64, pp *ProductPackaging) error {
	return c.Update(ProductPackagingModel, ids, pp, nil)
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
	return &((*pps)[0]), nil
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
	return &((*pps)[0]), nil
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
	return c.Search(ProductPackagingModel, criteria, options)
}

// FindProductPackagingId finds record id by querying it with criteria.
func (c *Client) FindProductPackagingId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProductPackagingModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
