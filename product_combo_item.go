package odoo

// ProductComboItem represents product.combo.item model.
type ProductComboItem struct {
	ComboId     *Many2One `xmlrpc:"combo_id,omitempty"`
	CompanyId   *Many2One `xmlrpc:"company_id,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	CurrencyId  *Many2One `xmlrpc:"currency_id,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	ExtraPrice  *Float    `xmlrpc:"extra_price,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	LstPrice    *Float    `xmlrpc:"lst_price,omitempty"`
	ProductId   *Many2One `xmlrpc:"product_id,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// ProductComboItems represents array of product.combo.item model.
type ProductComboItems []ProductComboItem

// ProductComboItemModel is the odoo model name.
const ProductComboItemModel = "product.combo.item"

// Many2One convert ProductComboItem to *Many2One.
func (pci *ProductComboItem) Many2One() *Many2One {
	return NewMany2One(pci.Id.Get(), "")
}

// CreateProductComboItem creates a new product.combo.item model and returns its id.
func (c *Client) CreateProductComboItem(pci *ProductComboItem) (int64, error) {
	ids, err := c.CreateProductComboItems([]*ProductComboItem{pci})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProductComboItem creates a new product.combo.item model and returns its id.
func (c *Client) CreateProductComboItems(pcis []*ProductComboItem) ([]int64, error) {
	var vv []interface{}
	for _, v := range pcis {
		vv = append(vv, v)
	}
	return c.Create(ProductComboItemModel, vv, nil)
}

// UpdateProductComboItem updates an existing product.combo.item record.
func (c *Client) UpdateProductComboItem(pci *ProductComboItem) error {
	return c.UpdateProductComboItems([]int64{pci.Id.Get()}, pci)
}

// UpdateProductComboItems updates existing product.combo.item records.
// All records (represented by ids) will be updated by pci values.
func (c *Client) UpdateProductComboItems(ids []int64, pci *ProductComboItem) error {
	return c.Update(ProductComboItemModel, ids, pci, nil)
}

// DeleteProductComboItem deletes an existing product.combo.item record.
func (c *Client) DeleteProductComboItem(id int64) error {
	return c.DeleteProductComboItems([]int64{id})
}

// DeleteProductComboItems deletes existing product.combo.item records.
func (c *Client) DeleteProductComboItems(ids []int64) error {
	return c.Delete(ProductComboItemModel, ids)
}

// GetProductComboItem gets product.combo.item existing record.
func (c *Client) GetProductComboItem(id int64) (*ProductComboItem, error) {
	pcis, err := c.GetProductComboItems([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*pcis)[0]), nil
}

// GetProductComboItems gets product.combo.item existing records.
func (c *Client) GetProductComboItems(ids []int64) (*ProductComboItems, error) {
	pcis := &ProductComboItems{}
	if err := c.Read(ProductComboItemModel, ids, nil, pcis); err != nil {
		return nil, err
	}
	return pcis, nil
}

// FindProductComboItem finds product.combo.item record by querying it with criteria.
func (c *Client) FindProductComboItem(criteria *Criteria) (*ProductComboItem, error) {
	pcis := &ProductComboItems{}
	if err := c.SearchRead(ProductComboItemModel, criteria, NewOptions().Limit(1), pcis); err != nil {
		return nil, err
	}
	return &((*pcis)[0]), nil
}

// FindProductComboItems finds product.combo.item records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductComboItems(criteria *Criteria, options *Options) (*ProductComboItems, error) {
	pcis := &ProductComboItems{}
	if err := c.SearchRead(ProductComboItemModel, criteria, options, pcis); err != nil {
		return nil, err
	}
	return pcis, nil
}

// FindProductComboItemIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductComboItemIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ProductComboItemModel, criteria, options)
}

// FindProductComboItemId finds record id by querying it with criteria.
func (c *Client) FindProductComboItemId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProductComboItemModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
