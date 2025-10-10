package odoo

// StockReplenishmentOption represents stock.replenishment.option model.
type StockReplenishmentOption struct {
	CreateDate          *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid           *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName         *String   `xmlrpc:"display_name,omitempty"`
	FreeQty             *Float    `xmlrpc:"free_qty,omitempty"`
	Id                  *Int      `xmlrpc:"id,omitempty"`
	LeadTime            *String   `xmlrpc:"lead_time,omitempty"`
	LocationId          *Many2One `xmlrpc:"location_id,omitempty"`
	ProductId           *Many2One `xmlrpc:"product_id,omitempty"`
	QtyToOrder          *Float    `xmlrpc:"qty_to_order,omitempty"`
	ReplenishmentInfoId *Many2One `xmlrpc:"replenishment_info_id,omitempty"`
	RouteId             *Many2One `xmlrpc:"route_id,omitempty"`
	Uom                 *String   `xmlrpc:"uom,omitempty"`
	WarehouseId         *Many2One `xmlrpc:"warehouse_id,omitempty"`
	WarningMessage      *String   `xmlrpc:"warning_message,omitempty"`
	WriteDate           *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid            *Many2One `xmlrpc:"write_uid,omitempty"`
}

// StockReplenishmentOptions represents array of stock.replenishment.option model.
type StockReplenishmentOptions []StockReplenishmentOption

// StockReplenishmentOptionModel is the odoo model name.
const StockReplenishmentOptionModel = "stock.replenishment.option"

// Many2One convert StockReplenishmentOption to *Many2One.
func (sro *StockReplenishmentOption) Many2One() *Many2One {
	return NewMany2One(sro.Id.Get(), "")
}

// CreateStockReplenishmentOption creates a new stock.replenishment.option model and returns its id.
func (c *Client) CreateStockReplenishmentOption(sro *StockReplenishmentOption) (int64, error) {
	ids, err := c.CreateStockReplenishmentOptions([]*StockReplenishmentOption{sro})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStockReplenishmentOption creates a new stock.replenishment.option model and returns its id.
func (c *Client) CreateStockReplenishmentOptions(sros []*StockReplenishmentOption) ([]int64, error) {
	var vv []interface{}
	for _, v := range sros {
		vv = append(vv, v)
	}
	return c.Create(StockReplenishmentOptionModel, vv, nil)
}

// UpdateStockReplenishmentOption updates an existing stock.replenishment.option record.
func (c *Client) UpdateStockReplenishmentOption(sro *StockReplenishmentOption) error {
	return c.UpdateStockReplenishmentOptions([]int64{sro.Id.Get()}, sro)
}

// UpdateStockReplenishmentOptions updates existing stock.replenishment.option records.
// All records (represented by ids) will be updated by sro values.
func (c *Client) UpdateStockReplenishmentOptions(ids []int64, sro *StockReplenishmentOption) error {
	return c.Update(StockReplenishmentOptionModel, ids, sro, nil)
}

// DeleteStockReplenishmentOption deletes an existing stock.replenishment.option record.
func (c *Client) DeleteStockReplenishmentOption(id int64) error {
	return c.DeleteStockReplenishmentOptions([]int64{id})
}

// DeleteStockReplenishmentOptions deletes existing stock.replenishment.option records.
func (c *Client) DeleteStockReplenishmentOptions(ids []int64) error {
	return c.Delete(StockReplenishmentOptionModel, ids)
}

// GetStockReplenishmentOption gets stock.replenishment.option existing record.
func (c *Client) GetStockReplenishmentOption(id int64) (*StockReplenishmentOption, error) {
	sros, err := c.GetStockReplenishmentOptions([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*sros)[0]), nil
}

// GetStockReplenishmentOptions gets stock.replenishment.option existing records.
func (c *Client) GetStockReplenishmentOptions(ids []int64) (*StockReplenishmentOptions, error) {
	sros := &StockReplenishmentOptions{}
	if err := c.Read(StockReplenishmentOptionModel, ids, nil, sros); err != nil {
		return nil, err
	}
	return sros, nil
}

// FindStockReplenishmentOption finds stock.replenishment.option record by querying it with criteria.
func (c *Client) FindStockReplenishmentOption(criteria *Criteria) (*StockReplenishmentOption, error) {
	sros := &StockReplenishmentOptions{}
	if err := c.SearchRead(StockReplenishmentOptionModel, criteria, NewOptions().Limit(1), sros); err != nil {
		return nil, err
	}
	return &((*sros)[0]), nil
}

// FindStockReplenishmentOptions finds stock.replenishment.option records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockReplenishmentOptions(criteria *Criteria, options *Options) (*StockReplenishmentOptions, error) {
	sros := &StockReplenishmentOptions{}
	if err := c.SearchRead(StockReplenishmentOptionModel, criteria, options, sros); err != nil {
		return nil, err
	}
	return sros, nil
}

// FindStockReplenishmentOptionIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockReplenishmentOptionIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(StockReplenishmentOptionModel, criteria, options)
}

// FindStockReplenishmentOptionId finds record id by querying it with criteria.
func (c *Client) FindStockReplenishmentOptionId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockReplenishmentOptionModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
