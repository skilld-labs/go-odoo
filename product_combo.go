package odoo

// ProductCombo represents product.combo model.
type ProductCombo struct {
	BasePrice      *Float    `xmlrpc:"base_price,omitempty"`
	ComboItemCount *Int      `xmlrpc:"combo_item_count,omitempty"`
	ComboItemIds   *Relation `xmlrpc:"combo_item_ids,omitempty"`
	CompanyId      *Many2One `xmlrpc:"company_id,omitempty"`
	CreateDate     *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid      *Many2One `xmlrpc:"create_uid,omitempty"`
	CurrencyId     *Many2One `xmlrpc:"currency_id,omitempty"`
	DisplayName    *String   `xmlrpc:"display_name,omitempty"`
	Id             *Int      `xmlrpc:"id,omitempty"`
	Name           *String   `xmlrpc:"name,omitempty"`
	Sequence       *Int      `xmlrpc:"sequence,omitempty"`
	WriteDate      *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid       *Many2One `xmlrpc:"write_uid,omitempty"`
}

// ProductCombos represents array of product.combo model.
type ProductCombos []ProductCombo

// ProductComboModel is the odoo model name.
const ProductComboModel = "product.combo"

// Many2One convert ProductCombo to *Many2One.
func (pc *ProductCombo) Many2One() *Many2One {
	return NewMany2One(pc.Id.Get(), "")
}

// CreateProductCombo creates a new product.combo model and returns its id.
func (c *Client) CreateProductCombo(pc *ProductCombo) (int64, error) {
	ids, err := c.CreateProductCombos([]*ProductCombo{pc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProductCombo creates a new product.combo model and returns its id.
func (c *Client) CreateProductCombos(pcs []*ProductCombo) ([]int64, error) {
	var vv []interface{}
	for _, v := range pcs {
		vv = append(vv, v)
	}
	return c.Create(ProductComboModel, vv, nil)
}

// UpdateProductCombo updates an existing product.combo record.
func (c *Client) UpdateProductCombo(pc *ProductCombo) error {
	return c.UpdateProductCombos([]int64{pc.Id.Get()}, pc)
}

// UpdateProductCombos updates existing product.combo records.
// All records (represented by ids) will be updated by pc values.
func (c *Client) UpdateProductCombos(ids []int64, pc *ProductCombo) error {
	return c.Update(ProductComboModel, ids, pc, nil)
}

// DeleteProductCombo deletes an existing product.combo record.
func (c *Client) DeleteProductCombo(id int64) error {
	return c.DeleteProductCombos([]int64{id})
}

// DeleteProductCombos deletes existing product.combo records.
func (c *Client) DeleteProductCombos(ids []int64) error {
	return c.Delete(ProductComboModel, ids)
}

// GetProductCombo gets product.combo existing record.
func (c *Client) GetProductCombo(id int64) (*ProductCombo, error) {
	pcs, err := c.GetProductCombos([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*pcs)[0]), nil
}

// GetProductCombos gets product.combo existing records.
func (c *Client) GetProductCombos(ids []int64) (*ProductCombos, error) {
	pcs := &ProductCombos{}
	if err := c.Read(ProductComboModel, ids, nil, pcs); err != nil {
		return nil, err
	}
	return pcs, nil
}

// FindProductCombo finds product.combo record by querying it with criteria.
func (c *Client) FindProductCombo(criteria *Criteria) (*ProductCombo, error) {
	pcs := &ProductCombos{}
	if err := c.SearchRead(ProductComboModel, criteria, NewOptions().Limit(1), pcs); err != nil {
		return nil, err
	}
	return &((*pcs)[0]), nil
}

// FindProductCombos finds product.combo records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductCombos(criteria *Criteria, options *Options) (*ProductCombos, error) {
	pcs := &ProductCombos{}
	if err := c.SearchRead(ProductComboModel, criteria, options, pcs); err != nil {
		return nil, err
	}
	return pcs, nil
}

// FindProductComboIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductComboIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ProductComboModel, criteria, options)
}

// FindProductComboId finds record id by querying it with criteria.
func (c *Client) FindProductComboId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProductComboModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
