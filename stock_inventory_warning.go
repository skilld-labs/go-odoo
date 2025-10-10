package odoo

// StockInventoryWarning represents stock.inventory.warning model.
type StockInventoryWarning struct {
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	QuantIds    *Relation `xmlrpc:"quant_ids,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// StockInventoryWarnings represents array of stock.inventory.warning model.
type StockInventoryWarnings []StockInventoryWarning

// StockInventoryWarningModel is the odoo model name.
const StockInventoryWarningModel = "stock.inventory.warning"

// Many2One convert StockInventoryWarning to *Many2One.
func (siw *StockInventoryWarning) Many2One() *Many2One {
	return NewMany2One(siw.Id.Get(), "")
}

// CreateStockInventoryWarning creates a new stock.inventory.warning model and returns its id.
func (c *Client) CreateStockInventoryWarning(siw *StockInventoryWarning) (int64, error) {
	ids, err := c.CreateStockInventoryWarnings([]*StockInventoryWarning{siw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStockInventoryWarning creates a new stock.inventory.warning model and returns its id.
func (c *Client) CreateStockInventoryWarnings(siws []*StockInventoryWarning) ([]int64, error) {
	var vv []interface{}
	for _, v := range siws {
		vv = append(vv, v)
	}
	return c.Create(StockInventoryWarningModel, vv, nil)
}

// UpdateStockInventoryWarning updates an existing stock.inventory.warning record.
func (c *Client) UpdateStockInventoryWarning(siw *StockInventoryWarning) error {
	return c.UpdateStockInventoryWarnings([]int64{siw.Id.Get()}, siw)
}

// UpdateStockInventoryWarnings updates existing stock.inventory.warning records.
// All records (represented by ids) will be updated by siw values.
func (c *Client) UpdateStockInventoryWarnings(ids []int64, siw *StockInventoryWarning) error {
	return c.Update(StockInventoryWarningModel, ids, siw, nil)
}

// DeleteStockInventoryWarning deletes an existing stock.inventory.warning record.
func (c *Client) DeleteStockInventoryWarning(id int64) error {
	return c.DeleteStockInventoryWarnings([]int64{id})
}

// DeleteStockInventoryWarnings deletes existing stock.inventory.warning records.
func (c *Client) DeleteStockInventoryWarnings(ids []int64) error {
	return c.Delete(StockInventoryWarningModel, ids)
}

// GetStockInventoryWarning gets stock.inventory.warning existing record.
func (c *Client) GetStockInventoryWarning(id int64) (*StockInventoryWarning, error) {
	siws, err := c.GetStockInventoryWarnings([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*siws)[0]), nil
}

// GetStockInventoryWarnings gets stock.inventory.warning existing records.
func (c *Client) GetStockInventoryWarnings(ids []int64) (*StockInventoryWarnings, error) {
	siws := &StockInventoryWarnings{}
	if err := c.Read(StockInventoryWarningModel, ids, nil, siws); err != nil {
		return nil, err
	}
	return siws, nil
}

// FindStockInventoryWarning finds stock.inventory.warning record by querying it with criteria.
func (c *Client) FindStockInventoryWarning(criteria *Criteria) (*StockInventoryWarning, error) {
	siws := &StockInventoryWarnings{}
	if err := c.SearchRead(StockInventoryWarningModel, criteria, NewOptions().Limit(1), siws); err != nil {
		return nil, err
	}
	return &((*siws)[0]), nil
}

// FindStockInventoryWarnings finds stock.inventory.warning records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockInventoryWarnings(criteria *Criteria, options *Options) (*StockInventoryWarnings, error) {
	siws := &StockInventoryWarnings{}
	if err := c.SearchRead(StockInventoryWarningModel, criteria, options, siws); err != nil {
		return nil, err
	}
	return siws, nil
}

// FindStockInventoryWarningIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockInventoryWarningIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(StockInventoryWarningModel, criteria, options)
}

// FindStockInventoryWarningId finds record id by querying it with criteria.
func (c *Client) FindStockInventoryWarningId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockInventoryWarningModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
