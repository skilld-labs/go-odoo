package odoo

// StockFixedPutawayStrat represents stock.fixed.putaway.strat model.
type StockFixedPutawayStrat struct {
	LastUpdate      *Time     `xmlrpc:"__last_update,omitempty"`
	CategoryId      *Many2One `xmlrpc:"category_id,omitempty"`
	CreateDate      *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid       *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName     *String   `xmlrpc:"display_name,omitempty"`
	FixedLocationId *Many2One `xmlrpc:"fixed_location_id,omitempty"`
	Id              *Int      `xmlrpc:"id,omitempty"`
	PutawayId       *Many2One `xmlrpc:"putaway_id,omitempty"`
	Sequence        *Int      `xmlrpc:"sequence,omitempty"`
	WriteDate       *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid        *Many2One `xmlrpc:"write_uid,omitempty"`
}

// StockFixedPutawayStrats represents array of stock.fixed.putaway.strat model.
type StockFixedPutawayStrats []StockFixedPutawayStrat

// StockFixedPutawayStratModel is the odoo model name.
const StockFixedPutawayStratModel = "stock.fixed.putaway.strat"

// Many2One convert StockFixedPutawayStrat to *Many2One.
func (sfps *StockFixedPutawayStrat) Many2One() *Many2One {
	return NewMany2One(sfps.Id.Get(), "")
}

// CreateStockFixedPutawayStrat creates a new stock.fixed.putaway.strat model and returns its id.
func (c *Client) CreateStockFixedPutawayStrat(sfps *StockFixedPutawayStrat) (int64, error) {
	ids, err := c.CreateStockFixedPutawayStrats([]*StockFixedPutawayStrat{sfps})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStockFixedPutawayStrat creates a new stock.fixed.putaway.strat model and returns its id.
func (c *Client) CreateStockFixedPutawayStrats(sfpss []*StockFixedPutawayStrat) ([]int64, error) {
	var vv []interface{}
	for _, v := range sfpss {
		vv = append(vv, v)
	}
	return c.Create(StockFixedPutawayStratModel, vv, nil)
}

// UpdateStockFixedPutawayStrat updates an existing stock.fixed.putaway.strat record.
func (c *Client) UpdateStockFixedPutawayStrat(sfps *StockFixedPutawayStrat) error {
	return c.UpdateStockFixedPutawayStrats([]int64{sfps.Id.Get()}, sfps)
}

// UpdateStockFixedPutawayStrats updates existing stock.fixed.putaway.strat records.
// All records (represented by ids) will be updated by sfps values.
func (c *Client) UpdateStockFixedPutawayStrats(ids []int64, sfps *StockFixedPutawayStrat) error {
	return c.Update(StockFixedPutawayStratModel, ids, sfps, nil)
}

// DeleteStockFixedPutawayStrat deletes an existing stock.fixed.putaway.strat record.
func (c *Client) DeleteStockFixedPutawayStrat(id int64) error {
	return c.DeleteStockFixedPutawayStrats([]int64{id})
}

// DeleteStockFixedPutawayStrats deletes existing stock.fixed.putaway.strat records.
func (c *Client) DeleteStockFixedPutawayStrats(ids []int64) error {
	return c.Delete(StockFixedPutawayStratModel, ids)
}

// GetStockFixedPutawayStrat gets stock.fixed.putaway.strat existing record.
func (c *Client) GetStockFixedPutawayStrat(id int64) (*StockFixedPutawayStrat, error) {
	sfpss, err := c.GetStockFixedPutawayStrats([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*sfpss)[0]), nil
}

// GetStockFixedPutawayStrats gets stock.fixed.putaway.strat existing records.
func (c *Client) GetStockFixedPutawayStrats(ids []int64) (*StockFixedPutawayStrats, error) {
	sfpss := &StockFixedPutawayStrats{}
	if err := c.Read(StockFixedPutawayStratModel, ids, nil, sfpss); err != nil {
		return nil, err
	}
	return sfpss, nil
}

// FindStockFixedPutawayStrat finds stock.fixed.putaway.strat record by querying it with criteria.
func (c *Client) FindStockFixedPutawayStrat(criteria *Criteria) (*StockFixedPutawayStrat, error) {
	sfpss := &StockFixedPutawayStrats{}
	if err := c.SearchRead(StockFixedPutawayStratModel, criteria, NewOptions().Limit(1), sfpss); err != nil {
		return nil, err
	}
	return &((*sfpss)[0]), nil
}

// FindStockFixedPutawayStrats finds stock.fixed.putaway.strat records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockFixedPutawayStrats(criteria *Criteria, options *Options) (*StockFixedPutawayStrats, error) {
	sfpss := &StockFixedPutawayStrats{}
	if err := c.SearchRead(StockFixedPutawayStratModel, criteria, options, sfpss); err != nil {
		return nil, err
	}
	return sfpss, nil
}

// FindStockFixedPutawayStratIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockFixedPutawayStratIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(StockFixedPutawayStratModel, criteria, options)
}

// FindStockFixedPutawayStratId finds record id by querying it with criteria.
func (c *Client) FindStockFixedPutawayStratId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockFixedPutawayStratModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
