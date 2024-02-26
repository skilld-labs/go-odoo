package odoo

// StockWarnInsufficientQtyScrap represents stock.warn.insufficient.qty.scrap model.
type StockWarnInsufficientQtyScrap struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	LocationId  *Many2One `xmlrpc:"location_id,omptempty"`
	ProductId   *Many2One `xmlrpc:"product_id,omptempty"`
	QuantIds    *Relation `xmlrpc:"quant_ids,omptempty"`
	ScrapId     *Many2One `xmlrpc:"scrap_id,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// StockWarnInsufficientQtyScraps represents array of stock.warn.insufficient.qty.scrap model.
type StockWarnInsufficientQtyScraps []StockWarnInsufficientQtyScrap

// StockWarnInsufficientQtyScrapModel is the odoo model name.
const StockWarnInsufficientQtyScrapModel = "stock.warn.insufficient.qty.scrap"

// Many2One convert StockWarnInsufficientQtyScrap to *Many2One.
func (swiqs *StockWarnInsufficientQtyScrap) Many2One() *Many2One {
	return NewMany2One(swiqs.Id.Get(), "")
}

// CreateStockWarnInsufficientQtyScrap creates a new stock.warn.insufficient.qty.scrap model and returns its id.
func (c *Client) CreateStockWarnInsufficientQtyScrap(swiqs *StockWarnInsufficientQtyScrap) (int64, error) {
	ids, err := c.CreateStockWarnInsufficientQtyScraps([]*StockWarnInsufficientQtyScrap{swiqs})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStockWarnInsufficientQtyScrap creates a new stock.warn.insufficient.qty.scrap model and returns its id.
func (c *Client) CreateStockWarnInsufficientQtyScraps(swiqss []*StockWarnInsufficientQtyScrap) ([]int64, error) {
	var vv []interface{}
	for _, v := range swiqss {
		vv = append(vv, v)
	}
	return c.Create(StockWarnInsufficientQtyScrapModel, vv, nil)
}

// UpdateStockWarnInsufficientQtyScrap updates an existing stock.warn.insufficient.qty.scrap record.
func (c *Client) UpdateStockWarnInsufficientQtyScrap(swiqs *StockWarnInsufficientQtyScrap) error {
	return c.UpdateStockWarnInsufficientQtyScraps([]int64{swiqs.Id.Get()}, swiqs)
}

// UpdateStockWarnInsufficientQtyScraps updates existing stock.warn.insufficient.qty.scrap records.
// All records (represented by ids) will be updated by swiqs values.
func (c *Client) UpdateStockWarnInsufficientQtyScraps(ids []int64, swiqs *StockWarnInsufficientQtyScrap) error {
	return c.Update(StockWarnInsufficientQtyScrapModel, ids, swiqs, nil)
}

// DeleteStockWarnInsufficientQtyScrap deletes an existing stock.warn.insufficient.qty.scrap record.
func (c *Client) DeleteStockWarnInsufficientQtyScrap(id int64) error {
	return c.DeleteStockWarnInsufficientQtyScraps([]int64{id})
}

// DeleteStockWarnInsufficientQtyScraps deletes existing stock.warn.insufficient.qty.scrap records.
func (c *Client) DeleteStockWarnInsufficientQtyScraps(ids []int64) error {
	return c.Delete(StockWarnInsufficientQtyScrapModel, ids)
}

// GetStockWarnInsufficientQtyScrap gets stock.warn.insufficient.qty.scrap existing record.
func (c *Client) GetStockWarnInsufficientQtyScrap(id int64) (*StockWarnInsufficientQtyScrap, error) {
	swiqss, err := c.GetStockWarnInsufficientQtyScraps([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*swiqss)[0]), nil
}

// GetStockWarnInsufficientQtyScraps gets stock.warn.insufficient.qty.scrap existing records.
func (c *Client) GetStockWarnInsufficientQtyScraps(ids []int64) (*StockWarnInsufficientQtyScraps, error) {
	swiqss := &StockWarnInsufficientQtyScraps{}
	if err := c.Read(StockWarnInsufficientQtyScrapModel, ids, nil, swiqss); err != nil {
		return nil, err
	}
	return swiqss, nil
}

// FindStockWarnInsufficientQtyScrap finds stock.warn.insufficient.qty.scrap record by querying it with criteria.
func (c *Client) FindStockWarnInsufficientQtyScrap(criteria *Criteria) (*StockWarnInsufficientQtyScrap, error) {
	swiqss := &StockWarnInsufficientQtyScraps{}
	if err := c.SearchRead(StockWarnInsufficientQtyScrapModel, criteria, NewOptions().Limit(1), swiqss); err != nil {
		return nil, err
	}
	return &((*swiqss)[0]), nil
}

// FindStockWarnInsufficientQtyScraps finds stock.warn.insufficient.qty.scrap records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockWarnInsufficientQtyScraps(criteria *Criteria, options *Options) (*StockWarnInsufficientQtyScraps, error) {
	swiqss := &StockWarnInsufficientQtyScraps{}
	if err := c.SearchRead(StockWarnInsufficientQtyScrapModel, criteria, options, swiqss); err != nil {
		return nil, err
	}
	return swiqss, nil
}

// FindStockWarnInsufficientQtyScrapIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockWarnInsufficientQtyScrapIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(StockWarnInsufficientQtyScrapModel, criteria, options)
}

// FindStockWarnInsufficientQtyScrapId finds record id by querying it with criteria.
func (c *Client) FindStockWarnInsufficientQtyScrapId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockWarnInsufficientQtyScrapModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
