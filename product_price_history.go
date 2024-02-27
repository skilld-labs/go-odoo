package odoo

// ProductPriceHistory represents product.price.history model.
type ProductPriceHistory struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	CompanyId   *Many2One `xmlrpc:"company_id,omitempty"`
	Cost        *Float    `xmlrpc:"cost,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	Datetime    *Time     `xmlrpc:"datetime,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	ProductId   *Many2One `xmlrpc:"product_id,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// ProductPriceHistorys represents array of product.price.history model.
type ProductPriceHistorys []ProductPriceHistory

// ProductPriceHistoryModel is the odoo model name.
const ProductPriceHistoryModel = "product.price.history"

// Many2One convert ProductPriceHistory to *Many2One.
func (pph *ProductPriceHistory) Many2One() *Many2One {
	return NewMany2One(pph.Id.Get(), "")
}

// CreateProductPriceHistory creates a new product.price.history model and returns its id.
func (c *Client) CreateProductPriceHistory(pph *ProductPriceHistory) (int64, error) {
	ids, err := c.CreateProductPriceHistorys([]*ProductPriceHistory{pph})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProductPriceHistorys creates a new product.price.history model and returns its id.
func (c *Client) CreateProductPriceHistorys(pphs []*ProductPriceHistory) ([]int64, error) {
	var vv []interface{}
	for _, v := range pphs {
		vv = append(vv, v)
	}
	return c.Create(ProductPriceHistoryModel, vv, nil)
}

// UpdateProductPriceHistory updates an existing product.price.history record.
func (c *Client) UpdateProductPriceHistory(pph *ProductPriceHistory) error {
	return c.UpdateProductPriceHistorys([]int64{pph.Id.Get()}, pph)
}

// UpdateProductPriceHistorys updates existing product.price.history records.
// All records (represented by ids) will be updated by pph values.
func (c *Client) UpdateProductPriceHistorys(ids []int64, pph *ProductPriceHistory) error {
	return c.Update(ProductPriceHistoryModel, ids, pph, nil)
}

// DeleteProductPriceHistory deletes an existing product.price.history record.
func (c *Client) DeleteProductPriceHistory(id int64) error {
	return c.DeleteProductPriceHistorys([]int64{id})
}

// DeleteProductPriceHistorys deletes existing product.price.history records.
func (c *Client) DeleteProductPriceHistorys(ids []int64) error {
	return c.Delete(ProductPriceHistoryModel, ids)
}

// GetProductPriceHistory gets product.price.history existing record.
func (c *Client) GetProductPriceHistory(id int64) (*ProductPriceHistory, error) {
	pphs, err := c.GetProductPriceHistorys([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*pphs)[0]), nil
}

// GetProductPriceHistorys gets product.price.history existing records.
func (c *Client) GetProductPriceHistorys(ids []int64) (*ProductPriceHistorys, error) {
	pphs := &ProductPriceHistorys{}
	if err := c.Read(ProductPriceHistoryModel, ids, nil, pphs); err != nil {
		return nil, err
	}
	return pphs, nil
}

// FindProductPriceHistory finds product.price.history record by querying it with criteria.
func (c *Client) FindProductPriceHistory(criteria *Criteria) (*ProductPriceHistory, error) {
	pphs := &ProductPriceHistorys{}
	if err := c.SearchRead(ProductPriceHistoryModel, criteria, NewOptions().Limit(1), pphs); err != nil {
		return nil, err
	}
	return &((*pphs)[0]), nil
}

// FindProductPriceHistorys finds product.price.history records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductPriceHistorys(criteria *Criteria, options *Options) (*ProductPriceHistorys, error) {
	pphs := &ProductPriceHistorys{}
	if err := c.SearchRead(ProductPriceHistoryModel, criteria, options, pphs); err != nil {
		return nil, err
	}
	return pphs, nil
}

// FindProductPriceHistoryIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductPriceHistoryIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ProductPriceHistoryModel, criteria, options)
}

// FindProductPriceHistoryId finds record id by querying it with criteria.
func (c *Client) FindProductPriceHistoryId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProductPriceHistoryModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
