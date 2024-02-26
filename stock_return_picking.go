package odoo

// StockReturnPicking represents stock.return.picking model.
type StockReturnPicking struct {
	LastUpdate         *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate         *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid          *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName        *String   `xmlrpc:"display_name,omptempty"`
	Id                 *Int      `xmlrpc:"id,omptempty"`
	LocationId         *Many2One `xmlrpc:"location_id,omptempty"`
	MoveDestExists     *Bool     `xmlrpc:"move_dest_exists,omptempty"`
	OriginalLocationId *Many2One `xmlrpc:"original_location_id,omptempty"`
	ParentLocationId   *Many2One `xmlrpc:"parent_location_id,omptempty"`
	PickingId          *Many2One `xmlrpc:"picking_id,omptempty"`
	ProductReturnMoves *Relation `xmlrpc:"product_return_moves,omptempty"`
	WriteDate          *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid           *Many2One `xmlrpc:"write_uid,omptempty"`
}

// StockReturnPickings represents array of stock.return.picking model.
type StockReturnPickings []StockReturnPicking

// StockReturnPickingModel is the odoo model name.
const StockReturnPickingModel = "stock.return.picking"

// Many2One convert StockReturnPicking to *Many2One.
func (srp *StockReturnPicking) Many2One() *Many2One {
	return NewMany2One(srp.Id.Get(), "")
}

// CreateStockReturnPicking creates a new stock.return.picking model and returns its id.
func (c *Client) CreateStockReturnPicking(srp *StockReturnPicking) (int64, error) {
	ids, err := c.CreateStockReturnPickings([]*StockReturnPicking{srp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStockReturnPicking creates a new stock.return.picking model and returns its id.
func (c *Client) CreateStockReturnPickings(srps []*StockReturnPicking) ([]int64, error) {
	var vv []interface{}
	for _, v := range srps {
		vv = append(vv, v)
	}
	return c.Create(StockReturnPickingModel, vv, nil)
}

// UpdateStockReturnPicking updates an existing stock.return.picking record.
func (c *Client) UpdateStockReturnPicking(srp *StockReturnPicking) error {
	return c.UpdateStockReturnPickings([]int64{srp.Id.Get()}, srp)
}

// UpdateStockReturnPickings updates existing stock.return.picking records.
// All records (represented by ids) will be updated by srp values.
func (c *Client) UpdateStockReturnPickings(ids []int64, srp *StockReturnPicking) error {
	return c.Update(StockReturnPickingModel, ids, srp, nil)
}

// DeleteStockReturnPicking deletes an existing stock.return.picking record.
func (c *Client) DeleteStockReturnPicking(id int64) error {
	return c.DeleteStockReturnPickings([]int64{id})
}

// DeleteStockReturnPickings deletes existing stock.return.picking records.
func (c *Client) DeleteStockReturnPickings(ids []int64) error {
	return c.Delete(StockReturnPickingModel, ids)
}

// GetStockReturnPicking gets stock.return.picking existing record.
func (c *Client) GetStockReturnPicking(id int64) (*StockReturnPicking, error) {
	srps, err := c.GetStockReturnPickings([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*srps)[0]), nil
}

// GetStockReturnPickings gets stock.return.picking existing records.
func (c *Client) GetStockReturnPickings(ids []int64) (*StockReturnPickings, error) {
	srps := &StockReturnPickings{}
	if err := c.Read(StockReturnPickingModel, ids, nil, srps); err != nil {
		return nil, err
	}
	return srps, nil
}

// FindStockReturnPicking finds stock.return.picking record by querying it with criteria.
func (c *Client) FindStockReturnPicking(criteria *Criteria) (*StockReturnPicking, error) {
	srps := &StockReturnPickings{}
	if err := c.SearchRead(StockReturnPickingModel, criteria, NewOptions().Limit(1), srps); err != nil {
		return nil, err
	}
	return &((*srps)[0]), nil
}

// FindStockReturnPickings finds stock.return.picking records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockReturnPickings(criteria *Criteria, options *Options) (*StockReturnPickings, error) {
	srps := &StockReturnPickings{}
	if err := c.SearchRead(StockReturnPickingModel, criteria, options, srps); err != nil {
		return nil, err
	}
	return srps, nil
}

// FindStockReturnPickingIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockReturnPickingIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(StockReturnPickingModel, criteria, options)
}

// FindStockReturnPickingId finds record id by querying it with criteria.
func (c *Client) FindStockReturnPickingId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockReturnPickingModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
