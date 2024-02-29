package odoo

// StockQuantityHistory represents stock.quantity.history model.
type StockQuantityHistory struct {
	LastUpdate    *Time      `xmlrpc:"__last_update,omitempty"`
	ComputeAtDate *Selection `xmlrpc:"compute_at_date,omitempty"`
	CreateDate    *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid     *Many2One  `xmlrpc:"create_uid,omitempty"`
	Date          *Time      `xmlrpc:"date,omitempty"`
	DisplayName   *String    `xmlrpc:"display_name,omitempty"`
	Id            *Int       `xmlrpc:"id,omitempty"`
	WriteDate     *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid      *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// StockQuantityHistorys represents array of stock.quantity.history model.
type StockQuantityHistorys []StockQuantityHistory

// StockQuantityHistoryModel is the odoo model name.
const StockQuantityHistoryModel = "stock.quantity.history"

// Many2One convert StockQuantityHistory to *Many2One.
func (sqh *StockQuantityHistory) Many2One() *Many2One {
	return NewMany2One(sqh.Id.Get(), "")
}

// CreateStockQuantityHistory creates a new stock.quantity.history model and returns its id.
func (c *Client) CreateStockQuantityHistory(sqh *StockQuantityHistory) (int64, error) {
	ids, err := c.CreateStockQuantityHistorys([]*StockQuantityHistory{sqh})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStockQuantityHistorys creates a new stock.quantity.history model and returns its id.
func (c *Client) CreateStockQuantityHistorys(sqhs []*StockQuantityHistory) ([]int64, error) {
	var vv []interface{}
	for _, v := range sqhs {
		vv = append(vv, v)
	}
	return c.Create(StockQuantityHistoryModel, vv, nil)
}

// UpdateStockQuantityHistory updates an existing stock.quantity.history record.
func (c *Client) UpdateStockQuantityHistory(sqh *StockQuantityHistory) error {
	return c.UpdateStockQuantityHistorys([]int64{sqh.Id.Get()}, sqh)
}

// UpdateStockQuantityHistorys updates existing stock.quantity.history records.
// All records (represented by ids) will be updated by sqh values.
func (c *Client) UpdateStockQuantityHistorys(ids []int64, sqh *StockQuantityHistory) error {
	return c.Update(StockQuantityHistoryModel, ids, sqh, nil)
}

// DeleteStockQuantityHistory deletes an existing stock.quantity.history record.
func (c *Client) DeleteStockQuantityHistory(id int64) error {
	return c.DeleteStockQuantityHistorys([]int64{id})
}

// DeleteStockQuantityHistorys deletes existing stock.quantity.history records.
func (c *Client) DeleteStockQuantityHistorys(ids []int64) error {
	return c.Delete(StockQuantityHistoryModel, ids)
}

// GetStockQuantityHistory gets stock.quantity.history existing record.
func (c *Client) GetStockQuantityHistory(id int64) (*StockQuantityHistory, error) {
	sqhs, err := c.GetStockQuantityHistorys([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*sqhs)[0]), nil
}

// GetStockQuantityHistorys gets stock.quantity.history existing records.
func (c *Client) GetStockQuantityHistorys(ids []int64) (*StockQuantityHistorys, error) {
	sqhs := &StockQuantityHistorys{}
	if err := c.Read(StockQuantityHistoryModel, ids, nil, sqhs); err != nil {
		return nil, err
	}
	return sqhs, nil
}

// FindStockQuantityHistory finds stock.quantity.history record by querying it with criteria.
func (c *Client) FindStockQuantityHistory(criteria *Criteria) (*StockQuantityHistory, error) {
	sqhs := &StockQuantityHistorys{}
	if err := c.SearchRead(StockQuantityHistoryModel, criteria, NewOptions().Limit(1), sqhs); err != nil {
		return nil, err
	}
	return &((*sqhs)[0]), nil
}

// FindStockQuantityHistorys finds stock.quantity.history records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockQuantityHistorys(criteria *Criteria, options *Options) (*StockQuantityHistorys, error) {
	sqhs := &StockQuantityHistorys{}
	if err := c.SearchRead(StockQuantityHistoryModel, criteria, options, sqhs); err != nil {
		return nil, err
	}
	return sqhs, nil
}

// FindStockQuantityHistoryIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockQuantityHistoryIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(StockQuantityHistoryModel, criteria, options)
}

// FindStockQuantityHistoryId finds record id by querying it with criteria.
func (c *Client) FindStockQuantityHistoryId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockQuantityHistoryModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
