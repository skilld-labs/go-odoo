package odoo

// StockRequestCount represents stock.request.count model.
type StockRequestCount struct {
	AccountingDate *Time      `xmlrpc:"accounting_date,omitempty"`
	CreateDate     *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid      *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName    *String    `xmlrpc:"display_name,omitempty"`
	Id             *Int       `xmlrpc:"id,omitempty"`
	InventoryDate  *Time      `xmlrpc:"inventory_date,omitempty"`
	QuantIds       *Relation  `xmlrpc:"quant_ids,omitempty"`
	SetCount       *Selection `xmlrpc:"set_count,omitempty"`
	UserId         *Many2One  `xmlrpc:"user_id,omitempty"`
	WriteDate      *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid       *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// StockRequestCounts represents array of stock.request.count model.
type StockRequestCounts []StockRequestCount

// StockRequestCountModel is the odoo model name.
const StockRequestCountModel = "stock.request.count"

// Many2One convert StockRequestCount to *Many2One.
func (src *StockRequestCount) Many2One() *Many2One {
	return NewMany2One(src.Id.Get(), "")
}

// CreateStockRequestCount creates a new stock.request.count model and returns its id.
func (c *Client) CreateStockRequestCount(src *StockRequestCount) (int64, error) {
	ids, err := c.CreateStockRequestCounts([]*StockRequestCount{src})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStockRequestCount creates a new stock.request.count model and returns its id.
func (c *Client) CreateStockRequestCounts(srcs []*StockRequestCount) ([]int64, error) {
	var vv []interface{}
	for _, v := range srcs {
		vv = append(vv, v)
	}
	return c.Create(StockRequestCountModel, vv, nil)
}

// UpdateStockRequestCount updates an existing stock.request.count record.
func (c *Client) UpdateStockRequestCount(src *StockRequestCount) error {
	return c.UpdateStockRequestCounts([]int64{src.Id.Get()}, src)
}

// UpdateStockRequestCounts updates existing stock.request.count records.
// All records (represented by ids) will be updated by src values.
func (c *Client) UpdateStockRequestCounts(ids []int64, src *StockRequestCount) error {
	return c.Update(StockRequestCountModel, ids, src, nil)
}

// DeleteStockRequestCount deletes an existing stock.request.count record.
func (c *Client) DeleteStockRequestCount(id int64) error {
	return c.DeleteStockRequestCounts([]int64{id})
}

// DeleteStockRequestCounts deletes existing stock.request.count records.
func (c *Client) DeleteStockRequestCounts(ids []int64) error {
	return c.Delete(StockRequestCountModel, ids)
}

// GetStockRequestCount gets stock.request.count existing record.
func (c *Client) GetStockRequestCount(id int64) (*StockRequestCount, error) {
	srcs, err := c.GetStockRequestCounts([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*srcs)[0]), nil
}

// GetStockRequestCounts gets stock.request.count existing records.
func (c *Client) GetStockRequestCounts(ids []int64) (*StockRequestCounts, error) {
	srcs := &StockRequestCounts{}
	if err := c.Read(StockRequestCountModel, ids, nil, srcs); err != nil {
		return nil, err
	}
	return srcs, nil
}

// FindStockRequestCount finds stock.request.count record by querying it with criteria.
func (c *Client) FindStockRequestCount(criteria *Criteria) (*StockRequestCount, error) {
	srcs := &StockRequestCounts{}
	if err := c.SearchRead(StockRequestCountModel, criteria, NewOptions().Limit(1), srcs); err != nil {
		return nil, err
	}
	return &((*srcs)[0]), nil
}

// FindStockRequestCounts finds stock.request.count records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockRequestCounts(criteria *Criteria, options *Options) (*StockRequestCounts, error) {
	srcs := &StockRequestCounts{}
	if err := c.SearchRead(StockRequestCountModel, criteria, options, srcs); err != nil {
		return nil, err
	}
	return srcs, nil
}

// FindStockRequestCountIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockRequestCountIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(StockRequestCountModel, criteria, options)
}

// FindStockRequestCountId finds record id by querying it with criteria.
func (c *Client) FindStockRequestCountId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockRequestCountModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
