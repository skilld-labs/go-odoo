package odoo

// StockScrap represents stock.scrap model.
type StockScrap struct {
	LastUpdate      *Time      `xmlrpc:"__last_update,omitempty"`
	CreateDate      *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid       *Many2One  `xmlrpc:"create_uid,omitempty"`
	DateExpected    *Time      `xmlrpc:"date_expected,omitempty"`
	DisplayName     *String    `xmlrpc:"display_name,omitempty"`
	Id              *Int       `xmlrpc:"id,omitempty"`
	LocationId      *Many2One  `xmlrpc:"location_id,omitempty"`
	LotId           *Many2One  `xmlrpc:"lot_id,omitempty"`
	MoveId          *Many2One  `xmlrpc:"move_id,omitempty"`
	Name            *String    `xmlrpc:"name,omitempty"`
	Origin          *String    `xmlrpc:"origin,omitempty"`
	OwnerId         *Many2One  `xmlrpc:"owner_id,omitempty"`
	PackageId       *Many2One  `xmlrpc:"package_id,omitempty"`
	PickingId       *Many2One  `xmlrpc:"picking_id,omitempty"`
	ProductId       *Many2One  `xmlrpc:"product_id,omitempty"`
	ProductUomId    *Many2One  `xmlrpc:"product_uom_id,omitempty"`
	ScrapLocationId *Many2One  `xmlrpc:"scrap_location_id,omitempty"`
	ScrapQty        *Float     `xmlrpc:"scrap_qty,omitempty"`
	State           *Selection `xmlrpc:"state,omitempty"`
	Tracking        *Selection `xmlrpc:"tracking,omitempty"`
	WriteDate       *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid        *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// StockScraps represents array of stock.scrap model.
type StockScraps []StockScrap

// StockScrapModel is the odoo model name.
const StockScrapModel = "stock.scrap"

// Many2One convert StockScrap to *Many2One.
func (ss *StockScrap) Many2One() *Many2One {
	return NewMany2One(ss.Id.Get(), "")
}

// CreateStockScrap creates a new stock.scrap model and returns its id.
func (c *Client) CreateStockScrap(ss *StockScrap) (int64, error) {
	ids, err := c.CreateStockScraps([]*StockScrap{ss})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStockScraps creates a new stock.scrap model and returns its id.
func (c *Client) CreateStockScraps(sss []*StockScrap) ([]int64, error) {
	var vv []interface{}
	for _, v := range sss {
		vv = append(vv, v)
	}
	return c.Create(StockScrapModel, vv, nil)
}

// UpdateStockScrap updates an existing stock.scrap record.
func (c *Client) UpdateStockScrap(ss *StockScrap) error {
	return c.UpdateStockScraps([]int64{ss.Id.Get()}, ss)
}

// UpdateStockScraps updates existing stock.scrap records.
// All records (represented by ids) will be updated by ss values.
func (c *Client) UpdateStockScraps(ids []int64, ss *StockScrap) error {
	return c.Update(StockScrapModel, ids, ss, nil)
}

// DeleteStockScrap deletes an existing stock.scrap record.
func (c *Client) DeleteStockScrap(id int64) error {
	return c.DeleteStockScraps([]int64{id})
}

// DeleteStockScraps deletes existing stock.scrap records.
func (c *Client) DeleteStockScraps(ids []int64) error {
	return c.Delete(StockScrapModel, ids)
}

// GetStockScrap gets stock.scrap existing record.
func (c *Client) GetStockScrap(id int64) (*StockScrap, error) {
	sss, err := c.GetStockScraps([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*sss)[0]), nil
}

// GetStockScraps gets stock.scrap existing records.
func (c *Client) GetStockScraps(ids []int64) (*StockScraps, error) {
	sss := &StockScraps{}
	if err := c.Read(StockScrapModel, ids, nil, sss); err != nil {
		return nil, err
	}
	return sss, nil
}

// FindStockScrap finds stock.scrap record by querying it with criteria.
func (c *Client) FindStockScrap(criteria *Criteria) (*StockScrap, error) {
	sss := &StockScraps{}
	if err := c.SearchRead(StockScrapModel, criteria, NewOptions().Limit(1), sss); err != nil {
		return nil, err
	}
	return &((*sss)[0]), nil
}

// FindStockScraps finds stock.scrap records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockScraps(criteria *Criteria, options *Options) (*StockScraps, error) {
	sss := &StockScraps{}
	if err := c.SearchRead(StockScrapModel, criteria, options, sss); err != nil {
		return nil, err
	}
	return sss, nil
}

// FindStockScrapIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockScrapIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(StockScrapModel, criteria, options)
}

// FindStockScrapId finds record id by querying it with criteria.
func (c *Client) FindStockScrapId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockScrapModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
