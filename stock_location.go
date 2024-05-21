package odoo

// StockLocation represents stock.location model.
type StockLocation struct {
	Active                   *Bool      `xmlrpc:"active,omitempty"`
	Barcode                  *String    `xmlrpc:"barcode,omitempty"`
	ChildIds                 *Relation  `xmlrpc:"child_ids,omitempty"`
	ChildInternalLocationIds *Relation  `xmlrpc:"child_internal_location_ids,omitempty"`
	Comment                  *String    `xmlrpc:"comment,omitempty"`
	CompanyId                *Many2One  `xmlrpc:"company_id,omitempty"`
	CompleteName             *String    `xmlrpc:"complete_name,omitempty"`
	CreateDate               *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                *Many2One  `xmlrpc:"create_uid,omitempty"`
	CyclicInventoryFrequency *Int       `xmlrpc:"cyclic_inventory_frequency,omitempty"`
	DisplayName              *String    `xmlrpc:"display_name,omitempty"`
	ForecastWeight           *Float     `xmlrpc:"forecast_weight,omitempty"`
	Id                       *Int       `xmlrpc:"id,omitempty"`
	IncomingMoveLineIds      *Relation  `xmlrpc:"incoming_move_line_ids,omitempty"`
	LastInventoryDate        *Time      `xmlrpc:"last_inventory_date,omitempty"`
	LocationId               *Many2One  `xmlrpc:"location_id,omitempty"`
	Name                     *String    `xmlrpc:"name,omitempty"`
	NetWeight                *Float     `xmlrpc:"net_weight,omitempty"`
	NextInventoryDate        *Time      `xmlrpc:"next_inventory_date,omitempty"`
	OutgoingMoveLineIds      *Relation  `xmlrpc:"outgoing_move_line_ids,omitempty"`
	ParentPath               *String    `xmlrpc:"parent_path,omitempty"`
	Posx                     *Int       `xmlrpc:"posx,omitempty"`
	Posy                     *Int       `xmlrpc:"posy,omitempty"`
	Posz                     *Int       `xmlrpc:"posz,omitempty"`
	PutawayRuleIds           *Relation  `xmlrpc:"putaway_rule_ids,omitempty"`
	QuantIds                 *Relation  `xmlrpc:"quant_ids,omitempty"`
	RemovalStrategyId        *Many2One  `xmlrpc:"removal_strategy_id,omitempty"`
	ReplenishLocation        *Bool      `xmlrpc:"replenish_location,omitempty"`
	ReturnLocation           *Bool      `xmlrpc:"return_location,omitempty"`
	ScrapLocation            *Bool      `xmlrpc:"scrap_location,omitempty"`
	StorageCategoryId        *Many2One  `xmlrpc:"storage_category_id,omitempty"`
	Usage                    *Selection `xmlrpc:"usage,omitempty"`
	ValuationInAccountId     *Many2One  `xmlrpc:"valuation_in_account_id,omitempty"`
	ValuationOutAccountId    *Many2One  `xmlrpc:"valuation_out_account_id,omitempty"`
	WarehouseId              *Many2One  `xmlrpc:"warehouse_id,omitempty"`
	WarehouseViewIds         *Relation  `xmlrpc:"warehouse_view_ids,omitempty"`
	WriteDate                *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                 *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// StockLocations represents array of stock.location model.
type StockLocations []StockLocation

// StockLocationModel is the odoo model name.
const StockLocationModel = "stock.location"

// Many2One convert StockLocation to *Many2One.
func (sl *StockLocation) Many2One() *Many2One {
	return NewMany2One(sl.Id.Get(), "")
}

// CreateStockLocation creates a new stock.location model and returns its id.
func (c *Client) CreateStockLocation(sl *StockLocation) (int64, error) {
	ids, err := c.CreateStockLocations([]*StockLocation{sl})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStockLocation creates a new stock.location model and returns its id.
func (c *Client) CreateStockLocations(sls []*StockLocation) ([]int64, error) {
	var vv []interface{}
	for _, v := range sls {
		vv = append(vv, v)
	}
	return c.Create(StockLocationModel, vv, nil)
}

// UpdateStockLocation updates an existing stock.location record.
func (c *Client) UpdateStockLocation(sl *StockLocation) error {
	return c.UpdateStockLocations([]int64{sl.Id.Get()}, sl)
}

// UpdateStockLocations updates existing stock.location records.
// All records (represented by ids) will be updated by sl values.
func (c *Client) UpdateStockLocations(ids []int64, sl *StockLocation) error {
	return c.Update(StockLocationModel, ids, sl, nil)
}

// DeleteStockLocation deletes an existing stock.location record.
func (c *Client) DeleteStockLocation(id int64) error {
	return c.DeleteStockLocations([]int64{id})
}

// DeleteStockLocations deletes existing stock.location records.
func (c *Client) DeleteStockLocations(ids []int64) error {
	return c.Delete(StockLocationModel, ids)
}

// GetStockLocation gets stock.location existing record.
func (c *Client) GetStockLocation(id int64) (*StockLocation, error) {
	sls, err := c.GetStockLocations([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*sls)[0]), nil
}

// GetStockLocations gets stock.location existing records.
func (c *Client) GetStockLocations(ids []int64) (*StockLocations, error) {
	sls := &StockLocations{}
	if err := c.Read(StockLocationModel, ids, nil, sls); err != nil {
		return nil, err
	}
	return sls, nil
}

// FindStockLocation finds stock.location record by querying it with criteria.
func (c *Client) FindStockLocation(criteria *Criteria) (*StockLocation, error) {
	sls := &StockLocations{}
	if err := c.SearchRead(StockLocationModel, criteria, NewOptions().Limit(1), sls); err != nil {
		return nil, err
	}
	return &((*sls)[0]), nil
}

// FindStockLocations finds stock.location records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockLocations(criteria *Criteria, options *Options) (*StockLocations, error) {
	sls := &StockLocations{}
	if err := c.SearchRead(StockLocationModel, criteria, options, sls); err != nil {
		return nil, err
	}
	return sls, nil
}

// FindStockLocationIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockLocationIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(StockLocationModel, criteria, options)
}

// FindStockLocationId finds record id by querying it with criteria.
func (c *Client) FindStockLocationId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockLocationModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
