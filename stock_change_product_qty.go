package odoo

// StockChangeProductQty represents stock.change.product.qty model.
type StockChangeProductQty struct {
	LastUpdate          *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate          *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid           *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName         *String   `xmlrpc:"display_name,omitempty"`
	Id                  *Int      `xmlrpc:"id,omitempty"`
	LocationId          *Many2One `xmlrpc:"location_id,omitempty"`
	LotId               *Many2One `xmlrpc:"lot_id,omitempty"`
	NewQuantity         *Float    `xmlrpc:"new_quantity,omitempty"`
	ProductId           *Many2One `xmlrpc:"product_id,omitempty"`
	ProductTmplId       *Many2One `xmlrpc:"product_tmpl_id,omitempty"`
	ProductVariantCount *Int      `xmlrpc:"product_variant_count,omitempty"`
	WriteDate           *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid            *Many2One `xmlrpc:"write_uid,omitempty"`
}

// StockChangeProductQtys represents array of stock.change.product.qty model.
type StockChangeProductQtys []StockChangeProductQty

// StockChangeProductQtyModel is the odoo model name.
const StockChangeProductQtyModel = "stock.change.product.qty"

// Many2One convert StockChangeProductQty to *Many2One.
func (scpq *StockChangeProductQty) Many2One() *Many2One {
	return NewMany2One(scpq.Id.Get(), "")
}

// CreateStockChangeProductQty creates a new stock.change.product.qty model and returns its id.
func (c *Client) CreateStockChangeProductQty(scpq *StockChangeProductQty) (int64, error) {
	ids, err := c.CreateStockChangeProductQtys([]*StockChangeProductQty{scpq})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStockChangeProductQtys creates a new stock.change.product.qty model and returns its id.
func (c *Client) CreateStockChangeProductQtys(scpqs []*StockChangeProductQty) ([]int64, error) {
	var vv []interface{}
	for _, v := range scpqs {
		vv = append(vv, v)
	}
	return c.Create(StockChangeProductQtyModel, vv, nil)
}

// UpdateStockChangeProductQty updates an existing stock.change.product.qty record.
func (c *Client) UpdateStockChangeProductQty(scpq *StockChangeProductQty) error {
	return c.UpdateStockChangeProductQtys([]int64{scpq.Id.Get()}, scpq)
}

// UpdateStockChangeProductQtys updates existing stock.change.product.qty records.
// All records (represented by ids) will be updated by scpq values.
func (c *Client) UpdateStockChangeProductQtys(ids []int64, scpq *StockChangeProductQty) error {
	return c.Update(StockChangeProductQtyModel, ids, scpq, nil)
}

// DeleteStockChangeProductQty deletes an existing stock.change.product.qty record.
func (c *Client) DeleteStockChangeProductQty(id int64) error {
	return c.DeleteStockChangeProductQtys([]int64{id})
}

// DeleteStockChangeProductQtys deletes existing stock.change.product.qty records.
func (c *Client) DeleteStockChangeProductQtys(ids []int64) error {
	return c.Delete(StockChangeProductQtyModel, ids)
}

// GetStockChangeProductQty gets stock.change.product.qty existing record.
func (c *Client) GetStockChangeProductQty(id int64) (*StockChangeProductQty, error) {
	scpqs, err := c.GetStockChangeProductQtys([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*scpqs)[0]), nil
}

// GetStockChangeProductQtys gets stock.change.product.qty existing records.
func (c *Client) GetStockChangeProductQtys(ids []int64) (*StockChangeProductQtys, error) {
	scpqs := &StockChangeProductQtys{}
	if err := c.Read(StockChangeProductQtyModel, ids, nil, scpqs); err != nil {
		return nil, err
	}
	return scpqs, nil
}

// FindStockChangeProductQty finds stock.change.product.qty record by querying it with criteria.
func (c *Client) FindStockChangeProductQty(criteria *Criteria) (*StockChangeProductQty, error) {
	scpqs := &StockChangeProductQtys{}
	if err := c.SearchRead(StockChangeProductQtyModel, criteria, NewOptions().Limit(1), scpqs); err != nil {
		return nil, err
	}
	return &((*scpqs)[0]), nil
}

// FindStockChangeProductQtys finds stock.change.product.qty records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockChangeProductQtys(criteria *Criteria, options *Options) (*StockChangeProductQtys, error) {
	scpqs := &StockChangeProductQtys{}
	if err := c.SearchRead(StockChangeProductQtyModel, criteria, options, scpqs); err != nil {
		return nil, err
	}
	return scpqs, nil
}

// FindStockChangeProductQtyIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockChangeProductQtyIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(StockChangeProductQtyModel, criteria, options)
}

// FindStockChangeProductQtyId finds record id by querying it with criteria.
func (c *Client) FindStockChangeProductQtyId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockChangeProductQtyModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
